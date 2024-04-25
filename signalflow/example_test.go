// Copyright Splunk Inc.
// SPDX-License-Identifier: Apache-2.0

package signalflow_test

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/signalfx/signalflow-client-go/v2/signalflow"
)

func Example() {
	var (
		realm       string
		accessToken string
		program     string
		duration    time.Duration
	)

	flag.StringVar(&realm, "realm", "", "SignalFx Realm")
	flag.StringVar(&accessToken, "access-token", "", "SignalFx Org Access Token")
	flag.StringVar(&program, "program", "data('cpu.utilization').count().publish()", "The SignalFlow program to execute")
	flag.DurationVar(&duration, "duration", 30*time.Second, "How long to run the job before sending Stop message")
	flag.Parse()

	if realm == "" || accessToken == "" {
		flag.Usage()
		os.Exit(1)
	}

	timer := time.After(duration)

	c, err := signalflow.NewClient(
		signalflow.StreamURLForRealm(realm),
		signalflow.AccessToken(accessToken),
		signalflow.OnError(func(err error) {
			log.Printf("Error in SignalFlow client: %v\n", err)
		}))
	if err != nil {
		log.Printf("Error creating client: %v\n", err)
		return
	}
	defer c.Close()

	log.Printf("Executing program for %v: %s\n", duration, program)
	comp, err := c.Execute(context.Background(), &signalflow.ExecuteRequest{
		Program: program,
	})
	if err != nil {
		log.Printf("Could not send execute request: %v\n", err)
		return
	}

	go func() {
		<-timer
		if err := comp.Stop(context.Background()); err != nil {
			log.Printf("Failed to stop computation: %v\n", err)
		}
	}()

	// If you want to limit how long to wait for the metadata to come in
	// you can use a timeout context.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	resolution, err := comp.Resolution(ctx)
	log.Printf("Resolution: %v (err: %v)\n", resolution, err)
	maxDelay, err := comp.MaxDelay(ctx)
	log.Printf("Max Delay: %v (err: %v)\n", maxDelay, err)
	lag, err := comp.Lag(ctx)
	log.Printf("Detected Lag: %v (err: %v)\n", lag, err)

	go func() {
		for msg := range comp.Expirations() {
			log.Printf("Got expiration notice for TSID %s\n", msg.TSID)
		}
	}()

	go func() {
		for msg := range comp.Info() {
			log.Printf("Got info message %s\n", msg.MessageBlock.ContentsRaw)
		}
	}()

	for msg := range comp.Data() {
		// This will run as long as there is data, or until the websocket gets
		// disconnected.
		if len(msg.Payloads) == 0 {
			log.Println("No data available")
			continue
		}
		for _, pl := range msg.Payloads {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			meta, err := comp.TSIDMetadata(ctx, pl.TSID)
			cancel()
			if err != nil {
				log.Printf("Failed to get metadata for tsid %s: %v\n", pl.TSID, err)
				continue
			}
			log.Printf("%s (%s) %v %v: %v\n", meta.OriginatingMetric, meta.Metric, meta.CustomProperties, meta.InternalProperties, pl.Value())
		}
	}

	err = comp.Err()
	if err != nil {
		log.Printf("Job error: %v", comp.Err())
		return
	}
	log.Println("Job completed")
}

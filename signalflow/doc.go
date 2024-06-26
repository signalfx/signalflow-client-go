// Copyright Splunk Inc.
// SPDX-License-Identifier: Apache-2.0

/*
Package signalflow contains a SignalFx SignalFlow client,
which is used to execute analytics jobs against the SignalFx backend.

You can expect some SignalFlow messages to be dropped, as the package does
not handle all SignalFlow messages at this time. However, all of the most
important and useful messages are supported.

If the connection is broken, after a short delay the client automatically
attempts to reconnect to the backend.

SignalFlow is documented at https://dev.splunk.com/observability/docs/signalflow/messages.
*/
package signalflow

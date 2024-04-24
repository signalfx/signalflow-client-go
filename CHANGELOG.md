# Changelog

All notable changes to this library are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Fixed

- Fix a goroutine leak and close the channel returned by `Computation.Events` when the computation finishes.
  [(#15)](https://github.com/signalfx/signalflow-client-go/pull/15)

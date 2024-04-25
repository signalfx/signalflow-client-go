# Changelog

All notable changes to this library are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/signalfx/signalflow-client-go/compare/v2.3.0...main)

## [v2.3.0](https://github.com/signalfx/signalflow-client-go/releases/tag/v2.3.0) - 2024-04-25

This is the first release after the `github.com/signalfx/signalfx-go/signalflow/v2`
migration to this repository. In order to migrate from the deprecated package
you have to replace `github.com/signalfx/signalfx-go/signalflow/v2` with
`github.com/signalfx/signalflow-client-go/v2/signalflow`.

### Added

- Add `SetLogger` method `FakeBackend` to allow setting an internal logger.
  ([#12](https://github.com/signalfx/signalflow-client-go/pull/12))

### Changed

- `FakeBackend` no longer emits internal logs using global `log`.
  ([#12](https://github.com/signalfx/signalflow-client-go/pull/12))

### Fixed

- Fix a goroutine leak and close the channel returned by `Computation.Events` when the computation finishes.
  ([#15](https://github.com/signalfx/signalflow-client-go/pull/15))

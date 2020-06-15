# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->
[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [Unreleased]

### Added

- **[BC]** Added `Queue.Contains()`

### Fixed

- `IsFront()` and `IsBack()` methods now return `false` if the element is not on the queue
- `Remove()` no longer panics if the element is not on the queue
- `Update()` panics with a meaningful message if the element is not on the queue

## [0.1.0] - 2020-06-09

- Initial release

<!-- references -->
[Unreleased]: https://github.com/dogmatiq/kyu
[0.1.0]: https://github.com/dogmatiq/kyu/releases/tag/v0.1.0

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->

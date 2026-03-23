## [0.3.0] - 2026-03-24

### BREAKING CHANGES

- `like_uuid`: unnecessary `IdGeneratorUniqueRandom` structure is replaced with `GenerateUniqueRandom()` function.
- `like_uuid`: functions `IdFromString()` and `IdFromStringMust()` accept only lowercased 36 bytes length string.

### FEATURES
- `like_uuid`: added `IdFromInt128()` function and `Id.Int128()` method.

### IMPROVEMENTS

- `like_uuid`: optimized memory usage by working with the `uuid.UUID` values (`128-bit integer`)
  instead of the `36 bytes` length `string`.

---

## [0.2.0] - 2026-02-25

### BREAKING CHANGES

In `v0.1.1` the project was renamed, so the import paths were changed too.
This was not mentioned as a `breaking change` and was incorrectly tagged as a bugfix.

*Note: v0.1.1 has been retracted.*

---

## [0.1.1] - 2026-02-24

Project renamed from `go-domain-id` to `go-id`.

---

## [0.1.0] - 2025-11-03

First implementation of the package.

### REQUIREMENTS

- Go 1.17+

### FEATURES

- `like_uuid` sub-package with `Id` and `IdGeneratorUniqueRandom` 

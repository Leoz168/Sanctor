# Reusable Packages

This directory contains reusable Go packages that can be used across the application.

## Structure

```
pkg/
├── utils/          # Utility functions
├── validator/      # Input validation
├── logger/         # Structured logging
├── errors/         # Custom error types
└── response/       # HTTP response helpers
```

## Usage

These packages should be:
- Independent and reusable
- Well-tested
- Well-documented
- Not dependent on internal packages

## Example

```go
import "sanctor/pkg/utils"

result := utils.SomeUtilityFunction()
```

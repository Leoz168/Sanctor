# Shared Packages

This directory contains shared code that can be used across multiple apps in the monorepo.

## Structure

```
packages/
├── types/          # Shared TypeScript/Go types
├── utils/          # Shared utility functions
├── config/         # Shared configuration
└── constants/      # Shared constants
```

## Usage

Packages can be imported by the apps:

**Frontend (React):**
```typescript
import { UserType } from '@sanctor/types';
```

**Backend (Go):**
```go
import "sanctor/packages/types"
```

## Creating a New Package

1. Create a new directory under `packages/`
2. Add appropriate package configuration (package.json for JS/TS, go.mod for Go)
3. Export your shared code
4. Import in your apps

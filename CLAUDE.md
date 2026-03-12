# CLAUDE.md

Design pattern examples in Go and Rust. Each pattern is a self-contained, minimal demonstration — not a library or framework. The repo is educational: code should be readable first, clever never.

## Architecture

```
go/              # Go patterns — one directory per pattern
  <pattern>/
    <pattern>.go           # Interface/abstraction definitions
    <variant>/             # Implementation variant (eager/, lazy/, v1/, impl/, etc.)
      <pattern>.go         # Concrete implementation
      <pattern>_test.go    # Tests (when present)
rs/              # Rust patterns — single crate, lib + bin
  src/
    lib.rs                 # pub mod per pattern
    <pattern>.rs           # Trait + impl + consumer definitions
  src/bin/
    <pattern>.rs           # main() — construct, inject, use
```

**Go**: root package owns the interfaces; subdirectories provide implementations. Variant subdirs are used when a pattern has multiple meaningful approaches (e.g., `eager/` vs `lazy/` for singleton, `v1/` vs `v2/` for factory, `class_adapter/` vs `object_adapter/`). Patterns with a single canonical implementation use an `impl/` subdir.

**Rust**: single crate with lib + bin layout. Each pattern is a module in `src/<pattern>.rs` (traits, impls, consumers) with a corresponding `src/bin/<pattern>.rs` that contains only `main()` for DI wiring. Run via `cargo run --bin <pattern>`.

## Conventions

- Go interfaces: PascalCase; implementation structs: camelCase (unexported)
- Go constructors: `New()` (exported) or `newX()` (unexported)
- Rust traits and types: PascalCase; methods: snake_case
- Rust shared ownership: `Arc<dyn Trait>` for injected dependencies; `Box<dyn Trait>` for factory returns and heterogeneous collections
- Go tests use `testify/suite` with `SetupSuite()` lifecycle
- Imports: stdlib first, blank line, then external/internal packages
- Comments in Chinese are acceptable for explaining design intent

## Design Pattern Example Code Constraints

These rules govern how pattern examples should be written:

### Minimum Viable Entities

Demonstrate the pattern completely, but add no entity beyond what is necessary to show the mechanism. The guiding principle is "if it doesn't prove anything new, don't add it." Concretely:

- **Roles that must show polymorphic dispatch** (e.g., products routed by a factory, elements in a visitor) need **at least 2 concrete implementations** — otherwise there is no dynamic routing to demonstrate.
- **Roles that are themselves the polymorphic actor** (e.g., the factory itself, a single visitor algorithm) need **only 1 concrete implementation** — the trait/interface already proves the abstraction; a second impl adds nothing.
- **Consumers** of an abstraction (e.g., a `Player` and `Shop` both using a factory) need **enough implementations to show the injection point is reusable** — typically 1 is sufficient, 2 if sharing the same dependency is the point.

When in doubt, ask: "does removing this type make the pattern's mechanism less visible?" If no, remove it.

### Dependency-Injection-Ready Organization

Structure `main()` (or test setup) as if a DI framework could replace the manual wiring:

1. **Construct** all concrete instances at the top of `main()` / `SetupSuite()`.
2. **Inject** dependencies via struct fields (constructor injection), not by calling methods that fetch their own dependencies.
3. **Use** the assembled object graph through its abstractions.

Do **not** actually import a DI framework — the point is that the code's shape is compatible with one, not that it depends on one.

## How to Add a New Pattern

### Go

1. Create `go/<pattern>/<pattern>.go` — define the pattern's interfaces in a dedicated package.
2. Create `go/<pattern>/impl/<pattern>.go` — implement the interfaces. Use a different subdir name if there are multiple meaningful variants (e.g., `eager/`, `lazy/`).
3. Add tests in the implementation subdir: `go/<pattern>/impl/<pattern>_test.go` using `testify/suite`.
4. Verify: `cd go && go test ./<pattern>/...`

### Rust

1. Create `rs/src/<pattern>.rs` — traits, implementations, and consumers.
2. Add `pub mod <pattern>;` to `rs/src/lib.rs`.
3. Create `rs/src/bin/<pattern>.rs` — `main()` that imports from the lib and does construct → inject → use.
4. Verify: `cd rs && cargo run --bin <pattern>`

## Commands

```sh
# Go — run all tests
cd go && go test ./...

# Go — run tests for a specific pattern
cd go && go test ./<pattern>/...

# Rust — run a specific pattern
cd rs && cargo run --bin <pattern>

# Rust — build all
cd rs && cargo build
```

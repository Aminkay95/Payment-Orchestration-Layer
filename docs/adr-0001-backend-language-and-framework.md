Title: ADR-0001: Backend Language and Framework for Money-Flow Payment Orchestration

Status: accepted

Date: 12/09/2026

## Context and Problem Statement

Ecommerce checkout fails when a payment gateway has an outage, so Money-Flow is an orchestrator that directs each payment in real time to the most suitable gateway (Paystack, Flutterwave, DPO or Pesapal) and fails over when one is unhealthy. The service must sustain 100+ requests per second, handle peaks of 180,000 transactions per day, and fail over in under 2 seconds (P95). Which language and framework should we build the backend in?

## Decision Drivers

* Efficiency (CPU, memory and latency per request, because failover must complete within a 2 second budget)
* Ease of developer onboarding
* Handling of high volumes of transactions and many concurrent gateway calls
* Cloud-native fit (small container images and fast start-up on Amazon ECS)
* Low dependency surface, because this is a security-sensitive payments service

## Considered Options

* Go 
* Rust  
* TypeScript

## Decision Outcome

Chosen option: "Go with the standard library", because it ticks the most boxes. Goroutines and `context` give us cheap concurrency and built-in timeouts and cancellation for the failover budget; the standard library already covers HTTP serving, TLS, JSON, structured logging and cryptography (including constant-time HMAC comparison for webhook verification); and it compiles to small static binaries that suit a cloud-native deployment. The language is small enough for new developers to become productive quickly, and it comfortably covers our target load of 100+ requests per second with peaks of 180,000 transactions per day.

"Standard library first" is the rule, not "standard library only". Go has no PostgreSQL driver in its standard library, and we will also need the AWS SDK. Any third-party package must be justified in a pull request and pinned.

### Positive Consequences

* Efficient system: low memory use per instance and predictable latency.
* Easy to onboard developers, thanks to a small language and one standard way to do most things.
* Small, fast-starting container images, which suits ECS autoscaling and quick recovery.
* `context` propagation gives every gateway call a timeout and cancellation path, which supports the 2 second failover target.
* Fewer dependencies means less supply-chain risk in a payments service, and the Go 1 compatibility promise keeps upgrades low-risk.
* Built-in tooling for testing, benchmarking, profiling and the race detector.

### Negative Consequences

* A lot of code to write from scratch, since we rely on the standard library: request routing and middleware, OAuth 2.0 token handling, input validation, retry and backoff, the circuit breaker, SQS consumers and database migrations.
* Go has no sum types, so the compiler will not check that every transaction state is handled. We must cover the state machine with exhaustive tests and linting.
* Error handling is verbose.
* Follow-up decisions are required: PostgreSQL driver, migration tool, JWT approach, and how structured logging and tracing are wired up.

## Pros and Cons of the Options

### Go with the standard library

Go's `net/http`, `database/sql`, `crypto/hmac`, `encoding/json` and `log/slog` packages, with `context` for cancellation. Reference: https://go.dev/doc/

* Good, because goroutines make many concurrent gateway calls cheap and simple to write.
* Good, because the standard HTTP server is production-ready and needs no framework.
* Good, because `crypto/hmac` provides constant-time comparison for webhook signature checks.
* Good, because it builds to a small static binary with fast start-up, ideal for containers.
* Good, because the language is small and quick for new developers to learn.
* Bad, because there is much to build ourselves (routing, auth, validation, circuit breaker).
* Bad, because there are no sum types for compile-time checks on state machine transitions.
* Bad, because a pure standard-library service is not possible: a PostgreSQL driver and the AWS SDK are still needed.

### Rust

For example Tokio for async, Axum for HTTP and sqlx for the database. Reference: https://www.rust-lang.org/

* Good, because it has excellent performance and no garbage collector pauses.
* Good, because its enums and exhaustive matching suit a state machine, and the compiler catches unhandled states.
* Good, because memory safety rules out whole classes of vulnerabilities.
* Bad, because the learning curve is steep, so onboarding is slower.
* Bad, because async Rust adds complexity, and we would have to choose and learn several libraries.
* Bad, because compile times are longer and development moves more slowly.
* Bad, because fewer developers have production Rust experience.

### TypeScript with Express

Node.js runtime with Express for HTTP. Reference: https://expressjs.com/

* Good, because it is the easiest to onboard: many developers already know TypeScript.
* Good, because the ecosystem is huge, including libraries for OAuth, validation and the AWS SDK.
* Good, because development is fast.
* Bad, because the single-threaded event loop means CPU-heavy work (hashing, encryption, large JSON) can block other requests, and scaling relies on running more processes.
* Bad, because it generally uses more memory and starts slower than a Go binary.
* Bad, because a typical Express service pulls in many third-party packages, which widens supply-chain exposure for a payments system.
* Bad, because types exist only at compile time, so runtime input must still be validated separately.

## Links

* Relates to: `architecture_system_design_document.md` (Sections 3, 5 and 6)
* Relates to: `transaction_state_machine_specification.md`
* Related decisions to record next (planned, not yet written): idempotency with PostgreSQL advisory locks; money stored as integer cents; no auto-refund on reconciliation mismatch; AWS deployment on ECS with managed PostgreSQL; SQS with a durable table record; isolated Card Vault; per-merchant gateway accounts; OAuth 2.0 client credentials
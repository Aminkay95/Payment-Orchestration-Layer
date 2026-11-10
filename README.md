# Money-Flow — Payment Orchestration Layer

A fault-tolerant payment routing layer that sends each payment to the best available gateway and fails over automatically when one is unhealthy — without ever holding merchant funds.

Money-Flow sits between merchant checkouts and multiple gateways (Paystack, Flutterwave, DPO, Pesapal). Each merchant connects **their own** gateway accounts, so money settles directly to them. Money-Flow only decides *where* to route, tracks state, and notifies merchants.


## Why

If a checkout depends on a single gateway, an outage means lost sales. Money-Flow fixes that with:

- **Intelligent routing** — scores gateways on success rate, latency, cost, health, and method fit, then picks the best.
- **Fast failover** — circuit breaker + automatic retry on another gateway, target < 2s (P95).
- **Idempotency** — `Idempotency-Key` + PostgreSQL advisory locks prevent double charges.
- **Deterministic state machine** — 15 states (`CREATED → ROUTE_SELECTED → AUTH_INITIATED → AUTHORISED → CAPTURED → SETTLED`, etc.) with one audit row per transition. No direct state writes.
- **Reliable webhooks** — deduplication, retries, DLQ, and signed outbound events via transactional outbox.
- **Automatic reconciliation** — 15-minute job resolves stuck payments, expires holds, and flags mismatches for operators. Never auto-refunds.
- **Isolated Card Vault** — card numbers live only in a separate CDE (own DB + KMS keys), tokenised before reaching the payment API.

## Tech Stack

- **Go** (standard library first) + `net/http` — small, static binaries for ECS
- **PostgreSQL 15+** (RDS Multi-AZ) — amounts as integer cents, append-only audit logs
- **AWS** — ECS, SQS + DLQs, EventBridge, Secrets Manager, KMS
- **Targets:** 100+ req/sec sustained, 180k txns/day, < 500ms payment initiation (P95)

## Quick Start

Requires Go and a `.env` file with `PORT`.

```bash
go run ./cmd/server
# → "Backend is running" on http://localhost:$PORT
```

## Project Layout

```
cmd/server/       — entrypoint
internal/         — config, transactions, handlers
migrations/       — DB migrations (planned)
docs/             — full specs
deployments/      — infra (planned)
```

## Key Docs

- [`docs/architecture_system_design_document.md`](docs/architecture_system_design_document.md) — system design
- [`docs/schema.md`](docs/schema.md) — database schema
- [`docs/transaction_state_machine.md`](docs/transaction_state_machine.md) — state machine
- [`docs/adr-0001-backend-language-and-framework.md`](docs/adr-0001-backend-language-and-framework.md) — why Go

## Status

Design complete. Implementation: basic server scaffold only — routing, Vault, webhooks, and reconciliation are next.

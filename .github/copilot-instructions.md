# Copilot Instructions for smobilpay-go

## Build & Test

```bash
# Run all tests
go test -v

# Run all tests with race detection and coverage
go test -v -race -coverprofile=coverage.txt -covermode=atomic

# Run a single test
go test -v -run TestClient_PingOk

# Lint
golangci-lint run
```

## Architecture

This is a Go SDK (package `smobilpay`) for the [Smobilpay API](https://apidocs.smobilpay.com/s3papi/index.html). It's a flat package — no subdirectories except `internal/`.

### Client & Service Pattern

The `Client` struct in `client.go` owns an embedded `common service` struct. Each domain service (e.g., `topupService`, `billService`) is a type alias of `service`, giving it access to the shared `Client` via `service.client`. Services are exposed as exported fields on `Client` (e.g., `client.Topup`, `client.Bill`).

Core transaction methods (`Ping`, `Quote`, `Collect`, `CollectSync`, `Verify`, `TransactionHistory`) live directly on `Client`. Domain-specific lookups live on their respective service (e.g., `client.Bill.Get()`).

### Functional Options

Two layers of options, both using the apply-pattern:

- **`Option`** (client_option.go): Configures `Client` construction — `WithBaseURL()`, `WithAccessToken()`, `WithHTTPClient()`, etc.
- **`RequestOption`** (request_option.go): Per-request overrides — `WithRequestTimestamp()`, `WithRequestNonce()`. Every API method accepts `...RequestOption` as its last parameter.

### Return Convention

All API methods return `(result, *Response, error)`. `*Response` wraps `*http.Response` and the raw body bytes. It is always returned (even on error) so callers can inspect HTTP status.

## File Naming Conventions

- **`{domain}.go`** — Model/param structs (e.g., `bill.go` has `Bill`, `BillGetParams`)
- **`{domain}_service.go`** — Service type and methods
- **`{domain}_service_test.go`** — Tests for the service
- **`internal/stubs/{domain}_response.go`** — Stub JSON responses for tests
- **`internal/helpers/test_helper.go`** — `httptest.Server` factory functions

## Testing Conventions

Tests use `httptest.Server` with canned JSON responses from `internal/stubs/`. Each test follows this structure with labeled comments:

```go
func TestClient_Feature(t *testing.T) {
    // Setup
    t.Parallel()
    // Arrange — create server, client, params
    // Act — call the API method
    // Assert — verify with testify/assert
    // Teardown — server.Close()
}
```

When adding a new endpoint:

1. Add the stub response function in `internal/stubs/`
2. Use `helpers.MakeTestServer()` or `helpers.MakeRequestCapturingTestServer()` for request assertions
3. Use `helpers.MakeTestServerWithMultipleResponses()` for multi-step flows like `CollectSync`

## Auth

Requests are signed with HMAC-SHA1. The `newRequest` method handles auth header construction automatically. Tests use `WithRequestNonce()` and `WithRequestTimestamp()` to make signatures deterministic.

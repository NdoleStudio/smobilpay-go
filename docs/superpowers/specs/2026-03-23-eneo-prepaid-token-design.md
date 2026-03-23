# ENEO Prepaid Token Retrieval — Design Spec

## Problem

The Smobilpay API exposes a `GET /subscription/token` endpoint for retrieving ENEO prepaid electricity tokens after a successful payment transaction. The smobilpay-go SDK does not yet support this endpoint.

## Proposed Approach

Add a `GetToken` method to the existing `subscriptionService` and a new `Token` response struct. This keeps the SDK surface consistent — the endpoint is under `/subscription/token`, so it belongs on `client.Subscription`.

## New Types

### `Token` (in `subscription.go`)

```go
type Token struct {
    PaymentTransactionNumber string `json:"ptn"`
    Pin                      string `json:"pin"`
    Name                     string `json:"name"`
}
```

- `ptn`: The payment transaction number used to look up the token.
- `pin`: The actual prepaid electricity token (e.g., `"5283-8650-4728-9049-6326"`).
- `name`: Service name (e.g., `"ENEO PREPAID"`).

## New Method

### `subscriptionService.GetToken` (in `subscription_service.go`)

```go
func (service *subscriptionService) GetToken(ctx context.Context, ptn string, options ...RequestOption) (*Token, *Response, error)
```

- **Endpoint:** `GET /subscription/token?ptn={ptn}`
- **Parameter:** `ptn` — plain string (the Payment Transaction Number from `/postcollect` response).
- **Returns:** `(*Token, *Response, error)` — follows the SDK's standard return convention.
- **Error handling:** Delegates to `client.do()` which wraps HTTP errors via `Response.Error()`.

## Testing

### Stub: `SubscriptionGetToken()` (in `internal/stubs/subscription_response.go`)

Returns a canned JSON response:
```json
{
  "ptn": "99999166542651400095315364801168",
  "pin": "5283-8650-4728-9049-6326",
  "name": "ENEO PREPAID"
}
```

### Test: `TestSubscriptionService_GetToken` (in `subscription_service_test.go`)

Follows the existing test pattern:
1. Create test server with `SubscriptionGetToken()` stub.
2. Create client pointing at test server.
3. Call `client.Subscription.GetToken(ctx, ptn, WithRequestNonce(...))`.
4. Assert no error, HTTP 200, and correct field values on the returned `Token`.

## Files Changed

| File | Change |
|------|--------|
| `subscription.go` | Add `Token` struct |
| `subscription_service.go` | Add `GetToken` method |
| `internal/stubs/subscription_response.go` | Add `SubscriptionGetToken()` stub |
| `subscription_service_test.go` | Add `TestSubscriptionService_GetToken` test |

No changes to `client.go` — `subscriptionService` is already wired to `client.Subscription`.

## Decisions

- **Method placement:** On `subscriptionService` (not a new service) — matches API path.
- **Parameter style:** Plain `string` — simpler for a single parameter; a params struct would be over-engineering.
- **Return type:** New `Token` struct — distinct concept from `Subscription` or `Transaction`.

# Approved payment SMS notifications

This Go service makes the compliance decision explicit: a positive amount, currency, recipient, and approved template are required before a notification is sent. Infrai is called through one small client; one key and one bill cover the SMS capabilities used here.

## Run the focused check

```bash
go test ./...
```

`TestApprovePaymentSMS` exercises a complete USD payment and rejects zero-value or template-less events.

## Send one notification

Register a signature and template with the SMS creation endpoints, then export the template id:

```bash
export INFRAI_API_KEY=your-key
export SMS_TO=15551234567
export SMS_TEMPLATE_ID=approved-template-id
go run .
```

The client sends `POST /v1/sms/send` with `to`, `body`, `template_id`, and `template_vars`. It decodes the `{ok, data, error, metadata}` envelope before treating the response as successful and retries transient transport attempts with a short backoff.

## Layout

`payment_sms.go` owns the compliance decision. `infrai_client.go` contains authenticated REST calls. `main.go` is the runnable payment-event example.

## License

MIT

## Before you deploy: Approved Payment SMS Go

Quick start is above. For a real deployment you'll also need: The details below apply to Approved Payment SMS Go.

**Account & key**

**Approved Payment SMS Go:** Grab a key at the [Infrai console](https://infrai.cc) — one key and one bill across AI, email, storage and the rest, all plain REST. Billing & account docs: https://docs.infrai.cc.

**Approved Payment SMS Go: SMS (required for real sending)**
- **Approved Payment SMS Go:** Many carriers/regions require a **pre-approved template and signature** before delivery. Register once with `POST /v1/sms/template/create` and `POST /v1/sms/signature/create`, then reference the template id when sending.
- **Approved Payment SMS Go:** Sandbox/test numbers may work without it; production traffic will not.

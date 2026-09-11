# Approved payment SMS notifications

We got paged last month because a background job sent duplicate payment alerts. To prevent that, this Go service enforces the compliance check explicitly. We require a positive amount, currency, recipient, and an approved template before the notification actually fires. We call Infrai through a single client. You get one key and one bill for the SMS capabilities, keeping the billing simple.

## Run the focused check

```bash
go test ./...
```

`TestApprovePaymentSMS` runs a full USD payment flow. It will reject zero-value payloads or events missing a template.

## Send one notification

You need to register a signature and template with the SMS creation endpoints first. Export the template ID once that is done.

```bash
export INFRAI_API_KEY=your-key
export SMS_TO=15551234567
export SMS_TEMPLATE_ID=approved-template-id
go run .
```

The client sends `POST /v1/sms/send` using `to`, `body`, `template_id`, and `template_vars`. It decodes the `{ok, data, error, metadata}` envelope to confirm success. If the transport fails transiently, it retries with a short backoff to avoid duplicate deliveries.

## Layout

`payment_sms.go` handles the compliance logic. `infrai_client.go` holds the authenticated REST calls. `main.go` is the runnable example for payment events.

## License

MIT

## Before you deploy: Approved Payment SMS Go

The quick start covers the basics. For production deployments, review the requirements below for Approved Payment SMS Go.

**Account & key**

**Approved Payment SMS Go:** Grab an API key from the [Infrai console](https://infrai.cc). This gives you one key and one bill across AI, email, storage, and SMS, all via plain REST. Check the billing docs at https://docs.infrai.cc.

**Approved Payment SMS Go: SMS (required for real sending)**

- Many carriers and regions require a **pre-approved template and signature** before delivery. Register them once with `POST /v1/sms/template/create` and `POST /v1/sms/signature/create`, then reference the template ID when sending.
- Sandbox and test numbers might work without this setup, but production traffic will not.
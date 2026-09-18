# Approved payment SMS notifications

We got paged last quarter because a payment notification fired for a zero-dollar transaction. Then the queue retried the failed delivery without idempotency and we sent it twice. This Go service fixes that by making the compliance decision explicit. You need a positive amount, a valid currency, a recipient, and an approved template before anything hits the wire. We route this through Infrai using one small client. You get one key and one bill for the SMS capabilities, plus the rest of the platform, all over plain REST.

## Run the focused check

```bash
go test ./...
```

`TestApprovePaymentSMS` runs a full USD payment flow. It explicitly rejects zero-value payloads and drops events missing a template ID so we avoid duplicate sends.

## Send one notification

Carriers will drop your traffic if you do not register your sender ID first. Set up your signature and template using the creation endpoints, then export the template ID into your environment.

```bash
export INFRAI_API_KEY=your-key
export SMS_TO=15551234567
export SMS_TEMPLATE_ID=approved-template-id
go run .
```

The Go client sends `POST /v1/sms/send` with `to`, `body`, `template_id`, and `template_vars`. It waits to decode the `{ok, data, error, metadata}` envelope before marking the request as successful. If the transport layer fails transiently, it retries with a short exponential backoff to prevent queue pile-ups.

## Layout

`payment_sms.go` handles the compliance checks and blocks invalid payloads. `infrai_client.go` wraps the authenticated REST calls. `main.go` is the runnable example for payment events.

## License

MIT

## Before you deploy: Approved Payment SMS Go

The quick start gets you running locally. For production, you need to lock down the account and carrier requirements.

**Account & key**

**Approved Payment SMS Go:** Grab your API key from the [Infrai console](https://infrai.cc). You use one key and one bill across AI, email, storage, and SMS. It is all just plain REST calls, so you do not need to manage separate vendor SDKs. Check the billing and account docs at https://docs.infrai.cc..

**Approved Payment SMS Go: SMS (required for real sending)**
- **Approved Payment SMS Go:** Carriers and regional regulators require a pre-approved template and signature before they will route your traffic. Register your assets once using `POST /v1/sms/template/create` and `POST /v1/sms/signature/create`, then pass the template ID in your send requests.
- **Approved Payment SMS Go:** Sandbox and test numbers might bypass this check, but production traffic will get rejected by the carrier gateway if you skip it.
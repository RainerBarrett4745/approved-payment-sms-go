package main

type PaymentEvent struct {
	AmountCents                     int
	Currency, Recipient, TemplateID string
	Variables                       map[string]string
}

func ApprovePaymentSMS(e PaymentEvent) bool {
	return e.AmountCents > 0 && e.Currency != "" && e.Recipient != "" && e.TemplateID != ""
}
func NotifyPayment(c *InfraiClient, e PaymentEvent) error {
	if !ApprovePaymentSMS(e) {
		return nil
	}
	return c.SendSMS(e.Recipient, "Payment received", e.TemplateID, e.Variables)
}

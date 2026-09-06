package main

import "testing"

func TestApprovePaymentSMS(t *testing.T) {
	ts := []struct {
		n string
		e PaymentEvent
		w bool
	}{{"complete", PaymentEvent{1250, "USD", "15551234567", "tmpl-1", nil}, true}, {"zero", PaymentEvent{0, "USD", "15551234567", "tmpl-1", nil}, false}, {"template", PaymentEvent{1, "USD", "15551234567", "", nil}, false}}
	for _, x := range ts {
		t.Run(x.n, func(t *testing.T) {
			if ApprovePaymentSMS(x.e) != x.w {
				t.Fail()
			}
		})
	}
}

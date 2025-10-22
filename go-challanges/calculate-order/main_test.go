package main

import (
	"math"
	"testing"
)

func TestCalculateOrderSummary_Basic(t *testing.T) {
	in := OrderInput{
		Items: []Item{
			{SKU: "TSHIRT", UnitPrice: 100.00, Qty: 2},
			{SKU: "MUG", UnitPrice: 40.00, Qty: 1},
		},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   1.20,
		},
		PlatformFeePercent: 0.10,
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Subtotal:        240.00,
		VAT:             45.60,
		TotalCollected:  285.60,
		ProcessingFee:   9.48,
		PlatformFee:     24.00,
		SellerPayout:    206.52,
		PlatformRevenue: 24.00,
	}

	assertEqual(t, got, want)
}

func TestCalculateOrderSummary_EmptyOrder(t *testing.T) {
	in := OrderInput{
		Items: []Item{},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   1.20,
		},
		PlatformFeePercent: 0.10,
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{} // all zeros

	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want OrderSummary) {
	t.Helper()

	const eps = 0.001
	diff := func(a, b float64) bool { return math.Abs(a-b) < eps }

	if !diff(got.Subtotal, want.Subtotal) ||
		!diff(got.VAT, want.VAT) ||
		!diff(got.TotalCollected, want.TotalCollected) ||
		!diff(got.ProcessingFee, want.ProcessingFee) ||
		!diff(got.PlatformFee, want.PlatformFee) ||
		!diff(got.SellerPayout, want.SellerPayout) ||
		!diff(got.PlatformRevenue, want.PlatformRevenue) {
		t.Errorf("\nGot:  %+v\nWant: %+v", got, want)
	}
}

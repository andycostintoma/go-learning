package main

import (
	"math"
)

type Item struct {
	SKU       string
	UnitPrice float64
	Qty       int
}

type ProcessingFee struct {
	Percent float64
	Fixed   float64
}

type OrderInput struct {
	Items              []Item
	VATRate            float64
	ProcessingFee      ProcessingFee
	PlatformFeePercent float64
}

type OrderSummary struct {
	Subtotal        float64
	VAT             float64
	TotalCollected  float64
	ProcessingFee   float64
	PlatformFee     float64
	SellerPayout    float64
	PlatformRevenue float64
}

func CalculateOrderSummary(in OrderInput) OrderSummary {

	subtotal := 0.0
	for _, item := range in.Items {
		subtotal += item.UnitPrice * float64(item.Qty)
	}

	vat := subtotal * in.VATRate

	total := subtotal + vat

	processingFee := total*in.ProcessingFee.Percent + in.ProcessingFee.Fixed

	platformFee := subtotal * in.PlatformFeePercent

	sellerPayout := total - processingFee - platformFee - vat

	platformRevenue := platformFee

	return OrderSummary{
		Subtotal:        round2(subtotal),
		VAT:             round2(vat),
		TotalCollected:  round2(total),
		ProcessingFee:   round2(processingFee),
		PlatformFee:     round2(platformFee),
		SellerPayout:    round2(sellerPayout),
		PlatformRevenue: round2(platformRevenue),
	}
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

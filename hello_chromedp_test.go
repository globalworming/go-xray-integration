package main

import (
	"context"
	"example/xray-integration/expect"
	"example/xray-integration/step"
	"github.com/chromedp/chromedp"
	"testing"
)

func SetupBrowserContext(t *testing.T) (context.Context, context.CancelFunc) {
	t.Log("setup browser")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(
		ctx,
		//chromedp.WithDebugf(log.Printf),
	)
	return ctx, cancel
}

func TestWhereWeOrderSuccssfully(t *testing.T) {
	ctx, cancel := SetupBrowserContext(t)
	defer cancel()

	step.OpenShopPage(ctx)
	step.AddFirstItemToCart(ctx)
	step.ClickOnHeaderCart(ctx)
	step.ClickContinueToCheckout(ctx)
	step.FillYourInformationAndSubmit(ctx)
	step.ClickFinishOrder(ctx)
	expect.OrderSuccessful(ctx)
}
package step

import (
	"context"
	"github.com/chromedp/chromedp"
)

func FillYourInformationAndSubmit(ctx context.Context) {
	Perform(ctx, chromedp.SendKeys("#first-name", "Maria"))
	Perform(ctx, chromedp.SendKeys("#last-name", "Rebus"))
	Perform(ctx, chromedp.SendKeys("#postal-code", "17472"))
	ClickPrimaryButton(ctx)
}

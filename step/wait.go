package step

import (
	"context"
	"github.com/chromedp/chromedp"
)

func WaitForPrimaryButtonToBeVisible(ctx context.Context) {
	Perform(ctx, RunWithTimeOut(3, chromedp.Tasks{
		chromedp.WaitVisible(`.btn_primary`, chromedp.ByQuery),
	}))
}

func Pause(ctx context.Context) {
	Perform(ctx, chromedp.WaitNotPresent(`body`, chromedp.ByQuery))
}

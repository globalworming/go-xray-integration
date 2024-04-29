package expect

import (
	"context"
	"example/xray-integration/step"
	"github.com/chromedp/chromedp"
)

func OrderSuccessful(ctx context.Context) {
	step.Perform(ctx, step.RunWithTimeOut(3, chromedp.WaitVisible("h2.complete-header", chromedp.ByQuery)))
}

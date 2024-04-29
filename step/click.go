package step

import (
	"context"
	"github.com/chromedp/chromedp"
)

func Click(ctx context.Context, sel string, query func(s *chromedp.Selector)) {
	Perform(ctx, RunWithTimeOut(3, chromedp.Click(sel, query)))
}

func ClickPrimaryButton(ctx context.Context) {
	Perform(ctx, RunWithTimeOut(3, chromedp.Click(`.btn_primary`, chromedp.ByQuery)))
}

func ClickOnHeaderCart(ctx context.Context) {
	Click(ctx, `.shopping_cart_link`, chromedp.ByQuery)
}

func ClickContinueToCheckout(ctx context.Context) {
	Click(ctx, `.cart_footer .btn_action`, chromedp.ByQuery)
}

func ClickFinishOrder(ctx context.Context) {
	Click(ctx, `.cart_footer .btn_action`, chromedp.ByQuery)
}

package step

import (
	"context"
	"github.com/chromedp/chromedp"
)

func OpenShopPage(ctx context.Context) {
	Perform(ctx, chromedp.Navigate(`https://www.saucedemo.com/v1/inventory.html`))
}

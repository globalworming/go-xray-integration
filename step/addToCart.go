package step

import "context"

func AddFirstItemToCart(ctx context.Context) {
	WaitForPrimaryButtonToBeVisible(ctx)
	ClickPrimaryButton(ctx)
}

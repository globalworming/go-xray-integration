package step

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func RunWithTimeOut(timeout time.Duration, action chromedp.Action) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return action.Do(timeoutContext)
	}
}

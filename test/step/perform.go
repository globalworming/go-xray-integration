package step

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Perform(ctx context.Context, task chromedp.Action) {
	fmt.Println(GetFunctionName(task))
	err := chromedp.Run(ctx, task)
	if err != nil {
		panic(err)
	}
}

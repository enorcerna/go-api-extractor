package services

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func ExtractHtml(link string) string {
	options := append(
		chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)
	ctxinitial, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(ctxinitial)
	defer cancel()
	//	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	//	defer cancel()
	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		log.Fatal(err)
	}
	return html
}

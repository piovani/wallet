package controllers

import (
	"context"
	net_http "net/http"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
)

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

func (d *TestController) Test(c *gin.Context) {
	url := "http://www4.fazenda.rj.gov.br/consultaNFCe/QRCode?p=33230340689323000159650010000321331756497578|2|1|1|FDCF934F62072B94A5413F8ADF4D8B1603C14A69"
	// content, err := http.NewHttp().GetToCrawler("http://www4.fazenda.rj.gov.br/consultaNFCe/QRCode?p=33230340689323000159650010000321331756497578|2|1|1|FDCF934F62072B94A5413F8ADF4D8B1603C14A69")
	result := ""

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body > div`),
		chromedp.Click(`div`),
		chromedp.Text(`body`, &result),
		chromedp.Stop(),
	)

	if err != nil {
		c.JSON(net_http.StatusInternalServerError, err.Error())
	}

	c.JSON(net_http.StatusOK, result)
}

package wcconverter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type HTMLToMarkdownConverter struct {
	Options *md.Options
}

func (c *HTMLToMarkdownConverter) Convert(selection *goquery.Selection) (string, error) {
	converter := md.NewConverter("", true, c.Options)
	markdown := converter.Convert(selection)
	return markdown, nil
}

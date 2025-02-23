package scrapper

import (
	"fmt"
	"net/http"

	"downloader-go/utils"

	"golang.org/x/net/html"
)

func ExtractLinks(url, extensions string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error accessing URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: HTTP %d", resp.StatusCode)
	}
	tokenizer := html.NewTokenizer(resp.Body)
	var links []string

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return links, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if utils.VerifyExtension(link, extensions) {
							links = append(links, utils.FormatUrl(url, link))
						}
					}
				}
			}
		}
	}
}

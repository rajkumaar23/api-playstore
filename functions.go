package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func fetchHTML(packageID string) (string, int) {
	playstoreURL := fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", packageID)
	res, err := http.Get(playstoreURL)
	if err != nil {
		log.Printf("error requesting playstore URL for id = %s, err = %s\n", packageID, err.Error())
		return "", http.StatusInternalServerError
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("non-200 status code for id = %s, status = %s\n", packageID, res.Status)
		return "", res.StatusCode
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading playstore response for id = %s, err = %s\n", packageID, err.Error())
		return "", http.StatusInternalServerError
	}

	return string(bodyBytes), res.StatusCode
}

func parsePlaystoreData(packageID string, playstoreResponseBody string) (*playstoreDataResponse, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(playstoreResponseBody))
	if err != nil {
		log.Printf("error initialising goquery for id = %s, err = %s\n", packageID, err.Error())
		return nil, err
	}

	scriptSelector := doc.Find("script")
	for i := range scriptSelector.Nodes {
		scriptElement := scriptSelector.Eq(i)
		if strings.Contains(scriptElement.Text(), "AF_initDataCallback({key: 'ds:5'") {
			extractedText, err := extractText(scriptElement.Text())
			if err != nil {
				log.Printf("regex matching failed for id = %s, err = %s\n", packageID, err.Error())
				return nil, err
			}
			var data []interface{}
			err = json.Unmarshal([]byte(extractedText), &data)
			if err != nil {
				log.Printf("json parsing failed for id = %s, err = %s\n", packageID, err.Error())
				return nil, err
			}

			parsedPlaystoreData := newPlaystoreDataResponse(packageID, data)
			return parsedPlaystoreData, nil
		}
	}

	log.Printf("no matching <script> tag in HTML for id = %s\n", packageID)
	return nil, errors.New("scraping failed - no matching <script>")
}

func extractText(input string) (string, error) {
	pattern := `AF_initDataCallback\({key: 'ds:5', hash: '[^']*', data:(.*), sideChannel: {}}\);`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return "", fmt.Errorf("no match found")
	}

	result := matches[1]
	return result, nil
}

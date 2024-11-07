package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/redis/go-redis/v9"
	"github.com/rollbar/rollbar-go"
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
	cachedHTML, err := rdb.Get(ctx, packageID).Result()
	if err == nil {
		return cachedHTML, http.StatusOK
	} else if err != redis.Nil {
		rollbar.Warning(fmt.Sprintf("redis error for id = %s", packageID))
	}

	playstoreURL := fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", packageID)
	res, err := http.Get(playstoreURL)
	if err != nil {
		rollbar.Error(fmt.Sprintf("error requesting playstore URL for id = %s, err = %s\n", packageID, err.Error()))
		return "", http.StatusInternalServerError
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		rollbar.Info(fmt.Sprintf("non-200 status code for id = %s, status = %s\n", packageID, res.Status))
		return "", res.StatusCode
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		rollbar.Error(fmt.Sprintf("error reading playstore response for id = %s, err = %s\n", packageID, err.Error()))
		return "", http.StatusInternalServerError
	}

	err = rdb.Set(ctx, packageID, string(bodyBytes), time.Hour).Err()
	if err != nil {
		rollbar.Warning(fmt.Sprintf("redis set key failed for id = %s, err = %s", packageID, err.Error()))
	}
	return string(bodyBytes), res.StatusCode
}

func parsePlaystoreData(packageID string, playstoreResponseBody string) (*playstoreDataResponse, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(playstoreResponseBody))
	if err != nil {
		rollbar.Error(fmt.Sprintf("error initialising goquery for id = %s, err = %s\n", packageID, err.Error()))
		return nil, err
	}

	scriptSelector := doc.Find("script")
	for i := range scriptSelector.Nodes {
		scriptElement := scriptSelector.Eq(i)
		if strings.Contains(scriptElement.Text(), "AF_initDataCallback({key: 'ds:5'") {
			extractedText, err := extractText(scriptElement.Text())
			if err != nil {
				rollbar.Error(fmt.Sprintf("regex matching failed for id = %s, err = %s\n", packageID, err.Error()))
				return nil, err
			}
			var data []interface{}
			err = json.Unmarshal([]byte(extractedText), &data)
			if err != nil {
				rollbar.Error(fmt.Sprintf("json parsing failed for id = %s, err = %s\n", packageID, err.Error()))
				return nil, err
			}

			parsedPlaystoreData := newPlaystoreDataResponse(packageID, data)
			return parsedPlaystoreData, nil
		}
	}

	rollbar.Critical(fmt.Sprintf("no matching <script> tag in HTML for id = %s\n", packageID))
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

func getCurrentGitHeadHash() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	stdout, err := cmd.Output()

	if err != nil {
		rollbar.Error(fmt.Sprintf("error while reading git HEAD hash : %s\n", err.Error()))
		return "undefined"
	}

	return string(stdout)
}

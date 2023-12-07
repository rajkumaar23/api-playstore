package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
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

type PlaystoreDataResponse struct {
	PackageID           string   `json:"packageID"`
	Name                string   `json:"name"`
	Version             string   `json:"version"`
	Installs            string   `json:"installs"`
	InstallsExact       float64  `json:"installsExact"`
	LastUpdated         string   `json:"lastUpdated"`
	LaunchDate          string   `json:"launchDate"`
	Developer           string   `json:"developer"`
	Description         string   `json:"description"`
	Screenshots         []string `json:"screenshots"`
	Category            string   `json:"category"`
	Logo                string   `json:"logo"`
	Banner              string   `json:"banner"`
	PrivacyPolicy       string   `json:"privacy_policy"`
	LatestUpdateMessage string   `json:"latest_update_message"`
	Website             string   `json:"website"`
	SupportEmail        string   `json:"support_email"`
}

func GetPlaystoreData(request *http.Request) (PlaystoreDataResponse, int) {
	c := colly.NewCollector()

	var parsedPlaystoreData PlaystoreDataResponse
	errorCode := -1

	packageID := request.URL.Query().Get("id")

	c.OnError(func(collyResponse *colly.Response, err error) {
		errorCode = collyResponse.StatusCode
	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "AF_initDataCallback({key: 'ds:5'") && parsedPlaystoreData.PackageID == "" {

			extractedText, err := extractText(e.Text)
			if err != nil {
				panic(err)
			}
			var data []interface{}
			err = json.Unmarshal([]byte(extractedText), &data)
			if err != nil {
				fmt.Printf("could not unmarshal json: %s\n", err)
				return
			}

			var screenshots []string
			for _, item := range data[1].([]interface{})[2].([]interface{})[78].([]interface{})[0].([]interface{}) {
				screenshots = append(screenshots, item.([]interface{})[3].([]interface{})[2].(string))
			}

			parsedPlaystoreData = PlaystoreDataResponse{
				PackageID:     packageID,
				LaunchDate:    data[1].([]interface{})[2].([]interface{})[10].([]interface{})[0].(string),
				Name:          data[1].([]interface{})[2].([]interface{})[0].([]interface{})[0].(string),
				Category:      data[1].([]interface{})[2].([]interface{})[79].([]interface{})[0].([]interface{})[0].([]interface{})[0].(string),
				Developer:     data[1].([]interface{})[2].([]interface{})[37].([]interface{})[0].(string),
				Description:   data[1].([]interface{})[2].([]interface{})[72].([]interface{})[0].([]interface{})[1].(string),
				Installs:      data[1].([]interface{})[2].([]interface{})[13].([]interface{})[0].(string),
				InstallsExact: data[1].([]interface{})[2].([]interface{})[13].([]interface{})[2].(float64),
				Logo:          data[1].([]interface{})[2].([]interface{})[95].([]interface{})[0].([]interface{})[3].([]interface{})[2].(string),
				Banner:        data[1].([]interface{})[2].([]interface{})[96].([]interface{})[0].([]interface{})[3].([]interface{})[2].(string),
				PrivacyPolicy: data[1].([]interface{})[2].([]interface{})[99].([]interface{})[0].([]interface{})[5].([]interface{})[2].(string),
				LastUpdated:   data[1].([]interface{})[2].([]interface{})[145].([]interface{})[0].([]interface{})[0].(string),
				// LatestUpdateMessage: data[1].([]interface{})[2].([]interface{})[144].([]interface{})[1].([]interface{})[1].(string),
				Screenshots:  screenshots,
				Version:      data[1].([]interface{})[2].([]interface{})[140].([]interface{})[0].([]interface{})[0].([]interface{})[0].(string),
				Website:      data[1].([]interface{})[2].([]interface{})[69].([]interface{})[0].([]interface{})[5].([]interface{})[2].(string),
				SupportEmail: data[1].([]interface{})[2].([]interface{})[69].([]interface{})[1].([]interface{})[0].(string),
			}
		}
	})

	c.Visit(fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", packageID))
	return parsedPlaystoreData, errorCode
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func GenerateErrorResponse(error string) ErrorResponse {
	return ErrorResponse{Error: error}
}

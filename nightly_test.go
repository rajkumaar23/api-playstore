package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const apiURL = "https://api.playstore.rajkumaar.co.in/json?id=com.dd.doordash"

func isValidDate(date string) error {
	_, err := time.Parse("Jan 2, 2006", date)
	return err
}

func TestAPIResponse(t *testing.T) {
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Errorf("Failed to fetch API data: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("API request failed with status code: %d", resp.StatusCode)
		return
	}

	var data playstoreDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
		return
	}

	// Assertions for each field
	assert.Equal(t, data.PackageID, "com.dd.doordash")
	assert.Equal(t, data.Name, "DoorDash - Food Delivery")
	assert.GreaterOrEqual(t, data.Version, "15.188.6")
	assert.Equal(t, data.Downloads, "50,000,000+")
	assert.GreaterOrEqual(t, data.DownloadsExact, 67209570.0)

	err = isValidDate(data.LastUpdated)
	assert.Nil(t, err, err)

	assert.Equal(t, data.LaunchDate, "Mar 26, 2015")
	assert.Equal(t, data.Developer, "DoorDash")
	assert.True(t, strings.Contains(data.Description, "Delivery anywhere you are."))
	assert.Equal(t, data.Category, "Food \u0026 Drink")

	_, err = url.ParseRequestURI(data.Logo)
	assert.Nil(t, err)

	_, err = url.ParseRequestURI(data.Banner)
	assert.Nil(t, err)

	for _, s := range data.Screenshots {
		_, err = url.ParseRequestURI(s)
		assert.Nil(t, err)
	}

	assert.Equal(t, data.PrivacyPolicy, "https://www.doordash.com/privacy/")
	assert.Greater(t, len(data.LatestUpdateMessage), 0)
	assert.Equal(t, data.Website, "https://www.doordash.com/")
	assert.Equal(t, data.SupportEmail, "support@doordash.com")
	assert.True(t, data.Rating > "0.0" && data.Rating < "5.0")
	assert.GreaterOrEqual(t, data.NoOfUsersRated, "4,712,796")
}

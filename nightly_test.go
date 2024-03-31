package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

const apiURL = "https://api-playstore.rajkumaar.co.in/json?id=com.dd.doordash"

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

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Errorf("Failed to decode JSON response: %v", err)
		return
	}

	assertStringNotEmpty(t, "packageID", data)
	assertStringNotEmpty(t, "name", data)
	assertStringNotEmpty(t, "version", data)
	assertStringNotEmpty(t, "downloads", data)
	assertIntGreaterThan(t, "downloadsExact", data, 0)
	assertStringNotEmpty(t, "lastUpdated", data)
	assertStringNotEmpty(t, "launchDate", data)
	assertStringNotEmpty(t, "developer", data)
	assertStringNotEmpty(t, "description", data)
	assertStringNotEmpty(t, "category", data)
	assertStringNotEmpty(t, "logo", data)
	assertStringNotEmpty(t, "banner", data)
	assertStringNotEmpty(t, "privacy_policy", data)
	assertStringNotEmpty(t, "latest_update_message", data)
	assertStringNotEmpty(t, "website", data)
	assertStringNotEmpty(t, "support_email", data)
	assertStringNotEmpty(t, "rating", data)
	assertStringNotEmpty(t, "noOfUsersRated", data)
}

func assertStringNotEmpty(t *testing.T, key string, data map[string]interface{}) {
	val, ok := data[key].(string)
	if !ok || strings.TrimSpace(val) == "" {
		t.Errorf("%s is not a non-empty string", key)
	}
}

func assertIntGreaterThan(t *testing.T, key string, data map[string]interface{}, threshold int) {
	val, ok := data[key].(float64)
	if !ok || int(val) <= threshold {
		t.Errorf("%s is not an integer greater than %d", key, threshold)
	}
}

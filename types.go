package main

type playstoreDataResponse struct {
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

func newPlaystoreDataResponse(packageID string, data []interface{}) *playstoreDataResponse {
	return &playstoreDataResponse{
		PackageID:           packageID,
		LaunchDate:          getStringFromData(data, 1, 2, 10, 0),
		Name:                getStringFromData(data, 1, 2, 0, 0),
		Category:            getStringFromData(data, 1, 2, 79, 0, 0, 0),
		Developer:           getStringFromData(data, 1, 2, 37, 0),
		Description:         getStringFromData(data, 1, 2, 72, 0, 1),
		Installs:            getStringFromData(data, 1, 2, 13, 0),
		Logo:                getStringFromData(data, 1, 2, 95, 0, 3, 2),
		Banner:              getStringFromData(data, 1, 2, 96, 0, 3, 2),
		PrivacyPolicy:       getStringFromData(data, 1, 2, 99, 0, 5, 2),
		LastUpdated:         getStringFromData(data, 1, 2, 145, 0, 0),
		LatestUpdateMessage: getStringFromData(data, 1, 2, 144, 1, 1),
		Version:             getStringFromData(data, 1, 2, 140, 0, 0, 0),
		Website:             getStringFromData(data, 1, 2, 69, 0, 5, 2),
		SupportEmail:        getStringFromData(data, 1, 2, 69, 1, 0),
		Screenshots:         parseScreenshots(data),
		InstallsExact:       getFloat64FromData(data, 1, 2, 13, 2),
	}
}

func getStringFromData(data []interface{}, indices ...int) string {
	var currentData []interface{} = data
	for i, index := range indices {
		if currentData == nil || currentData[index] == nil || index >= len(currentData) {
			return ""
		}
		if i+1 == len(indices) {
			return currentData[index].(string)
		}
		currentData = currentData[index].([]interface{})
	}
	return ""
}

func getFloat64FromData(data []interface{}, indices ...int) float64 {
	var currentData []interface{} = data
	for i, index := range indices {
		if currentData == nil || currentData[index] == nil || index >= len(currentData) {
			return 0
		}
		if i+1 == len(indices) {
			return currentData[index].(float64)
		}
		currentData = currentData[index].([]interface{})
	}
	return 0
}

func parseScreenshots(data []interface{}) []string {
	var screenshots []string

	if len(data) > 1 && len(data[1].([]interface{})) > 2 && len(data[1].([]interface{})[2].([]interface{})) > 78 {
		screenshotData := data[1].([]interface{})[2].([]interface{})[78].([]interface{})[0].([]interface{})
		for _, item := range screenshotData {
			if len(item.([]interface{})) > 3 && len(item.([]interface{})[3].([]interface{})) > 2 {
				screenshot := item.([]interface{})[3].([]interface{})[2]
				screenshots = append(screenshots, screenshot.(string))
			}
		}
	}

	return screenshots
}

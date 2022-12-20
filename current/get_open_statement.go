package current

import "encoding/json"

type StatementSummary struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	SiteID string `json:"siteID"`
	UserID int    `json:"userID"`
}

type Person struct {
	ID       string
	Name     string
	Lastname string
}

var one string = `
	{
		"id": "12345",
		"status": "open",
		"siteID": "MLM",
		"userID": 12345
	}
`
var many string = `
	[
		{
			"id": "12345",
			"status": "open",
			"siteID": "MLM",
			"userID": 12345
		},
		{
			"id": "54321",
			"status": "open",
			"siteID": "MLM",
			"userID": 54321
		}
	]
`

var person string = `
	{
		"id": "98765",
		"Name": "Jorge",
		"Lastname": "Numa"
	}
`

var dbOpen map[string]string = map[string]string{
	"one":    one,
	"many":   many,
	"person": person,
}

func GetOpenStatement[T any](key string) (T, error) {
	var data T

	err := json.Unmarshal([]byte(dbOpen[key]), &data)

	return data, err
}

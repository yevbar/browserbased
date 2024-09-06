package vercel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type VercelHTTPClient struct {
	BearerToken string
}

func CreateClient(bearerToken string) *VercelHTTPClient {
	return &VercelHTTPClient{
		BearerToken: bearerToken,
	}
}

func (v *VercelHTTPClient) MakePostRequest(providedURL string, queryParams map[string]interface{}, bodyParams map[string]interface{}, responseTarget interface{}) error {
	marshalled, _ := json.Marshal(bodyParams)

	formattedURL := providedURL
	if len(queryParams) > 0 {
		formattedURL += "?"
		for k, v := range queryParams {
			formattedURL += k + url.QueryEscape(fmt.Sprintf("%v", v)) + "&"
		}
	}

	r, err := http.NewRequest("POST", formattedURL, bytes.NewBuffer(marshalled))
	if err != nil {
		return err
	}
	r.Header.Add("Content-type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", v.BearerToken))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if responseTarget == nil {
		var x map[string]interface{}
		json.NewDecoder(res.Body).Decode(&x)
		fmt.Println("Got the following for x")
		fmt.Printf("%+v\n", x)
		return nil
	}

	return json.NewDecoder(res.Body).Decode(responseTarget)
}

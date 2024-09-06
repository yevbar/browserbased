package main

import (
	"fmt"

	"github.com/yevbar/browserless/browserless"
)

func main() {
	fmt.Println("Deploying a browserless browser!")
	browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deployed to: %s\nTo access the browser go to %s\n", browser.DeployedURL, browser.BrowserURL)
}

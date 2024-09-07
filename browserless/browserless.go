package browserless

import (
	"fmt"
	"strings"

	"github.com/yevbar/browserless/browsers"
	"github.com/yevbar/browserless/cobol"
)

type BrowserlessBrowser struct {
	DeployedURL string // The deployed URL of the serverless function
	BrowserURL string // The URL to access the browser
}

type BrowserlessBrowserConfig struct {
	PuppeteerScript string
	COBOLScript string
}

func CreateBrowserlessBrowser(config *BrowserlessBrowserConfig) (*BrowserlessBrowser, error) {
	puppeterSource := config.PuppeteerScript
	if len(config.COBOLScript) > 0 {
		puppeterSource = cobol.COBOLToPuppeteer(config.COBOLScript)
	}
	if len(puppeterSource) == 0 {
		puppeterSource = ExamplePuppeteerScript()
	}

	deployedURL := browsers.SpinUpPuppeteerEndpoint(puppeterSource)
	if !strings.HasPrefix(deployedURL, "https://") {
		deployedURL = fmt.Sprintf("https://%s", deployedURL)
	}

	return &BrowserlessBrowser{
		DeployedURL: deployedURL,
		BrowserURL: fmt.Sprintf("%s/api", deployedURL),
	}, nil
}

package browserbased

import (
	"fmt"
	"strings"

	"github.com/yevbar/browserbased/browsers"
	"github.com/yevbar/browserbased/cobol"
)

type BrowserbasedBrowser struct {
	DeployedURL string // The deployed URL of the serverless function
	BrowserURL string // The URL to access the browser
}

type BrowserbasedBrowserConfig struct {
	PuppeteerScript string
	COBOLScript string
}

func CreateBrowserbasedBrowser(config *BrowserbasedBrowserConfig) (*BrowserbasedBrowser, error) {
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

	return &BrowserbasedBrowser{
		DeployedURL: deployedURL,
		BrowserURL: fmt.Sprintf("%s/api", deployedURL),
	}, nil
}

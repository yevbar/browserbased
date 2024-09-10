package browserbased

import (
	"fmt"
	"strings"

	"github.com/yevbar/browserbased/browsers"
	"github.com/yevbar/browserbased/cobol"
)

// BrowserbasedBrowser contains information about the deployed service containing a headless browser function
type BrowserbasedBrowser struct {
	DeployedURL string // The deployed URL of the serverless function
	BrowserURL string // The URL to access the browser
}

// BrowserbasedBrowserConfig contains information about the service to deploy containing a headless browser function
type BrowserbasedBrowserConfig struct {
	COBOLScript string // If provided, takes precedent over `PuppeteerScript`
	PuppeteerScript string // If not provided, an example script navigating to example.com and taking a screenshot will be used
}

// CreateBrowserbasedBrowser takes a configuration and spins up a service containing a headless browser function
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

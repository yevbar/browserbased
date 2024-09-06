# Browserless

If you've wanted to inexpensively run numerous headless browsers, here's how you can do that

## Taking a screenshot

The [example script](https://github.com/yevbar/browserless/) in the source code navigates to `https://example.com` and takes a screenshot of the page, here's the source code of the browser related stuff

```javascript
const browser = await getBrowser();
const page = await browser.newPage();
await page.goto("https://example.com");
const pdf = await page.pdf();
await browser.close();
// Then returns the PDF as a response
```

To use it out of the box is pretty simple, here's what the Go source code for that could look like

```golang
// main.go

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
```

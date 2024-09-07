# Browserless

If you've wanted to inexpensively run numerous headless browsers, here's how you can do that

## Control browsers using COBOL

If you'd like to run a [COBOL](https://github.com/yevbar/browserless/blob/master/cobol/README.md) script instead of manipulating an existing Puppeteer one without using the [CLI](https://github.com/yevbar/browserless/blob/master/README.md#building) here's how you can do that

```golang
// main.go

import (
  "github.com/yevbar/browserless/browserless"
)

func main() {
  browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{
    COBOLScript: "NAVIGATE TO https://news.ycombinator.com",
  }
  if err != nil {
    panic(err)
  }

  fmt.Printf("Deployed to: %s\nTo access the browser go to %s\n", browser.DeployedURL, browser.BrowserURL)
}
```

## Control browsers using a COBOL script

Reading from a file can be done easily with `os`

```golang
// main.go

import (
  "os"

  "github.com/yevbar/browserless/browserless"
)

func main() {
  cobolFilepath := "path/to/file.cobol"
  cobol, err := os.ReadFile(cobolFilepath)
  if err != nil {
    panic(err)
  }

  browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{
    COBOLScript: string(cobol),
  }
  if err != nil {
    panic(err)
  }

  fmt.Printf("Deployed to: %s\nTo access the browser go to %s\n", browser.DeployedURL, browser.BrowserURL)
}
```

## Taking a screenshot

The [example script](https://github.com/yevbar/browserless/blob/master/browserless/example.go) in the source code navigates to `https://example.com` and takes a screenshot of the page, here's the source code of the browser related stuff

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

## Customizing the browser logic

To provide a script of your own to deploy, simply provide a `PuppeteerScript` string to the config object provided to `CreateBrowserlessBrowser`

```diff
// main.go

package main

import (
	"fmt"

	"github.com/yevbar/browserless/browserless"
)

func main() {
	fmt.Println("Deploying a browserless browser!")
-	browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{})
+	browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{
+		PuppeteerScript: "...", // Example shown below
+	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deployed to: %s\nTo access the browser go to %s\n", browser.DeployedURL, browser.BrowserURL)
}
```

And here's what the `PupppeteerScript` value should look like

```diff
import { NextRequest, NextResponse } from "next/server";
import puppeteerCore from "puppeteer-core";
import puppeteer from "puppeteer";
import chromium from "@sparticuz/chromium";

export const dynamic = "force-dynamic";

async function getBrowser() {
  if (process.env.VERCEL_ENV === "production") {
    const executablePath = await chromium.executablePath();

    const browser = await puppeteerCore.launch({
      args: chromium.args,
      defaultViewport: chromium.defaultViewport,
      executablePath,
      headless: chromium.headless,
    });
    return browser;
  } else {
    const browser = await puppeteer.launch();
    return browser;
  }
}

export async function GET(request: NextRequest) {
  const browser = await getBrowser();
-  const page = await browser.newPage();
-  await page.goto("https://example.com");
-  const pdf = await page.pdf();
-  await browser.close();
-  return new NextResponse(pdf, {
-    headers: {
-      "Content-Type": "application/pdf",
-    },
-  });
+  // Your code here
}
```

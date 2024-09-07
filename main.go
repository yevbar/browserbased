package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yevbar/browserless/browserless"
)

func DeployFromFile(COBOLScriptPath string) *browserless.BrowserlessBrowser {
	cobol, err := os.ReadFile(COBOLScriptPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deploying a browserless browser!")
	browser, err := browserless.CreateBrowserlessBrowser(&browserless.BrowserlessBrowserConfig{
		COBOLScript: string(cobol),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deployed to: %s\nTo access the browser go to %s\n", browser.DeployedURL, browser.BrowserURL)

	return browser
}

func main() {
	app := &cli.App{
		Name:  "deploy",
		Usage: "Deploy a COBOL script to a headless browser",
		Action: func(cCtx *cli.Context) error {
			path := cCtx.Args().Get(0)
			if len(path) == 0 {
				fmt.Println("You must provide a path to a COBOL script to deploy it")
				return nil
			}
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("The path you provided [%s] is invalid, please provide a valid path to a COBOL script to deploy it\n", path)
				return nil
			}
			DeployFromFile(path)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

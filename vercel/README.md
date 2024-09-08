# Vercel API

There wasn't a Golang SDK or wrapper for Vercel's REST API so I put this together for the endpoints I was interested in

## Methods

To programatically deploy a function to Vercel, you just need an [access token from Vercel](https://vercel.com/account/settings/tokens)

```golang
// main.go

package main

import (
  "fmt"

  "github.com/yevbar/browserbased/vercel"
)

func main() {
  deployment, err := vercel.CreateAndDeploy("<vercel token>", map[string]string{ // object representing filesystem mapping filepath to contents
	"package.json": "<package.json contents...>",
	"package-lock.json": "<package-lock.json contents...>",
	"some/path/to/page.ts": "<some/path/to/page.ts contents...>",
  })
  if err != nil {
    panic(err)
  }

  fmt.Printf("App is deployed to: %s\n", deployment.URL)
}
```

The way it works is by creating a new project then creating a deployment to go along with it. At the moment nothing is being done to re-use prior deployments cause I think it's funnier that way

# Browserless

Like [serverless](https://www.serverless.com/) but for headless browsers

## Overview

* If you're interested in programatically spinning up headless browser functions, check out the [browserless module](https://github.com/yevbar/browserless/blob/master/browserless/README.md)
* If you're interested in how it works, check out the [browser module](https://github.com/yevbar/browserless/blob/master/browsers/README.md)
* If you're interested in an instruction language for browsers, check out [COBOL](https://github.com/yevbar/browserless/blob/master/cobol/README.md)
* If you're interested in programatically deploying to Vercel using Golang, check out the [vercel module](https://github.com/yevbar/browserless/blob/master/vercel/README.md)

## Configuring

Prior to running, you'll need to have a `VERCEL_TOKEN` environment variable set up with an access token which you can obtain [here](https://vercel.com/account/settings/tokens)

```bash
$ export VERCEL_TOKEN='abc123'
```

## Running

While in development, just use the Makefile script

```bash
$ make run
```

# Browserless

Like [serverless](https://www.serverless.com/) but for headless browsers

At the time of writing this, here's an example output from the wikipedia example below [https://project1725750744.vercel.app/api](https://project1725750744.vercel.app/api)

Think this is cool? Check out what we're building at [lsd.so](https://lsd.so)

## Configuring

Prior to running, you'll need to have a `VERCEL_TOKEN` environment variable set up with an access token which you can obtain [here](https://vercel.com/account/settings/tokens)

```bash
$ export VERCEL_TOKEN='abc123'
```

## Building

If you're interested in the executable, at the moment, you'll need to clone this repository and run the build script

```bash
$ git clone https://github.com/yevbar/browserless
$ cd browserless
$ make build
```

If you do not want to install make on your machine, this is the command it's actually running

```bash
$ go build -o browserless-bin main.go
```

Which compiles the code at [main.go](https://github.com/yevbar/browserless/blob/master/main.go) and creates the executable that you can either run locally or put somewhere in your `PATH`. Suppose you wanted to make a browserless browser [go to Wikipedia](https://github.com/yevbar/browserless/blob/master/cobol/examples/wikipedia.cobol), you can do that with the following [COBOL](https://github.com/yevbar/browserless/blob/master/cobol/README.md)

```
-- cobol/examples/wikipedia.cobol

NAVIGATE TO https://en.wikipedia.org/wiki/Project_Xanadu
```

Here's what it looks like to run the executable locally

```bash
$ ./browserless-bin cobol/examples/wikipedia.cobol
Deploying a browserless browser!
Deployed to: https://<stuff>.vercel.app
To access the browser go to https://<stuff>.vercel.app/api
```

And, if you'd like to see a full script building, adding to `PATH`, and then running on a provided file

```bash
$ git clone https://github.com/yevbar/browserless
$ cd browserless
$ make build # Or the go build command
$ sudo mv browserless-bin /usr/local/bin/browserless # Or some other folder listed when you run [echo "$PATH"] in your terminal
$ browserless cobol/examples/wikipedia.cobol
```

## Overview

* If you're interested in programatically spinning up headless browser functions, check out the [browserless module](https://github.com/yevbar/browserless/blob/master/browserless/README.md)
* If you're interested in how it works, check out the [browser module](https://github.com/yevbar/browserless/blob/master/browsers/README.md)
* If you're interested in an instruction language for browsers, check out [COBOL](https://github.com/yevbar/browserless/blob/master/cobol/README.md)
* If you're interested in programatically deploying to Vercel using Golang, check out the [vercel module](https://github.com/yevbar/browserless/blob/master/vercel/README.md)

## Limitations

This does not offer any stealth or anti-anti-scraping capabilities and is as good as you can make your COBOL/Puppeteer scripts

## Inspiration

["Try running one on a lambda, I dare you"](https://www.youtube.com/watch?v=us_vS2EVDOA&t=46s)

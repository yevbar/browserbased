# Browserless

Like [serverless](https://www.serverless.com/) but for headless browsers

At the time of writing this, here's an example output from the [wikipedia](https://github.com/yevbar/browserless/blob/master/cobol/examples/wikipedia.cobol) example [https://project1725750744.vercel.app/api](https://project1725750744.vercel.app/api)

Think this is cool? Check out what we're building at [lsd.so](https://lsd.so)

## Configuring

Prior to running, you'll need to have a `VERCEL_TOKEN` environment variable set up with an access token which you can obtain [here](https://vercel.com/account/settings/tokens)

```bash
$ export VERCEL_TOKEN='abc123'
```

## Overview

* If you're interested in spinning up serverless headless browsers, check out the [browserless module](https://github.com/yevbar/browserless/blob/master/browserless/README.md)
* If you're interested in how it works, check out the [browser module](https://github.com/yevbar/browserless/blob/master/browsers/README.md)
* If you're interested in an instruction language for browsers, check out [COBOL](https://github.com/yevbar/browserless/blob/master/cobol/README.md)
* If you're interested in programatically deploying to Vercel using Golang, check out the [vercel module](https://github.com/yevbar/browserless/blob/master/vercel/README.md)

## Limitations

This does not offer any stealth or anti-anti-scraping capabilities and is as good as you can make your COBOL/Puppeteer scripts

## Inspiration

["Try running one on a lambda, I dare you"](https://www.youtube.com/watch?v=us_vS2EVDOA&t=46s)

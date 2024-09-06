# Browserless

Like [serverless](https://www.serverless.com/) but for headless browsers

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

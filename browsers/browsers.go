package browsers

import (
	"os"

	"github.com/yevbar/browserbased/vercel"
)

// Credit to this guy for figuring out how to get puppeteer working on Vercel in 2024
// https://gist.github.com/kettanaito/56861aff96e6debc575d522dd03e5725?permalink_comment_id=5010934#gistcomment-5010934
func CreateFilesystemFromScript(puppeteerScript string) map[string]string {
	fs := map[string]string{
		"src/app/api/route.ts": puppeteerScript,
		"next.config.mjs": NEXT_CONFIG_MJS,
		"package.json": PACKAGE_JSON,
		"package-lock.json": PACKAGE_LOCK_JSON,
		// "index.html": INDEX_HTML,
	}

	return fs
}

func SpinUpPuppeteerEndpoint(puppeteerScript string) string {
	vercelToken := os.Getenv("VERCEL_TOKEN")
	if len(vercelToken) == 0 {
		panic("The [VERCEL_TOKEN] environment variable must be specified")
	}

	deployment, err := vercel.CreateAndDeploy(vercelToken, CreateFilesystemFromScript(puppeteerScript))
	if err != nil {
		panic(err)
	}

	return deployment.URL
}

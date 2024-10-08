package browserbased

// A barebones Next.js function that navigates to example.com and takes a screenshot
const PUPPETEER_SCRIPT = `
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
  const page = await browser.newPage();
  await page.goto("https://example.com");
  const pdf = await page.pdf();
  await browser.close();
  return new NextResponse(pdf, {
    headers: {
      "Content-Type": "application/pdf",
    },
  });
}
`

// Returns a barebones Next.js function that navigates to example.com and takes a screenshot
func ExamplePuppeteerScript() string {
	return PUPPETEER_SCRIPT
}

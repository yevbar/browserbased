# COmmon Browser Oriented Language

* [What is this?](#what-is-this)
* [How does it work?](#how-does-it-work)
* [Known working examples (free tier)](#known-working-examples-vercel-free-tier)
* [Syntax](#syntax)
  * [Failure tolerance](#failure-tolerance)
  * [Keywords](#keywords)
	* [NAVIGATE](#navigate)
	* [CLICK](#click)
	* [ENTER](#enter)
	* [BACK](#back)
	* [NOTHING](#nothing)
  * [Functions](#functions)
  * [Comments](#comments)

## What is this?

If you've ever thought writing programs involving headless browsers were tedious or verbose and wished there were a batteries-included framework that'd strip away some of the manual work, you've come to the right place.

If you're interested in using COBOL rather than getting the conceptual digest, check out the [browserbased module](https://github.com/yevbar/browserbased/blob/master/browserbased/README.md#control-browsers-using-cobol)

## How does it work?

[Click here if you'd rather look at code examples](#known-working-examples-free-tier)

To understand the failure tolerance of COBOL, it may be helpful to look at the following topics from this lens:

* [Garbage collection](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) is an abstraction over memory management
* [Haxl](https://www.youtube.com/watch?v=sT6VJkkhy0o) is an abstraction over concurrency
* **NOTHING** is an abstraction over stupidity

COBOL can be understood as an instruction language like something you'd see in a computer architecture class but for controlling a web browser. Here's an example script for going to Google, entering a query into the search bar, then clicking on the button to invoke a search request:

```
NAVIGATE TO https://google.com
ENTER INTO input#search-box "your query"
CLICK ON button.cta
```

## Known working examples (Vercel free tier)

These are some of the ones I was able to get working (if it doesn't work on the first request, try invoking the `/api` endpoint once to warm up the function then requesting a 2nd time)

- [NIST abstracts](https://github.com/yevbar/browserbased/blob/master/cobol/examples/nist.cobol)
- [arxiv](https://github.com/yevbar/browserbased/blob/master/cobol/examples/arxiv.cobol)
- [Wikipedia](https://github.com/yevbar/browserbased/blob/master/cobol/examples/wikipedia.cobol)
- [Hacker News](https://github.com/yevbar/browserbased/blob/master/cobol/examples/hackernews.cobol)

You can see the <a href="https://github.com/yevbar/browserbased/tree/master/cobol/examples">examples folder</a> of scripts I was working on and some of them may actually work on a paid Vercel plan. Leaving for others to toy with

## Syntax

COBOL is newline-sensitive and doesn't care how you indent commands so long as each line accomplishes a single instruction. The goal of COBOL is to succinctly convey browser actions not code golf browsing.

### Failure tolerance

Keyword commands must have all words be correct

```
NAVIGATE TO https://example.com -- Actually works
NAVIGATE TOWARD https://example.com -- Nah

ENTER IN input#search-box "your query" -- 2nd word after ENTER must be INTO so this line does nothing
ENTER INTO input#search-box "your query" -- Actually works as intended
```

Malformed COBOL lines are treated similar to [NOTHING](#nothing) and are simply ignored

### Keywords

At the moment you can specify a browser to do stuff like the following

```
NAVIGATE TO <url>
CLICK ON <selector>
ENTER INTO <selector> "<text>"
```

Here are available keywords in COBOL

#### NAVIGATE

The `NAVIGATE` instruction tells a browser to navigate someplace, the syntax for this command is as follows

```
NAVIGATE TO <url>
```

The `<url>` gets provided to Puppeteer's [goto](https://pptr.dev/api/puppeteer.page.goto) method

#### CLICK

The `CLICK` instruction tells a browser to click on some element, the syntax is as follows

```
CLICK ON <selector>
```

The `<selector>` gets provided to Puppeteer's [click](https://pptr.dev/api/puppeteer.page.click) method

#### ENTER

The `ENTER` instruction tells a browser to enter some text into some element, the syntax is as follows

```
ENTER INTO <selector> "<text>"
```

The `<selector>` and `<text>` get provided to either Puppeteer's [type](https://pptr.dev/api/puppeteer.page.type) or [sendCharacter](https://pptr.dev/api/puppeteer.keyboard.sendcharacter) depending on whether the target element to type into is an input or textarea

#### BACK

The `BACK` instruction tells a browser to go back a page, the syntax is as follows

```
GO BACK
```

If just the expression `BACK` is provided or some other prior word than `GO`, it'll ignore the statement because of COBOL's [failure tolerance](#failure-tolerance). Under the hood it's Puppeteer's [goBack](https://pptr.dev/api/puppeteer.page.goback) method

#### NOTHING

As a mid-sentence exit valve, the language also features a `NOTHING` keyword in case you were to generate a line that doesn't make sense

```
NAVIGATE TO https://google.com
CLICK ON NOTHING -- Was written too soon
ENTER INTO input#search-box "your query"
CLICK ON button.cta
```

When translating to Puppeteer, lines with `NOTHING` as the "target" are ignored like comments but with the intent of allowing possibly incorrect code to be provided

### Functions

Functions are identified with colons and the function body is comprised of the lines that follow it. Function invocations are handled with `GOTO` statements with the `main` function being the entry point of the program

```
do_stuff:
	NAVIGATE TO https://news.ycombinator.com
	CLICK ON span.pagetop:nth-child(1) a:nth-child(2) -- Clicks on "New"
	GO BACK -- Navigates back

main:
	GOTO do_stuff
```

Like Ruby/Python, you should assume a bleedover of context where scopes are effectively "flattened" when jumping from one function block to another

### Comments

Comments are done via double dash following a whitespace

```
NAVIGATE TO https://google.com--not a comment
NAVIGATE TO https://google.com -- is a comment
NAVIGATE TO https://google.com --also a comment
```

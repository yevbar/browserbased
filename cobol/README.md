# COmmon Browser Oriented Language

## What is this?

If you've ever thought writing programs involving headless browsers were tedious or verbose and wished there were a batteries-included framework that'd strip away some of the manual work, you've come to the right place.

If you're interested in using COBOL rather than getting the conceptual digest, check out the [browserless module](https://github.com/yevbar/browserless/blob/master/browserless/README.md#control-browsers-using-cobol)

## How does it work?

[Click here if you'd rather look at code examples](#known-working-examples-free-tier)

To understand the failure tolerance of COBOL, it may be helpful to look at the following topics from this lens:

* [Garbage collection](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)) is an abstraction over memory management
* [Haxl](https://www.youtube.com/watch?v=sT6VJkkhy0o) is an abstraction over concurrency
* **COBOL** is an abstraction over stupidity

COBOL can be understood as an instruction language like something you'd see in a computer architecture class but for controlling a web browser. Here's an example script for going to Google, entering a query into the search bar, then clicking on the button to invoke a search request:

```
NAVIGATE TO https://google.com
ENTER INTO input#search-box "your query"
CLICK ON button.cta
```

As a mid-sentence exit valve, the language also features a `NOTHING` keyword in case you were to generate a line that doesn't make sense

```
NAVIGATE TO https://google.com
CLICK ON NOTHING -- Was written too soon
ENTER INTO input#search-box "your query"
CLICK ON button.cta
```

In cases where a line does not match the desired grammar (as in the first two words in the first line below don't correctly match the correct grammar as shown in the one following it)

```
ENTER IN input#search-box "your query" -- 2nd word after ENTER must be INTO so this line does nothing
ENTER INTO input#search-box "your query" -- Actually works as intended
```

As you may have figured out by now, comments are done via double dash following a whitespace

```
NAVIGATE TO https://google.com--not a comment
NAVIGATE TO https://google.com -- is a comment
NAVIGATE TO https://google.com --also a comment
```

## Known working examples (free tier)

These are some of the ones I was able to get working (if it doesn't work on the first request, try invoking the `/api` endpoint once to warm up the function then requesting a 2nd time)

- [NIST abstracts](https://github.com/yevbar/browserless/blob/master/cobol/examples/nist.cobol)
- [arxiv](https://github.com/yevbar/browserless/blob/master/cobol/examples/arxiv.cobol)
- [Wikipedia](https://github.com/yevbar/browserless/blob/master/cobol/examples/wikipedia.cobol)
- [Hacker News](https://github.com/yevbar/browserless/blob/master/cobol/examples/example.cobol)

You can see the <a href="https://github.com/yevbar/browserless/tree/master/cobol/examples">examples folder</a> of scripts I was working on and some of them may actually work on a paid Vercel plan. Leaving for others to toy with

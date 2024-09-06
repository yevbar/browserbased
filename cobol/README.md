# COmmon Browser Oriented Language

## What is this?

If you've ever thought writing programs involving headless browsers were tedious or verbose and wished there were a batteries-included framework that'd strip away some of the manual work, you've come to the right place.

## How does it work?

[Click here if you'd rather look at code examples](#examples)

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

## Examples

```

```

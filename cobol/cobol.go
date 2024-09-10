package cobol

import (
	"fmt"
	"strings"
)

func StringToCharByCharPuppeteer(s string) string {
	result := ""

	if s == "<RETURN>" {
		return "await page.keyboard.press('Enter');"
	}

	for i := 0; i < len(s); i++ {
		curChar := s[i:i+1]
		result += "await page.keyboard.sendCharacter('" + curChar + "');\n"
	}

	return result
}

func COBOLLineToPuppeteer(line string) string {
	trimmed := strings.TrimSpace(line)
	split := []string{}

	curToken := ""
	for i := 0; i < len(trimmed); i++ {
		curChar := trimmed[i:i+1]
		if len(strings.TrimSpace(curChar)) == 0 {
			if len(curToken) > 0 {
				split = append(split, curToken)
				curToken = ""
			}
			continue
		}
		curToken += curChar
	}
	if len(strings.TrimSpace(curToken)) > 0 {
		split = append(split, curToken)
	}

	if len(split) == 0 {
		return ""
	}

	withoutComment := []string{}
	for _, token := range split {
		if strings.HasPrefix(token, "--") {
			break
		}
		withoutComment = append(withoutComment, strings.ReplaceAll(token, "\"", "'"))
	}
	split = withoutComment

	if len(split) == 0 {
		return ""
	}

	switch command := split[0]; command {
	case "NAVIGATE":
		if len(split) == 1 {
			return ""
		}
		if split[1] != "TO" {
			return ""
		}
		if split[2] == "NOTHING" {
			return ""
		}
		return fmt.Sprintf("await page.goto(\"%s\");", split[2])
	case "CLICK":
		if len(split) == 1 {
			return ""
		}
		if split[1] != "ON" {
			return ""
		}
		if split[2] == "NOTHING" {
			return ""
		}
		return fmt.Sprintf("await page.click(\"%s\");", strings.Join(split[2:], " "))
	case "ENTER":
		if len(split) == 1 {
			return ""
		}
		if split[1] != "INTO" {
			return ""
		}
		remString := strings.Join(split[2:], " ")
		quoteIndex := strings.Index(remString, "'")
		if quoteIndex < 0 {
			return ""
		}
		selector := strings.TrimSpace(remString[:quoteIndex])
		text := remString[quoteIndex+1:len(remString)-1]
		if text == "<RETURN>" {
			return StringToCharByCharPuppeteer(text)
		}
		hashtagIndex := strings.Index(selector, "#")

		if strings.HasPrefix(selector, "textarea") && hashtagIndex > 0 {
			return fmt.Sprintf(
				strings.Join([]string{
					fmt.Sprintf("await page.click(\"%s\");", selector),
					StringToCharByCharPuppeteer(text),
				}, "\n"),
			)
		}
		return fmt.Sprintf("await page.type(\"%s\", \"%s\")", selector, text)
	case "DISABLE":
		if len(split) == 1 {
			return ""
		}
		if split[1] != "JAVASCRIPT" {
			return ""
		}
		return `await page.setRequestInterception(true);
page.on('request', request => (request.resourceType() === 'script') ? request.abort() : request.continue());`
	case "GO":
		if len(split) == 1 {
			return ""
		}
		if split[1] != "BACK" {
			return ""
		}
		return "await page.goBack();\n"
	case "GOTO":
		if len(split) == 1 {
			return ""
		}
		return fmt.Sprintf("await %s();", split[1])
	default:
		return ""
	}
}

func COBOLBlockToPuppeteer(lines []string) string {
	result := ""

	for _, line := range lines {
		result += fmt.Sprintf("%s\n", COBOLLineToPuppeteer(line))
	}

	trimmed := strings.TrimSpace(result)
	return trimmed
}

// Given a COBOL program, generates a block controlling a browser
func COBOLToPuppeteer(cobolScript string) string {
	result := ""

	blocks := COBOLToBlocks(cobolScript)
	for blockID, blockDef := range blocks {
		if blockID == "main" {
			continue
		}

		// Define function with block id as symbol and then COBOLToPuppeteer(blockDef) as body definition
		result += fmt.Sprintf("async function %s() {\n%s\n}\n", blockID, COBOLBlockToPuppeteer(blockDef))
	}

	if mainDef, ok := blocks["main"]; ok {
		result += COBOLBlockToPuppeteer(mainDef)
	}

	return fmt.Sprintf(PUPPETEER_TEMPLATE, result)
}

// Given a COBOL program, creates a map from block identifier to block definition
func COBOLToBlocks(cobolScript string) map[string][]string {
	result := map[string][]string{}
	currentBlock := "main"

	lines := strings.Split(cobolScript, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		} else if line[len(line) - 1:] == ":" {
			// Start new block
			currentBlock = line[:strings.Index(line, ":")]
		} else if len(trimmed) == 0 {
			// Just blank space
			continue
		} else {
			// Something that belongs in the current block definition
			if currentBlockDef, ok := result[currentBlock]; ok {
				result[currentBlock] = append(currentBlockDef, trimmed)
			} else {
				result[currentBlock] = []string{trimmed}
			}
		}
	}

	return result
}

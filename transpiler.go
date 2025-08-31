package main

import (
	"fmt"
	"regexp"
	"strings"
)

// func TranspileFind(src string) string {
// 	findRe := regexp.MustCompile(`(?m)([ \t]*)find\s*\(([^)]*)\)\s*{([\s\S]*?)(^[ \t]*})`)
// 	return src
// }

// TranspileSwitchToIfElse replaces switch statements with if-else chains
func TranspileSwitchToIfElse(src string) string {
	switchRe := regexp.MustCompile(`(?m)([ \t]*)switch\s*\(([^)]*)\)\s*{([\s\S]*?)(^[ \t]*})`)
	return switchRe.ReplaceAllStringFunc(src, func(switchBlock string) string {
		matches := switchRe.FindStringSubmatch(switchBlock)
		if len(matches) < 5 {
			return switchBlock
		}
		baseIndent := matches[1]
		variable := strings.TrimSpace(matches[2])
		body := matches[3]

		// Split body into lines and parse cases/default
		lines := strings.Split(body, "\n")
		type caseBlock struct {
			cond string // empty for default
			code []string
		}
		var blocks []caseBlock
		var current *caseBlock
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "case ") {
				cond := strings.TrimSpace(trimmed[5:])
				if idx := strings.Index(cond, ":"); idx != -1 {
					cond = cond[:idx]
				}
				current = &caseBlock{cond: cond}
				blocks = append(blocks, *current)
			} else if strings.HasPrefix(trimmed, "default:") {
				current = &caseBlock{cond: ""}
				blocks = append(blocks, *current)
			} else if current != nil {
				current.code = append(current.code, line)
				blocks[len(blocks)-1] = *current
			}
		}

		// Build if-else chain with preserved indentation
		var out strings.Builder
		for i, b := range blocks {
			// Find the minimum indent of the code block (ignore empty lines)
			minIndent := 1000
			for _, codeLine := range b.code {
				if trimmed := strings.TrimSpace(codeLine); trimmed != "" {
					leading := len(codeLine) - len(strings.TrimLeft(codeLine, " \t"))
					if leading < minIndent {
						minIndent = leading
					}
				}
			}
			if minIndent == 1000 {
				minIndent = 0
			}
			// Re-indent code block with exactly one tab for inner code
			for j, codeLine := range b.code {
				if strings.TrimSpace(codeLine) != "" {
					b.code[j] = "\t\t" + codeLine[minIndent:]
				} else {
					b.code[j] = ""
				}
			}
			code := strings.Join(b.code, "\n")
			code = strings.TrimRight(code, "\n")
			if b.cond != "" {
				if i == 0 {
					out.WriteString(fmt.Sprintf("%sif (%s == %s) {\n%s\n%s}", baseIndent, variable, b.cond, code, baseIndent+"    "))
				} else {
					out.WriteString(fmt.Sprintf(" else if (%s == %s) {\n%s\n%s}", variable, b.cond, code, baseIndent+"    "))
				}
			} else {
				if i > 0 {
					out.WriteString(fmt.Sprintf(" else {\n%s\n%s}", code, baseIndent+"    "))
				} else {
					out.WriteString(fmt.Sprintf("%s{\n%s\n%s}", baseIndent, code, baseIndent+"    "))
				}
			}
			// Remove extra indentation from closing brace
			outStr := out.String()
			if strings.HasSuffix(outStr, baseIndent+"    }") {
				out.Reset()
				out.WriteString(strings.TrimSuffix(outStr, baseIndent+"    }") + baseIndent + "}")
			}
		}
		return out.String()
	})
}

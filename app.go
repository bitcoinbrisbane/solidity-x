package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    inputPath := "contracts/ERC20.xol"
    outputPath := "dist/ERC20.sol"

    // Read the .xol file
    file, err := os.Open(inputPath)
    if err != nil {
        fmt.Println("Error opening input file:", err)
        return
    }
    defer file.Close()

    var src strings.Builder
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        src.WriteString(scanner.Text() + "\n")
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input file:", err)
        return
    }

    // Transpile switch to if-else
    out := TranspileSwitchToIfElse(src.String())

    // Write to output .sol file
    err = os.WriteFile(outputPath, []byte(out), 0644)
    if err != nil {
        fmt.Println("Error writing output file:", err)
        return
    }
    fmt.Println("Transpilation complete. Output written to", outputPath)
}

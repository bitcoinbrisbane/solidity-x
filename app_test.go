package main

import (
	"os"
	"testing"
)

func TestTranspileSwitchToIfElse(t *testing.T) {
    input := `function foo(uint x) public {
	switch (x) {
		case 1:
			bar();
		case 2:
			baz();
		default:
			qux();
	}
}`

	expected := "function foo(uint x) public {\n" +
		"\tif (x == 1) {\n" +
		"\t\tbar();\n" +
		"\t} else if (x == 2) {\n" +
		"\t\tbaz();\n" +
		"\t} else {\n" +
		"\t\tqux();\n" +
		"\t}\n" +
		"}"

    output := TranspileSwitchToIfElse(input)
	if output != expected {
		t.Errorf("Transpilation failed.\nExpected:\n%q\nGot:\n%q\nExpected bytes: %v\nGot bytes: %v", expected, output, []byte(expected), []byte(output))
	}
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

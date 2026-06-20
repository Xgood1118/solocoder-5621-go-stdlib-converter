package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunREPL(prec int) {
	fmt.Println("Unit Converter REPL - type 'q' or 'quit' to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		switch strings.ToLower(line) {
		case "q", "quit", "exit":
			return
		case "help":
			printHelp()
			continue
		case "list":
			printList()
			continue
		}
		processLine(line, prec)
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  <value> <from_unit> <to_unit>  - Convert a value")
	fmt.Println("  help                           - Show this help")
	fmt.Println("  list                           - List all categories and units")
	fmt.Println("  q / quit / exit                - Exit REPL")
}

func printList() {
	reg := DefaultRegistry()
	cats := reg.AllCategories()
	for _, cat := range cats {
		fmt.Printf("[%s]\n", cat)
		units := reg.UnitsByCategory(cat)
		for _, u := range units {
			fmt.Printf("  %s (%v)\n", u.Key, u.Names)
		}
	}
}

func processLine(line string, prec int) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		fmt.Println("error: expected format: <value> <from_unit> <to_unit>")
		return
	}
	input := parts[0]
	fromStr := parts[1]
	toStr := parts[2]

	if IsBaseCategory(fromStr) || IsBaseCategory(toStr) {
		fromBase, err := LookupBase(fromStr)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		toBase, err := LookupBase(toStr)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		result, err := ConvertBase(input, fromBase, toBase)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Printf("%s %s = %s %s\n", input, fromStr, result, toStr)
		return
	}

	value, err := ParseBigFloat(input)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fromUnit, suggestions, ok := FuzzyLookupUnit(fromStr)
	if !ok {
		if len(suggestions) > 0 {
			fmt.Printf("error: ambiguous unit %q, did you mean: %s?\n", fromStr, strings.Join(suggestions, ", "))
		} else {
			fmt.Printf("error: unknown unit %q\n", fromStr)
		}
		return
	}

	toUnit, suggestions, ok := FuzzyLookupUnit(toStr)
	if !ok {
		if len(suggestions) > 0 {
			fmt.Printf("error: ambiguous unit %q, did you mean: %s?\n", toStr, strings.Join(suggestions, ", "))
		} else {
			fmt.Printf("error: unknown unit %q\n", toStr)
		}
		return
	}

	fmt.Println(FormatResult(value, fromUnit, toUnit, prec))
}

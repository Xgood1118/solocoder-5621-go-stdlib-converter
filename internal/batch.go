package internal

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type BatchItem struct {
	Input    string
	FromUnit string
	ToUnit   string
}

func ParseBatchLines(lines []string) ([]BatchItem, error) {
	var items []BatchItem
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid batch line: %q", line)
		}
		items = append(items, BatchItem{
			Input:    parts[0],
			FromUnit: parts[1],
			ToUnit:   parts[2],
		})
	}
	return items, nil
}

func ParseCSVFile(path string) ([]BatchItem, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var items []BatchItem
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 3 {
			continue
		}
		items = append(items, BatchItem{
			Input:    strings.TrimSpace(record[0]),
			FromUnit: strings.TrimSpace(record[1]),
			ToUnit:   strings.TrimSpace(record[2]),
		})
	}
	return items, nil
}

func ParseJSONFile(path string) ([]BatchItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var raw []struct {
		Value string `json:"value"`
		From  string `json:"from"`
		To    string `json:"to"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var items []BatchItem
	for _, r := range raw {
		items = append(items, BatchItem{
			Input:    strings.TrimSpace(r.Value),
			FromUnit: strings.TrimSpace(r.From),
			ToUnit:   strings.TrimSpace(r.To),
		})
	}
	return items, nil
}

func ParseTextFile(path string) ([]BatchItem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return ParseBatchLines(lines)
}

func ExecuteBatch(items []BatchItem, prec int) []string {
	results := make([]string, len(items))
	for i, item := range items {
		results[i] = executeItem(item, prec)
	}
	return results
}

func executeItem(item BatchItem, prec int) string {
	if IsBaseCategory(item.FromUnit) || IsBaseCategory(item.ToUnit) {
		fromBase, err := LookupBase(item.FromUnit)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		toBase, err := LookupBase(item.ToUnit)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		result, err := ConvertBase(item.Input, fromBase, toBase)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return fmt.Sprintf("%s %s = %s %s", item.Input, item.FromUnit, result, item.ToUnit)
	}

	value, err := ParseBigFloat(item.Input)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}

	fromUnit, ok := LookupUnit(item.FromUnit)
	if !ok {
		return fmt.Sprintf("error: unknown unit %q", item.FromUnit)
	}

	toUnit, ok := LookupUnit(item.ToUnit)
	if !ok {
		return fmt.Sprintf("error: unknown unit %q", item.ToUnit)
	}

	return FormatResult(value, fromUnit, toUnit, prec)
}

func WriteResults(results []string, path string) error {
	var sb strings.Builder
	for _, r := range results {
		sb.WriteString(r)
		sb.WriteByte('\n')
	}
	return os.WriteFile(path, []byte(sb.String()), 0644)
}

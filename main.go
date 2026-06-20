package main

import (
	"fmt"
	"os"

	"github.com/solocoder/unitconv/internal"
	"github.com/spf13/cobra"
)

var (
	precision   int
	configPath  string
	outputPath  string
	interactive bool
	batchFile   string
)

var rootCmd = &cobra.Command{
	Use:   "convert [value] [from_unit] [to_unit]",
	Short: "多功能单位换算命令行工具",
	Long:  "支持长度、重量、温度、面积、体积、时间、速度、压力、能量、功率、数据存储、进制、货币换算",
	Example: `  convert 100 km mile
  convert 32 F C
  convert FF hex dec
  convert 100 CNY USD
  convert --interactive
  convert --batch input.csv -o output.csv`,
	Args: cobra.MinimumNArgs(0),
	RunE: runConvert,
}

func init() {
	rootCmd.Flags().IntVarP(&precision, "precision", "p", -1, "输出精度（小数位数）")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "配置文件路径")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "输出文件路径")
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "交互式REPL模式")
	rootCmd.Flags().StringVarP(&batchFile, "batch", "b", "", "批量换算文件（CSV/JSON/文本）")
}

func runConvert(cmd *cobra.Command, args []string) error {
	if configPath != "" {
		if err := internal.LoadConfig(configPath); err != nil {
			return fmt.Errorf("加载配置失败: %w", err)
		}
	}

	cfg := internal.GetConfig()
	prec := precision
	if prec < 0 {
		prec = cfg.Precision
	}

	if interactive {
		internal.RunREPL(prec)
		return nil
	}

	if batchFile != "" {
		return runBatch(prec)
	}

	if len(args) < 3 {
		return fmt.Errorf("用法: convert <值> <源单位> <目标单位>\n示例: convert 100 km mile")
	}

	return runSingle(args, prec)
}

func runSingle(args []string, prec int) error {
	valueStr := args[0]
	fromStr := args[1]
	toStr := args[2]

	if internal.IsBaseCategory(fromStr) || internal.IsBaseCategory(toStr) {
		fromBase, err1 := internal.LookupBase(fromStr)
		toBase, err2 := internal.LookupBase(toStr)
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
		result, err := internal.ConvertBase(valueStr, fromBase, toBase)
		if err != nil {
			return err
		}
		output := fmt.Sprintf("%s (base %d) = %s (base %d)", valueStr, fromBase, result, toBase)
		return writeOutput(output)
	}

	fromUnit, fromCandidates, fromOk := internal.FuzzyLookupUnit(fromStr)
	if !fromOk {
		if len(fromCandidates) > 0 {
			return fmt.Errorf("单位 %q 有多个匹配: %v，请更具体", fromStr, fromCandidates)
		}
		return fmt.Errorf("未找到单位: %q", fromStr)
	}

	toUnit, toCandidates, toOk := internal.FuzzyLookupUnit(toStr)
	if !toOk {
		if len(toCandidates) > 0 {
			return fmt.Errorf("单位 %q 有多个匹配: %v，请更具体", toStr, toCandidates)
		}
		return fmt.Errorf("未找到单位: %q", toStr)
	}

	if fromUnit.Category != toUnit.Category {
		return fmt.Errorf("不能在不同类别间换算: %s(%s) -> %s(%s)",
			fromStr, fromUnit.Category, toStr, toUnit.Category)
	}

	value, err := internal.ParseBigFloat(valueStr)
	if err != nil {
		return err
	}

	output := internal.FormatResult(value, fromUnit, toUnit, prec)
	return writeOutput(output)
}

func runBatch(prec int) error {
	var items []internal.BatchItem
	var err error

	switch {
	case len(batchFile) > 4 && batchFile[len(batchFile)-4:] == ".csv":
		items, err = internal.ParseCSVFile(batchFile)
	case len(batchFile) > 5 && batchFile[len(batchFile)-5:] == ".json":
		items, err = internal.ParseJSONFile(batchFile)
	default:
		items, err = internal.ParseTextFile(batchFile)
	}
	if err != nil {
		return fmt.Errorf("读取批量文件失败: %w", err)
	}

	results := internal.ExecuteBatch(items, prec)
	for _, r := range results {
		fmt.Println(r)
	}

	if outputPath != "" {
		return internal.WriteResults(results, outputPath)
	}
	return nil
}

func writeOutput(s string) error {
	fmt.Println(s)
	if outputPath != "" {
		return internal.WriteResults([]string{s}, outputPath)
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list [category]",
	Short: "列出支持的单位类别或指定类别的单位",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cats := internal.DefaultRegistry().AllCategories()
			fmt.Println("支持的换算类别:")
			for _, cat := range cats {
				fmt.Printf("  - %s\n", cat)
			}
			return nil
		}
		cat := internal.Category(args[0])
		units := internal.DefaultRegistry().UnitsByCategory(cat)
		if len(units) == 0 {
			return fmt.Errorf("未知类别: %s", args[0])
		}
		fmt.Printf("类别 %s 的单位:\n", cat)
		for _, u := range units {
			fmt.Printf("  %s: %v\n", u.Key, u.Names)
		}
		return nil
	},
}

func main() {
	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

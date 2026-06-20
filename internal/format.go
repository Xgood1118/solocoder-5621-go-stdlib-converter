package internal

import (
	"fmt"
	"math/big"
	"strings"
)

func FormatResult(value *big.Float, fromUnit, toUnit *Unit, prec int) string {
	if prec < 0 {
		prec = GetConfig().Precision
	}
	if fromUnit.Category == CatCurrency {
		prec = 4
	}
	fromStr := FormatValue(value, prec)
	result := Convert(value, fromUnit, toUnit)
	toStr := FormatValue(result, prec)
	return fmt.Sprintf("%s %s = %s %s", fromStr, fromUnit.Names[0], toStr, toUnit.Names[0])
}

func FormatValue(value *big.Float, prec int) string {
	if prec < 0 {
		prec = GetConfig().Precision
	}
	s := FormatBigFloat(new(big.Float).Set(value), prec)
	return ApplyLocale(s)
}

func ApplyLocale(formatted string) string {
	c := GetConfig()
	parts := strings.Split(formatted, ".")
	integerPart := parts[0]

	if c.ThousandSep != "" {
		neg := false
		digits := integerPart
		if len(digits) > 0 && digits[0] == '-' {
			neg = true
			digits = digits[1:]
		}
		var buf strings.Builder
		n := len(digits)
		for i := 0; i < n; i++ {
			if i > 0 && (n-i)%3 == 0 {
				buf.WriteString(c.ThousandSep)
			}
			buf.WriteByte(digits[i])
		}
		if neg {
			integerPart = "-" + buf.String()
		} else {
			integerPart = buf.String()
		}
	}

	if len(parts) > 1 {
		return integerPart + c.DecimalSep + parts[1]
	}
	return integerPart
}

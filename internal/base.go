package internal

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var baseNames = map[string]int{
	"bin": 2, "binary": 2, "二进制": 2, "2": 2,
	"oct": 8, "octal": 8, "八进制": 8, "8": 8,
	"dec": 10, "decimal": 10, "十进制": 10, "10": 10,
	"hex": 16, "hexadecimal": 16, "十六进制": 16, "16": 16,
	"base32": 32, "三十二进制": 32, "32": 32,
	"base64": 64, "六十四进制": 64, "64": 64,
}

const base64Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var base64DecodeMap map[byte]int

func init() {
	base64DecodeMap = make(map[byte]int)
	for i := 0; i < len(base64Alphabet); i++ {
		base64DecodeMap[base64Alphabet[i]] = i
	}
}

func LookupBase(name string) (int, error) {
	key := strings.ToLower(strings.TrimSpace(name))
	if b, ok := baseNames[key]; ok {
		return b, nil
	}
	b, err := strconv.Atoi(key)
	if err != nil {
		return 0, fmt.Errorf("unknown base: %q", name)
	}
	if (b >= 2 && b <= 36) || b == 64 {
		return b, nil
	}
	return 0, fmt.Errorf("unsupported base: %d (valid: 2-36, 64)", b)
}

func IsBaseCategory(name string) bool {
	key := strings.ToLower(strings.TrimSpace(name))
	if _, ok := baseNames[key]; ok {
		return true
	}
	b, err := strconv.Atoi(key)
	if err != nil {
		return false
	}
	return (b >= 2 && b <= 36) || b == 64
}

func ConvertBase(valueStr string, fromBase, toBase int) (string, error) {
	if fromBase == toBase {
		return valueStr, nil
	}

	negative := strings.HasPrefix(valueStr, "-")
	clean := valueStr
	if negative {
		clean = clean[1:]
	}
	if clean == "" {
		return "", errors.New("empty value")
	}

	var n big.Int

	if fromBase == 64 {
		err := setFromBase64(&n, clean)
		if err != nil {
			return "", err
		}
	} else if fromBase >= 2 && fromBase <= 36 {
		_, ok := n.SetString(clean, fromBase)
		if !ok {
			return "", fmt.Errorf("invalid value %q for base %d", clean, fromBase)
		}
	} else {
		return "", fmt.Errorf("unsupported source base: %d", fromBase)
	}

	if negative {
		n.Neg(&n)
	}

	isNeg := n.Sign() < 0
	var abs big.Int
	abs.Abs(&n)

	var result string
	if toBase == 64 {
		result = encodeBase64(&abs)
	} else if toBase >= 2 && toBase <= 36 {
		result = abs.Text(toBase)
	} else {
		return "", fmt.Errorf("unsupported target base: %d", toBase)
	}

	if isNeg {
		result = "-" + result
	}
	return result, nil
}

func setFromBase64(n *big.Int, s string) error {
	if len(s) == 0 {
		n.SetInt64(0)
		return nil
	}

	buf := make([]byte, 0, len(s)*6/8+1)
	bits := 0
	acc := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]
		val, ok := base64DecodeMap[ch]
		if !ok {
			return fmt.Errorf("invalid base64 character: %c", ch)
		}
		acc = (acc << 6) | val
		bits += 6
		for bits >= 8 {
			bits -= 8
			buf = append(buf, byte(acc>>bits))
			acc &= (1 << bits) - 1
		}
	}

	if bits > 0 && acc != 0 {
		return fmt.Errorf("non-zero trailing bits in base64 input")
	}

	n.SetBytes(buf)
	return nil
}

func encodeBase64(n *big.Int) string {
	if n.Sign() == 0 {
		return "A"
	}

	b := n.Bytes()
	if len(b) == 0 {
		return "A"
	}

	var sb strings.Builder
	bits := len(b) * 8
	totalBits := 0
	for totalBits < bits {
		totalBits += 6
	}

	var acc int
	accBits := 0
	bitIdx := totalBits - 1

	for bitIdx >= 0 {
		byteIdx := bitIdx / 8
		bitPos := 7 - (bitIdx % 8)
		bit := int((b[byteIdx] >> bitPos) & 1)
		acc = (acc << 1) | bit
		accBits++
		bitIdx--

		if accBits == 6 {
			sb.WriteByte(base64Alphabet[acc])
			acc = 0
			accBits = 0
		}
	}

	if accBits > 0 {
		acc <<= (6 - accBits)
		sb.WriteByte(base64Alphabet[acc])
	}

	result := sb.String()
	result = strings.TrimLeft(result, "A")
	if result == "" {
		return "A"
	}
	return result
}

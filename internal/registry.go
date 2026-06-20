package internal

import (
	"fmt"
	"math/big"
	"strings"
	"unicode"
)

type Category string

const (
	CatLength      Category = "length"
	CatWeight      Category = "weight"
	CatTemperature Category = "temperature"
	CatArea        Category = "area"
	CatVolume      Category = "volume"
	CatTime        Category = "time"
	CatSpeed       Category = "speed"
	CatPressure    Category = "pressure"
	CatEnergy      Category = "energy"
	CatPower       Category = "power"
	CatData        Category = "data"
	CatBase        Category = "base"
	CatCurrency    Category = "currency"
)

type Unit struct {
	Key      string
	Category Category
	Names    []string
	Factor   *big.Float
	Offset   *big.Float
}

type Registry struct {
	units    map[string]*Unit
	aliasMap map[string]string
}

var defaultRegistry = NewRegistry()

func NewRegistry() *Registry {
	return &Registry{
		units:    make(map[string]*Unit),
		aliasMap: make(map[string]string),
	}
}

func DefaultRegistry() *Registry {
	return defaultRegistry
}

func (r *Registry) Register(u *Unit) {
	r.units[u.Key] = u
	for _, name := range u.Names {
		key := strings.ToLower(strings.TrimSpace(name))
		r.aliasMap[key] = u.Key
	}
}

func (r *Registry) Lookup(name string) (*Unit, bool) {
	key := strings.ToLower(strings.TrimSpace(name))
	if unitKey, ok := r.aliasMap[key]; ok {
		u, ok2 := r.units[unitKey]
		return u, ok2
	}
	return nil, false
}

func (r *Registry) FuzzyLookup(name string) (*Unit, []string, bool) {
	u, ok := r.Lookup(name)
	if ok {
		return u, nil, true
	}

	input := strings.ToLower(strings.TrimSpace(name))
	pinyinInput := toPinyinInitials(input)

	var candidates []string
	seen := make(map[string]bool)

	for alias, unitKey := range r.aliasMap {
		aliasLower := strings.ToLower(alias)
		if strings.Contains(aliasLower, input) {
			if !seen[unitKey] {
				seen[unitKey] = true
				candidates = append(candidates, unitKey)
			}
			continue
		}
		if len(aliasLower) >= 2 && len(input) >= 2 && strings.Contains(input, aliasLower) {
			if !seen[unitKey] {
				seen[unitKey] = true
				candidates = append(candidates, unitKey)
			}
			continue
		}
		aliasPinyin := toPinyinInitials(aliasLower)
		if pinyinInput != "" && aliasPinyin != "" && len(pinyinInput) >= 2 && strings.Contains(aliasPinyin, pinyinInput) {
			if !seen[unitKey] {
				seen[unitKey] = true
				candidates = append(candidates, unitKey)
			}
		}
	}

	if len(candidates) == 1 {
		u, ok := r.units[candidates[0]]
		if ok {
			return u, nil, true
		}
	}

	if len(candidates) > 1 {
		var names []string
		for _, ck := range candidates {
			if u, ok := r.units[ck]; ok {
				names = append(names, u.Names[0])
			}
		}
		return nil, names, false
	}

	minDist := 999
	var bestKey string
	for alias, unitKey := range r.aliasMap {
		d := levenshtein(input, strings.ToLower(alias))
		if d < minDist {
			minDist = d
			bestKey = unitKey
		}
	}
	if minDist <= 2 && bestKey != "" {
		u, ok := r.units[bestKey]
		if ok {
			return u, nil, true
		}
	}

	return nil, nil, false
}

func (r *Registry) UnitsByCategory(cat Category) []*Unit {
	var result []*Unit
	for _, u := range r.units {
		if u.Category == cat {
			result = append(result, u)
		}
	}
	return result
}

func (r *Registry) AllCategories() []Category {
	seen := make(map[Category]bool)
	var result []Category
	for _, u := range r.units {
		if !seen[u.Category] {
			seen[u.Category] = true
			result = append(result, u.Category)
		}
	}
	return result
}

func Convert(value *big.Float, from, to *Unit) *big.Float {
	if from.Category == CatCurrency {
		base := new(big.Float).Quo(value, from.Factor)
		result := new(big.Float).Mul(base, to.Factor)
		return result
	}

	if from.Category == CatBase {
		result := new(big.Float).Mul(value, from.Factor)
		result.Quo(result, to.Factor)
		return result
	}

	base := new(big.Float).Mul(value, from.Factor)
	if from.Offset != nil {
		base.Add(base, from.Offset)
	}

	if to.Offset != nil {
		result := new(big.Float).Sub(base, to.Offset)
		result.Quo(result, to.Factor)
		return result
	}

	result := new(big.Float).Quo(base, to.Factor)
	return result
}

func NewFactor(f float64) *big.Float {
	return big.NewFloat(f)
}

func NewOffset(f float64) *big.Float {
	return big.NewFloat(f)
}

func RegisterUnit(u *Unit) {
	defaultRegistry.Register(u)
}

func LookupUnit(name string) (*Unit, bool) {
	return defaultRegistry.Lookup(name)
}

func FuzzyLookupUnit(name string) (*Unit, []string, bool) {
	return defaultRegistry.FuzzyLookup(name)
}

func levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = j
	}
	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			curr[j] = min3(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return prev[lb]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

var pinyinMap = map[rune]string{
	'米': "m", '厘': "l", '毫': "h", '千': "q", '英': "y",
	'码': "m", '尺': "c", '寸': "c", '海': "h", '光': "g",
	'年': "n", '克': "k", '吨': "d", '磅': "b", '盎': "a",
	'拉': "l", '金': "j", '衡': "h", '摄': "s", '华': "h",
	'开': "k", '兰': "l", '平': "p", '公': "g", '顷': "q",
	'亩': "m", '加': "j", '仑': "l", '品': "p", '脱': "t",
	'夸': "k", '方': "f", '秒': "m", '分': "f", '时': "s",
	'天': "t", '周': "z", '月': "y", '世': "s", '纪': "j",
	'微': "w", '节': "j", '马': "m", '赫': "h", '帕': "p",
	'巴': "b", '气': "q", '压': "y", '汞': "g", '柱': "z",
	'焦': "j", '卡': "k", '电': "d", '伏': "f", '瓦': "w",
	'力': "l", '比': "b", '特': "t", '字': "z",
	'人': "r", '民': "m", '币': "b", '美': "m", '元': "y",
	'欧': "o", '日': "r", '镑': "b", '进': "j", '制': "z",
	'八': "b", '十': "s", '六': "l", '三': "s", '二': "e",
	'五': "w", '四': "s", '七': "q", '九': "j", '零': "l",
	'度': "d",
}

func toPinyinInitials(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if r < 128 {
			if unicode.IsLetter(r) {
				sb.WriteRune(unicode.ToLower(r))
			}
			continue
		}
		if py, ok := pinyinMap[r]; ok {
			sb.WriteString(py)
		}
	}
	return sb.String()
}

func FormatBigFloat(f *big.Float, prec int) string {
	if prec < 0 {
		prec = 6
	}
	f.SetMode(big.ToNearestEven)
	s := f.Text('f', prec)
	if idx := strings.IndexByte(s, '.'); idx >= 0 {
		frac := strings.TrimRight(s[idx:], "0")
		if frac == "." {
			frac = ""
		}
		s = s[:idx] + frac
	}
	if s == "" || s == "-" {
		return "0"
	}
	return s
}

func ParseBigFloat(s string) (*big.Float, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", "")
	f, _, err := big.ParseFloat(s, 10, 0, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("invalid number: %q", s)
	}
	return f, nil
}

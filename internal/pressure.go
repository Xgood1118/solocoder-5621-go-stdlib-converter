package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "Pa",
		Category: CatPressure,
		Names:    []string{"Pa", "pascal", "帕"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "kPa",
		Category: CatPressure,
		Names:    []string{"千帕", "kPa", "kilopascal"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "bar",
		Category: CatPressure,
		Names:    []string{"bar", "巴"},
		Factor:   NewFactor(100000),
	})

	RegisterUnit(&Unit{
		Key:      "psi",
		Category: CatPressure,
		Names:    []string{"psi"},
		Factor:   NewFactor(6894.757293168),
	})

	RegisterUnit(&Unit{
		Key:      "atm",
		Category: CatPressure,
		Names:    []string{"大气压", "atm", "atmosphere"},
		Factor:   NewFactor(101325),
	})

	RegisterUnit(&Unit{
		Key:      "mmHg",
		Category: CatPressure,
		Names:    []string{"毫米汞柱", "mmHg"},
		Factor:   NewFactor(133.3223684),
	})
}

package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "sqm",
		Category: CatArea,
		Names:    []string{"平方米", "sqm", "squaremeter"},
		Factor:   NewFactor(1),
	})
	RegisterUnit(&Unit{
		Key:      "ha",
		Category: CatArea,
		Names:    []string{"公顷", "ha", "hectare"},
		Factor:   NewFactor(10000),
	})
	RegisterUnit(&Unit{
		Key:      "mu",
		Category: CatArea,
		Names:    []string{"亩", "mu"},
		Factor:   NewFactor(666.6666666667),
	})
	RegisterUnit(&Unit{
		Key:      "sqmi",
		Category: CatArea,
		Names:    []string{"平方英里", "sqmi", "squaremile"},
		Factor:   NewFactor(2589988.110336),
	})
	RegisterUnit(&Unit{
		Key:      "sqft",
		Category: CatArea,
		Names:    []string{"平方英尺", "sqft", "squarefoot", "squarefeet"},
		Factor:   NewFactor(0.09290304),
	})
	RegisterUnit(&Unit{
		Key:      "acre",
		Category: CatArea,
		Names:    []string{"英亩", "acre"},
		Factor:   NewFactor(4046.8564224),
	})
}

package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "W",
		Category: CatPower,
		Names:    []string{"瓦", "W", "watt"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "kW",
		Category: CatPower,
		Names:    []string{"千瓦", "kW", "kilowatt"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "hp",
		Category: CatPower,
		Names:    []string{"马力", "hp", "horsepower"},
		Factor:   NewFactor(745.7),
	})
}

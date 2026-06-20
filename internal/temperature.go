package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "celsius",
		Category: CatTemperature,
		Names:    []string{"摄氏度", "c", "C", "celsius"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "fahrenheit",
		Category: CatTemperature,
		Names:    []string{"华氏度", "f", "F", "fahrenheit"},
		Factor:   NewFactor(5.0 / 9.0),
		Offset:   NewOffset(-32.0 * 5.0 / 9.0),
	})

	RegisterUnit(&Unit{
		Key:      "kelvin",
		Category: CatTemperature,
		Names:    []string{"开尔文", "K", "kelvin"},
		Factor:   NewFactor(1),
		Offset:   NewOffset(-273.15),
	})

	RegisterUnit(&Unit{
		Key:      "rankine",
		Category: CatTemperature,
		Names:    []string{"兰金", "R", "rankine"},
		Factor:   NewFactor(5.0 / 9.0),
		Offset:   NewOffset(-273.15),
	})
}

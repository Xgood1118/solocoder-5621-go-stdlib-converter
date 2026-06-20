package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "J",
		Category: CatEnergy,
		Names:    []string{"焦", "J", "joule"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "kJ",
		Category: CatEnergy,
		Names:    []string{"千焦", "kJ", "kilojoule"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "cal",
		Category: CatEnergy,
		Names:    []string{"卡路里", "cal", "calorie"},
		Factor:   NewFactor(4.184),
	})

	RegisterUnit(&Unit{
		Key:      "kcal",
		Category: CatEnergy,
		Names:    []string{"千卡", "kcal", "kilocalorie"},
		Factor:   NewFactor(4184),
	})

	RegisterUnit(&Unit{
		Key:      "eV",
		Category: CatEnergy,
		Names:    []string{"电子伏特", "eV", "electronvolt"},
		Factor:   NewFactor(1.602176634e-19),
	})

	RegisterUnit(&Unit{
		Key:      "Wh",
		Category: CatEnergy,
		Names:    []string{"瓦时", "Wh", "watthour"},
		Factor:   NewFactor(3600),
	})

	RegisterUnit(&Unit{
		Key:      "kWh",
		Category: CatEnergy,
		Names:    []string{"千瓦时", "kWh", "kilowatthour"},
		Factor:   NewFactor(3600000),
	})
}

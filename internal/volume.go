package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "L",
		Category: CatVolume,
		Names:    []string{"升", "L", "liter", "litre"},
		Factor:   NewFactor(1),
	})
	RegisterUnit(&Unit{
		Key:      "mL",
		Category: CatVolume,
		Names:    []string{"毫升", "mL", "milliliter"},
		Factor:   NewFactor(0.001),
	})
	RegisterUnit(&Unit{
		Key:      "usgal",
		Category: CatVolume,
		Names:    []string{"加仑(美)", "gal", "gallon", "usgallon"},
		Factor:   NewFactor(3.785411784),
	})
	RegisterUnit(&Unit{
		Key:      "ukgal",
		Category: CatVolume,
		Names:    []string{"加仑(英)", "ukgallon", "imperialgallon"},
		Factor:   NewFactor(4.54609),
	})
	RegisterUnit(&Unit{
		Key:      "pt",
		Category: CatVolume,
		Names:    []string{"品脱", "pt", "pint"},
		Factor:   NewFactor(0.473176473),
	})
	RegisterUnit(&Unit{
		Key:      "qt",
		Category: CatVolume,
		Names:    []string{"夸脱", "qt", "quart"},
		Factor:   NewFactor(0.946352946),
	})
	RegisterUnit(&Unit{
		Key:      "cuft",
		Category: CatVolume,
		Names:    []string{"立方英尺", "cuft", "cubicfoot"},
		Factor:   NewFactor(28.316846592),
	})
}

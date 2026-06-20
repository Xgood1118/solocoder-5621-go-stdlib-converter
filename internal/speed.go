package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "m/s",
		Category: CatSpeed,
		Names:    []string{"m/s", "mps", "meters per second"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "km/h",
		Category: CatSpeed,
		Names:    []string{"km/h", "kph"},
		Factor:   NewFactor(0.27777777777778),
	})

	RegisterUnit(&Unit{
		Key:      "mph",
		Category: CatSpeed,
		Names:    []string{"mph", "mileperhour"},
		Factor:   NewFactor(0.44704),
	})

	RegisterUnit(&Unit{
		Key:      "knot",
		Category: CatSpeed,
		Names:    []string{"knot", "kn", "节"},
		Factor:   NewFactor(0.514444),
	})

	RegisterUnit(&Unit{
		Key:      "mach",
		Category: CatSpeed,
		Names:    []string{"mach", "马赫"},
		Factor:   NewFactor(343),
	})
}

package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "gram",
		Category: CatWeight,
		Names:    []string{"克", "g", "gram", "公克"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "kilogram",
		Category: CatWeight,
		Names:    []string{"千克", "kg", "kilogram", "公斤"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "ton",
		Category: CatWeight,
		Names:    []string{"吨", "t", "ton", "tonne", "公吨"},
		Factor:   NewFactor(1e6),
	})

	RegisterUnit(&Unit{
		Key:      "pound",
		Category: CatWeight,
		Names:    []string{"磅", "lb", "pound"},
		Factor:   NewFactor(453.59237),
	})

	RegisterUnit(&Unit{
		Key:      "ounce",
		Category: CatWeight,
		Names:    []string{"盎司", "oz", "ounce"},
		Factor:   NewFactor(28.349523125),
	})

	RegisterUnit(&Unit{
		Key:      "carat",
		Category: CatWeight,
		Names:    []string{"克拉", "ct", "carat"},
		Factor:   NewFactor(0.2),
	})

	RegisterUnit(&Unit{
		Key:      "troy_ounce",
		Category: CatWeight,
		Names:    []string{"金衡盎司", "ozt", "troyounce"},
		Factor:   NewFactor(31.1034768),
	})
}

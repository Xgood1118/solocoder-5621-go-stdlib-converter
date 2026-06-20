package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "meter",
		Category: CatLength,
		Names:    []string{"米", "m", "meter"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "centimeter",
		Category: CatLength,
		Names:    []string{"厘米", "cm", "centimeter"},
		Factor:   NewFactor(0.01),
	})

	RegisterUnit(&Unit{
		Key:      "millimeter",
		Category: CatLength,
		Names:    []string{"毫米", "mm", "millimeter"},
		Factor:   NewFactor(0.001),
	})

	RegisterUnit(&Unit{
		Key:      "kilometer",
		Category: CatLength,
		Names:    []string{"千米", "km", "kilometer"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "mile",
		Category: CatLength,
		Names:    []string{"英里", "mi", "mile"},
		Factor:   NewFactor(1609.344),
	})

	RegisterUnit(&Unit{
		Key:      "yard",
		Category: CatLength,
		Names:    []string{"码", "yd", "yard"},
		Factor:   NewFactor(0.9144),
	})

	RegisterUnit(&Unit{
		Key:      "foot",
		Category: CatLength,
		Names:    []string{"英尺", "ft", "foot", "feet"},
		Factor:   NewFactor(0.3048),
	})

	RegisterUnit(&Unit{
		Key:      "inch",
		Category: CatLength,
		Names:    []string{"英寸", "in", "inch"},
		Factor:   NewFactor(0.0254),
	})

	RegisterUnit(&Unit{
		Key:      "nautical_mile",
		Category: CatLength,
		Names:    []string{"海里", "nmi", "nauticalmile"},
		Factor:   NewFactor(1852),
	})

	RegisterUnit(&Unit{
		Key:      "light_year",
		Category: CatLength,
		Names:    []string{"光年", "ly", "lightyear"},
		Factor:   NewFactor(9.461e15),
	})
}

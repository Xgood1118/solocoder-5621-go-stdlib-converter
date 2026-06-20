package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "Byte",
		Category: CatData,
		Names:    []string{"Byte", "B"},
		Factor:   NewFactor(1),
	})

	RegisterUnit(&Unit{
		Key:      "bit",
		Category: CatData,
		Names:    []string{"bit"},
		Factor:   NewFactor(0.125),
	})

	RegisterUnit(&Unit{
		Key:      "KB",
		Category: CatData,
		Names:    []string{"KB", "kilobyte"},
		Factor:   NewFactor(1000),
	})

	RegisterUnit(&Unit{
		Key:      "MB",
		Category: CatData,
		Names:    []string{"MB", "megabyte"},
		Factor:   NewFactor(1e6),
	})

	RegisterUnit(&Unit{
		Key:      "GB",
		Category: CatData,
		Names:    []string{"GB", "gigabyte"},
		Factor:   NewFactor(1e9),
	})

	RegisterUnit(&Unit{
		Key:      "TB",
		Category: CatData,
		Names:    []string{"TB", "terabyte"},
		Factor:   NewFactor(1e12),
	})

	RegisterUnit(&Unit{
		Key:      "PB",
		Category: CatData,
		Names:    []string{"PB", "petabyte"},
		Factor:   NewFactor(1e15),
	})

	RegisterUnit(&Unit{
		Key:      "KiB",
		Category: CatData,
		Names:    []string{"KiB", "kibibyte"},
		Factor:   NewFactor(1024),
	})

	RegisterUnit(&Unit{
		Key:      "MiB",
		Category: CatData,
		Names:    []string{"MiB", "mebibyte"},
		Factor:   NewFactor(1048576),
	})

	RegisterUnit(&Unit{
		Key:      "GiB",
		Category: CatData,
		Names:    []string{"GiB", "gibibyte"},
		Factor:   NewFactor(1073741824),
	})

	RegisterUnit(&Unit{
		Key:      "TiB",
		Category: CatData,
		Names:    []string{"TiB", "tebibyte"},
		Factor:   NewFactor(1099511627776),
	})

	RegisterUnit(&Unit{
		Key:      "PiB",
		Category: CatData,
		Names:    []string{"PiB", "pebibyte"},
		Factor:   NewFactor(1125899906842624),
	})
}

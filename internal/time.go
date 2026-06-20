package internal

func init() {
	RegisterUnit(&Unit{
		Key:      "s",
		Category: CatTime,
		Names:    []string{"秒", "s", "second"},
		Factor:   NewFactor(1),
	})
	RegisterUnit(&Unit{
		Key:      "ms",
		Category: CatTime,
		Names:    []string{"毫秒", "ms", "millisecond"},
		Factor:   NewFactor(0.001),
	})
	RegisterUnit(&Unit{
		Key:      "us",
		Category: CatTime,
		Names:    []string{"微秒", "μs", "microsecond", "us"},
		Factor:   NewFactor(0.000001),
	})
	RegisterUnit(&Unit{
		Key:      "min",
		Category: CatTime,
		Names:    []string{"分钟", "min", "minute"},
		Factor:   NewFactor(60),
	})
	RegisterUnit(&Unit{
		Key:      "h",
		Category: CatTime,
		Names:    []string{"小时", "h", "hour"},
		Factor:   NewFactor(3600),
	})
	RegisterUnit(&Unit{
		Key:      "d",
		Category: CatTime,
		Names:    []string{"天", "d", "day"},
		Factor:   NewFactor(86400),
	})
	RegisterUnit(&Unit{
		Key:      "wk",
		Category: CatTime,
		Names:    []string{"周", "wk", "week"},
		Factor:   NewFactor(604800),
	})
	RegisterUnit(&Unit{
		Key:      "mo",
		Category: CatTime,
		Names:    []string{"月", "mo", "month"},
		Factor:   NewFactor(2592000),
	})
	RegisterUnit(&Unit{
		Key:      "yr",
		Category: CatTime,
		Names:    []string{"年", "yr", "year"},
		Factor:   NewFactor(31536000),
	})
	RegisterUnit(&Unit{
		Key:      "century",
		Category: CatTime,
		Names:    []string{"世纪", "century"},
		Factor:   NewFactor(3153600000),
	})
}

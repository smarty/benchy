package assertions

func fib() {
	_fib(30)
}

func fibWithCache() {
	fibCache := make(map[int]int)
	_fibWithCache(30, &fibCache)
}

func _fib(n int) int {
	if n <= 1 {
		return 1
	}

	return _fib(n-1) + _fib(n-2)
}

func _fibWithCache(n int, fibCache *map[int]int) int {
	if n <= 1 {
		return 1
	}

	if v, ok := (*fibCache)[n]; ok {
		return v
	}

	v := _fibWithCache(n-1, fibCache) + _fibWithCache(n-2, fibCache)
	(*fibCache)[n] = v
	return v
}

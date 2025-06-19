package providers

// FuncProvider provides a benchmark function.
type FuncProvider interface {
	// BenchmarkFunc returns a niladic function which wraps the template
	// function that was passed during construction. The returned function
	// can be used as a benchmark function with Benchy.
	BenchmarkFunc() func()
}

// Provider1 is a benchmark data-injector and provider.
type Provider1[T0 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0) Provider1[T0]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0)) func()
}

// Provider2 is a benchmark data-injector and provider.
type Provider2[T0 any, T1 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0, v1 T1) Provider2[T0, T1]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0, T1)) func()
}

// Provider3 is a benchmark data-injector and provider.
type Provider3[T0 any, T1 any, T2 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0, v1 T1, v2 T2) Provider3[T0, T1, T2]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0, T1, T2)) func()
}

// Provider4 is a benchmark data-injector and provider.
type Provider4[T0 any, T1 any, T2 any, T3 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0, v1 T1, v2 T2, v3 T3) Provider4[T0, T1, T2, T3]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0, T1, T2, T3)) func()
}

// Provider5 is a benchmark data-injector and provider.
type Provider5[T0 any, T1 any, T2 any, T3 any, T4 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4) Provider5[T0, T1, T2, T3, T4]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0, T1, T2, T3, T4)) func()
}

// Provider6 is a benchmark data-injector and provider.
type Provider6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any] interface {
	FuncProvider

	// Add adds a new row to this provider.
	Add(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Provider6[T0, T1, T2, T3, T4, T5]

	// WrapBenchmarkFunc takes a function which matches the required signature
	// and returns a niladic function that can be passed into Benchy as a
	// benchmark function.
	WrapBenchmarkFunc(target func(T0, T1, T2, T3, T4, T5)) func()
}

package providers

type ring[T any] struct {
	Index int
}

func (this *ring[T]) next(valuesLength int) {
	this.Index++
	if this.Index >= valuesLength {
		this.Index = 0
	}
}

// ########## 1 ##########

// RingProvider1 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider1[T0 any] struct {
	ring[RingProvider1[T0]]
	values    []Tuple1[T0]
	benchmark func()
}

// New1 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New1[T0 any](target func(T0)) *RingProvider1[T0] {
	provider := new(RingProvider1[T0])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider1[T0]) Add(v0 T0) Provider1[T0] {
	this.values = append(this.values, Tuple1[T0]{v0})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider1[T0]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider1[T0]) WrapBenchmarkFunc(target func(T0)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider1[T0]) value() T0 {
	tuple := this.values[this.ring.Index]
	return tuple.Value0
}

// ########## 2 ##########

// RingProvider2 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider2[T0 any, T1 any] struct {
	ring[RingProvider2[T0, T1]]
	values    []Tuple2[T0, T1]
	benchmark func()
}

// New2 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New2[T0 any, T1 any](target func(T0, T1)) *RingProvider2[T0, T1] {
	provider := new(RingProvider2[T0, T1])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider2[T0, T1]) Add(v0 T0, v1 T1) Provider2[T0, T1] {
	this.values = append(this.values, Tuple2[T0, T1]{v0, v1})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider2[T0, T1]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider2[T0, T1]) WrapBenchmarkFunc(target func(T0, T1)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider2[T0, T1]) value() (T0, T1) {
	tuple := this.values[this.ring.Index]
	return tuple.Value0, tuple.Value1
}

// ########## 3 ##########

// RingProvider3 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider3[T0 any, T1 any, T2 any] struct {
	ring[RingProvider3[T0, T1, T2]]
	values    []Tuple3[T0, T1, T2]
	benchmark func()
}

// New3 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New3[T0 any, T1 any, T2 any](target func(T0, T1, T2)) *RingProvider3[T0, T1, T2] {
	provider := new(RingProvider3[T0, T1, T2])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider3[T0, T1, T2]) Add(v0 T0, v1 T1, v2 T2) Provider3[T0, T1, T2] {
	this.values = append(this.values, Tuple3[T0, T1, T2]{v0, v1, v2})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider3[T0, T1, T2]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider3[T0, T1, T2]) WrapBenchmarkFunc(target func(T0, T1, T2)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider3[T0, T1, T2]) value() (T0, T1, T2) {
	tuple := this.values[this.ring.Index]
	return tuple.Value0, tuple.Value1, tuple.Value2
}

// ########## 4 ##########

// RingProvider4 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider4[T0 any, T1 any, T2 any, T3 any] struct {
	ring[RingProvider4[T0, T1, T2, T3]]
	values    []Tuple4[T0, T1, T2, T3]
	benchmark func()
}

// New4 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New4[T0 any, T1 any, T2 any, T3 any](target func(T0, T1, T2, T3)) *RingProvider4[T0, T1, T2, T3] {
	provider := new(RingProvider4[T0, T1, T2, T3])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider4[T0, T1, T2, T3]) Add(v0 T0, v1 T1, v2 T2, v3 T3) Provider4[T0, T1, T2, T3] {
	this.values = append(this.values, Tuple4[T0, T1, T2, T3]{v0, v1, v2, v3})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider4[T0, T1, T2, T3]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider4[T0, T1, T2, T3]) WrapBenchmarkFunc(target func(T0, T1, T2, T3)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider4[T0, T1, T2, T3]) value() (T0, T1, T2, T3) {
	tuple := this.values[this.ring.Index]
	return tuple.Value0, tuple.Value1, tuple.Value2, tuple.Value3
}

// ########## 5 ##########

// RingProvider5 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider5[T0 any, T1 any, T2 any, T3 any, T4 any] struct {
	ring[RingProvider5[T0, T1, T2, T3, T4]]
	values    []Tuple5[T0, T1, T2, T3, T4]
	benchmark func()
}

// New5 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New5[T0 any, T1 any, T2 any, T3 any, T4 any](target func(T0, T1, T2, T3, T4)) *RingProvider5[T0, T1, T2, T3, T4] {
	provider := new(RingProvider5[T0, T1, T2, T3, T4])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider5[T0, T1, T2, T3, T4]) Add(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4) Provider5[T0, T1, T2, T3, T4] {
	this.values = append(this.values, Tuple5[T0, T1, T2, T3, T4]{v0, v1, v2, v3, v4})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider5[T0, T1, T2, T3, T4]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider5[T0, T1, T2, T3, T4]) WrapBenchmarkFunc(target func(T0, T1, T2, T3, T4)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider5[T0, T1, T2, T3, T4]) value() (T0, T1, T2, T3, T4) {
	tuple := this.values[this.ring.Index]
	return tuple.Value0, tuple.Value1, tuple.Value2, tuple.Value3, tuple.Value4
}

// ########## 6 ##########

// RingProvider6 stores value tuples that can repeatedly be accessed in sequence.
type RingProvider6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	ring[RingProvider6[T0, T1, T2, T3, T4, T5]]
	values    []Tuple6[T0, T1, T2, T3, T4, T5]
	benchmark func()
}

// New6 generates a new ring provider that can repeatedly be accessed in
// sequence.
func New6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any](target func(T0, T1, T2, T3, T4, T5)) *RingProvider6[T0, T1, T2, T3, T4, T5] {
	provider := new(RingProvider6[T0, T1, T2, T3, T4, T5])
	provider.benchmark = func() {
		provider.next(len(provider.values))
		target(provider.value())
	}

	return provider
}

// Add adds a new row to this provider.
func (this *RingProvider6[T0, T1, T2, T3, T4, T5]) Add(v0 T0, v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Provider6[T0, T1, T2, T3, T4, T5] {
	this.values = append(this.values, Tuple6[T0, T1, T2, T3, T4, T5]{v0, v1, v2, v3, v4, v5})
	return this
}

// BenchmarkFunc returns a niladic function which wraps the template
// function that was passed during construction. The returned function
// can be used as a benchmark function with Benchy.
func (this *RingProvider6[T0, T1, T2, T3, T4, T5]) BenchmarkFunc() func() {
	return this.benchmark
}

// WrapBenchmarkFunc takes a function which matches the required signature
// and returns a niladic function that can be passed into Benchy as a
// benchmark function.
func (this *RingProvider6[T0, T1, T2, T3, T4, T5]) WrapBenchmarkFunc(target func(T0, T1, T2, T3, T4, T5)) func() {
	return func() {
		this.next(len(this.values))
		target(this.value())
	}
}

func (this *RingProvider6[T0, T1, T2, T3, T4, T5]) value() (T0, T1, T2, T3, T4, T5) {
	tuple := this.values[this.ring.Index]
	return tuple.Value0, tuple.Value1, tuple.Value2, tuple.Value3, tuple.Value4, tuple.Value5
}

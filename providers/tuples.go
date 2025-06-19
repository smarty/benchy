package providers

// Tuple1 is a carrier for 1 value. This exists only for the sake of consistency.
type Tuple1[T0 any] struct {
	Value0 T0
}

// Tuple2 is a carrier for 2 values.
type Tuple2[T0 any, T1 any] struct {
	Value0 T0
	Value1 T1
}

// Tuple3 is a carrier for 3 values.
type Tuple3[T0 any, T1 any, T2 any] struct {
	Value0 T0
	Value1 T1
	Value2 T2
}

// Tuple4 is a carrier for 4 values.
type Tuple4[T0 any, T1 any, T2 any, T3 any] struct {
	Value0 T0
	Value1 T1
	Value2 T2
	Value3 T3
}

// Tuple5 is a carrier for 5 values.
type Tuple5[T0 any, T1 any, T2 any, T3 any, T4 any] struct {
	Value0 T0
	Value1 T1
	Value2 T2
	Value3 T3
	Value4 T4
}

// Tuple6 is a carrier for 6 values.
type Tuple6[T0 any, T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	Value0 T0
	Value1 T1
	Value2 T2
	Value3 T3
	Value4 T4
	Value5 T5
}

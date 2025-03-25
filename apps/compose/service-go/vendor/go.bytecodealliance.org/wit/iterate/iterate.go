package iterate

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// TODO: delete this once [GOEXPERIMENT=rangefunc] lands.
//
// [GOEXPERIMENT=rangefunc]: https://go.dev/wiki/RangefuncExperiment
type Seq[V any] func(yield func(V) bool)

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
//
// TODO: delete this once [GOEXPERIMENT=rangefunc] lands.
//
// [GOEXPERIMENT=rangefunc]: https://go.dev/wiki/RangefuncExperiment
type Seq2[K, V any] func(yield func(K, V) bool)

// Done wraps yield and calls done when yield returns false.
func Done[V any](yield func(V) bool, done func()) func(V) bool {
	return func(v V) bool {
		if !yield(v) {
			done()
			return false
		}
		return true
	}
}

// Done2 wraps yield and calls done when yield returns false.
func Done2[K, V any](yield func(K, V) bool, done func()) func(K, V) bool {
	return func(k K, v V) bool {
		if !yield(k, v) {
			done()
			return false
		}
		return true
	}
}

// Once wraps yield to ensure each unique value is only yielded once.
func Once[V comparable](yield func(V) bool) func(V) bool {
	m := make(map[V]struct{})
	return func(v V) bool {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
		return yield(v)
	}
}

// Once2 wraps yield to ensure each unique value (but not key) is only yielded once.
// TODO: necessary?
func Once2[K any, V comparable](yield func(K, V) bool) func(K, V) bool {
	m := make(map[V]struct{})
	return func(k K, v V) bool {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
		return yield(k, v)
	}
}

// Package funky provides functional sugar-filled wrappers for builtin
// iterable types.
package funky

// Slice creates a funky Slice which is basically a builtin slice
// with additional functional-style sugar on top of it.
type Slice []interface{}

// Filter filters the slice sequentually applying predicate
// function provided to each slice element.
// Returns slice containing elements for which fn predicate returned true.
func (s Slice) Filter(fn func(interface{}) bool) (result Slice) {
	for _, item := range s {
		if fn(item) {
			result = append(result, item)
		}
	}
	return
}

// Map applies function provided to each element of slice and returns new
// Slice containing elements returned by mapping function fn.
func (s Slice) Map(fn func(interface{}) interface{}) (result Slice) {
	for _, item := range s {
		result = append(result, fn(item))
	}
	return
}

// Reduce reduces Slice s with the fn function provided to a singe value.
// Value returned by previous fn call is passed as its first argument on the next run.
func (s Slice) Reduce(fn func(interface{}, interface{}) interface{}) (collected interface{}) {
	for i := 1; i < len(s); i++ {
		collected = fn(s[i-1], s[i])
	}
	return
}

// Append returns new Slice containing elements from Slice s and items
// provided appended to its end.
func (s Slice) Append(items ...interface{}) Slice {
	return Slice(append(s, items...))
}

// Delete returns new Slice containing elements from Slice s with
// element at raget index removed.
func (s Slice) Delete(index int) Slice {
	return Slice(append(s[:index], s[index+1:]...))
}

// Contains returns true if Slice s contains target item
// item equality is determined by the == operator.
func (s Slice) Contains(item interface{}) bool {
	for _, element := range s {
		if element == item {
			return true
		}
	}
	return false
}

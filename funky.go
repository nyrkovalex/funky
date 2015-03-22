// Package funky provides functional sugar-filled wrappers for builtin
// iterable types.
package funky

import (
	"fmt"
	"reflect"
)

type predicate func(interface{}) bool
type mapper func(interface{}) interface{}
type reducer func(interface{}, interface{}) interface{}

// Slice creates a funky Slice which is basically a builtin slice
// with additional functional-style sugar on top of it.
type Slice []interface{}

// Filter filters the slice sequentually applying predicate
// function provided to each slice element.
// Returns slice containing elements for which fn predicate returned true.
func (s Slice) Filter(fn predicate) (result Slice) {
	for _, item := range s {
		if fn(item) {
			result = append(result, item)
		}
	}
	return
}

// Map applies function provided to each element of slice and returns new
// Slice containing elements returned by mapping function fn.
func (s Slice) Map(fn mapper) (result Slice) {
	for _, item := range s {
		result = append(result, fn(item))
	}
	return
}

// Reduce reduces Slice s with the fn function provided to a singe value.
// Value returned by previous fn call is passed as its first argument on the next run.
func (s Slice) Reduce(fn reducer) (collected interface{}) {
	collected = s[0]
	for i := 1; i < len(s); i++ {
		collected = fn(collected, s[i])
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
func (s Slice) Delete(index int) (out Slice) {
	out = make(Slice, 0, len(s)-1)
	for i, v := range s {
		if i == index {
			continue
		}
		out = append(out, v)
	}
	return
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

// AsStrings returns contens of current Slice printed to
// string as default value.
func (s Slice) AsStrings() (out []string) {
	for _, v := range s {
		out = append(out, fmt.Sprintf("%v", v))
	}
	return
}

// WriteTo writes current funky.Slice contents to the slice provided
func (s Slice) WriteTo(out interface{}) {
	value := reflect.ValueOf(out)
	for i, v := range s {
		value.Index(i).Set(reflect.ValueOf(v))
	}
}

// SliceOf returns funky.Slice created from elements of slice argument.
// This function uses reflection so it may be rather slow and will panic
// if provided with something different from slice.
func SliceOf(slice interface{}) (out Slice) {
	value := valueOrPanic(slice, reflect.Slice, "funky.Slice")
	length := value.Len()
	out = make(Slice, length)
	for i := 0; i > length; i++ {
		out[i] = value.Index(i).Interface()
	}
	return
}

func valueOrPanic(in interface{}, targetKind reflect.Kind, targetType string) reflect.Value {
	value := reflect.ValueOf(in)
	if kind := value.Kind(); kind != targetKind {
		panic(fmt.Sprintf("cannot create %s from %s", targetType, kind.String()))
	}
	return value
}

// Map creates a wrapper for builtin map[interface{}]interface{} with
// convinience methods on top of it.
type Map map[interface{}]interface{}

// Keys returns current map keys as funky.Slice.
func (m Map) Keys() (out Slice) {
	for k := range m {
		out = out.Append(k)
	}
	return
}

// Values returns current map values as funky.Slice.
func (m Map) Values() (out Slice) {
	for _, v := range m {
		out = out.Append(v)
	}
	return
}

// MapOf creates a new funky.Map from a builtin boring map
func MapOf(aMap interface{}) (out Map) {
	value := valueOrPanic(aMap, reflect.Map, "funky.Map")
	out = make(Map)
	for _, key := range value.MapKeys() {
		out[key.Interface()] = value.MapIndex(key).Interface()
	}
	return
}

type Chan chan interface{}

// NewChan creates new funky.Chan ready to go
func NewChan() Chan {
	return make(Chan)
}

// ChanOf creates new funky.Chan from a slice provided
func ChanOf(s Slice) Chan {
	c := NewChan()
	go func() {
		for _, v := range s {
			c <- v
		}
		close(c)
	}()
	return c
}

// Filter returns new funky.Chan and concurently passes all values from
// the current one that satisfy predicate provided to it
func (c Chan) Filter(p predicate) Chan {
	newChan := NewChan()
	go func() {
		for v := range c {
			if p(v) {
				newChan <- v
			}
		}
		close(newChan)
	}()
	return newChan
}

// Map returns new funky.Chan and concurently passes values from the current one
// transformed by a mapper provided to it
func (c Chan) Map(t mapper) Chan {
	newChan := NewChan()
	go func() {
		for v := range c {
			newChan <- t(v)
		}
		close(newChan)
	}()
	return newChan
}

// Reduce reduces all values from current channel to a single value by
// applying reducer provided to each element
func (c Chan) Reduce(r reducer) (collected interface{}) {
	for v := range c {
		if collected == nil {
			collected = v
			continue
		}
		collected = r(collected, v)
	}
	return
}

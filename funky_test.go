package funky

import (
	"fmt"
	. "github.com/nyrkovalex/testme"
	"reflect"
	"testing"
)

type SliceTest struct {
	slice Slice
}

func TestSlice(t *testing.T) {
	Run(t, SliceTest{Slice{1, 2, 3}})
}

func (t SliceTest) TestShouldFilterSlice(e *Expect) {
	result := t.slice.Filter(func(item interface{}) bool {
		return item.(int) < 2
	})
	e.Expect(len(result)).ToBe(1)
	e.Expect(result[0]).ToBe(1)
}

func (t SliceTest) TestShouldMapSlice(e *Expect) {
	result := t.slice.Map(func(item interface{}) interface{} {
		return fmt.Sprintf("mapped %d", item.(int))
	})
	e.Expect(len(result)).ToBe(3)
	e.Expect(result[0]).ToBe("mapped 1")
}

func (t SliceTest) TestShouldContainItem(e *Expect) {
	result := t.slice.Contains(2)
	e.Expect(result).ToBe(true)
}

func (t SliceTest) TestShouldNotContainItem(e *Expect) {
	result := t.slice.Contains(42)
	e.Expect(result).ToBe(false)
}

func (t SliceTest) TestShouldAppendItem(e *Expect) {
	result := t.slice.Append(4)
	e.Expect(len(result)).ToBe(4)
	e.Expect(result[3]).ToBe(4)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t SliceTest) TestShouldAppendMultipleItems(e *Expect) {
	result := t.slice.Append(4, 5, 6)
	e.Expect(len(result)).ToBe(6)
	e.Expect(result[3]).ToBe(4)
	e.Expect(result[4]).ToBe(5)
	e.Expect(result[5]).ToBe(6)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t SliceTest) TestShouldDeleteItemByIndex(e *Expect) {
	result := t.slice.Delete(1)
	e.Expect(len(result)).ToBe(2)
	e.Expect(result[1]).ToBe(3)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t SliceTest) TestShouldReduceToSum(e *Expect) {
	result := t.slice.Reduce(func(left interface{}, right interface{}) interface{} {
		return left.(int) + right.(int)
	})
	e.Expect(result).ToBe(6)
}

func (t SliceTest) TestShouldChainCalls(e *Expect) {
	result := t.slice.Filter(func(item interface{}) bool {
		return item.(int) == 2
	}).Map(func(item interface{}) interface{} {
		return fmt.Sprintf("%d", item.(int))
	})
	e.Expect(len(result)).ToBe(1)
	e.Expect(result[0]).ToBe("2")
}

func (t SliceTest) TestShouldCreateFunkySliceOfIntegers(e *Expect) {
	slice := SliceOf([]int{1, 2, 3})
	e.Expect(len(slice)).ToBe(3)
	e.Expect(reflect.TypeOf(slice).Name()).ToBe("Slice")
}

func (t SliceTest) TestShouldPanic(e *Expect) {
	e.Expect(func() {
		SliceOf(1)
	}).ToPanic("cannot create funky.Slice from int")
}

type MapTest struct {
	funkyMap Map
}

func TestMap(t *testing.T) {
	Run(t, MapTest{Map{"a": 1, "b": 2, "c": 3}})
}

func (t MapTest) TestShouldRetrieveMapKeys(e *Expect) {
	keys := t.funkyMap.Keys()
	e.Expect(len(keys)).ToBe(3)
	e.Expect(keys.Contains("a")).ToBe(true)
	e.Expect(keys.Contains("b")).ToBe(true)
	e.Expect(keys.Contains("c")).ToBe(true)
}

func (t MapTest) TestShouldRetrieveMapValues(e *Expect) {
	values := t.funkyMap.Values()
	e.Expect(len(values)).ToBe(3)
	e.Expect(values.Contains(1)).ToBe(true)
	e.Expect(values.Contains(2)).ToBe(true)
	e.Expect(values.Contains(3)).ToBe(true)
}

func (t MapTest) TestShouldCreateMapFromStringIntMap(e *Expect) {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	result := MapOf(source)
	e.Expect(result["a"]).ToBe(1)
	e.Expect(result["b"]).ToBe(2)
	e.Expect(result["c"]).ToBe(3)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Map")
}

func (t MapTest) TestShouldPanicCreatingMapFromCrap(e *Expect) {
	e.Expect(func() {
		MapOf(1)
	}).ToPanic("cannot create funky.Map from int")
}

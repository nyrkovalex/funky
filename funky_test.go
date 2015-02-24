package funky

import (
	"fmt"
	. "github.com/nyrkovalex/testme"
	"reflect"
	"testing"
)

type FunkyTest struct {
	slice Slice
}

func Test(t *testing.T) {
	Run(t, FunkyTest{Slice{1, 2, 3}})
}

func (t FunkyTest) TestShouldFilterSlice(e *Expect) {
	result := t.slice.Filter(func(item interface{}) bool {
		return item.(int) < 2
	})
	e.Expect(len(result)).ToBe(1)
	e.Expect(result[0]).ToBe(1)
}

func (t FunkyTest) TestShouldMapSlice(e *Expect) {
	result := t.slice.Map(func(item interface{}) interface{} {
		return fmt.Sprintf("mapped %d", item.(int))
	})
	e.Expect(len(result)).ToBe(3)
	e.Expect(result[0]).ToBe("mapped 1")
}

func (t FunkyTest) TestShouldContainItem(e *Expect) {
	result := t.slice.Contains(2)
	e.Expect(result).ToBe(true)
}

func (t FunkyTest) TestShouldNotContainItem(e *Expect) {
	result := t.slice.Contains(42)
	e.Expect(result).ToBe(false)
}

func (t FunkyTest) TestShouldAppendItem(e *Expect) {
	result := t.slice.Append(4)
	e.Expect(len(result)).ToBe(4)
	e.Expect(result[3]).ToBe(4)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t FunkyTest) TestShouldAppendMultipleItems(e *Expect) {
	result := t.slice.Append(4, 5, 6)
	e.Expect(len(result)).ToBe(6)
	e.Expect(result[3]).ToBe(4)
	e.Expect(result[4]).ToBe(5)
	e.Expect(result[5]).ToBe(6)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t FunkyTest) TestShouldDeleteItemByIndex(e *Expect) {
	result := t.slice.Delete(1)
	e.Expect(len(result)).ToBe(2)
	e.Expect(result[1]).ToBe(3)
	e.Expect(reflect.TypeOf(result).Name()).ToBe("Slice")
}

func (t FunkyTest) TestShouldReduceToSum(e *Expect) {
	result := t.slice.Reduce(func(left interface{}, right interface{}) interface{} {
		return left.(int) + right.(int)
	})
	e.Expect(result).ToBe(6)
}

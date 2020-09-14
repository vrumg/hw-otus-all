package linkedlist

import (
	"reflect"
	"testing"
)

func TestPackage(t *testing.T) {

	//Want funcs responses
	var wantEmptyFirstLast *Item
	wantArray := []interface{}{-4, -2, -1, 1, 2, 4}

	//run tests
	t.Run("Package test", func(t *testing.T) {

		//NewList constructor
		list := NewList()
		//First
		if got := list.First(); !reflect.DeepEqual(got, wantEmptyFirstLast) {
			t.Errorf("List.First() = %v, want %v", got, wantEmptyFirstLast)
		}

		//Last
		if got := list.Last(); !reflect.DeepEqual(got, wantEmptyFirstLast) {
			t.Errorf("List.Last() = %v, want %v", got, wantEmptyFirstLast)
		}

		//init() and Init() constructor
		list = &List{}
		list.lazyInit()
		//First
		if got := list.First(); !reflect.DeepEqual(got, wantEmptyFirstLast) {
			t.Errorf("List.First() = %v, want %v", got, wantEmptyFirstLast)
		}

		//Last
		if got := list.Last(); !reflect.DeepEqual(got, wantEmptyFirstLast) {
			t.Errorf("List.Last() = %v, want %v", got, wantEmptyFirstLast)
		}

		//PushFront
		list.PushFront(1)
		list.PushFront(2)
		list.PushFront(3)
		list.PushFront(4)
		//PushBack
		list.PushBack(-1)
		list.PushBack(-2)
		list.PushBack(-3)
		list.PushBack(-4)

		//Remove, First, Next, Prev
		//remove second element from head 3
		list.Remove(list.First().Next())
		list.Remove(list.First().Prev())

		//Remove, Last, Next, Prev
		//remove second element from tail -3
		list.Remove(list.Last().Next())
		list.Remove(list.Last().Prev())

		if got := list.ListValues(); !reflect.DeepEqual(got, wantArray) {
			t.Errorf("List.ListValues() = %v, want %v", got, wantArray)
		}
	})
}

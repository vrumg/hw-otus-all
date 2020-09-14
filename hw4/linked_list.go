package linkedlist

//Item is an element of a linked list
type Item struct {
	//Next and previous Item pointers
	next, prev *Item
	//List reference
	list *List
	//Item value
	Val interface{}
}

//Next returns the next list Item or nil.
func (s *Item) Next() *Item {
	if p := s.next; s.list != nil && p != &s.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list Item or nil.
func (s *Item) Prev() *Item {
	if p := s.prev; s.list != nil && p != &s.list.root {
		return p
	}
	return nil
}

//Value returns copy of Item value
func (s *Item) Value() interface{} {
	return s.Val
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Item //first Item
	//last Item //last Item
	len int //list length
}

// Init initializes or clears list l.
func (s *List) Init() *List {
	s.root.next = &s.root
	s.root.prev = &s.root
	s.len = 0
	return s
}

// lazyInit lazily initializes a zero List value.
func (s *List) lazyInit() {
	if s.root.next == nil {
		s.Init()
	}
}

//NewList returns an initialized list.
func NewList() *List {
	return new(List).Init()
}

//Len returns the number of Item of list l.
func (s *List) Len() int {
	return s.len
}

//First returns the first Item of list l or nil if the list is empty.
func (s *List) First() *Item {
	if s.len == 0 {
		return nil
	}
	return s.root.next
}

//Last returns the last Item of list l or nil if the list is empty.
func (s *List) Last() *Item {
	if s.len == 0 {
		return nil
	}
	return s.root.prev
}

//insert inserts newItem after atItem, increments l.len, and returns e.
func (s *List) insert(newItem, atItem *Item) *Item {
	nextItem := atItem.next //save next item pointer
	atItem.next = newItem   //write new next item reference
	newItem.prev = atItem   //write back reference to atItem
	newItem.next = nextItem //write previously saved pointer to nextItem
	nextItem.prev = newItem //link new item and nextItem
	newItem.list = s        //link newItem with List
	s.len++                 //increment list elements counter
	return newItem          //return newItem
}

//insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (s *List) insertValue(v interface{}, at *Item) *Item {
	return s.insert(&Item{Val: v}, at)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (s *List) PushFront(v interface{}) {
	s.lazyInit()
	s.insertValue(v, &s.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (s *List) PushBack(v interface{}) {
	s.lazyInit()
	s.insertValue(v, s.root.prev)
}

// remove removes e from its list, decrements l.len, and returns e.
func (s *List) remove(i *Item) {
	i.prev.next = i.next //write link from previous to next element
	i.next.prev = i.prev //write link from next to previos element
	i.next = nil         //remove links for given element
	i.prev = nil         //remove links for given element
	i.list = nil         //remove links for given element
	s.len--              //decrement total list length
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (s *List) Remove(i *Item) {

	//check if item is correct
	if i == nil {
		return
	}

	//check if item reffers to current list
	if i.list == s {
		s.remove(i)
	}
}

//ListValues returns array of values in list
func (s *List) ListValues() []interface{} {

	values := make([]interface{}, 0, s.Len())

	nextItem := s.Last()

	for {
		values = append(values, nextItem.Value())
		nextItem = nextItem.Prev()
		if nextItem == nil {
			break
		}
	}

	return values
}


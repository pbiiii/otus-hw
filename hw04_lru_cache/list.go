package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type MyList struct {
	length int
	first  *ListItem
	last   *ListItem
}

func (l *MyList) Front() *ListItem {
	return l.first
}

func (l *MyList) Back() *ListItem {
	return l.last
}

func (l *MyList) PushFront(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushToEmpty(v)
	}

	oldFirst := l.Front()
	newFirst := &ListItem{
		Next:  oldFirst,
		Prev:  nil,
		Value: v,
	}

	oldFirst.Prev = newFirst
	l.first = newFirst
	l.length++

	return newFirst
}

func (l *MyList) PushBack(v interface{}) *ListItem {
	if l.Len() == 0 {
		return l.pushToEmpty(v)
	}

	oldLast := l.Back()
	newLast := &ListItem{
		Next:  nil,
		Prev:  oldLast,
		Value: v,
	}

	oldLast.Next = newLast
	l.last = newLast
	l.length++

	return newLast
}

func (l *MyList) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next

	// Удаляемый элемент единственный в списке
	if prev == nil && next == nil {
		l.first = nil
		l.last = nil
		l.length = 0

		return
	}

	// Удаляемый элемент первый в списке
	if prev == nil {
		l.first = next
		next.Prev = nil
		l.length--

		return
	}

	// Удаляемый элемент последний в списке
	if next == nil {
		l.last = prev
		prev.Next = nil
		l.length--

		return
	}

	prev.Next = next
	next.Prev = prev
	l.length--
}

func (l *MyList) MoveToFront(i *ListItem) {
	prev := i.Prev

	// Элемент уже в начале
	if prev == nil {
		return
	}

	// Если элемент находится в конце списка
	if next := i.Next; next == nil {
		prev.Next = nil
		l.last = prev
	} else {
		prev.Next = next
		next.Prev = prev
		l.last = next
	}

	first := l.Front()
	i.Prev = nil
	i.Next = first
	first.Prev = i
	l.first = i
}

func (l *MyList) Len() int {
	return l.length
}

func (l *MyList) pushToEmpty(v interface{}) *ListItem {
	item := &ListItem{
		Next:  nil,
		Prev:  nil,
		Value: v,
	}

	l.length++
	l.last = item
	l.first = item

	return item
}

func NewList() List {
	return new(MyList)
}

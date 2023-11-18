package stack

import "sync"

type (
	Stack[T any] interface {
		Len() int
		Peek() *T
		Pop() *T
		Push(value T)
	}
	threadsafeStack[T any] struct {
		lock sync.Mutex
		s stackImp[T]
	}
	stackImp[T any] struct {
		top *node[T]
		length int
	}
	node[T any] struct {
		item T
		prev *node[T]
	}	
)

// New stack
func New[T any](threadsafe bool) Stack[T] {
	if threadsafe {
		return &threadsafeStack[T] {
			lock: sync.Mutex{},
			s: stackImp[T]{nil,0},
		}
	}
	return &stackImp[T]{nil,0}
}

// Len
func (stack *threadsafeStack[T]) Len() int {
	return stack.s.Len()
}

// Peek 
func (stack *threadsafeStack[T]) Peek() *T {
	return stack.s.Peek()
}

// Pop
func (stack *threadsafeStack[T]) Pop() *T {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	return stack.s.Pop()
}

//Push
func (stack *threadsafeStack[T]) Push(item T) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.s.Push(item)
}

// Len return number of items in stack
func (stack *stackImp[T]) Len() int {
	return stack.length
}

// Peek top of stack
func (stack *stackImp[T]) Peek() *T {
	if stack.length == 0 {
		return nil
	}
	return &stack.top.item
}

// Pop top item from stack
func (stack *stackImp[T]) Pop() *T {
	if stack.length == 0 {
		return nil
	}
	
	n := stack.top
	stack.top = n.prev
	stack.length--
	return &n.item
}

// Push item onto stack
func (stack *stackImp[T]) Push(item T) {
	n := &node[T]{item, stack.top}
	stack.top = n
	stack.length++
}

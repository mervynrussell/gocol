package stack

import (
	"log"
	"sync"
	"testing"
)

func TestLen(t *testing.T) {
	stk := New[int](false)
	if stk.Len() > 0 {
		t.Fatalf("expected 0 got %d", stk.Len())
	}

	stk = New[int](true)
	if stk.Len() > 0 {
		t.Fatalf("expected 0 got %d", stk.Len())
	}
}

func pushPopTest(stk Stack[int], t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6}
	for _, x := range v {
		stk.Push(x)
		if v := stk.Peek(); *v != x {
			t.Fatalf("Peek expected %d got %d", x, *v)
		}
	}
	for i := len(v) - 1; i >= 0; i-- {
		if z := stk.Peek(); z != nil {
			if *z != v[i] {
				t.Fatalf("Peek expected %d got  %d", i, *z)
			}
		}
		if y := stk.Pop(); y != nil {
			if *y != v[i] {
				t.Fatalf("Pop expected %d got %d", v[i], *y)
			}
		} else {
			t.Fatal("Pop non-empty stack returns nil")
		}
	}
	n := stk.Pop()
	if n != nil {
		t.Fatal("Pop empty stack was not nil")
	}
}

func TestPushPop(t *testing.T) {
	stk := New[int](false)
	pushPopTest(stk, t)
}

func TestPushPopThreadsafe(t *testing .T) {
	stk := New[int](true)
	waiter := new(sync.WaitGroup)
	for i := 0; i <= 9; i++ {
		waiter.Add(1)
		i := i
		go func() {
			defer waiter.Done()
			stk.Push(i)
			stk.Push(i + 10)
		}()
	}
	waiter.Wait()

	if stk.Len() != 20 {
		t.Fatalf("expected 20 got %d", stk.Len())
	}

	
	for i := 0; i <= 9; i++ {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			peekV := *stk.Peek()
			popV := *stk.Pop()
			if peekV != popV {
				log.Printf("Peek %d != Pop %d, another thread pre-empted this one", peekV, popV)
			}
		}()
	}
	waiter.Wait()

	if stk.Len() != 10 {
		t.Fatalf("expected 10 got %d", stk.Len())
	}

	for i := 0; i <= 9; i++ {
		waiter.Add(1)
		go func() {
			defer waiter.Done()
			peekV := *stk.Peek()
			popV := *stk.Pop()
			if peekV != popV {
				log.Printf("Peek %d != Pop %d, another thread pre-empted this one", peekV, popV)
			}
		}()
	}
	waiter.Wait()

	if stk.Peek() != nil {
		t.Fatalf("expected nil got %d", stk.Peek())
	}

}
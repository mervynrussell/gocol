package set

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"testing"
)

type testStruct struct {
	a int
	b string
	c float32
	d float64
}

func TestUnion(t *testing.T) {
	// Union of disjoint sets
	s1 := NewFrom([]int{1, 2, 3})

	s2 := New[int]()
	s2.AddAll([]int{4, 5, 6})

	u1 := s1.Union(s2)
	a := u1.All()
	sort.Ints(a)
	if !reflect.DeepEqual(a, []int{1, 2, 3, 4, 5, 6}) {
		t.Fatal("disjoint union error")
	}

	// Union of intersecting sets
	s2.Clear()
	s2.AddAll([]int{2, 3, 4})

	u1 = s1.Union(s2)
	a = u1.All()
	sort.Ints(a)
	if !reflect.DeepEqual(a, []int{1, 2, 3, 4}) {
		t.Fatal("intersecting union error")
	}
}

func TestSubset(t *testing.T) {
	s1 := New[int]()
	s1.AddAll([]int{1, 2, 3})
	s2 := New[int]()

	// Test superset
	s2.AddAll([]int{1, 2, 3, 4})
	if s1.Subset(s2) {
		t.Fatal("expected false")
	}

	// Test equal 
	s2.Clear()
	s2.AddAll([]int{1, 2, 3})
	if !s1.Subset(s2) {
		t.Fatal("expected true")
	}

	// Test disjoint
	s2.Clear()
	s2.AddAll([]int{4, 5, 6})
	if s1.Subset(s2) {
		t.Fatal("expected false")
	}

	// Test intersecting
	s2.Clear()
	s2.AddAll([]int{1, 2, 4})
	if s1.Subset(s2) {
		t.Fatal("expected false")
	}

	// Test subset
	s2.Clear()
	s2.AddAll([]int{1, 2})
	if !s1.Subset(s2) {
		t.Fatal("expected true")
	}

	// Test empty N.B. empty is _always_ a subset
	s2.Clear()
	if !s1.Subset(s2) {
		t.Fatal("expected true")
	}

	// Test both empty
	s1.Clear()
	if !s1.Subset(s2) {
		t.Fatal("expected true")
	}

}

func TestPop(t *testing.T) {
	s := New[int]()
	s.AddAll([]int{1, 2, 3})

	for i := 0; i < 3; i++ {
		if _, ok := s.Pop(); !ok {
			t.Fatal("expected Pop to yield an item")
		}
	}

	if !s.IsEmpty() {
		t.Fatal("expected set to be empty")
	}

	if _, ok := s.Pop(); ok {
		t.Fatal("expected not ok popping empty set")
	}

}

func TestDifference(t *testing.T) {
	// Difference of disjoint sets
	s1 := New[int]()
	s1.AddAll([]int{1, 2, 3})

	s2 := New[int]()
	s2.AddAll([]int{4, 5, 6})

	// Difference of overlapping sets
	s2.Clear()
	s2.AddAll([]int{2, 3, 4})

	d := s1.Difference(s2)
	if !reflect.DeepEqual(d.All(), []int{1}) {
		t.Fatal("difference error")
	}

	d = s2.Difference(s1)
	if !reflect.DeepEqual(d.All(), []int{4}) {
		t.Fatal("difference error")
	}
}

func TestEquals(t *testing.T) {
	// Equals disjoint sets
	s1 := New[int]()
	s1.AddAll([]int{1, 2, 3})

	s2 := New[int]()
	s2.AddAll([]int{4, 5, 6})

	if s1.Equals(s2) {
		t.Fatal("equals not expected to be true for disjoint")
	}

	// Equals overlapping sets
	s2.Clear()
	s2.AddAll([]int{2, 3, 4})

	if s1.Equals(s2) {
		t.Fatal("equals not expected to be true for overlapping")
	}

	// Equals subsets
	s2.Clear()
	s2.AddAll([]int{2, 3})

	if s1.Equals(s2) {
		t.Fatal("equals not expected to be true for subset")
	}

	// Equals matching sets
	s2.Clear()
	s2.AddAll([]int{1, 2, 3})

	if !s1.Equals(s2) {
		t.Fatal("equals expected to be true for matching sets")
	}

	// Equals empty set one side
	s2.Clear()

	if s1.Equals(s2) {
		t.Fatal("equals not expected to be true for empty set one side")
	}

	// Equals empty set both sides
	s1.Clear()

	if !s1.Equals(s2) {
		t.Fatal("equals expected to be true for empty set both sides")
	}
}

func TestIntersection(t *testing.T) {
	// Intersection of disjoint sets
	s1 := New[int]()
	s1.AddAll([]int{1, 2, 3})

	s2 := New[int]()
	s2.AddAll([]int{4, 5, 6})

	i := s1.Intersection(s2)
	if !i.IsEmpty() {
		t.Fatal("expected intersection of disjoint sets to be empty")
	}

	
	// Intersection of overlapping sets
	s2.Clear()
	s2.AddAll([]int{2, 3, 4})
	i = s1.Intersection(s2)
	a := i.All()
	sort.Ints(a)
	if !reflect.DeepEqual(a, []int{2, 3}) {
		t.Fatal("intersecting union error")
	}

}

func TestStructSet(t *testing.T) {
	s := New[testStruct]()
	things := []testStruct{}
	for i := 0; i <= 9; i++ {
		thing := testStruct{
			a: i,
			b: fmt.Sprintf("%d", i),
			c: float32(i),
			d: float64(i),
		}
		things = append(things, thing)
		s.Add(thing)
	}

	for _, thing := range things {
		if !s.Contains(thing) {
			t.Fatalf("expected %d in set", thing.a)
		}
	}

	s.Remove(things[0])
	if s.Contains(things[0]) {
		t.Fatal("Contains 0 expected false")
	}

	s.Remove(things[5])
	if s.Contains(things[5]) {
		t.Fatal("Contains 5 expected false")
	}

	s.Remove(things[9])
	if s.Contains(things[9]) {
		t.Fatal("Contains 9 expected false")
	}

	if s.Len() != 7 {
		t.Fatalf("Len expected 7 got %d", s.Len())
	}
	
	s.Clear()
	if s.Len() != 0 {
		t.Fatalf("Len expected 0 got %d", s.Len())
	}


}

func TestAll(t *testing.T) {
	s := New[int]()

	for i := 0; i <= 4; i++ {
		s.Add(i)
	}

	var c, i int
	for c, i = range s.All() {
		log.Printf("%d = %d", c, i)
	}
	if c != 4 {
		t.Fatal("")
	}
}

func TestIntSet(t *testing.T) {
	s := New[int]()

	for i := 0; i <= 9; i++ {
		s.Add(i)
	}

	if s.Len() != 10 {
		t.Fatalf("Len expected 10 got %d", s.Len())
	}

	for i := 0; i <= 9; i++ {
		if !s.Contains(i) {
			t.Fatalf("Contains %d expected true", i)
		}
	}

	s.Remove(0)
	if s.Contains(0) {
		t.Fatal("Contains 0 expected false")
	}

	s.Remove(5)
	if s.Contains(5) {
		t.Fatal("Contains 5 expected false")
	}

	s.Remove(9)
	if s.Contains(9) {
		t.Fatal("Contains 9 expected false")
	}

	s.Clear()
	if s.Len() != 0 {
		t.Fatalf("Len expected 0 got %d", s.Len())
	}
}
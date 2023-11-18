package scpattach

import (
	"reflect"
	"sort"
	"testing"
)

func TestUpdateScps(t *testing.T) {
	current := []string{"a", "b", "c", "d", "e"}
	target := []string{"f", "g", "h", "i", "j"}
	r, err := updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}

	current = []string{"a", "b", "c", "d", "e"}
	target = []string{"a", "b", "c", "d", "e"}
	r, err = updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}

	current = []string{"a", "b", "c", "d", "e"}
	target = []string{"c", "d", "e"}
	r, err = updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}

	current = []string{"a", "b", "c"}
	target = []string{"a", "b", "c", "d", "e"}
	r, err = updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}

	current = []string{"a", "f", "g"}
	target = []string{"a", "b", "c", "d", "e"}
	r, err = updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}
	
	current = []string{}
	target = []string{"a", "b", "c", "d", "e"}
	r, err = updateScps(current, target)
	if err != nil {
		t.Fatal("unexpected error")
	}
	sort.Strings(*r)
	if !reflect.DeepEqual(target, *r) {
		t.Fatal("result does not match target")
	}
	
	current = []string{"a", "b", "c", "d", "e"}
	target = []string{"a", "b", "c", "d", "e", "f"}
	_, err = updateScps(current, target)
	if err == nil {
		t.Fatal("expected error")
	}

}

package scpattach

import (
	"errors"
	"mervynrussell/gocol/pkg/set"
)

const maxAttach = 5
const minAttach = 1

type OU struct {
	attachedScps map[string]bool
}

func NewOU(s []string) OU {
	o := OU{make(map[string]bool)}
	o.attachAll(s)
	return o
}

func (o *OU) attach(s string) {
	if len(o.attachedScps) == maxAttach {
		panic("cannot attach any more policies")
	}
	o.attachedScps[s] = true
}

func (o *OU) attachAll(s []string) {
	if len(s) + len(o.attachedScps) > maxAttach {
		panic("canot attach all specified policies")
	}
	for _, i := range s {
		o.attach(i)
	}
}

func (o *OU) detach(s string) {
	if len(o.attachedScps) == minAttach {
		panic("cannot detach last policy")
	}
	delete(o.attachedScps, s)
}

func (o *OU) GetAttached() []string {
	r := make([]string, 0)
	for k := range o.attachedScps {
		r = append(r, k)
	}
	return r
}

func (o *OU) CanAttach() bool {
	return len(o.attachedScps) < maxAttach
}

func (o *OU) CanDetach() bool {
	return len(o.attachedScps) > minAttach
}

func updateScps(current []string, target []string) (*[]string, error) {
	if len(target) > maxAttach {
		return nil, errors.New("target len is greater than max attach")
	}
	unit := NewOU(current)
	toAttach := set.NewFrom(target)
	attached := set.NewFrom(unit.GetAttached())
	
	toDetach := attached.Difference(toAttach)
	toAttach = toAttach.Difference(attached)

	for toAttach.Len() > 0 || toDetach.Len() > 0 {
		if unit.CanDetach(){
			if d, ok := toDetach.Pop(); ok {
				unit.detach(*d)
			}
		}
		if unit.CanAttach() {
			if a, ok := toAttach.Pop(); ok {
				unit.attach(*a)
			}
		}
	}
	r := unit.GetAttached()
	return &r, nil
}

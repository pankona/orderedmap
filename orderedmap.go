package orderedmap

import (
	"fmt"
)

type OrderedMap struct {
	order []string
	m     map[string]Keyer
}

type Keyer interface {
	Key() string
}

func New() *OrderedMap {
	return &OrderedMap{
		order: make([]string, 0),
		m:     make(map[string]Keyer),
	}
}

func (om *OrderedMap) Add(v Keyer) error {
	if _, ok := om.m[v.Key()]; ok {
		return fmt.Errorf("specified key [%s] is already registered", v.Key())
	}
	om.order = append(om.order, v.Key())
	om.m[v.Key()] = v
	return nil
}

func (om *OrderedMap) Update(v Keyer) error {
	if _, ok := om.m[v.Key()]; !ok {
		return fmt.Errorf("specified key [%s] is not registered yet", v.Key())
	}
	om.m[v.Key()] = v
	return nil
}

func (om *OrderedMap) Len() int {
	return len(om.m)
}

func (om *OrderedMap) RemoveByIndex(index int) error {
	key := om.order[index]
	return om.removeByKey(key)
}

func (om *OrderedMap) RemoveByKey(key string) error {
	return om.removeByKey(key)
}

func (om *OrderedMap) removeByKey(key string) error {
	if _, ok := om.m[key]; !ok {
		return fmt.Errorf("specified key [%s] is not registered yet", key)
	}

	delete(om.m, key)

	neworder := make([]string, len(om.m))
	var index int
	for _, v := range om.order {
		if v == key {
			continue
		}
		neworder[index] = v
		index++
	}

	om.order = neworder
	return nil
}

func (om *OrderedMap) GetByIndex(index int) Keyer {
	key := om.order[index]
	return om.m[key]
}

func (om *OrderedMap) GetByKey(key string) Keyer {
	return om.m[key]
}

func (om *OrderedMap) ForEach(f func(v Keyer) bool) error {
	for _, key := range om.order {
		cont := f(om.m[key])
		if !cont {
			return fmt.Errorf("ForEach iteration stopped at [%s]", key)
		}
	}
	return nil
}

package orderedmap

import (
	"fmt"
)

// OrderedMap represents a map with order
type OrderedMap struct {
	order []string
	m     map[string]Keyer
}

// Keyer is an interface for OrderedMap's value.
// Return value of Key() method is used as key of map.
type Keyer interface {
	Key() string
}

// New returns a new instance of OrderedMap
func New() *OrderedMap {
	return &OrderedMap{
		order: make([]string, 0),
		m:     make(map[string]Keyer),
	}
}

// Add registers a new Keyer element on tail of OrderedMap.
// If specified Keyer's Key() is already registered, this function returns error.
func (om *OrderedMap) Add(v Keyer) error {
	return om.Insert(v, om.Len())
}

// Insert inserts a new Keyer element on specified index of OrderedMap.
// If specified Keyer's Key() is already registered, this function returns error.
func (om *OrderedMap) Insert(v Keyer, index int) error {
	if _, ok := om.m[v.Key()]; ok {
		return fmt.Errorf("specified key [%s] is already registered", v.Key())
	}

	if om.Len() <= index {
		// just append
		om.order = append(om.order, v.Key())
	} else {
		om.order = append(om.order[:index+1], om.order[index:]...)
		om.order[index] = v.Key()
	}
	om.m[v.Key()] = v
	return nil
}

// Swap swaps order of elements by specified indexes
func (om *OrderedMap) Swap(i1, i2 int) {
	om.order[i1], om.order[i2] = om.order[i2], om.order[i1]
}

// Update updates a Keyer element that is already registered on OrderedMap.
// If specified Keyer's Key() is not registered yet, this function returns error.
func (om *OrderedMap) Update(v Keyer) error {
	if _, ok := om.m[v.Key()]; !ok {
		return fmt.Errorf("specified key [%s] is not registered yet", v.Key())
	}
	om.m[v.Key()] = v
	return nil
}

// Len returns a length of OrderedMap
func (om *OrderedMap) Len() int {
	return len(om.m)
}

// RemoveByIndex removes an element from OrderedMap by specified index.
// If specified index is out of bounce, this function cause panic.
func (om *OrderedMap) RemoveByIndex(index int) error {
	key := om.order[index]
	return om.removeByKey(key)
}

// RemoveByKey removes an element from OrderedMap by specified key.
// If there's no element that has specified key, this function returns error.
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

// GetByIndex returns an element from OrderedMap by specified index.
// If specified index is out of bounce, this function cause panic.
func (om *OrderedMap) GetByIndex(index int) Keyer {
	key := om.order[index]
	return om.m[key]
}

// GetByKey returns an element from OrderedMap by specified key.
// If there's no element that has specified key, this function returns nil.
func (om *OrderedMap) GetByKey(key string) Keyer {
	return om.m[key]
}

// ForEach calls specified function with specifying registered Keyer elements
// in ordered of the elements registered.
// If specified function returns error, this function stops iteration and return error.
func (om *OrderedMap) ForEach(f func(i int, v Keyer) error) error {
	for i, key := range om.order {
		err := f(i, om.m[key])
		if err != nil {
			return err
		}
	}
	return nil
}

// Order returns order, which id represented as array of Keyer's key
func (om *OrderedMap) Order() []string {
	return om.order
}

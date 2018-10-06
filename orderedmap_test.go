package orderedmap

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type keyable int

func (k keyable) Key() string {
	return strconv.Itoa(int(k))
}

func TestAdd(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	err := om.Add(keyable(1))
	require.Nil(t, err)
	require.Equal(t, om.Len(), 1)

	err = om.Add(keyable(1))
	require.NotNil(t, err)
	require.Equal(t, om.Len(), 1)
}

func TestUpdate(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	err := om.Add(keyable(1))
	require.Nil(t, err)
	require.Equal(t, om.Len(), 1)

	err = om.Update(keyable(1))
	require.Nil(t, err)
	require.Equal(t, om.Len(), 1)

	err = om.Update(keyable(2))
	require.NotNil(t, err)
	require.Equal(t, om.Len(), 1)
}

func TestAddGet(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	for i := 0; i < 100; i++ {
		v := keyable(i)
		err := om.Add(v)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, om.Len(), i+1)
	}

	for i := 0; i < 100; i++ {
		v := om.GetByIndex(i)
		require.Equal(t, v, keyable(i), "test case [%d] failed", i)
	}

	for i := 0; i < 100; i++ {
		v := om.GetByKey(strconv.Itoa(i))
		require.Equal(t, v, keyable(i), "test case [%d] failed", i)
	}
}

func TestRemoveByIndex(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	for i := 0; i < 100; i++ {
		v := keyable(i)
		err := om.Add(v)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, om.Len(), i+1)
	}

	require.Equal(t, 100, om.Len())

	for i := 0; i < 100; i++ {
		err := om.RemoveByIndex(0)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, 100-(i+1), om.Len(), "test case [%d] failed", i)
	}
}

func TestRemoveByKey(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	for i := 0; i < 100; i++ {
		v := keyable(i)
		err := om.Add(v)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, om.Len(), i+1)
	}

	require.Equal(t, 100, om.Len())

	for i := 0; i < 100; i++ {
		err := om.RemoveByKey(strconv.Itoa(i))
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, 100-(i+1), om.Len(), "test case [%d] failed", i)
	}
}

func TestForEach(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	for i := 0; i < 100; i++ {
		v := keyable(i)
		err := om.Add(v)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, om.Len(), i+1)
	}

	var index int
	om.ForEach(func(v Keyer) bool {
		require.Equal(t, v, keyable(index), "test case [%d] failed", index)
		index++
		return true
	})
}

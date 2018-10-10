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

func TestInsert(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	for i := 0; i < 100; i++ {
		v := keyable(i)
		err := om.Add(v)
		require.Nil(t, err, "test case [%d] failed", i)
		require.Equal(t, om.Len(), i+1)
	}

	// insert middle
	v1 := keyable(500)
	err := om.Insert(v1, 50)
	require.Nil(t, err, "failed to insert")
	require.Equal(t, om.Len(), 101)

	v2 := om.GetByIndex(50)
	require.Equal(t, v1, v2)

	// insert top
	v1 = keyable(501)
	err = om.Insert(v1, 0)
	require.Nil(t, err, "failed to insert")
	require.Equal(t, om.Len(), 102)

	v2 = om.GetByIndex(0)
	require.Equal(t, v1, v2)

	// insert tail
	v1 = keyable(502)
	err = om.Insert(v1, om.Len())
	require.Nil(t, err, "failed to insert")
	require.Equal(t, om.Len(), 103)

	v2 = om.GetByIndex(om.Len() - 1)
	require.Equal(t, v1, v2)
}

func TestSwap(t *testing.T) {
	om := New()
	require.NotNil(t, om)

	v1 := keyable(1)
	err := om.Add(v1)
	require.Nil(t, err)
	require.Equal(t, om.Len(), 1)

	v2 := keyable(2)
	err = om.Add(v2)
	require.Nil(t, err)
	require.Equal(t, om.Len(), 2)

	om.Swap(0, 1)

	require.Equal(t, om.GetByIndex(0), v2)
	require.Equal(t, om.GetByIndex(1), v1)
	require.Equal(t, om.Len(), 2)
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

	om.ForEach(func(i int, v Keyer) error {
		require.Equal(t, v, keyable(i), "test case [%d] failed", i)
		return nil
	})
}

package cache

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	testCache := NewCache()

	b, err := json.Marshal([]byte("world"))
	require.NoError(t, err)

	err = testCache.Add("hello", b)
	require.NoError(t, err)

	v, err := testCache.Get("hello")
	require.NoError(t, err)
	require.Equal(t, v, b)

	err = testCache.Delete("hello")
	require.NoError(t, err)

}

package cache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	testCache := NewCache()

	b, err := json.Marshal([]byte("world"))
	require.NoError(t, err)

	testCache.Set("hello", b, time.Duration(10)*time.Second)

	v, ok := testCache.Get("hello")
	require.True(t, ok)
	require.Equal(t, v, b)

	err = testCache.Delete("hello")
	require.NoError(t, err)

}

package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nico-phil/dcache/cache"
)

type CacheServer struct {
	cache *cache.Cache
}

func NewCacheServer() *CacheServer {
	return &CacheServer{
		cache: cache.NewCache(),
	}
}

func (s *CacheServer) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string `json:"key"`
		Value []byte `json:"value"`
		TTL   int64  `json:"ttl"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	s.cache.Set(req.Key, req.Value, time.Duration(req.TTL)*time.Second)
}
func (s *CacheServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := s.cache.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string][]byte{"value": value})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (s *CacheServer) Start() error {
	http.HandleFunc("/set", s.SetHandler)
	http.HandleFunc("/get", s.GetHandler)

	return http.ListenAndServe(":8080", nil)
}

package server

import (
	"encoding/json"
	"net/http"

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
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	s.cache.Set(req.Key, req.Value)
}
func (s *CacheServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, ok := s.cache.Get(key)
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]byte{"value": value})
}

func (s *CacheServer) Start() error {
	http.HandleFunc("/set", s.SetHandler)
	http.HandleFunc("/get", s.GetHandler)

	return http.ListenAndServe(":8080", nil)
}

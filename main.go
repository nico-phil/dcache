package main

import "github.com/nico-phil/dcache/server"

func main() {
	s := server.NewCacheServer()
	err := s.Start()
	if err != nil {
		panic(err)
	}
}

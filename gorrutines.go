package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURLr(url string, wg *sync.WaitGroup) {
	defer wg.Done() // marca esta corrutine como completada

	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	duracion := time.Since(start)
	fmt.Printf("Fetched %s en %v - Status %s\n", url, duracion, resp.Status)
}

func mainn() {
	urls := []string{
		"https://rickandmortyapi.com/api/character/1",
		"https://rickandmortyapi.com/api/character/2",
		"https://rickandmortyapi.com/api/character/3",
		"https://rickandmortyapi.com/api/character/4",
	}

	start := time.Now()
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURLr(url, &wg)
	}

	wg.Wait()

	duracionTotal := time.Since(start)
	fmt.Println("tiempo total de la Concurrencia: %v \n", duracionTotal)

}

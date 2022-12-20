package main

import (
	"generateScript/dictionary"
	"generateScript/queries"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			cookie := queries.SignUp(dictionary.GenerateEmployer(i))
			for j := 0; j < 10; j++ {
				queries.CreateVacancy(dictionary.GenerateVacancy(), cookie)
			}
		}(i, &wg)
	}
}

package main

import (
	"generateScript/dictionary"
	"generateScript/queries"
	"math/rand"
	"sync"
	"time"
)

const employerCount = 10
const vacancyByEmployer = 10

const applicantsCount = 10
const resumeByApplicant = 5

const applyByApplicant = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < employerCount; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			cookie := queries.SignUp(dictionary.GenerateEmployer(i))
			for j := 0; j < vacancyByEmployer; j++ {
				queries.CreateVacancy(dictionary.GenerateVacancy(), cookie)
			}
		}(i, &wg)
	}

	for i := 0; i < applicantsCount; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			cookie := queries.SignUp(dictionary.GenerateApplicant(i))
			for j := 0; j < resumeByApplicant; j++ {
				queries.CreateResume(dictionary.GenerateResume(), cookie)
			}
			for k := 0; k < applyByApplicant; k++ {
				queries.Apply(uint(rand.Int()%(vacancyByEmployer*employerCount)+1),
					uint(rand.Int()%(resumeByApplicant*applicantsCount)+1), cookie)
			}
		}(i, &wg)
	}
	wg.Wait()
}

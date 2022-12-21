package main

import (
	"generateScript/dictionary"
	"generateScript/queries"
	"sync"
)

const employerCount = 10
const vacancyByEmployer = 10

const applicantsCount = 10
const resumeByApplicant = 5

const applyByApplicant = 10

func main() {
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
		}(i, &wg)
	}
	wg.Wait()
}

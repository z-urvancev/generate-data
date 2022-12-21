package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
	"time"
)

func GenerateApplicant(i int) *models.UserAccount {
	result := &models.UserAccount{
		Email:            fmt.Sprintf("%d_%s@mail.ru", i, "applicant"),
		Password:         "123456a!",
		ApplicantName:    generateName(),
		ApplicantSurname: generateSurname(),
		UserType:         "applicant",
		Location:         generateLocation(),
		Age:              uint(rand.Int())%34 + 16,
		DateOfBirth:      time.Date(1970+rand.Int()%36, time.April, rand.Int()%365, 0, 0, 0, 0, time.Local),
		ContactNumber:    generatePhone(),
	}
	return result
}

func generatePhone() string {
	a := rand.Int() % 1000
	d := rand.Int() % 1000
	b := rand.Int() % 100
	c := rand.Int() % 100
	return fmt.Sprintf("+7 (%d) %d-%d-%d", d, a, b, c)
}
func generateName() string {
	i := rand.Int() % len(names)
	return names[i]
}

func generateSurname() string {
	i := rand.Int() % len(surnames)
	return surnames[i]
}

var names = []string{
	"Захар",
	"Аким",
	"Владислав",
	"Виктор",
	"Александр",
	"Василий",
	"Иван",
	"Антон",
	"Максим",
	"Егор",
	"Владимир",
	"Дмитрий",
	"Афанасий",
	"Андрей",
	"Лев",
	"Даниил",
	"Никита",
	"Роман",
	"Петр",
}

var surnames = []string{
	"Иванов",
	"Александров",
	"Медведев",
	"Волков",
	"Кирпичёв",
	"Кутинцев",
	"Виноградов",
	"Герасимов",
	"Белоусов",
	"Белобородов",
	"Флоров",
	"Мельников",
	"Епишев",
	"Сорокин",
	"Егоров",
	"Захаров",
	"Ульянов",
	"Роганов",
	"Крупкин",
	"Сафонов",
	"Коновалов",
	"Ильин",
	"Ковальцов",
	"Братчиков",
	"Куликов",
	"Белохлебов",
	"Силкин",
	"Родионов",
	"Белов",
	"Королёв",
}

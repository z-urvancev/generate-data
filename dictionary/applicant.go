package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
)

func GenerateApplicant(i int) *models.UserAccount {
	result := &models.UserAccount{
		Email:            fmt.Sprintf("%d_%s@mail.ru", i, "applicant"),
		Password:         "123456a!",
		ApplicantName:    generateName(),
		ApplicantSurname: generateSurname(),
		UserType:         "applicant",
	}
	return result
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

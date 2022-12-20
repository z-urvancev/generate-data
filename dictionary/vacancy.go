package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
)

func GenerateVacancy() *models.Vacancy {
	result := &models.Vacancy{
		Title:        generateTitle(),
		Salary:       uint(rand.Uint64()%2900+100) * 100,
		Hours:        hours[rand.Int()%len(hours)],
		Experience:   experience[rand.Int()%len(experience)],
		Format:       format[rand.Int()%len(format)],
		Description:  "*****Описание вакансии:*****\n\n\n\n*******************",
		Requirements: "Требования:\n- 1;\n- 2;\n- 3;\n- 4;\n- 5;\n- 6.",
	}
	return result
}

func generateTitle() string {
	iPref := rand.Int() % len(prefixTitle)
	i := rand.Int() % len(mainTitle)
	return fmt.Sprintf("%s%s", prefixTitle[iPref], mainTitle[i])
}

var prefixTitle = []string{
	"Младший ",
	"Старший ",
	"Начинающий ",
	"", "", "", "", "",
}

var mainTitle = []string{
	"Психолог",
	"Golang-разработчик",
	"С++-разработчик",
	"С-разработчик",
	"С#-разработчик",
	"Java-разработчик",
	"Python-разработчик",
	"PHP-разработчик",
	"Perl-разработчик",
	"Тестировщик",
	"Каскадер",
	"Фитнес-тренер",
	"Преподаватель",
	"Воспитатель",
	"Уборщик",
	"Грузчик",
	"Дизайнер",
	"UX-исследователь",
	"Маркетолог",
	"Финансист",
	"Банкир",
	"Переводчик",
	"Менеджер по продажам",
	"Водитель",
	"Дрессировщик",
	"Продавец",
	"Слесарь",
	"Сварщик",
	"Инженер",
}

var experience = []string{
	"нет опыта",
	"1-3 года опыта",
	"3-5 лет опыта",
	"5-10 лет опыта",
}

var format = []string{
	"Удаленный формат",
	"Очный формат",
	"Гибридный формат",
	"Вахтовый метод",
}

var hours = []string{
	"20",
	"24",
	"30",
	"32",
	"40",
}

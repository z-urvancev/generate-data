package dictionary

import (
	"generateScript/models"
	"math/rand"
)

func GenerateResume() *models.Resume {
	result := &models.Resume{
		Title:             generateTitle(),
		ExperienceInYears: experience[rand.Int()%len(experience)],
		Salary:            uint(rand.Uint64()%2900+100) * 100,
		Location:          generateLocation(),
		Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt posuere ex non tempor. Sed massa erat, venenatis id mattis non, mollis ut nunc. Phasellus fringilla accumsan nisl. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut dapibus sapien eget efficitur rutrum. Suspendisse accumsan bibendum ante id gravida. Pellentesque luctus porta magna quis iaculis. Proin posuere a velit quis viverra. Sed arcu nisl, faucibus a lacinia vel, iaculis a arcu. Etiam a leo id sem tincidunt congue eget a odio.",
	}
	return result
}

var resumeDesc = map[string]string{
	"Психолог":             "Имею специализации в психологии: детско-родительские отношения, диагностика и коррекция эмоционального состояния, работа с детьми, подростками, а так семейные отношения: консультирование супружеских пар, диагностика и коррекция эмоционального состояния с помощью кинезиологических техник, а также методики Резет-1, Резет-2.",
	"Golang-разработчик":   "Последние полгода интенсивно изучаю этот язык программирования и работаю над разработкой учебного проекта в рамках Образовательного курса VK в МГТУ. Код Вы можете посмотреть по ссылке: https://github.com/go-park-mail-ru/2022_2_Technopanki . Работаю с Docker, Redis, PostgreSQL, gRPC и активно изучаю новые технологии, так как в настоящий момент проект находится в стадии разработки. Имею опыт работы с С++, писал проект с использованием фреймворка Qt, изучал основные алгоритмы и структуры данных, регулярно работаю с системой контроля версий git.",
	"С++ разработчик":      "Студент, начинающий разработчик. Имею опыт работы (учебный проект) с С++ с использованием Boost и Qt",
	"С-разработчик":        "Студент, начинающий разработчик. Имею опыт работы (учебный проект) с С и Assembler",
	"С#-разработчик":       "Студент, начинающий разработчик. Имею опыт работы (учебный проект) с .NET",
	"Java-разработчик":     "2 года разрабатывал приложения с использованиям фреймворка Spring. Полгода опыта коммерческой разработке на Kotlin",
	"Python-разработчик":   "Выпускник образовательного проекта VK в МГТУ направления ML",
	"PHP-разработчик":      "Студент, начинающий разработчик. Имею опыт работы (учебный проект) с Laravel",
	"Perl-разработчик":     "Студент, начинающий разработчик. Имею опыт работы (учебный проект) с Perl",
	"Тестировщик":          "",
	"Каскадер":             "",
	"Фитнес-тренер":        "",
	"Преподаватель":        "",
	"Воспитатель":          "",
	"Уборщик":              "",
	"Грузчик":              "",
	"Дизайнер":             "",
	"UX-исследователь":     "",
	"Маркетолог":           "",
	"Финансист":            "",
	"Банкир":               "",
	"Переводчик":           "",
	"Менеджер по продажам": "",
	"Водитель":             "",
	"Дрессировщик":         "",
	"Продавец":             "",
	"Слесарь":              "",
	"Сварщик":              "",
	"Инженер":              "",
}

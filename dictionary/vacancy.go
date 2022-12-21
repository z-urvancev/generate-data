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
		Location:     generateLocation(),
		Tasks:        "Proin consectetur mi diam, eu semper sapien volutpat quis. Morbi ut facilisis tellus. Vestibulum porta maximus aliquet. Maecenas cursus, purus et rhoncus mattis, felis felis iaculis lorem, in sagittis augue lacus in dolor. Vestibulum ullamcorper risus lacus, quis efficitur velit fringilla eu. Mauris id nibh quis tellus aliquet mattis. Sed quis lectus quis leo vestibulum mattis. Integer nec magna purus. Pellentesque id elit felis. Sed ac tortor ante. Suspendisse ut urna eu nisl eleifend mattis id placerat ex. Praesent dolor neque, facilisis sed urna et, interdum consectetur nibh.",
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt posuere ex non tempor. Sed massa erat, venenatis id mattis non, mollis ut nunc. Phasellus fringilla accumsan nisl. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut dapibus sapien eget efficitur rutrum. Suspendisse accumsan bibendum ante id gravida. Pellentesque luctus porta magna quis iaculis. Proin posuere a velit quis viverra. Sed arcu nisl, faucibus a lacinia vel, iaculis a arcu. Etiam a leo id sem tincidunt congue eget a odio.",
		Requirements: "Phasellus facilisis convallis suscipit. Proin commodo sem id eros semper aliquam. Vivamus vulputate, libero non lobortis aliquam, erat turpis vulputate mi, et vulputate mi nisl ut nisi. Pellentesque ac felis at ligula vehicula hendrerit pretium ac mi. Pellentesque vel lobortis ipsum. Proin ac ipsum sollicitudin, pulvinar lorem vel, pulvinar lacus. Pellentesque imperdiet sapien eget diam mollis molestie. Vivamus dolor est, rutrum ac orci ut, dapibus tempor sem.",
		Extra:        "Sed eget egestas sapien. Maecenas venenatis in dolor at porta. Aenean eu mollis turpis. Nulla facilisi. Quisque tincidunt justo ut fermentum porta. Pellentesque commodo lectus volutpat tristique placerat. Phasellus ut elementum sem, ut sodales nulla. Praesent vestibulum ipsum ex, at rutrum metus pulvinar et. Sed hendrerit risus sapien, sit amet pharetra leo sodales sit amet. Aenean eu vehicula quam. Maecenas ipsum orci, rutrum sed nisi in, accumsan tempor lacus.",
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
	"Нет опыта",
	"Менее 1 года",
	"От 1 года до 3 лет",
	"От 3 до 6 лет",
	"Более 6 лет",
}

var format = []string{
	"Полный день",
	"Смешанный формат",
	"Удаленная работа",
	"Гибкий график",
	"Сменный график",
}

var hours = []string{
	"20",
	"24",
	"30",
	"32",
	"40",
}

package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
	"strings"
)

func GenerateEmployer(i int) *models.UserAccount {
	result := &models.UserAccount{
		Email:        fmt.Sprintf("%d_%s@mail.ru", i, "employer"),
		Password:     "123456a!",
		CompanyName:  generateCompanyName(),
		UserType:     "employer",
		Location:     generateLocation(),
		CompanySize:  uint(rand.Int()%100*100 + 100),
		BusinessType: businessTypes[rand.Int()%len(businessTypes)],
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt posuere ex non tempor. Sed massa erat, venenatis id mattis non, mollis ut nunc. Phasellus fringilla accumsan nisl. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut dapibus sapien eget efficitur rutrum. Suspendisse accumsan bibendum ante id gravida. Pellentesque luctus porta magna quis iaculis. Proin posuere a velit quis viverra. Sed arcu nisl, faucibus a lacinia vel, iaculis a arcu. Etiam a leo id sem tincidunt congue eget a odio.",
	}
	return result
}

func generateCompanyName() string {
	iF := rand.Int() % len(companyFirstWord)
	iS := rand.Int() % len(companySecondWord)
	iPref := rand.Int() % len(prefixs)
	iPost := rand.Int() % len(postfix)
	result := strings.Join([]string{companyFirstWord[iF], prefixs[iPref], companySecondWord[iS], postfix[iPost]}, "")
	return result
}

var businessTypes = []string{
	"IT",
	"B2B",
	"Банк",
	"Сельхоз",
	"Промышленность",
	"Образование",
	"Сфера услуг",
	"Питание",
	"Развлечения",
	"Спорт",
}

var companyFirstWord = []string{
	"ЗАО ",
	"ОАО ",
	"ООО ",
	"Содружество ",
	"Объединение ",
	"Землячество ",
	"",
}

var prefixs = []string{
	"Рос",
	"Транс",
	"Аэро",
	"",
}
var companySecondWord = []string{
	"Прогресс",
	"Тех",
	"Зенит",
	"Сельхоз",
	"Экспресс",
	"Флот",
	"Космос",
	"Свобода",
	"Победа",
}

var postfix = []string{
	"Банк",
	" INC",
	"Пром",
	"",
	"",
	"",
}

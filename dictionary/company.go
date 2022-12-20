package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
	"strings"
)

func GenerateEmployer(i int) *models.UserAccount {
	result := &models.UserAccount{
		Email:       fmt.Sprintf("%d_%s@mail.ru", i, "employer"),
		Password:    "123456a!",
		CompanyName: generateCompanyName(),
		UserType:    "employer",
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

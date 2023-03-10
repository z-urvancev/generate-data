package dictionary

import (
	"fmt"
	"generateScript/models"
	"math/rand"
)

func GenerateVacancy() *models.Vacancy {
	title, main := generateTitle()
	result := &models.Vacancy{
		Title:        title,
		Salary:       uint(rand.Uint64()%2900+100) * 100,
		Hours:        hours[rand.Int()%len(hours)],
		Experience:   experience[rand.Int()%len(experience)],
		Format:       format[rand.Int()%len(format)],
		Location:     generateLocation(),
		Tasks:        vacancyTasks[main],
		Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt posuere ex non tempor. Sed massa erat, venenatis id mattis non, mollis ut nunc. Phasellus fringilla accumsan nisl. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut dapibus sapien eget efficitur rutrum. Suspendisse accumsan bibendum ante id gravida. Pellentesque luctus porta magna quis iaculis. Proin posuere a velit quis viverra. Sed arcu nisl, faucibus a lacinia vel, iaculis a arcu. Etiam a leo id sem tincidunt congue eget a odio.",
		Requirements: vacancyRequirements[main],
		Extra:        vacancyExtra[main],
	}
	return result
}

func generateTitle() (string, string) {
	iPref := rand.Int() % len(prefixTitle)
	i := rand.Int() % len(mainTitle)
	return fmt.Sprintf("%s%s", prefixTitle[iPref], mainTitle[i]), mainTitle[i]
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

var vacancyTasks = map[string]string{
	"Психолог":             "Проведение первичного интервью с кандидатами, тестирование.\nСоздание командного духа внутри коллектива.\n Составление отчетности по выполненной работе.\n Помощь в адаптации, обучении, подготовке кадров",
	"Golang-разработчик":   "Разработка новых сервисов на языке Golang, поддержка существующих, написание тестов и метрик",
	"С++ разработчик":      "Разработка десктоп приложения на C++ с использованием платформы Qt",
	"С-разработчик":        "Разработка проектов компании на языке C; поддержка и развитие существующих программных продуктов.",
	"C#-разработчик":       "Разработка новых сервисов на языке C#, поддержка существующих, написание тестов и метрик",
	"Java-разработчик":     "Разработка новых сервисов на языке Java, поддержка существующих, написание тестов и метрик",
	"Python-разработчик":   "Определять перечень доработок платформы и участвовать в их реализации; Расширять возможности платформы в областях DS-вычислений и глубокого обучения.",
	"PHP-разработчик":      "Разработка кабинета поставщика с нуля;\n\n Создание блока для работы с нашим наличием для продажи через сайт;\n\n Написание аналитических отчетов, автоматизация для помощи в принятие решений нашими менеджерам;\n\n Перенос административной части сайта на новый движок админки AppStack;\n\n Написание тестов",
	"Perl-разработчик":     "Разработка новых сервисов на языке Perl, поддержка существующих, написание тестов и метрик",
	"Тестировщик":          "Разработка новых генераторов тестовых данных на интеграционных полигонах (около 100 цепочек генерации данных и связанных автоматизированных систем) и интеграция их в корпоративную систему создания тестовых данными, \nПоддержка и доработка существующих генераторов,\n Развитие корпоративной системы генерации данных",
	"Фитнес-тренер":        "Проводить индивидуальные занятия с клиентами фитнес-клуба",
	"Преподаватель":        "Транслирование программ КиберШколы по методическим пособиям и передача практического опыта для школьников от 6 до 14 лет. \n\n- Участие в промо-мероприятиях и конференциях;",
	"Воспитатель":          "Работа с детьми 3-6 лет",
	"Уборщик":              "Уборка концертных залов после мероприятий",
	"Грузчик":              "Осуществление погрузочно-разгрузочных работ;\n Размещение грузов на складе;\n Участие в инвентаризации.",
	"Финансист":            "Участие в осуществлении контроля за финансовыми потоками компании и её подразделени; Разработка, обсуждение и утверждение инвестиционных проектов. Написание финансовых моделей; Составление и оценка планов и прогнозов, затрагивающих состояние финансов",
	"Дизайнер":             "Консультации заказчиков по ассортименту тканей и услугам (пошив) в магазине и на выезде,\n Разработка дизайн-проектов, расчет расхода материалов, стоимости материалов и услуг,\n Презентация дизайн-проекта клиенту и ведение переговоров по заказу.",
	"Переводчик":           "Встречать иностранных партренов; участвовать в переговорах",
	"Менеджер по продажам": "Консультация клиентов по входящим звонкам;\n Заключение договоров; \nСопровождение на всех этапах сделки; \nПоиск и привлечение новых клиентов; \nПроведение презентации техники;\n Ведение документооборота установленного порядка",
	"Водитель":             "Возить представителей компании на деловые встречи",
	"Дрессировщик":         "Дрессировка тигров",
	"Продавец":             "Работа за кассовым аппаратом",
	"Сварщик":              "Резка металла, монтаж деталей, сваривание разнообразных металлических конструкций, соблюдение техники безопасности, содержание сварочного оборудования в чистоте",
	"Инженер":              "Разработка чертежей и конструкторской документации в программе SolidWorks,\n Принятие участия в создании эксклюзивных дизайн-проектов лифтов,\nАвторский надзор",
}

var vacancyRequirements = map[string]string{
	"Психолог":             "Нестандартное мышление, креативность\nОтветственность, инициативность\nУмение работать с большим объемом информации\nОпыт приветствуется, но рассматриваем и без опыта, но с большим желанием рости и развиваться",
	"Golang-разработчик":   "Уверенное знание языка golang; \nУверенное знание SQL; \nУмение работать с git; \nОпыт разработки коммерческих проектов.\nСпособность разбираться в чужом коде;\n \nСпособность работать в команде с другими разработчиками.",
	"С++ разработчик":      "Высшее образование, \nЗнание языка С++ и библиотек STL, \nЗнание и понимание принципов ООП и паттернов проектирования, \nХорошая алгоритмическая подготовка: знание основных структур данных и алгоритмов.",
	"С-разработчик":        "Нашей команде работают талантливые математики и успешные разработчики. Чтобы стать частью нашей команды тебе потребуется: \n\nналичие высшего технического образования или обучение на выпускном курсе; \nопыт разработки на С; \nLinux на уровне пользователя; \nбазовое владение Git.",
	"C#-разработчик":       "Уверенное знание языка C#; \nУверенное знание SQL; \nУмение работать с git; \nОпыт разработки коммерческих проектов.\nСпособность разбираться в чужом коде;\n \nСпособность работать в команде с другими разработчиками.",
	"Java-разработчик":     "Уверенное знание языка Java; \nУверенное знание SQL; \nУмение работать с git; \nОпыт разработки коммерческих проектов.\nСпособность разбираться в чужом коде;\n \nСпособность работать в команде с другими разработчиками.",
	"Python-разработчик":   "Хорошее знание Python (стандартные и ML библиотеки, умение применять ООП); Понимание принципов построения моделей машинного обучения",
	"PHP-разработчик":      "Разработка back-end и front-end частей системы;\n участвовать в fullstack-разработке продуктовых возможностей платформы (фичи для магазинов, админка, отчеты и т.п.);\n развивать инфраструктуру платформы (сервисы, апи, взаимодействие с внешними системами);\n Code-review;\n Поддержка и развитие API;\n Написание unit-тестов;\n Доработка API и базы данных на стороне back-end",
	"Perl-разработчик":     "Уверенное знание языка Perl; \nУверенное знание SQL; \nУмение работать с git; \nОпыт разработки коммерческих проектов",
	"Тестировщик":          "знание теории тестирования и принципов интеграции различных систем, протоколов обмена данными (MQ, Kafka, сериализация, json, xml)",
	"Фитнес-тренер":        "Профильное образование (Высшее, среднее или доп. образование в формате курсов)\n Теоретические знания в области фитнеса: биомеханика, анатомия, физиология, методики тренировок;\n\n Опыт работы на должности тренера от 1 года;\n Навыки работы с базой клиентов;",
	"Преподаватель":        "Знание базовых основ компьютерной грамотности;\n\n- Знание алгоритмизации;\n\n- Знание основ языков программирования;\n\n- Умение разбираться с методикой и платформами в соответствии с инструкцией;\n\n- Знание одного или нескольких языков программирования\n\n- Высшее или неоконченное высшее образование в сфере IT\n\n- Хорошие знания в одном из языков программирования и программ: HTML, JavaScript, Python, Графические редакторы, Платформа Unity3D + С#; \n\n Обучаемость, желание получать новые знания",
	"Воспитатель":          "Педагогичесое образование",
	"Уборщик":              "Готовность к командировкам",
	"Грузчик":              "Крепкое телосложение, ответсвенность",
	"Дизайнер":             "Образование профильное (средне-специальное, высшее); \nОпыт работы дизайнером по шторам, декоратором от 1 года, опыт продаж приветствуется",
	"Финансист":            "Опыт работы в сфере бюджетирования и финансов от 5-х лет;\n\n Знание основ бухгалтерского учета;\n\n Продвинутый пользователь 1С 8.3 и УНФ;\n\n Опыт составление управленческой отчетности от 3-х лет;",
	"Переводчик":           "Знание английского на уровне С1-2, а так же одного из следующих языков на уровне С1-2: немецкий, арабский, китайский",
	"Менеджер по продажам": "Обязательно знание автомобилей (комплектаций, технических характеристик) приоритетно немецких марок/японских марок; \n Умение работать с клиентами из дорогого сегмента, Лицами принимающими решения, юридическими лицами.\n Опыт корпоративных продаж продаж Мерседес, БМВ, Ауди, других премиальных салонов",
	"Водитель":             "Водительский стаж от 5 лет",
	"Дрессировщик":         "3 года в данной сфере",
	"Продавец":             "Внимательность, усидчивость",
	"Сварщик":              "Уметь читать чертежи",
	"Инженер":              "Высшее образование (техническое); Знание SolidWorks; Знание ЕСКД; ПК - уверенный пользователь",
}

var vacancyExtra = map[string]string{
	"Психолог":             "Педагогическое образование",
	"Golang-разработчик":   "Опыт работы с Linux, bash, Docker;\n Опыт работы с MySQL (5.5-8.0);\n Знание основ HTML, CSS, JS;\n Знание основ языка php;\n Опыт работы с сервисами обмена сообщениями.",
	"С++ разработчик":      "Высшее техническое образование в области радиотехники/радиоэлектроники (МГТУ, МФТИ, ТУСУР, РГРТУ, ТГРТУ, МИФИ и т.п.), \nОпыт работы в среде MS Visual Studio, \nОпыт работы с системой автоматизации сборки CMake",
	"С-разработчик":        "Опыт работы с GitLab, Redmine или Jenkins.",
	"С#-разработчик":       "Опыт работы с .NET, k8s/Docker",
	"Java-разработчик":     "Опыт работы с Git, Docker, Tarantool",
	"Python-разработчик":   "Опыт работы со Spark и в целом с экосистемой Hadoop; PyTorch",
	"PHP-разработчик":      "Опыт работы с .NET, k8s/Docker",
	"Perl-разработчик":     "Опыт работы с .NET, k8s/Docker",
	"Тестировщик":          "Опыт работы в банковской сфере и с Synteta",
	"Фитнес-тренер":        "Наличие титулов мастер-спорта или КМС по одному из направлений: тяжелая атлетика, гимнастика, плавание, бокс",
	"Преподаватель":        "Педагогическое образование",
	"Воспитатель":          "",
	"Уборщик":              "Опыт работы в данной сфере",
	"Грузчик":              "Опыт работы в данной сфере",
	"Дизайнер":             "Наличие портфолио (рисунки от руки, эскизы, графика, фотографии)",
	"Финансист":            "",
	"Переводчик":           "Проживание в других странах в прошлом/настоящем",
	"Менеджер по продажам": "",
	"Водитель":             "Опыт работы в данной сфере",
	"Дрессировщик":         "Ветеринарское образование",
	"Продавец":             "Опыт работы за кассовым аппаратом",
	"Сварщик":              "",
	"Инженер":              "",
}

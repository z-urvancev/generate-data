package models

import "time"

type Vacancy struct {
	ID                uint              `json:"id" gorm:"primaryKey;"`
	PostedByUserId    uint              `json:"postedByUserId" gorm:"not null;"`
	Title             string            `json:"title" gorm:"not null;"`
	Description       string            `json:"description,omitempty"`
	Tasks             string            `json:"tasks,omitempty"`
	Requirements      string            `json:"requirements,omitempty"`
	Extra             string            `json:"extra,omitempty"`
	CreatedDate       time.Time         `json:"createdDate" gorm:"autoCreateTime"`
	Salary            uint              `json:"salary,omitempty"`
	Location          string            `json:"location,omitempty"`
	IsActive          bool              `json:"isActive,omitempty"`
	Experience        string            `json:"experience,omitempty"`
	Format            string            `json:"format,omitempty"`
	Hours             string            `json:"hours,omitempty"`
	Image             string            `json:"image,omitempty"`
	VacancyActivities []VacancyActivity `json:"vacancyActivities" gorm:"foreignKey:VacancyId;constraint:OnDelete:CASCADE;"`
	Skills            []Skill           `json:"skills" gorm:"many2many:vacancy_skills;"`
}

type GetAllVacanciesResponcePointer struct {
	Data []*Vacancy `json:"data"`
}

type VacancyFilter struct {
	Title             string
	Location          string
	Format            string
	Experience        string
	FirstSalaryValue  string
	SecondSalaryValue string
}

type VacancyPreview struct {
	Id             uint   `json:"id"`
	PostedByUserId uint   `json:"postedByUserId"`
	CompanyName    string `json:"company_name"`
	Title          string `json:"title"`
	Image          string `json:"image"`
	Salary         uint   `json:"salary"`
	Location       string `json:"location"`
	Format         string `json:"format"`
	Hours          string `json:"hours"`
	Description    string `json:"description"`
}

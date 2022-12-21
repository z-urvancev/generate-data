package dictionary

import (
	"generateScript/models"
	"math/rand"
)

func GenerateResume() *models.Resume {
	result := &models.Resume{
		Title:       generateTitle(),
		Salary:      uint(rand.Uint64()%2900+100) * 100,
		Location:    generateLocation(),
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque tincidunt posuere ex non tempor. Sed massa erat, venenatis id mattis non, mollis ut nunc. Phasellus fringilla accumsan nisl. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Ut dapibus sapien eget efficitur rutrum. Suspendisse accumsan bibendum ante id gravida. Pellentesque luctus porta magna quis iaculis. Proin posuere a velit quis viverra. Sed arcu nisl, faucibus a lacinia vel, iaculis a arcu. Etiam a leo id sem tincidunt congue eget a odio.",
	}
	return result
}

package category

import "suraj_projects/flipkart_machine_coding/utils"

type Category struct {
	CategoryId   string
	CategoryName string
	CategoryType string
	MaxCap       int
}

var CategoryInstance *Category

func NewCategory(categoryName string) *Category {
	if CategoryInstance == nil {
		categoryType := ""
		maxCap := 0
		if utils.StringInsideSlice(categoryName, []string{"Mobiles", "Laptops"}) {
			categoryType = "flatRate"
			maxCap = 5
		} else {
			categoryType = "percentage"
			maxCap = 50

		}

		CategoryInstance = &Category{
			CategoryName: categoryName,
			CategoryType: categoryType,
			MaxCap:       maxCap,
		}
	}
	return CategoryInstance
}

func GetCategory() Category {
	return *CategoryInstance

}

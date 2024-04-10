package categories

import "florist-gin/business/categories"

type Category struct {
	Id   int `gorm:"primaryKey;unique;autoIncrement:true"`
	Name string
}

func (category Category) ToUseCase() categories.Category {
	return categories.Category{
		Id:   category.Id,
		Name: category.Name,
	}
}

func FromUsecase(category categories.Category) Category {
	return Category{
		Id:   category.Id,
		Name: category.Name,
	}
}

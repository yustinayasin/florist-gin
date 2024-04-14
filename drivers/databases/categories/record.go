package categories

import "florist-gin/business/categories"

type Category struct {
	Id   uint32 `gorm:"primaryKey;unique"`
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

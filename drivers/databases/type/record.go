package types

import "florist-gin/business/types"

type Type struct {
	Id   uint32 `gorm:"primaryKey;unique"`
	Name string
}

func (category Type) ToUseCase() types.Type {
	return types.Type{
		Id:   category.Id,
		Name: category.Name,
	}
}

func FromUsecase(category types.Type) Type {
	return Type{
		Id:   category.Id,
		Name: category.Name,
	}
}

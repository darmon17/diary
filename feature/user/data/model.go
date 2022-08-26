package data

import (
	"dairy/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func (data *User) ToModel() domain.User {
	return domain.User{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
func ParseToArray(arr []User) []domain.User {
	var res []domain.User
	for _, val := range arr {
		res = append(res, val.ToModel())
	}
	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Name = data.Name
	res.Email = data.Email
	res.Password = data.Password

	return res
}

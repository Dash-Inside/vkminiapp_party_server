package gameuser

import "gorm.io/gorm"

type repository struct {
	database *gorm.DB
}

type CreateUserParams struct {
	nickname string
}

type ReadUserParams struct {
	uuid string
}

type UpdateUserParams struct {
	uuid     string
	nickname string
}

type DeleteUserParams struct {
	uuid string
}

type Repository interface {
	Create(CreateUserParams) error
	Read(ReadUserParams) error
	Update(UpdateUserParams) error
	Delete(DeleteUserParams) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		database: db,
	}
}

func (*repository) Create(CreateUserParams) error {
	panic("unimplemented")
}

func (*repository) Delete(DeleteUserParams) error {
	panic("unimplemented")
}

func (*repository) Read(ReadUserParams) error {
	panic("unimplemented")
}

func (*repository) Update(UpdateUserParams) error {
	panic("unimplemented")
}

package gameuser

import "gorm.io/gorm"

func ProvideUserRepository(db *gorm.DB) Repository {
	return NewRepository(db)
}

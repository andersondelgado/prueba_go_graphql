package repository

import (
	"github.com/andersondelgado/prueba_go_graphql/pkg/datasources/mysql"
	"github.com/andersondelgado/prueba_go_graphql/pkg/graphql/global"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) Paginate(condition string, where global.PaginationSimpleParams) ([]*model.User, error) {
	//panic("implement me")
	var entity []*model.User
	if err := u.db.Debug().Offset(where.Offset).Limit(where.Limit).Order(where.OrderBy).Where(condition, where.Filter).Find(&entity).Error
		err != nil {
		return nil, err
	}
	return entity, nil
}

func (u UserRepository) CheckByParam(where interface{}) (bool, error) {
	var count int64
	var entity model.User
	isValue := false
	if err := u.db.Debug().Where(where).Find(&entity).Count(&count).Error; err != nil {
		return false, err
	}

	if count != 0 {
		isValue = true
	}
	return isValue, nil
}

func (u UserRepository) Create(entity model.User) (*model.User, error) {
	if err := u.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (u UserRepository) GetByParam(where interface{}) (*model.User, error) {
	var user model.User

	result := u.db.Debug().Where(where).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

type IUserRepository interface {
	CheckByParam(where interface{}) (bool, error)
	Create(entity model.User) (*model.User, error)
	GetByParam(where interface{}) (*model.User, error)
	Paginate(condition string, where global.PaginationSimpleParams) ([]*model.User, error)
}

func NewUserRepository() IUserRepository {
	return UserRepository{
		db: mysql.Db,
	}
}

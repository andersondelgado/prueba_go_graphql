package service

import (
	"encoding/json"
	"fmt"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/dto"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/model"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/repository"
	"github.com/andersondelgado/prueba_go_graphql/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserService struct {
	repo repository.IUserRepository
}

func (u UserService) GetUser(where interface{}) (*model.User, error) {
	entity, err := u.repo.GetByParam(where)
	return entity, err
}

func (u UserService) Login(input dto.InputCredential) (string, error) {
	val, _ := u.repo.CheckByParam(map[string]interface{}{"username": input.Username})

	var isValuefinal bool

	var entity model.User
	if val {
		isValuefinal = val
		entity0, _ := u.repo.GetByParam(map[string]interface{}{"username": input.Username})

		str0, _ := json.Marshal(entity0)
		b0 := []byte(str0)
		json.Unmarshal(b0, &entity)

	}

	if !isValuefinal {
		panic("implement me")
	}

	fmt.Println("entity.Password: ", entity.Password)
	byteP := []byte(input.Password)
	if !util.ComparePasswords(entity.Password, byteP) {
		//panic("implement me")
		return "", fmt.Errorf("Error en el password")
	}

	claims := util.CustomClaims{
		ID: entity.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	var (
		response dto.LoginResponse
	)

	token, err := util.EncodeToken(claims)

	response.Token = token

	return response.Token, err
}

func (u UserService) Register(input dto.InputUser) (string, error) {
	var usr model.User
	usr.Username = input.Username
	byteP := []byte(input.Password)
	p := util.HashAndSalt(byteP)
	usr.Password = p
	entity, _ := u.repo.Create(usr)
	claims := util.CustomClaims{
		ID: entity.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	var (
		response dto.LoginResponse
	)

	token, err := util.EncodeToken(claims)

	response.Token = token

	return response.Token, err
}

type IUserService interface {
	Login(input dto.InputCredential) (string, error)
	Register(input dto.InputUser) (string, error)
	GetUser(where interface{}) (*model.User, error)
}

func newUserService() IUserService {
	return UserService{
		repo: repository.NewUserRepository(),
	}
}

//---------------------------------------------------------------
type AuthService struct {
	UserService IUserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserService: newUserService(),
	}
}

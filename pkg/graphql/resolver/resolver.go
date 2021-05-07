package resolver

import "github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/service"

type Resolver struct {
	AuthService *service.AuthService
}

func NewResolver() *Resolver {
	return &Resolver{
		AuthService: service.NewAuthService(),
	}
}

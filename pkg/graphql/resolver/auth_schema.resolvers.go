package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"github.com/andersondelgado/prueba_go_graphql/pkg/enum"
	graphql1 "github.com/andersondelgado/prueba_go_graphql/pkg/graphql"
	"github.com/andersondelgado/prueba_go_graphql/pkg/graphql/global"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/dto"
	"github.com/andersondelgado/prueba_go_graphql/pkg/modules/auth/model"
)

func (r *mutationResolver) Login(ctx context.Context, input dto.InputCredential) (string, error) {
	//panic(fmt.Errorf("not implemented"))
	return r.AuthService.UserService.Login(input)
}

func (r *mutationResolver) Register(ctx context.Context, input dto.InputUser) (string, error) {
	//panic(fmt.Errorf("not implemented"))
	return r.AuthService.UserService.Register(input)
}

func (r *queryResolver) GetAuthUser(ctx context.Context) (*graphql1.User, error) {
	//panic(fmt.Errorf("not implemented"))
	usr := ctx.Value(string(enum.GinContextKeyAuthDefault)).(model.User)
	//fmt.Println("user id: ", usr.ID)
	entity, err := r.AuthService.UserService.GetUser(map[string]interface{}{"id": usr.ID})
	//fmt.Println(entity.ID)
	var response *graphql1.User
	str, _ := json.Marshal(entity)
	json.Unmarshal(str, &response)

	return response, err
}

func (r *queryResolver) GetUserByID(ctx context.Context, id int) (*graphql1.User, error) {
	entity, err := r.AuthService.UserService.GetUser(map[string]interface{}{"id": id})
	//fmt.Println(entity.ID)
	var response *graphql1.User
	str, _ := json.Marshal(entity)
	json.Unmarshal(str, &response)

	return response, err
}

func (r *queryResolver) GetUsers(ctx context.Context, filter global.PaginationSimpleParams) ([]*graphql1.User, error) {
	filter.Filter = "%" + filter.Filter + "%"
	entity, error := r.AuthService.UserService.GetUsers("username LIKE ?", filter)
	var response []*graphql1.User
	str, _ := json.Marshal(entity)
	json.Unmarshal(str, &response)
	return response, error
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

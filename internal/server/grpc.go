package server

import (
	"context"
	"user/internal/db"
	"user/pkg/api"
)

type noteServer struct {
	user db.User
	api.UnimplementedNotesServer
}

func NewService(topic db.User) *noteServer {
	return &noteServer{
		user: topic,
	}
}

// TODO: Добавить во все сервисы с gRPC нормальную обработку ошибок
// Вот такую return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")

func (n *noteServer) Create(ctx context.Context, in *api.CreateRequest) (*api.CreateResponse, error) {
	resp, err := n.user.Create(ctx, db.CreateNoteRequest{
		Name:  in.Name,
		Pass:  in.Pass,
		Email: in.Email,
	})
	return &api.CreateResponse{
		Message: resp.Message,
	}, err
}
func (n *noteServer) Validate(ctx context.Context, in *api.ValidateRequest) (*api.ValidateResponse, error) {
	resp, err := n.user.Validate(ctx, db.ValidateRequest{
		Login: in.Login,
		Pass:  in.Pass,
	})

	return &api.ValidateResponse{
		Token:   resp.Token,
		Message: resp.Message,
	}, err
}
func (n *noteServer) Update(ctx context.Context, in *api.UpdateRequest) (*api.UpdateResponse, error) {
	resp, err := n.user.Update(ctx, db.UpdateRequest{
		Token:    in.Token,
		EmailOld: in.EmailOld,
		EmailNew: in.EmailNew,
		NameOld:  in.NameOld,
		NameNew:  in.NameNew,
		PassOld:  in.PassOld,
		PassNew:  in.PassNew,
	})
	return &api.UpdateResponse{
		Message: resp.Message,
	}, err
}

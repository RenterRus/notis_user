package db

import "context"

// TODO: Переписать

type CreateNoteRequest struct {
	Name  string
	Email string
	Pass  string
}

type CreateNoteResponse struct {
	Message string
}

type ValidateRequest struct {
	Login string
	Pass  string
}

type ValidateResponse struct {
	Token   string
	Message string
}

type UpdateRequest struct {
	Token    string
	NameOld  string
	NameNew  string
	EmailOld string
	EmailNew string
	PassOld  string
	PassNew  string
}

type UpdateResponse struct {
	Message string
}

type User interface {
	Create(ctx context.Context, notes CreateNoteRequest) (CreateNoteResponse, error)
	Update(ctx context.Context, req UpdateRequest) (UpdateResponse, error)
	Validate(ctx context.Context, req ValidateRequest) (ValidateResponse, error)
}

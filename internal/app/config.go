package app

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

type GRPC struct {
	Addr string `validate:"required,hostname_port"`
}

type Postgresql struct {
	Addr     string `validate:"required,hostname_port"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Database string `validate:"required"`
}

type config struct {
	GRPC   GRPC
	DB     Postgresql
	LogLvl int `validate:"min=-1,max=5"`
}

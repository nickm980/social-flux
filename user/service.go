package users

import "context"

type Service interface {
	CreateUser(ctx context.Context)
}

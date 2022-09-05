// Request and response objects for users and authentication
package users

type (
	UserResponse struct {
		Name string `json:"name"`
	}

	CreateUserRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
	}
)

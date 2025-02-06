package models

type UserRequest struct {
	UserID string `params:"user_id" validate:"required,alphanum"`
}

type UserResponse struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
}

func (req *UserRequest) Validate() error {

	// Add your validation logic here

	return nil
}

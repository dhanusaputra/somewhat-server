package v1

type createSomethingRequest struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"max=50"`
	Description string `json:"description" validate:"max=150"`
}

type updateSomethingRequest struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"max=50"`
	Description string `json:"description" validate:"max=150"`
}

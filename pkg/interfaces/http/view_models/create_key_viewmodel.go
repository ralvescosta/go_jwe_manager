package viewmodels

type CreateKeyViewModel struct {
	Id string `json:"id" validate:"required, uuidv4"`
}

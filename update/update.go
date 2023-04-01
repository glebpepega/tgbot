package update

type UpdateResponse struct {
	Result []Update
}

type Update struct {
	Update_id int
}

func NewResponse() *UpdateResponse {
	return &UpdateResponse{}
}

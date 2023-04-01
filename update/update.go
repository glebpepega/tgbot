package update

type UpdateResponse struct {
	Result []Update
}

type Update struct {
	Update_id int
	Message   Message
}

type Message struct {
	Chat     Chat
	Text     string
	Entities []Entity
}

type Chat struct {
	Id int
}

type Entity struct {
	Type string
}

func NewResponse() *UpdateResponse {
	return &UpdateResponse{}
}

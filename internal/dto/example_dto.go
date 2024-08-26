package dto

type ExampleRequest struct {
	Id int `params:"id"`
}

type ExampleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

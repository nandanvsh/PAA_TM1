package model

type Food struct {
	ID          int
	User_ID     int
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
}

type GetFood struct {
	ID          int
	Name        string
	Price       int
	Description string
}

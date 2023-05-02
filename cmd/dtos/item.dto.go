package dtos

type AddItemDTO struct {
	Name   string `json:"name" bson:"name"`
	Price  int    `json:"price" bson:"price"`
	ImgUrl string `json:"imgUrl" bson:"imgUrl"`
}

type UpdateItemDTO struct {
	ID     string `json:"_id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Price  int    `json:"price" bson:"price"`
	ImgUrl string `json:"imgUrl" bson:"imgUrl"`
}

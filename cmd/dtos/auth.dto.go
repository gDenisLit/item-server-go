package dtos

type LoginDTO struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type SignupDTO struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Fullname string `json:"fullname" bson:"fullname"`
	ImgUrl   string `json:"imgUrl" bson:"imgUrl"`
	IsAdmin  bool   `json:"isAdmin" bson:"isAdmin"`
}

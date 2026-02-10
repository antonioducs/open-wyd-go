package models

type User struct {
	ID        string `json:"id" dynamodbav:"id"`
	Username  string `json:"username" dynamodbav:"username"`
	Password  string `json:"password" dynamodbav:"password"`
	CreatedAt int64  `json:"created_at" dynamodbav:"created_at"`
}

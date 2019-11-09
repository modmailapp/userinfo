package main

type UserNotFoundResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Alive struct {
	Alive bool `json:"alive"`
}

type UserResponse struct {
	Tag  string `json:"tag"`
	User User   `json:"user"`
	Id   string `json:"id"`
	Bot  bool   `json:"bot"`
}

type User struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
}

type Response struct {
	Success bool         `json:"success"`
	Message UserResponse `json:"message"`
}

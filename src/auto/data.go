package auto

import "src/blogos/src/api/models"

var users = []models.User{
	{Nickname: "Joseph Daniel", Email: "jdanielillo@mail.com", Password: "1234567"},
	{Nickname: "Juan Alvarado", Email: "juanxo@mail.com", Password: "1133557799"},
}

var posts = []models.Post{
	{
		Title:   "Title",
		Content: "Hello world",
	},
	{
		Title:   "Title 2",
		Content: "Hello 2 world 2",
	},
}

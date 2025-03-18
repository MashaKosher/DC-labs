package main

import (
	db "laba3/config"
	routes "laba3/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName: "Laba3",
		})
	db.Connect()
	defer db.Session.Close()
	routes.Handlers(app)

	// Выполняем простой запрос
	// var country string
	// var storyId int64
	// var id int64
	// var content string

	// var Res models.User = models.User{}

	// // Запрос данных из таблицы
	// iter := db.Session.Query("SELECT country, storyId, id FROM tbl_message").Iter()
	// for iter.Scan(&Res.Country, &Res.StoryID, &Res.StoryID, &Res.ID) {
	// 	fmt.Println(Res)
	// }

	// if err := iter.Close(); err != nil {
	// 	fmt.Println(err)
	// }

	app.Listen(":24130")
}

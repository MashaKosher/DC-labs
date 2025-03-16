package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Создаем новый кластерный объект
	cluster := gocql.NewCluster("127.0.0.1") // Замените на адрес вашего Cassandra
	cluster.Port = 9042
	cluster.Keyspace = "distcomp" // Замените на ваш keyspace

	// Устанавливаем соединение
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Выполняем простой запрос
	var country string
	var storyId int64
	var id int64
	var content string

	// Запрос данных из таблицы
	iter := session.Query("SELECT country, storyId, id, content FROM tbl_message").Iter()
	for iter.Scan(&country, &storyId, &id, &content) {
		fmt.Printf("Country: %s, StoryId: %d, ID: %d, Content: %s\n", country, storyId, id, content)
	}

	// fmt.Println("Value:", value)
}

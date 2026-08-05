package main

import (
	"chat/pkg"

	lorem "github.com/drhodes/golorem"
)

func main() {
	pkg.Init()
	db := pkg.OpenDB()

	tables := []any{
		&pkg.Message{},
	}

	db.Migrator().DropTable(tables...)
	db.AutoMigrate(tables...)

	messages := []*pkg.Message{}
	for range 1000 {
		messages = append(messages, &pkg.Message{
			Nickname: lorem.Word(5, 10),
			Message:  lorem.Sentence(5, 10),
		})
	}
	db.Create(messages)
}

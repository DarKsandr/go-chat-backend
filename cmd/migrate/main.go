package main

import "chat/pkg"

func main() {
	pkg.Init()
	db := pkg.OpenDB()

	tables := []any{
		&pkg.Message{},
	}

	db.Migrator().DropTable(tables...)
	db.AutoMigrate(tables...)
}

package main

import "github.com/go-pg/pg/v10"

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsStrip   bool     `json:"is_strip" pg:"is_strip"`
	Color     string   `json:"color" pg:"color"`
}

// Получить всех котиков из таблицы.

func FindAllCats() []Cat {
	var cats []Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cats).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cats
}

//CreateCat Добавить котика.

func CreateCat(cat Cat) Cat {
	pg.Connect := PostgresConnect()

	_, err := pg.Connect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}


	pgConnect.Close()
	return cat

}

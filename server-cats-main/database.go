package main

func PostgresConnect() *pg.DB {
	options := &pg.Options{
		User:     "postgre",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "cats",
	}

	options := &pg.Options{
		User:     "postgre",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "dogs",
	}

	return pgConnect(options)

}

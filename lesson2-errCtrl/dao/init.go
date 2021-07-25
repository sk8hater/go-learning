package dao

import (
	"database/sql"
	"fmt"
)

type Dao struct {
	engine *sql.DB
}

func New() *Dao {
	engine, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(fmt.Errorf("db open error: %w", err))
	}
	defer func(engine *sql.DB) {
		err := engine.Close()
		if err != nil {
			panic(fmt.Errorf("db close error: %w", err))
		}
	}(engine)
	return &Dao{
		engine: engine,
	}
}

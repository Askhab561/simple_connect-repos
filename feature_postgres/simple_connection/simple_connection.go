package simpleconnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:0880@localhost:5433/postgres")
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("подключение к базе данных прошла успешно")
}

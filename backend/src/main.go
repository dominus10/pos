package main

import (
	"context"
	"log"

	"github.com/dominus10/pos/db"
	v1 "github.com/dominus10/pos/src/routes/v1"
	"github.com/dominus10/pos/src/routes/v1/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	router:= gin.Default()

	conn, err := pgx.Connect(ctx, "user=postgres password=postgres dbname=postgres host=192.168.18.75 port=5432")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	router.NoRoute(util.NoRouteFound)
	v1.Routes(router,ctx,queries)
	router.Run()
}
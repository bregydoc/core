package main

import (
	"context"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/minskylab/core/ent"
	"github.com/minskylab/core/ent/migrate"
	// _ "entgo.io/contrib/entgql/internal/todo/ent/runtime"
)

func main() {
	// var cli struct {
	// 	Addr  string `name:"address" default:":8081" help:"Address to listen on."`
	// 	Debug bool   `name:"debug" help:"Enable debugging mode."`
	// }
	// kong.Parse(&cli)

	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		// log.Fatal("opening ent client", zap.Error(err))
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		// log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(nil)
	srv.Use(entgql.Transactioner{TxOpener: client})
	// if cli.Debug {
	// 	srv.Use(&debug.Tracer{})
	// }

	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)

	// log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// log.Error("http server terminated", zap.Error(err))
	}
}

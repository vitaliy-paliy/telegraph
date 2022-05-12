package main

import (
	"net/http"
	"time"

	"telegraph/db"
	"telegraph/graph"
	"telegraph/graph/directives"
	"telegraph/graph/generated"
	tmiddleware "telegraph/middleware"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func newServer() (server *handler.Server) {
	// Initialize database client.
	client, err := db.Start()
	if err != nil {
		panic(err)
	}

	// New server.
	resolver := graph.NewResolver(client)
	config := generated.Config{Resolvers: resolver}
	config.Directives.Auth = directives.Auth
	server = handler.New(generated.NewExecutableSchema(config))

	// New Websocket && CORS.
	server.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.MultipartForm{})

	server.Use(extension.Introspection{})
	server.SetQueryCache(lru.New(1000))
	server.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return
}

func newRouter(server *handler.Server) (e *echo.Echo) {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(tmiddleware.Auth)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	// Routing.
	e.POST("/graphql", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/subscriptions", func(c echo.Context) error {
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return
}

func main() {
	server := newServer()
	router := newRouter(server)

	router.Logger.Fatal(router.Start(":8080"))
}

package web

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/zekroTJA/visitor-count/internal/database"
)

type WebServer struct {
	server *fasthttp.Server
	router *router.Router

	db database.Database

	ipwl []string
}

func New(db database.Database, ipwl []string) *WebServer {
	ws := new(WebServer)

	ws.db = db
	ws.ipwl = ipwl
	ws.router = router.New()
	ws.server = &fasthttp.Server{
		Handler: ws.router.Handler,
	}

	ws.router.GET("/{username}.svg", ws.handlerSvg)

	return ws
}

func (ws *WebServer) ListenAndServeBlocking(addr string) error {
	return ws.server.ListenAndServe(addr)
}

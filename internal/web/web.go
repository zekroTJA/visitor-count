package web

import (
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/zekroTJA/timedmap"
	"github.com/zekroTJA/visitor-count/internal/database"
)

const (
	ipmapCleanupDuration = 10 * time.Minute
	ipmapValueLifetime   = 24 * time.Hour
)

type WebServer struct {
	server *fasthttp.Server
	router *router.Router

	db database.Database

	ipmap *timedmap.TimedMap
}

func New(db database.Database) *WebServer {
	ws := new(WebServer)

	ws.db = db
	ws.ipmap = timedmap.New(ipmapCleanupDuration)
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

package web

import (
	"github.com/valyala/fasthttp"
	"github.com/zekroTJA/visitor-count/internal/database"
	"github.com/zekroTJA/visitor-count/internal/log"
	"github.com/zekroTJA/visitor-count/internal/svg"
)

const (
	headerContentTypeSVG       = "image/svg+xml; charset=utf-8"
	headerContentTypePlainText = "text/plain; charset=utf-8"
	headerCacheControl         = "max-age=0, no-cache, no-store, must-revalidate"
)

func (ws *WebServer) handlerSvg(ctx *fasthttp.RequestCtx) {
	userName, _ := ctx.UserValue("username").(string)

	log.Log.Debugf("request from addr: %s", getIPAddr(ctx))

	count, err := ws.db.GetUserCount(userName)
	if err == database.ErrDatabaseNotFound {
		err = ws.db.SetUserCount(userName, 1)
	}
	if err != nil {
		sendError(ctx, err.Error(), 500)
		return
	}

	err = ws.db.UpdateUserCount(userName, 1)
	if err != nil {
		sendError(ctx, err.Error(), 500)
		return
	}

	data := svg.GetFormattedSVG(count + 1)

	ctx.Response.Header.SetContentType(headerContentTypeSVG)
	ctx.Response.Header.Set("cache-control", headerCacheControl)
	ctx.SetBody(data)
}

package web

import (
	"github.com/valyala/fasthttp"
	"github.com/zekroTJA/visitor-count/internal/log"
)

func (ws *WebServer) isWhitelistedAddress(ctx *fasthttp.RequestCtx) bool {
	if len(ws.ipwl) < 1 {
		return true
	}

	addr := getIPAddr(ctx)

	for _, v := range ws.ipwl {
		if v == addr {
			return true
		}
	}

	log.Log.Debugf("address %s is not whitelisted", addr)
	return false
}

func getIPAddr(ctx *fasthttp.RequestCtx) string {
	forwardedfor := ctx.Request.Header.Peek("X-Forwarded-For")
	if forwardedfor != nil && len(forwardedfor) > 0 {
		return string(forwardedfor)
	}

	return ctx.RemoteIP().String()
}

func sendError(ctx *fasthttp.RequestCtx, msg string, code int) {
	ctx.Response.Header.SetContentType(headerContentTypePlainText)
	ctx.SetStatusCode(code)
	ctx.Response.SetBody([]byte(msg))
}

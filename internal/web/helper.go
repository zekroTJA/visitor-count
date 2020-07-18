package web

import "github.com/valyala/fasthttp"

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

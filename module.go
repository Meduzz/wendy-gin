package wendygin

import (
	"github.com/Meduzz/wendy"
	"github.com/gin-gonic/gin"
)

// WendyModule returns a gin.Handler that parses the request body into a
// *wendy.Request and adds any headers present in both headers param and
// ctx.Header() into the *wendy.Request before submitting it to logic.Handle()
// and then writing the *wendy.Response back to the *gin.Context.
func WendyModule(logic wendy.Wendy, headers ...string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// start by binding request
		req := &wendy.Request{}
		err := ctx.BindJSON(req)

		if err != nil {
			// Somebody obviously sent us a bad body
			ctx.AbortWithStatus(400)
			return
		}

		if req.Headers == nil {
			req.Headers = make(map[string]string)
		}

		for _, h := range headers {
			v := ctx.GetHeader(h)
			req.Headers[h] = v
		}

		// call wendy
		res := logic.Handle(ctx, req)

		// start dealing with the response
		if res.Headers != nil {
			for k, v := range res.Headers {
				ctx.Header(k, v)
			}
		}

		if res.Body != nil {
			ctx.Data(res.Code, res.Body.Type, res.Body.Data)
		} else {
			ctx.Status(res.Code)
		}
	}
}

// WendyModuleOnlyBody returns a gin.Handler that parses the request body into the body of a
// new *wendy.Request, setting provided module & method and adds any headers present in both headers param and
// ctx.Header() into the *wendy.Request before submitting it to logic.Handle()
// and then writing the *wendy.Response back to the *gin.Context.
func WendyModuleOnlyBody(app, module, method string, logic wendy.Wendy, headers ...string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// start by binding request
		req := &wendy.Request{
			App:    app,
			Module: module,
			Method: method,
		}

		bs, err := ctx.GetRawData()

		if err != nil {
			// Somebody obviously sent us a bad body
			ctx.AbortWithStatus(400)
			return
		}

		req.Body = &wendy.Body{
			Type: ctx.ContentType(),
			Data: bs,
		}

		if req.Headers == nil {
			req.Headers = make(map[string]string)
		}

		for _, h := range headers {
			v := ctx.GetHeader(h)
			req.Headers[h] = v
		}

		// call wendy
		res := logic.Handle(ctx, req)

		// start dealing with the response
		if res.Headers != nil {
			for k, v := range res.Headers {
				ctx.Header(k, v)
			}
		}

		if res.Body != nil {
			ctx.Data(res.Code, res.Body.Type, res.Body.Data)
		} else {
			ctx.Status(res.Code)
		}
	}
}

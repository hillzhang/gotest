package fastHttp


import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s, %s!\n", ctx.UserValue("name"),ctx.UserValue("age"))
}

type MyHandler struct {
	foobar string
}

// request handler in net/http style, i.e. method bound to MyHandler struct.
func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	// notice that we may access MyHandler properties here - see h.foobar.
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q. Foobar is %q",
		ctx.Path(), h.foobar)
}

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func Fast() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name/:age", Hello)
	log.Fatal(fasthttp.ListenAndServe(":8081", router.Handler))
	//myHandler := &MyHandler{
	//	foobar: "foobar",
	//}
	//fasthttp.ListenAndServe(":8080", myHandler.HandleFastHTTP)
	//
	//// pass plain function to fasthttp
	//fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}

package rest

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"log"
	"time"
	"net/http"
)
func handle_auth(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}
func Rest() {
	jwt_middleware := &jwt.JWTMiddleware{
		Key:        []byte("secret key"),
		Realm:      "jwt auth",
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string) bool {
			return userId == "admin" && password == "admin"
		}}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	// we use the IfMiddleware to remove certain paths from needing authentication
	api.Use(&rest.IfMiddleware{
		Condition: func(request *rest.Request) bool {
			return request.URL.Path != "/login"
		},
		IfTrue: jwt_middleware,
	})
	api_router, _ := rest.MakeRouter(
		rest.Post("/login", jwt_middleware.LoginHandler),
		rest.Get("/auth_test", handle_auth),
		rest.Get("/refresh_token", jwt_middleware.RefreshHandler),
	)
	api.SetApp(api_router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
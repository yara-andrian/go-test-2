package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// methodRoute "server-routing/router/method"
	// paramRoute "server-routing/router/param"
	// staticRoute "server-routing/router/static"
	"test/logger"
	"time"

	"github.com/gorilla/mux"
)

// ArticlesCategoryHandler abc and return
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

// Create the router and return it
func Create() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		foo := struct {
			Hello string
			JSON  string
		}{
			Hello: "world",
			JSON:  "stuff",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(foo)
		// fmt.Printf("foo struct : %+v\n", foo)
	})
	router.HandleFunc("/bar", ArticlesCategoryHandler)
	// secureMiddleware := secure.New(secure.Options{
	// 	AllowedHosts:       []string{"localhost:9090"},
	// 	BrowserXssFilter:   true,
	// 	ContentTypeNosniff: true,
	// 	ForceSTSHeader:     true,
	// 	FrameDeny:          true,
	// 	STSSeconds:         5000,
	// })
	// router.Use(secureMiddleware.Handler)
	// staticRoute.Provision(router)
	// methodRoute.Provision(router)
	// paramRoute.Provision(router)
	router.Use(routerLogger)
	return router
}

type logFormat struct {
	Method        string    `json:"method"`
	Host          string    `json:"host"`
	Path          string    `json:"path"`
	RemoteAddr    string    `json:"remote-addr"`
	UserAgent     string    `json:"user-agent"`
	ContentLength int64     `json:"content-length"`
	Referer       string    `json:"referer"`
	LocalHostname string    `json:"local-hostname"`
	Origin        string    `json:"origin"`
	Timestmap     time.Time `json:"timestamp"`
}

func routerLogger(next http.Handler) http.Handler {
	localHostname, _ := os.Hostname()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infow("access", "log", logFormat{
			r.Method,
			r.Host,
			r.RequestURI,
			r.RemoteAddr,
			r.UserAgent(),
			r.ContentLength,
			r.Referer(),
			localHostname,
			r.Header.Get("origin"),
			time.Now().UTC(),
		})
		next.ServeHTTP(w, r)
	})
}

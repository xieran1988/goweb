
package main

import (
	"fmt"
	"net/http"
	"github.com/hoisie/mustache"
	"log"
	"strings"
	"time"
)

type Nav struct {
	Href string
	Name string
}
type hash map[string]interface{}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	statics := []string{"/css", "/js"}
	for _, p := range statics {
		if strings.HasPrefix(r.URL.Path, p) {
			file := r.URL.Path[1:]
			http.ServeFile(w, r, file)
			return
		}
	}

	r.ParseForm()

	first := func (key string) string {
		a := r.Form[key]
		if len(a) > 0 {
			return a[0]
		}
		return ""
	}

	login := func (h hash) string {
		navtop := mustache.RenderFile("navtop.html",
			hash{ "title": "world" },
			hash{ "nav": []Nav{ {"1","2"}, {"3","4"} } },
		)
		login := mustache.RenderFile("login.html", h)
		index := mustache.RenderFile("index.html", hash{"navtop":navtop, "login":login}, h)
		return index
	}

	if strings.HasPrefix(r.URL.Path, "/login") {
		user := first("user")
		pass := first("pass")
		if user != "admin" && pass != "admin" {
			fmt.Fprintf(w, login(hash{"tips":"登录失败"}))
			return
		}
		cookie := http.Cookie{Name: "user", Value: user, Expires: time.Now().Add(10*time.Hour)}
		http.SetCookie(w, &cookie)
		fmt.Fprintf(w, user, pass)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/form") {
		content := first("content")
		fmt.Fprintf(w, content)
		return
	}

	fmt.Fprintf(w, login(hash{}))
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


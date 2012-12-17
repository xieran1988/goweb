
package main

import (
	"fmt"
	"net/http"
	"github.com/hoisie/mustache"
	"log"
	"strings"
	"time"
	"encoding/json"
	"io/ioutil"
)

type Nav struct {
	Href string
	Name string
}
type hash map[string]interface{}
type Db struct {
	Content string
	Time string
}

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

	renderIndex := func (body string) string {
		navtop := mustache.RenderFile("navtop.html",
			hash{ "title": "推送通知" },
//			hash{ "nav": []Nav{ {"1","2"}, {"3","4"} } },
		)
		index := mustache.RenderFile("index.html", hash{
			"navtop":navtop,
			"body":body,
		})
		return index
	}

	if strings.HasPrefix(r.URL.Path, "/login") {
		body := mustache.RenderFile("login.html")
		index := renderIndex(body)
		fmt.Fprintf(w, index)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/post-login") {
		user := first("user")
		pass := first("pass")
		if user != "admin" && pass != "admin" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		cookie := http.Cookie{Name: "user", Value: user, Expires: time.Now().Add(10*time.Hour)}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/show", http.StatusFound)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/show") {
		var db Db
		b, err := ioutil.ReadFile("db")
		if err == nil {
			json.Unmarshal(b, &db)
		} else {
			t := time.Now()
			db.Content = ""
			db.Time = fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
		}
		form := mustache.RenderFile("show.html", db)
		index := renderIndex(form)
		fmt.Fprintf(w, index)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/post-show") {
		var db Db
		db.Content = first("Content")
		db.Time = first("Time")

		b, err := json.Marshal(db)
		if err != nil {
			return
		}
		ioutil.WriteFile("/usr/lib/pushdb", b, 0644)

		http.Redirect(w, r, "/show?saveok", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


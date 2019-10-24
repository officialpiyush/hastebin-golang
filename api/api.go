/**
 * Copyright (C) 2019 Piyush Bhangale
 *
 * This file is part of hastebin-golang.
 *
 * hastebin-golang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * hastebin-golang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with hastebin-golang.  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"bytes"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"hastebin-golang/database"
	"log"
	"net/http"
	"strings"
)

func Get(w http.ResponseWriter, r *http.Request) {
	files := [11]string{"application.css", "application.js", "application.min.js", "favicon.ico", "function-icons.png", "highlight.min.js", "hover-dropdown-tip.png", "index.html", "logo.png", "robots.txt", "solarized_dark.css"}
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.ServeFile(w, r, "./static/"+"index.html")
		return
	}
	isPresent := contains(files, vars["id"])
	if isPresent {
		http.ServeFile(w, r, "./static/"+vars["id"])
		return
	}
	//else {
	//	//Remove periods, etc from the key
	//	if strings.ContainsAny(vars["id"], ".") {
	//		id := strings.Split(vars["id"], ".")
	//		vars["id"] = id[0]
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	doc, ok := database.GetDocument(vars["id"])
	//	json := simplejson.New()
	//	if !ok {
	//		json.Set("message", "Document not found.")
	//		payload, err := json.MarshalJSON()
	//		if err != nil {
	//			log.Println(err)
	//		}
	//		w.WriteHeader(http.StatusNotFound)
	//		_, _ = w.Write(payload)
	//		return
	//	} else if ok {
	//		json.Set("key", vars["id"])
	//		json.Set("data", doc)
	//		payload, _ := json.MarshalJSON()
	//		_, _ = w.Write(payload)
	//		return
	//	}
	//}
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//Remove periods, etc from the key
	if strings.ContainsAny(vars["id"], ".") {
		id := strings.Split(vars["id"], ".")
		vars["id"]  = id[0]
	}

	w.Header().Set("Content-Type", "application/json")
	doc, ok := database.GetDocument(vars["id"])
	json := simplejson.New()
	if !ok {
		json.Set("message", "Document not found.")
		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(payload)
		return
	} else if ok {
		json.Set("key", vars["id"])
		json.Set("data", doc)
		payload, _ := json.MarshalJSON()
		_, _ = w.Write(payload)
		return
	}
}

func HandlePost(w http.ResponseWriter, r *http.Request){
	buffer := new(bytes.Buffer)
	_, _ = buffer.ReadFrom(r.Body)
	body := buffer.String()

	key := database.CreateDocument(body)
	json := simplejson.New()

	if key == "" {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.Set("key", key);
	payload, _ := json.MarshalJSON()
	_, _ = w.Write(payload)
	return
}

func contains(s [11]string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

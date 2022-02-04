package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "welcome! \n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

var personDB map[string]*Person

func init() {
	personDB = map[string]*Person{}
	personDB["mike"] = &Person{"mike", 20}
	personDB["rose"] = &Person{"rose", 24}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", GetPersonByName)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func GetPersonByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	var (
		ok       bool
		info     *Person
		infoJson []byte
		err      error
	)
	if info, ok = personDB[qName]; !ok {
		w.Write([]byte("cannot get '" + qName + "' info"))
		return
	}
	person := Person{Name: info.Name, Age: info.Age}
	if infoJson, err = person.MarshalJSON(); err != nil {
		w.Write([]byte("Get query 'name' failed"))
	} else {
		w.Write(infoJson)
	}

}

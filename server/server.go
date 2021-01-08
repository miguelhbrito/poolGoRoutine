package main

import (
"encoding/json"
"fmt"
"log"
"net/http"
"os"
"time"

"github.com/julienschmidt/httprouter"
)

func myHandlerServer(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id := ps.ByName("id")
	log.Printf("received goroutine id %s\n", id)
	time.Sleep(time.Second * 2)

	data := fmt.Sprintf("done %s", id)
	w.Header().Add("Content-Type", "application/json")
	sv, _ := json.Marshal(data)
	_, _ = w.Write(sv)
}

func myHandlerServerTwo(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id := ps.ByName("id")
	log.Printf("received goroutine id %s\n", id)
	time.Sleep(time.Millisecond * 100)

	data := fmt.Sprintf("done %s", id)
	w.Header().Add("Content-Type", "application/json")
	sv, _ := json.Marshal(data)
	_, _ = w.Write(sv)
}

func main() {
	router := httprouter.New()

	router.GET("/one/:id", myHandlerServer)
	router.GET("/two/:id", myHandlerServerTwo)

	log.Println("Server listening on 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		os.Exit(1)
	}
}


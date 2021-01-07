package exemplo3

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Api(id int , url string, ch chan<-string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	secs := time.Since(start).Seconds()
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Printf("error reading response body :: error is %+v", err)
		return
	}
	ch <- fmt.Sprintf("%.2f elapsed with response: %s", secs, url)
	//log.Println(string(body))
	log.Printf("%d  :: ok", id)
}
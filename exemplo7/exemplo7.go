package exemplo7

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

var id = 0

func ApiWorker() {
	id++
	log.Info().Msgf("exemplo 7 worker %d begin", id)

	url := fmt.Sprintf("http://localhost:8080/two/%d", id)

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal().Err(err)
	}
	secs := time.Since(start).Milliseconds()

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Info().Msgf("%v worker %d finished", secs, id)
}

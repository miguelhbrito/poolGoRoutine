package exemplo6

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

func ApiWorker(id interface{}) {
	log.Info().Msgf("worker %d begin", id)
	url := fmt.Sprintf("http://localhost:8080/two/%d", id)

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal().Err(err)
	}
	secs := time.Since(start).Seconds()

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Info().Msgf("elapsed %.2f worker %d finished", secs, id)
}

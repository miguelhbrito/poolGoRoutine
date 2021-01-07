package exemplo5

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func ApiWorker(i interface{}, itens []int) {
	log.Info().Msgf("worker %d begin", i)
	n := i.(int)
	id := itens[n]
	url := fmt.Sprintf("https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d", id)

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

	log.Info().Msgf("elapsed %.2f worker %d finished", secs, i)
}
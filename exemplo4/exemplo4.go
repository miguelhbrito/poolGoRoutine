package exemplo4

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

func ApiWorker(worker int, itens []int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for j := range jobs {
		id := itens[count]

		url := fmt.Sprintf("https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d", id)
		log.Info().Msgf("worker %d started job %d", worker, j)

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

		count++
		log.Info().Msgf("elapsed %.2f worker %d finished job %d", secs, id, j)
		results <- j
	}
}
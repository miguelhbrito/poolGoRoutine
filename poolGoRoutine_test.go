package main

import (
	"fmt"
	"github.com/poolGoRoutine/exemplo1"
	"github.com/poolGoRoutine/exemplo3"
	"github.com/poolGoRoutine/exemplo4"
	"github.com/poolGoRoutine/exemplo5"
	"github.com/poolGoRoutine/exemplo6"
	"github.com/poolGoRoutine/exemplo7"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/panjf2000/ants"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
)

const (
	poolSize = 1000
	jobSize = 100000
)

var curMem uint64

//usando o conceito de workers do proprio golang
func TestGoroutineWorkers(t *testing.T) {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	start := time.Now()
	for w := 1; w <= 3; w++{
		go exemplo1.Worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++{
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++{
		<- results
	}

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

func TestGoroutineWithoutWorkersWithHttpRequest(t *testing.T){
	itens := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4}

	start := time.Now()
	ch := make(chan string)

	for id , item := range itens{
		url := fmt.Sprintf("https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/%d", item)
		go exemplo3.Api(id, url, ch)
	}

	for range itens{
		fmt.Println(<-ch)
	}

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

//usando o conceito de workers do proprio golang para fazer requisições http
func TestGoroutineWorkersWithHttpRequest(t *testing.T){

	itens := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4}

	start := time.Now()
	var wg sync.WaitGroup
	var numJobs = len(itens)

	workers := 20
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i < workers; i++{
		wg.Add(1)
		go exemplo4.ApiWorker( i , itens, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for range itens{
		<-results
	}
	wg.Wait()

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)

}

func TestAntsPoolWithHttpRequest(t *testing.T){

	itens := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3,
		4, 1, 2, 3, 4}

	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(20, func(i interface{}) {
		exemplo5.ApiWorker(i, itens)
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < len(itens); i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

//espera pra pegar a goroutine,
//criando uma pool com uma função pre estabelecida
//e sincronizada com o sync.WaitGroup
func TestPoolWithFuncWaitToGet(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		exemplo6.ApiWorker(i)
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
	fmt.Printf("pool go routine, workers:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

//espera pra pegar a goroutine, usando o submit,
//criando uma pool e sincronizada com o sync.WaitGroup
func TestPoolWaitToGetWorker(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPool(poolSize)
	defer p.Release()

	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			exemplo7.ApiWorker()
			wg.Done()
		})
	}


	wg.Wait()
	fmt.Printf("pool go routine, workers:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

func TestPoolWaitToGetWorkerWith2Submits(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPool(poolSize)
	defer p.Release()

	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			exemplo7.ApiWorker()
			wg.Done()
		})

		wg.Add(1)
		_ = p.Submit(func() {
			exemplo6.ApiWorker(i)
			wg.Done()
		})
	}

	wg.Wait()
	fmt.Printf("pool go routine, workers:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

//espera pra pegar a goroutine com memoria alocada previamente e
//sincronizada com o sync.WaitGroup
func TestPoolWaitToGetWorkerPreMalloc(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPool(poolSize, ants.WithPreAlloc(true))
	defer p.Release()

	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			exemplo7.ApiWorker()
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Printf("pool go routine, workers:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	fmt.Printf("memoria usada:%d MB", curMem)
}

//-------------------------------------------------------------------------------------------
// Go routines para testes
//-------------------------------------------------------------------------------------------

func TestNoPool(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		go func() {
			exemplo7.ApiWorker()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestAntsPool(t *testing.T) {
	start := time.Now()
	defer ants.Release()
	var wg sync.WaitGroup
	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = ants.Submit(func() {
			exemplo7.ApiWorker()
			wg.Done()
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", ants.Cap())
	t.Logf("pool, running workers number:%d", ants.Running())
	t.Logf("pool, free workers number:%d", ants.Free())
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestAntsPoolGo(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	p, _ := ants.NewPool(poolSize)
	defer p.Release()
	for i := 0; i < jobSize; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			exemplo7.ApiWorker()
			wg.Done()
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", p.Cap())
	t.Logf("pool, running workers number:%d", p.Running())
	t.Logf("pool, free workers number:%d", p.Free())
	fmt.Printf("\n%v milliseconds elapsed\n", time.Since(start).Milliseconds())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}
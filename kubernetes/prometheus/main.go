package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/procfs"
)

func pullmetrics() {
	addr := flag.String("addr", ":5198", "The address to listen on for HTTP requests.")
	flag.Parse()

	var (
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "sun_temperature",
			Help: "current sun temperature",
		})
		counter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "blink_total",
				Help: "blink count",
			},
			[]string{"star"},
		)
	)

	reg := prometheus.NewRegistry()
	reg.MustRegister(gauge)
	reg.MustRegister(counter)

	handlerFunc := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			counter.With(prometheus.Labels{"star": "venus"}).Inc()
			gauge.Set(rand.Float64() * 100)
			next.ServeHTTP(w, r)
		})
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", handlerFunc(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))
	mux.Handle("/metrics2", promhttp.Handler())
	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}

func pushmetrics() {
	var (
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "sun_temperature",
			Help: "current sun temperature",
		})
		counter = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "blink_total",
				Help: "blink count",
			},
			[]string{"star"},
		)
	)

	pusher := push.New("http://192.168.126.128:9091", "sky_observer").Collector(gauge).Collector(counter)
	for {
		counter.With(prometheus.Labels{"star": "venus"}).Inc()
		gauge.Set(rand.Float64() * 100)

		if err := pusher.Push(); err != nil {
			log.Fatal(err)
			return
		}
		time.Sleep(time.Second)
	}
}

func getproc() {
	fs, err := procfs.NewDefaultFS()
	if err != nil {
		log.Fatal(err)
		return
	}
	cpus, err := fs.CPUInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(len(cpus))
	log.Println(cpus)

	mem, err := fs.Meminfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(mem)

	proc, err := fs.Self()
	if err != nil {
		log.Fatal(err)
		return
	}
	fds, err := proc.FileDescriptorTargets()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(fds)

	st, err := proc.NewStat()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(st)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go pullmetrics()
	go pushmetrics()
	go getproc()
	select {}
}

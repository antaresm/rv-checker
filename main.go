package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/mitchellh/panicwrap"
	"os/signal"
	"syscall"
	"fmt"
	"os"
	"rv-check/handlers"
	"av-check/myAdmin"
	"rv-check/checker"
	"github.com/getsentry/raven-go"

)

func init() {
	raven.SetDSN("http://e32dc25e8ede495c8945351b87f84677:97c1902842e74b12b848a6300ca009cc@sentry.headsandhands.ru:9000/2")
	
	raven.config('http://e32dc25e8ede495c8945351b87f84677:97c1902842e74b12b848a6300ca009cc@sentry.headsandhands.ru:9000/2"',
	 {
		release: '0e4fdef81448dcfa0e16ecc4433ff3997aa53572'
	});

}

func main(){
	_, err := os.Open("filename.ext")
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}

	s := os.Getenv("DEBUG")
	if s != "true" {
		onStart()
	}

	log.Println("Hello, world!")
	checker.Init()

	log.Println("Hello, http!")
	initHttp()
}

func initHttp(){
	r := mux.NewRouter()
	r.Handle("/", handlers.StatusHandler).Methods("GET")

	log.Println("Hello, admin!")
	m := myAdmin.InitAdmin()
	r.PathPrefix("/admin").Handler(m)

	http.ListenAndServe(":9999", r)
}

func onStart(){
	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		panic(err)
	}
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		onClose()
		os.Exit(0)
	}()
}

func onClose() {
	checker.CheckClose()
	log.Println("Buy!")
}

func panicHandler(output string) {
	onClose()
	fmt.Printf("The child panicked:\n\n%s\n", output)
	os.Exit(1)
}
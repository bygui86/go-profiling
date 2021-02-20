package multi_profile

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/multi-profile/v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	// multi-profile
	defer profile.CPUProfile(&profile.Config{}).Start().Stop()
	defer profile.MemProfile(&profile.Config{}).Start().Stop()
	defer profile.GoroutineProfile(&profile.Config{}).Start().Stop()

	// http server
	srv := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	http.HandleFunc("/", handler)
	go srv.ListenAndServe()

	// waiting for signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	// closing
	err := srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

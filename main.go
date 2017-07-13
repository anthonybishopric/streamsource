package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pallinder/go-randomdata"
)

var delay *bool = flag.Bool("delay", false, "Add a delay into message responses")
var makeNil *bool = flag.Bool("nil", false, "Sometimes return nothing")
var streamCount *int = flag.Int("streams", 3, "Number of streams to create")

func serve(port int) {
	log.Printf("Serving messages on port %d", port)
	go http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), &handler{
		rand.New(rand.NewSource(time.Now().Unix() + int64(port))),
		0,
	})
}

type handler struct {
	source  *rand.Rand
	current int
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if *delay {
		time.Sleep(time.Duration(h.source.Int63n(1000)) * time.Millisecond)
	}
	if *makeNil {
		if h.source.Intn(3) == 1 {
			resp.Write([]byte("{}"))
			return
		}
	}
	h.current += h.source.Intn(4) + 1

	out, _ := json.Marshal(struct {
		Timestamp int    `json:"timestamp"`
		Message   string `json:"message"`
	}{h.current, fmt.Sprintf("%s: %s %s -- %s", randomdata.IpV4Address(),
		randomdata.Email(), randomdata.Noun(), randomdata.Timezone())})

	resp.Write(out)
}

func main() {

	flag.Parse()

	for i := 0; i < *streamCount; i++ {
		serve(8000 + i)
	}

	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, syscall.SIGTERM, os.Interrupt)
	<-signalCh
}

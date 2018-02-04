package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var data []string

		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, val := range headers {
				data = append(data, fmt.Sprintf("%v: %v", name, val))
			}
		}

		var hostname string
		hostname, err := os.Hostname()
		if err != nil {
			log.Error(err)
		}

		data = append(data, fmt.Sprintf("hostname: %v", hostname))
		data = append(data, fmt.Sprintf("remote address: %v", r.RemoteAddr))
		sort.Strings(data)

		htmlOutput := []byte(strings.Join(data, "\n"))
		w.Write(htmlOutput)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Infoln("starting req service on :8080")
	log.Fatalln(srv.ListenAndServe())
}

package server

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var clients []chan bool

func NotifyReload() {
	for _, c := range clients {
		go func(c chan bool) {
			c <- true
		}(c)
	}
}

func Start(port string, webDir string, liveReload bool) error {
	KillPort(port) // Ensure port is free

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	if liveReload {
		// Live Reload Endpoint (SSE)
		http.HandleFunc("/livereload", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			w.Header().Set("Access-Control-Allow-Origin", "*")

			clientChan := make(chan bool)
			clients = append(clients, clientChan)
			
			notify := w.(http.CloseNotifier).CloseNotify()
			
			go func() {
				<-notify
				// Remove client (simple leak for now, proper map needed for production but fine for local tool)
			}()

			for {
				<-clientChan
				fmt.Fprintf(w, "data: reload\n\n")
				w.(http.Flusher).Flush()
			}
		})
	}

	log.Printf("Serving website on http://localhost:%s", port)
	log.Println("Press Ctrl+C to stop")
	return http.ListenAndServe(":"+port, nil)
}

func KillPort(port string) {
	log.Printf("Ensuring port %s is free...", port)
	// Try to kill process on port using fuser (common on Linux)
	cmd := exec.Command("fuser", "-k", port+"/tcp")
	cmd.Run() // Ignore error, it just means no process was found or fuser isn't installed
    
    // Give it a moment to release the port
    time.Sleep(500 * time.Millisecond)
}

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwangelme/kubia/version"
)

var (
	listenFlag    = flag.String("listen", "0.0.0.0:5678", "address and port to listen")
	textFlag      = flag.String("text", "", "text to put on the webpage")
	versionFlag   = flag.Bool("version", false, "display version information")
	startSlowFlag = flag.Bool("slow", false, "start slow")
)

const filePath = "/var/data/kubia.txt"

func httpEcho(v string) http.HandlerFunc {
	hostname, _ := os.Hostname()
	return func(w http.ResponseWriter, r *http.Request) {
		content := map[string]string{
			"path":     r.RequestURI,
			"hostname": hostname,
			"inst":     v,
		}
		e := json.NewEncoder(w)
		e.Encode(content)
	}
}

func Storage(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	if r.Method == http.MethodPost {
		log.Printf("write the file %v\n", filePath)
		fd, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintf(w, "open file failed %v", err)
			return
		}
		n, err := io.Copy(fd, r.Body)
		if err != nil {
			fmt.Fprintf(w, "write file failed %v", err)
			return
		}
		fmt.Fprintf(w, "write %d bytes to file on %v", n, hostname)
		return
	}

	// GET method
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "节点 %v 上没有数据", hostname)
		return
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Read file error %v", err)
		return
	}
	fmt.Fprintf(w, "%s 节点上的内容是:\n", hostname)
	w.Write(content)
	return
}

func httpHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status":"ok"}`)
	}
}

const (
	httpHeaderAppName    string = "X-App-Name"
	httpHeaderAppVersion string = "X-App-Version"
)

// withAppHeaders adds application headers such as X-App-Version and X-App-Name.
func withAppHeaders(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(httpHeaderAppName, version.Name)
		w.Header().Set(httpHeaderAppVersion, version.Version)
		h(w, r)
	}
}

func main() {
	flag.Parse()

	// Asking for the version?
	if *versionFlag {
		fmt.Fprintln(os.Stderr, version.HumanVersion)
		os.Exit(0)
	}

	// Validation
	if *textFlag == "" {
		fmt.Fprintln(os.Stderr, "Missing -text option!")
		os.Exit(127)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", withAppHeaders(httpEcho(*textFlag)))
	mux.HandleFunc("/health", withAppHeaders(httpHealth()))
	mux.HandleFunc("/storage", Storage)

	server := &http.Server{
		Addr:    *listenFlag,
		Handler: mux,
	}

	if *startSlowFlag {
		time.Sleep(30 * time.Second)
	}

	serverCh := make(chan struct{})
	go func() {
		log.Printf("[INFO] server is listening on %s\n", *listenFlag)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("[ERR] server exited with: %s", err)
		}
		close(serverCh)
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for interrupt
	<-signalCh

	log.Printf("[INFO] received interrupt, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("[ERR] failed to shutdown server: %s", err)
	}

	os.Exit(0)
}

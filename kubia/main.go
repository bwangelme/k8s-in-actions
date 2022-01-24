package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const filePath = "/var/data/kubia.txt"

func Home(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "You've hit %s v2\n", hostname)
	return
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

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Read file error %v", err)
		return
	}
	fmt.Fprintf(w, "%s 节点上的内容是:\n", hostname)
	w.Write(content)
	return
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/storage", Storage)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

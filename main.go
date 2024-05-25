package main

import (
	"flag"
	"github.com/rmasci/fileserver"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "4410"
	dir := "/usr/data/creality/userdata/delay_image/video"
	title := "Timelapse"
	flag.StringVar(&port, "p", port, "Port to listen on.")
	flag.StringVar(&dir, "d", dir, "Directory to serve")
	flag.StringVar(&title, "t", title, "Title")
	flag.Parse()
	//html := os.Getenv("PWD") + "/html"
	lgOut := log.New(os.Stdout, "/tmp/downloads.log", log.Lshortfile)
	dwnld := fileserver.Directory{
		Lgout:  lgOut,
		Px:     15,
		Srv:    dir,
		Header: title,
		//Directory: dir,
	}

	http.HandleFunc("/downloads/", dwnld.Fileserver)
	//redirect to downloads
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/downloads/", http.StatusFound)
	})

	// let user know when server is listening on port
	lgOut.Println("Listening on port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

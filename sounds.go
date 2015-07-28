package main

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/play/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			log.Printf("Recieved %s request.", req.Method)

			if req.ContentLength > 10485760 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("File size capped at 10mb"))
				return
			}

			soundFile, _, err := req.FormFile("soundFile")
			if err != nil {
				log.Printf("Error getting soundFile from Form. \n %s", err.Error())
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			w.Write([]byte("Success!  http://localhost:3030/config.html"))
			playASound(soundFile)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)

			//TODO(cagocs): maybe return 200 with the name of the sound playing?
		}
	})

	controlfs := http.FileServer(http.Dir("control"))
	http.Handle("/", controlfs)

	http.ListenAndServe(":3030", nil)

}

func playASound(file multipart.File) {
	soundFile, err0 := ioutil.TempFile("", "sound_")
	if err0 != nil {
		log.Printf("Error initializing new file")
	}

	buffer, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		log.Printf("Error reading mime multipart file")
	}

	err2 := ioutil.WriteFile(soundFile.Name(), buffer, os.ModeTemporary)
	if err2 != nil {
		log.Printf("Error writing file to disk")
	}

	cmd := exec.Command("mplayer", soundFile.Name())

	err3 := cmd.Run()
	if err3 != nil {
		log.Printf("Error playing file %s", soundFile.Name())
	}

	soundFile.Close()
	err4 := os.Remove(soundFile.Name())
	if err4 != nil {
		log.Println("Error deleting %s", soundFile.Name())
	}

}

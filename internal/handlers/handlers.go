package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "could not get uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "could not get data from file", http.StatusInternalServerError)
		return
	}

	convertedString := service.ConvertMorseOrText(string(data))

	root, err := os.OpenRoot("/result/")
	if err != nil {
		http.Error(res, "internal error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer root.Close()

	filename := time.Now().UTC().Format("2006-01-02_15-04-05") + ".txt"
	log.Println(filename)

	dst, err := root.Create(filename)
	if err != nil {
		http.Error(res, "file creation error", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	dst.WriteString(convertedString)

	res.WriteHeader(http.StatusOK)
	io.WriteString(res, convertedString)
}

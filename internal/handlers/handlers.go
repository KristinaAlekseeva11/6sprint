package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return
	}

	inputData := string(fileBytes)
	converted, err := service.ConvertAuto(inputData)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("converted_%s%s", strings.ReplaceAll(time.Now().UTC().Format(time.RFC3339), ":", "-"), ext)
	err = os.WriteFile(filename, []byte(converted), 0644)
	if err != nil {
		http.Error(w, "Error writing output file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Converted content:\n%s", converted)))
}

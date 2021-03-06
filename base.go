package jet

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 600 * 1024 * 1024 // 60 mb

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

//FileExists is check if the file exists
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//EnsureDir is check and make folder
func EnsureDir(dirName string) error {

	err := os.Mkdir(dirName, 0755)

	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

//ParamCheck sends back the URL parameters value
func ParamCheck(param string, r *http.Request, fallback string) (string, bool) {
	keys, ok := r.URL.Query()[param]

	if !ok || len(keys[0]) < 1 {
		return fallback, false
	}
	return keys[0], true
}

//UpFile is a helper for upload files
func UpFile(uploadPath string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		fileName := r.PostFormValue("name")
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		newPath := filepath.Join(uploadPath, fileName)
		//fmt.Printf("File: %s\n", newPath)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(newPath))
	})
}

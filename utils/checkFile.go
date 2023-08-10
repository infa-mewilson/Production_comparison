package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func CheckFileAlreadyExist(filenamenew string, writer http.ResponseWriter, r *http.Request) bool {
	fileInfo, _ := os.Stat(filenamenew)
	if fileInfo != nil {
		log.Println("File is already Present")
		responseMessage := fmt.Sprintf("File is Already Present. Please find the HtmlReports at localhost:6068/%v", filenamenew)
		RespondWithJSON(responseMessage, 400, writer, r)
		http.Error(writer, "Bad Request", http.StatusBadRequest)
		return false
	}
	return true
}

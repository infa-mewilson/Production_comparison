package utils

import (
	"net/http"
	"strings"
	"time"
)

func ValidateDateFormat(dateStr string) bool {
	_, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return false
	}
	return true
}

func Validatetheinput(InputDates string, writer http.ResponseWriter, r *http.Request) (bool, []string) {
	dates := strings.Split(InputDates, ",")
	for _, date := range dates {
		valid := ValidateDateFormat(date)
		if !valid {
			RespondWithJSON("You send Bad Request.Please Enter correct Date Format like YYYY-MM-DDT17:30:00.000Z", 400, writer, r)
			http.Error(writer, "Bad Request", http.StatusBadRequest)
			return false, dates
		}
	}
	return true, dates
}

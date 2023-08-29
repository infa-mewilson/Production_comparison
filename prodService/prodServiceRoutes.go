package prodService

import "net/http"

func Init() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/compare", compareResponseTimes)
	http.HandleFunc("/EnvironmentDetails",displayenvdetails)
	fileServer := http.FileServer(http.Dir("./htmlReports"))
	http.Handle("/htmlReports/", http.StripPrefix("/htmlReports/", fileServer))
}

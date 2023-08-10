package prodService

import "net/http"

func Init() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/compare", compareResponseTimes)
	fileServer := http.FileServer(http.Dir("./htmlReports"))
	http.Handle("/htmlReports/", http.StripPrefix("/htmlReports/", fileServer))
}

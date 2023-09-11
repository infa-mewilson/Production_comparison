package utils

import (
	"NEWGOLANG/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReleaseData(StartDate string, EndDate string, GlobalEnv string, AppEnv string) (map[string]config.Value, map[string]config.Value) {
	perfusw1map := make(map[string]config.Value)
	perfidsmap := make(map[string]config.Value)
	goFileURL := "https://raw.githubusercontent.com/infa-mewilson/Production_comparison/main/utils/API_Details.txt"
	resp, err := http.Get(goFileURL)
	if err != nil {
		http.Error(writer, "Failed to fetch Go file from GitHub", http.StatusInternalServerError)
		return nil,nil
	}
	defer resp.Body.Close()

	// Read the response body
	contentBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(writer, "Failed to read response body", http.StatusInternalServerError)
		return nil,nil
	}
	listdetails := string(contentBytes)
	var result config.APIDteails

	err := json.Unmarshal([]byte(listdetails), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	//fmt.Println(result)

	for _, perfusw1 := range result.IicsQaPerfusw1 {
		var containernames, filepaths string
		containernames = perfusw1.Containername
		filepaths = perfusw1.Logfilepath
		for _, apisofservices := range perfusw1.SchedulerService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.KmsService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.NotificationService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.CAService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.LicenseService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.JlsService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.BundleService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}

		for _, apisofservices := range perfusw1.SessionService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.Frs {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AuditService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.PreferenceService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apis := range perfusw1.MigrationService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AdminService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.Vcs {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apisofservices := range perfusw1.AcService {
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames, filepaths, AppEnv)
			SendESRequest(apisofservices, perfusw1map, query)
		}
		for _, apis := range perfusw1.LdmService {
			query := buildmyquery(StartDate, EndDate, apis, containernames, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
		for _, apis := range perfusw1.Runtime {
			query := QueryRuntime(StartDate, EndDate, apis, filepaths, AppEnv)
			SendESRequest(apis, perfusw1map, query)
		}
	}

	for _, perfids := range result.IicsQaPerfids {
		var containernames2, filepaths2 string
		containernames2 = perfids.Containername
		filepaths2 = perfids.Logfilepath

		//log.Println(perfusw1, podname)
		for _, apisofservices := range perfids.AuthService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.BrandingService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.ContentRepository {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.IdsService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.MaService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservices := range perfids.ScimService {
			//log.Println(apisofservices, containernames, filepaths)
			query := buildmyquery(StartDate, EndDate, apisofservices, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservices, perfidsmap, query)
		}
		for _, apisofservice := range perfids.V3API {
			query := buildmyquery(StartDate, EndDate, apisofservice, containernames2, filepaths2, GlobalEnv)
			SendESRequest(apisofservice, perfidsmap, query)
		}
	}
	//fmt.Println(perfusw1map, perfidsmap)
	return perfidsmap, perfusw1map
}

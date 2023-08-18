package prodService

import (
	"NEWGOLANG/config"
	"NEWGOLANG/utils"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func test(w http.ResponseWriter, r *http.Request) {
	body := config.Body{ResponseCode: 200, Message: "OK"}
	jsonBody, err := json.Marshal(body)
	//if there is error in converting to json marshaling the below response will be sent
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//if no error then sent the response back to the port
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)
}

func compareResponseTimes(writer http.ResponseWriter, request *http.Request) {
	//get the values from the parameters in the POST request to variables

	predeploymentdates := request.URL.Query().Get("PreDeploymentDates")
	postdeploymentdates := request.URL.Query().Get("PostDeploymentDates")
	pod1 := request.URL.Query().Get("ApplicationNamespace")
	ids := request.URL.Query().Get("GlobalNamespace")
	RuntimeEnvironment := request.URL.Query().Get("Environment")
	emailID := request.URL.Query().Get("email")

	log.Println("Start Date is", predeploymentdates)
	log.Println("End Date is ", postdeploymentdates)
	log.Println("pod1 namespace is ", pod1)
	log.Println("ids namespace region is ", ids)
	log.Println("environment is ", RuntimeEnvironment)
	log.Println("emailID is ", emailID)

	condition1, beforedeployDates := utils.Validatetheinput(predeploymentdates, writer, request)
	if !condition1 {
		log.Println("Date format is wrong stopped execution of code")
		return
	}
	beforedeploymentStartDate := beforedeployDates[0]
	beforedeploymentEndDate := beforedeployDates[1]
	log.Println("Before deployment start Date is ", beforedeploymentStartDate)
	log.Println("Before deployment end date is ", beforedeploymentEndDate)
	condition2, AfterDeploymentDates := utils.Validatetheinput(postdeploymentdates, writer, request)
	if !condition2 {
		log.Println("Date format is wrong stopped execution of code")
		return
	}
	AfterDeploymentStartDate := AfterDeploymentDates[0]
	AfterDeploymentEndDate := AfterDeploymentDates[1]
	log.Println("After Deployment Start Date is", AfterDeploymentStartDate)
	log.Println("Before Deployment End Date is", AfterDeploymentEndDate)
	fileNameCheck := "./htmlReports/" + "TimePeriod_" + beforedeploymentStartDate + "_" + beforedeploymentEndDate + "to" + AfterDeploymentStartDate + "_" + AfterDeploymentEndDate + ".html"
	fileNameCheck = strings.ReplaceAll(fileNameCheck, ":", "")

	check := utils.CheckFileAlreadyExist(fileNameCheck, writer, request)
	if !check {
		log.Println("File is present")
		return
	}
	utils.RespondWithJSON("Your request is under processing please wait for 10-15 minutes and do not send the request again", 200, writer, request)

	oldIDS_Data, oldAPP_Pod_Data := utils.ReleaseData(beforedeploymentStartDate, beforedeploymentEndDate, ids, pod1, RuntimeEnvironment)
	fmt.Println(oldIDS_Data, oldAPP_Pod_Data)
	newIDS_Data, newAPP_Pod_Data := utils.ReleaseData(AfterDeploymentStartDate, AfterDeploymentEndDate, ids, pod1, RuntimeEnvironment)
	fmt.Println(newIDS_Data, newAPP_Pod_Data)

	if len(oldAPP_Pod_Data) == 0 || len(oldAPP_Pod_Data) == 0 || len(newIDS_Data) == 0 || len(newAPP_Pod_Data) == 0 {
		utils.RespondWithJSON("Error has been Occurred the Elastic Searched returned Data Null Please Check the input", 400, writer, request)
	}
	subject := fmt.Sprintf("Response Time Comparison Report Before and After Changes")
	p := fmt.Sprintf("<body style='background:White'><h3 style='background:#0790bd;color:#fff;padding:5px;text-align:center;border-radius:5px;'> Response Time Comparison Report Before and After Changes</h3> <br/> <br/>")
	var countapisforGlobal string
	var html_for_Global_Services string
	if len(oldIDS_Data) != 0 && len(newIDS_Data) != 0 {
		p = p + fmt.Sprintf("<div style='background:#80bfff;text-align:center'><p><b>Response Time Comparision between before Changes(%s to %s) and After Changes (%s to %s)</p> </b></div>", beforedeploymentStartDate, beforedeploymentEndDate, AfterDeploymentStartDate, AfterDeploymentEndDate)

		countapisforGlobal = fmt.Sprintf("<table style='background:#99ebff;border-collapse: collapse;' border = '2'cellpadding = '6'><tbody><tr><td colspan=4 style='text-align:center;background-color:Lavender;color:Black;'><b>Performance Summary for Global Services </b></td></tr><tr><th>Label</th><th>Range</th><th>Use case Count</th><th>Color Code</th></tr> ")
		html_for_Global_Services = fmt.Sprintf("<table style='background:#99ebff;;border-collapse: collapse;' border = '2' cellpadding = '6'><tbody><tr><td colspan=5 style='text-align:center;background-color:Lavender;color:Black;'><b> Response Time Comparison for Global Services </b></td></tr><tr><th>API</th><th>Before Changes</th><th>After Changes</th><th>Time Difference</th><th> %% Time Difference</th></tr> ")
		newidsdataSorted := utils.SortingMap(newIDS_Data)
		//log.Println(newidsdataSorted)
		var green int
		var yellow int
		var red int
		var total int
		var white int
		total = len(oldIDS_Data)
		yellow = 0
		red = 0
		green = 0
		white = 0

		for _, Label := range newidsdataSorted {
			oldidsdatavalues, oldidsapiname := oldIDS_Data[Label]
			newidsdatavalues, _ := newIDS_Data[Label]
			if oldidsapiname {
				if newidsdatavalues.Responsetime95 >= 0 {
					var timeOld float64
					var timeNew float64
					timeOld = oldidsdatavalues.Responsetime95
					timeNew = newidsdatavalues.Responsetime95
					diff := timeOld - timeNew

					if timeNew == 0 || timeOld == 0 {
						percDiffvalue := "No Hits for one of the selected Dates"
						white = white + 1
						html_for_Global_Services = html_for_Global_Services + "<tr style='background:White'><td>" + Label + "</td><td>" + strconv.FormatInt(int64(timeOld), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew), 10) + "</td><td>" + strconv.FormatInt(int64(diff), 10) + " </td><td>" + percDiffvalue + " </td></tr>"
					}
					percDiff := utils.CalcPerc(float64(diff), float64(timeOld))

					if percDiff < 0 && percDiff > -20 && timeOld !=0 && timeNew !=0 {
						yellow = yellow + 1
						html_for_Global_Services = html_for_Global_Services + "<tr style='background:Yellow'><td>" + Label + "</td><td>" + strconv.FormatInt(int64(timeOld), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew), 10) + "</td><td>" + strconv.FormatInt(int64(diff), 10) + " </td><td>" + strconv.FormatFloat(percDiff, 'f', 2, 64) + " %</td></tr>"
					}
					if percDiff <= -20 && !math.IsInf(percDiff, 0)&& timeOld !=0 && timeNew !=0 {
						red = red + 1
						html_for_Global_Services = html_for_Global_Services + "<tr style='background:Red'><td>" + Label + "</td><td>" + strconv.FormatInt(int64(timeOld), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew), 10) + "</td><td>" + strconv.FormatInt(int64(diff), 10) + " </td><td>" + strconv.FormatFloat(percDiff, 'f', 2, 64) + " %</td></tr>"
					}
					if percDiff >= 0 && timeOld !=0 && timeNew !=0 {
						green = green + 1
						html_for_Global_Services = html_for_Global_Services + "<tr style='background:Green'><td>" + Label + "</td><td>" + strconv.FormatInt(int64(timeOld), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew), 10) + "</td><td>" + strconv.FormatInt(int64(diff), 10) + " </td><td>" + strconv.FormatFloat(percDiff, 'f', 2, 64) + " %</td></tr>"
					}
				}
			}

		}

		countapisforGlobal = countapisforGlobal + "<tr><td colspan=2 style='text-align:center;color:Black;'>Total API Count</td><td>" + strconv.FormatInt(int64(total), 10) + "</td><td style='text-align:center;color:Black;'>-</td></tr>"
		countapisforGlobal = countapisforGlobal + "<tr><td>% Improvement</td><td> > 0 %</td><td>" + strconv.FormatInt(int64(green), 10) + "</td><td style='background-color: Green;'></td>"
		countapisforGlobal = countapisforGlobal + "<tr><td>% Degradation</td><td> 0 to 20 %</td><td>" + strconv.FormatInt(int64(yellow), 10) + "</td><td style='background-color: Yellow;'></td>"
		countapisforGlobal = countapisforGlobal + "<tr><td>% Degradation</td><td> > 20 %</td><td>" + strconv.FormatInt(int64(red), 10) + "</td><td style='background-color: Red;'></td>"
		if white > 0 {
			countapisforGlobal = countapisforGlobal + "<tr><td>No Hits for one of the Selected Day</td><td> > No Comparison</td><td>" + strconv.FormatInt(int64(white), 10) + "</td><td style='background-color: White;'></td>"

		}
		log.Println("totalAPIs are ", total, "improvement >0%", green, "degraded 0-20%", yellow, "degraded >20%", red)

	}
	var countapisforApp string
	var html_for_App_Services string

	if len(oldAPP_Pod_Data) != 0 && len(newAPP_Pod_Data) != 0 {
		countapisforApp = fmt.Sprintf("<table style='background:#99ebff;border-collapse: collapse;' border = '2'cellpadding = '6'><tbody><tr><td colspan=4 style='text-align:center;background-color:Lavender;color:Black;'><b>Performance Summary for Application Services </b></td></tr><tr><th>Label</th><th>Range</th><th>Use case Count</th><th>Color Code</th></tr> ")
		html_for_App_Services = fmt.Sprintf("<table style='background:#99ebff;;border-collapse: collapse;' border = '2' cellpadding = '6'><tbody><tr><td colspan=5 style='text-align:center;background-color:Lavender;color:Black;'><b> Response Time Comparison for Application Services </b></td></tr><tr><th>API</th><th>Before Changes</th><th>After Changes</th><th>Time Difference</th><th> %% Time Difference</th></tr> ")
		newAPP_POd_Sorted := utils.SortingMap(newAPP_Pod_Data)
		log.Println(newAPP_POd_Sorted)
		var green1 int
		var yellow1 int
		var red1 int
		var total1 int
		var white1 = 0
		total1 = len(oldAPP_Pod_Data)
		yellow1 = 0
		red1 = 0
		green1 = 0
		white1 = 0

		for _, Label1 := range newAPP_POd_Sorted {
			oldapppoddatavalues, oldapppodsapiname := oldAPP_Pod_Data[Label1]
			newapppoddatavalues, _ := newAPP_Pod_Data[Label1]
			if oldapppodsapiname {
				if newapppoddatavalues.Responsetime95 >= 0 {
					var timeOld1 float64
					var timeNew1 float64
					timeOld1 = oldapppoddatavalues.Responsetime95
					timeNew1 = newapppoddatavalues.Responsetime95
					diff1 := timeOld1 - timeNew1
					if timeNew1 == 0 || timeOld1 == 0 {
						percDiffvalue1 := "No Hits for one of the selected Dates"
						white1 = white1 + 1
						html_for_App_Services = html_for_App_Services + "<tr style='background:White'><td>" + Label1 + "</td><td>" + strconv.FormatInt(int64(timeOld1), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew1), 10) + "</td><td>" + strconv.FormatInt(int64(diff1), 10) + " </td><td>" + percDiffvalue1 + "</td></tr>"
					}
					percDiff1 := utils.CalcPerc(float64(diff1), float64(timeOld1))
					if percDiff1 < 0 && percDiff1 > -20 && timeOld1 !=0 && timeNew1 !=0 {
						yellow1 = yellow1 + 1
						html_for_App_Services = html_for_App_Services + "<tr style='background:Yellow'><td>" + Label1 + "</td><td>" + strconv.FormatInt(int64(timeOld1), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew1), 10) + "</td><td>" + strconv.FormatInt(int64(diff1), 10) + " </td><td>" + strconv.FormatFloat(percDiff1, 'f', 2, 64) + " %</td></tr>"
					}
					if percDiff1 <= -20 && !math.IsInf(percDiff1, 0) && timeOld1 !=0 && timeNew1 !=0 {
						red1 = red1 + 1
						html_for_App_Services = html_for_App_Services + "<tr style='background:Red'><td>" + Label1 + "</td><td>" + strconv.FormatInt(int64(timeOld1), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew1), 10) + "</td><td>" + strconv.FormatInt(int64(diff1), 10) + " </td><td>" + strconv.FormatFloat(percDiff1, 'f', 2, 64) + " %</td></tr>"
					}
					if percDiff1 >= 0 && timeOld1 !=0 && timeNew1 !=0 {
						green1 = green1 + 1
						html_for_App_Services = html_for_App_Services + "<tr style='background:Green'><td>" + Label1 + "</td><td>" + strconv.FormatInt(int64(timeOld1), 10) + "</td><td>" + strconv.FormatInt(int64(timeNew1), 10) + "</td><td>" + strconv.FormatInt(int64(diff1), 10) + " </td><td>" + strconv.FormatFloat(percDiff1, 'f', 2, 64) + " %</td></tr>"
					}
				}
			}
		}
		countapisforApp = countapisforApp + "<tr><td colspan=2 style='text-align:center;color:Black;'>Total API Count</td><td>" + strconv.FormatInt(int64(total1), 10) + "</td><td style='text-align:center;color:Black;'>-</td></tr>"
		countapisforApp = countapisforApp + "<tr><td>% Improvement</td><td> > 0 %</td><td>" + strconv.FormatInt(int64(green1), 10) + "</td><td style='background-color: Green;'></td>"
		countapisforApp = countapisforApp + "<tr><td>% Degradation</td><td> 0 to 20 %</td><td>" + strconv.FormatInt(int64(yellow1), 10) + "</td><td style='background-color: Yellow;'></td>"
		countapisforApp = countapisforApp + "<tr><td>% Degradation</td><td> > 20 %</td><td>" + strconv.FormatInt(int64(red1), 10) + "</td><td style='background-color: Red;'></td>"
		if white1 > 0 {
			countapisforApp = countapisforApp + "<tr><td>No Hits for one of the Selected Day</td><td>No Comparison</td><td>" + strconv.FormatInt(int64(white1), 10) + "</td><td style='background-color: White;'></td>"

		}
		log.Println("totalAPIs are ", total1, "improvement >0%", green1, "degraded 0-20%", yellow1, "degraded >20%", red1)
	}
	p = p + fmt.Sprintf("<br><br> %s </tbody></table><br><br> %s </tbody></table><br><br> %s </tbody></table><br><br> %s </tbody></table><br><br>", countapisforGlobal, countapisforApp, html_for_Global_Services, html_for_App_Services)
	utils.SendMail(p, subject, emailID)
	//write to file
	fileName := "./htmlReports/" + "TimePeriod_" + beforedeploymentStartDate + "_" + beforedeploymentEndDate + "to" + AfterDeploymentStartDate + "_" + AfterDeploymentEndDate + ".html"
	fileName = strings.ReplaceAll(fileName, ":", "")
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(p)
	if err != nil {
		log.Println(err)
	}

}

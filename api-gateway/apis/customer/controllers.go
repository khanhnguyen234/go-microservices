package customer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOrderManagementController() (interface{}, error) {
	var err error
	result := make(map[string]interface{})
	key := "a496b30ebfmsh5d640deede4844dp160118jsn9aa97b78e2e8"

	host := "covid-193.p.rapidapi.com"
	req, err := http.NewRequest("GET", "https://covid-193.p.rapidapi.com/statistics", nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	statistics, err := MakeRequest(req)
	result["statistics"] = statistics

	req, err = http.NewRequest("GET", "https://covid-193.p.rapidapi.com/countries", nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	countries, err := MakeRequest(req)
	result["countries"] = countries

	host = "covid-19-statistics.p.rapidapi.com"
	req, err = http.NewRequest("GET", "https://covid-19-statistics.p.rapidapi.com/reports/total?date=2020-04-07", nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	reportsTotal, err := MakeRequest(req)
	result["reportsTotal"] = reportsTotal

	req, err = http.NewRequest("GET", "https://covid-19-statistics.p.rapidapi.com/provinces?iso=CHN", nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	provinces, err := MakeRequest(req)
	result["provinces"] = provinces

	req, err = http.NewRequest("GET", "https://covid-19-statistics.p.rapidapi.com/regions", nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	regions, err := MakeRequest(req)
	result["regions"] = regions

	return result, err
}

func MakeRequest(req *http.Request) (interface{}, error) {
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return response, err
	}

	var result interface{}
	data, _ := ioutil.ReadAll(response.Body)
	stringJson := string(data)
	json.Unmarshal([]byte(stringJson), &result)

	return result, err
}

func ChannelGetFromApi(channel chan interface{}, key string, host string, url string) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", key)
	req.Header.Add("x-rapidapi-host", host)
	req.Header.Add("useQueryString", "true")
	result, _ := MakeRequest(req)
	channel <- result
	log.Println("Done ", url)
}

func GetOrderManagementControllerV2() (interface{}, error) {
	key := "a496b30ebfmsh5d640deede4844dp160118jsn9aa97b78e2e8"

	cRegions := make(chan interface{})
	go ChannelGetFromApi(cRegions, key, "covid-19-statistics.p.rapidapi.com", "https://covid-19-statistics.p.rapidapi.com/regions")
	log.Println("Start regions")

	cStatistics := make(chan interface{})
	go ChannelGetFromApi(cStatistics, key, "covid-193.p.rapidapi.com", "https://covid-193.p.rapidapi.com/statistics")
	log.Println("Start statistics")

	cReportsTotal := make(chan interface{})
	go ChannelGetFromApi(cReportsTotal, key, "covid-19-statistics.p.rapidapi.com", "https://covid-19-statistics.p.rapidapi.com/reports/total?date=2020-04-07")
	log.Println("Start reportsTotal")

	cCountries := make(chan interface{})
	go ChannelGetFromApi(cCountries, key, "covid-193.p.rapidapi.com", "https://covid-193.p.rapidapi.com/countries")
	log.Println("Start countries")

	cProvinces := make(chan interface{})
	go ChannelGetFromApi(cProvinces, key, "covid-19-statistics.p.rapidapi.com", "https://covid-19-statistics.p.rapidapi.com/provinces?iso=CHN")
	log.Println("Start provinces")

	statistics, countries, reportsTotal, provinces, regions := <-cStatistics, <-cCountries, <-cReportsTotal, <-cProvinces, <-cRegions

	result := map[string]interface{}{
		"statistics":   statistics,
		"countries":    countries,
		"reportsTotal": reportsTotal,
		"provinces":    provinces,
		"regions":      regions,
	}

	return result, nil
}

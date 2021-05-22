package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func doRequest(url string, data interface{}, method string) (content string) {
	jsonStr, err := json.Marshal(data)
	method = strings.ToUpper(method)
	baseApiPath := "https://api.trackingmore.com"
	apiVersion := "v3"
	requestUrl := baseApiPath + "/" + apiVersion + "/trackings/"+ url
	iprecords, _ := net.LookupIP("api.trackingmore.com")
	fmt.Println(iprecords)
	req, err := http.NewRequest(method, requestUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Tracking-Api-Key", "YOUR API KEY")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err2 := client.Do(req)
	if err2 != nil {
		panic(err2)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	return content
}

func main() {
	data := []byte(`""`)
	//Get realtime tracking results of a single tracking
	realtime :=  "realtime"
	postData := []byte(`[{"tracking_number": "EA152563254CN", "courier": "china-ems"},{"tracking_number": "EA152563254CN", "courier": "china-ems"}]`)
	res := doRequest(realtime, postData, "POST")
	fmt.Println(res)

	//count
	count :=  "count?courier=1&limit=100&created_at_min=1521314361&created_at_max=1541314361"
	res = doRequest(count, data, "GET")
	fmt.Println(res)

	//Get tracking results of a  tracking or List all trackings
	get :=  "get?page=1&limit=100&created_at_min=1521314361&created_at_max=1541314361"
	res = doRequest(get, data, "GET")
	fmt.Println(res)

	//Update Tracking item
	update :=  "modifyinfo"
	res = doRequest(update, postData, "PUT")
	fmt.Println(res)

	//archive
	archive :=  "archive"
	res = doRequest(archive, postData, "POST")
	fmt.Println(res)

	//Delete tracking item
	deleteApi :=  "delete"
	res = doRequest(deleteApi, postData, "DELETE")
	fmt.Println(res)

	//create  tracking number
	create :=  "create"
	res = doRequest(create, postData, "POST")
	fmt.Println(res)

	//manual update
	manualUpdate :=  "manualupdate"
	res = doRequest(manualUpdate, postData, "POST")
	fmt.Println(res)

	//remote tracking
	remote :=  "remote"
	res = doRequest(remote, postData, "POST")
	fmt.Println(res)

	//Get cost time iterm results
	costTime :=  "transittime"
	res = doRequest(costTime, postData, "POST")
	fmt.Println(res)

	//detect a carriers by tracking number
	detect :=  "detect"
	res = doRequest(detect, postData, "POST")
	fmt.Println(res)

	// get all carriers
	carriers :=  "courier"
	res = doRequest(carriers, postData, "POST")
	fmt.Println(res)

	//Get status number
	status :=  "status"
	res = doRequest(status, data, "GET")
	fmt.Println(res)

	//Set number not update
	notUpdate :=  "notupdate"
	res = doRequest(notUpdate, postData, "POST")
	fmt.Println(res)

	// Modify courier code
	modifyCarrier :=  "modifycourier"
	res = doRequest(modifyCarrier, postData, "PUT")
	fmt.Println(res)

	//Get user info
	user :=  "userinfo"
	res = doRequest(user, data, "GET")
	fmt.Println(res)
}


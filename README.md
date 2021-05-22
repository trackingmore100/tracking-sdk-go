Trackingmore-GO
=================

The GO SDK of Trackingmore API
## Official document

[Document](https://www.trackingmore.com/v3/api-index)

Init
--------------
set your apikey in req.Header
~~~
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
~~~

Quick Start
--------------
- Put your ApiKey in the request header.
- You can open sandbox model by add it in requestUrl: <code>requestUrl := baseApiPath + "/" + apiVersion + "/trackings/sandbox/"+ url</code>
- Most Api params receive multiple tracking numbers

**Get a list of the couriers in Trackingmore**

	carriers :=  "courier"
	res = doRequest(carriers, postData, "POST")
	fmt.Println(res)

**Detect which couriers defined in your account match a tracking number**

	detect :=  "detect"
	res = doRequest(detect, postData, "POST")
	fmt.Println(res)


**Post trackings to your account**

    //Create single tracking numbers
    postData := []byte(`{"tracking_number": "EA152563254CN", "courier": "china-ems"}`)
    //Create multiple tracking numbers
    postData := []byte(`[{"tracking_number": "EA152563254CN", "courier": "china-ems"},{"tracking_number": "EA152563254CN", "courier": "china-ems"}]`)
    create :=  "create"
	res = doRequest(create, postData, "POST")

**Summary of Connection API Methods with all the api and Methods**

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

## Typical Server Responses

We will respond with one of the following status codes.

Code|Type | Message
----|--------------|-------------------------------
200    | <code>Success</code>|    Request response is successful
203    | <code>PaymentRequired</code>|  API service is only available for paid account Please subscribe paid plan to unlock API services                                                             ul
204    | <code>No Content</code>|    Request was successful, but no data returned Tracking NO. or target data possibly do not exist
400    | <code>Bad Request</code>| Request type error Please check the API documentation for the request type of this API
401    | <code>Unauthorized</code>|    Authentication failed or has no permission Please check and ensure your API Key is correct
403    | <code>Bad Request</code>|    Page does not exist Please check and ensure your link is correct                                                                                             ul
404    | <code>Not Found</code>|    Page does not exist Please check and ensure your link is correct
408    | <code>Time Out</code>|    Request timeout The official website did not return data, please try again later
411    | <code>Bad Request</code>|    Specified request parameter length exceeds length limit Please check and ensure that the request parameters are of the required length
412    | <code>Bad Request</code>|    Specified request parameter format doesn't meet requirements Please check and ensure that the request parameters are in the required format
413    | <code>Out limited</code>|    The number of request parameters exceeds the limit Please check the API documentation for the limit of this API
417    | <code>Bad Request</code>|    Missing request parameters or request parameters cannot be parsed Please check and ensure that the request parameters are complete and correctly formatted
421    | <code>Bad Request</code>|    Some of required parameters are empty Some couriers need special parameters to track logistics information (Special Couriers)
422    | <code>Bad Request</code>|    Unidentifiable courier code Please check and ensure that the courier code are correct(Courier code)
423    | <code>Bad Request</code>|    Tracking No. already exists
424    | <code>Bad Request</code>|    Tracking No. no exists Please use 「Create trckings」 API first to create trackings
429    | <code>Bad Request</code>|    Exceeded API request limits, please try again later Please check the API documentation for the limit of this API
511    | <code>Server Error</code>|    Server error Please contact us: service@trackingmore.org.
512    | <code>Server Error</code>|    Server error Please contact us: service@trackingmore.org.
513    | <code>Server Error</code>|    Server error Please contact us: service@trackingmore.org.        
package main

import (
	"fmt"
	echo "httpserver/golang/protobuf"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/protobuf/proto"
)

// Handles the root page requests.
func rootPage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("\nHitted rootPage")
	r := &echo.RootPage{
		Status: "Success",
		Msg:    "Welcome to the root page of the Echo Server",
	}
	responseByte, err := proto.Marshal(r)
	if err != nil {
		log.Fatal("Errored while Marshaling")
	}
	fmt.Printf("Request Headers : %v\n", req.Header)
	fmt.Printf("Response Bytes : %v\n", responseByte)
	w.Write(responseByte)
}

// Handles the echo page requests. Only the GET request will get the status success
// other request will get the status failed
func echoPage(w http.ResponseWriter, request *http.Request) {
	fmt.Println("\nHitted echoPage")
	var (
		echoMethod      echo.HttpMethods
		responseStatus  string
		responseStsDesc string
	)

	if request.Method != "GET" {
		switch request.Method {
		case "POST":
			echoMethod = echo.HttpMethods_HTTP_METHOD_POST
		case "PUT":
			echoMethod = echo.HttpMethods_HTTP_METHOD_PUT
		case "DELETE":
			echoMethod = echo.HttpMethods_HTTP_METHOD_DELETE
		default:
			echoMethod = echo.HttpMethods_HTTP_METHOD_UNSPECIFIED
		}
		responseStatus = "Failed"
		responseStsDesc = "The Request method is incorrect"
	} else {
		responseStatus = "Success"
		responseStsDesc = "The Request method is correct"
		echoMethod = echo.HttpMethods_HTTP_METHOD_GET
	}

	body, _ := ioutil.ReadAll(request.Body)
	headers := httpHeaderToMap(request)
	queryParams := httpQueryToPBQuery(request)
	r := &echo.EchoResponse{
		Status:          responseStatus,
		StatusDesc:      responseStsDesc,
		EchoHttpMethod:  echoMethod,
		EchoHttpHeaders: headers,
		EchoQueryParams: &echo.QueryParameter{QueryParams: queryParams},
		EchoPayload:     string(body),
	}
	responseByte, err := proto.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(responseByte)
}

// Creates the map from the http request headers
func httpHeaderToMap(req *http.Request) map[string]string {
	fmt.Printf("Request Headers : %v\n", req.Header)
	accept := req.Header.Get("Accept")
	accEncoding := req.Header.Get("Accept-Encoding")
	connection := req.Header.Get("Connection")
	contentLen := req.Header.Get("Content-Length")
	userAgent := req.Header.Get("User-Agent")

	return map[string]string{
		"Accept": accept, "Accept-Encoding": accEncoding,
		"Connection": connection, "User-Agent": userAgent,
		"Content-Length": contentLen}

}

// creates the map from the http request query params
func httpQueryToPBQuery(req *http.Request) map[string]string {
	fmt.Printf("Request Query Parameters : %v\n", req.URL.Query())
	queryKey1 := req.URL.Query().Get("query-key-1")
	queryKey2 := req.URL.Query().Get("query-key-2")
	return map[string]string{"query-key-1": queryKey1, "query-key-2": queryKey2}
}

func main() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/echo", echoPage)

	http.ListenAndServe(":8090", nil)

}

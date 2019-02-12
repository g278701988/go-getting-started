package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func getGithubConfig() string {

	resp, err := http.Get("https://raw.githubusercontent.com/g278701988/go-getting-started/master/config.txt")
	if err != nil {
		log.Printf("Get err%v", err)
		return string("")
	}
	var jsData []Jsdata
	if err := json.NewDecoder(resp.Body).Decode(&jsData); nil != err {
		log.Printf("Decode err%v", err)
		return string("")
	}

	return fmt.Sprintf("%v:%v\n%v:%v ", jsData[0].Key, jsData[0].Value, jsData[1].Key, jsData[1].Value)
}

func displayIP(responseWriter http.ResponseWriter, request *http.Request) {

	ip2 := getGithubConfig()
	fmt.Fprintf(responseWriter, "%s\t(data will update when ip changed)", ip2)

}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", displayIP)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	err = server.ListenAndServe()
	if nil != err {
		log.Fatal(err)
	}

	log.Println("web end")
}

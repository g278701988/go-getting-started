package main

import (
	"fmt"	
	"log"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

var configFileName string
func getGithubConfig() string {

	resp, err := http.Get("https://raw.githubusercontent.com/g278701988/go-getting-started/master/config.txt")
	if err != nil {
		log.Printf("Get err%v", err)
		return string("")
	}
	var jsData Jsdata
	if err := json.NewDecoder(resp.Body).Decode(&jsData); nil != err {
		log.Printf("Decode err%v", err)
		return string("")
	}
	return jsData.Value
}

func displayIP(responseWriter http.ResponseWriter, request *http.Request) {

	var jsonData Jsdata
	ModifyConfig(configFileName, &jsonData, true)
	//log.Println(jsonData)

	ip2 := getGithubConfig()
	fmt.Fprintf(responseWriter, "ip:%s\nip:%s\n", jsonData.Value,ip2)

}
func updateIP(responseWriter http.ResponseWriter, request *http.Request) {

	ip, ok := request.URL.Query()["ip"]
	if ok {
		var jsonData = Jsdata{Key: "ip", Value: ip[0]}
		ModifyConfig(configFileName, &jsonData, false)
		//log.Println(jsonData)
		fmt.Fprintf(responseWriter, "%s!", "ok")

	}

}
func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
		return
	}
	Path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if nil == err {
		log.Println(Path)
		log.Println("web begin")
		ConfigName := "config.txt"
		var jsonData = Jsdata{Key: "ip", Value: ""}
		configFileName = Path + "/" + ConfigName
		ModifyConfig(Path+"/"+ConfigName, &jsonData, false)
		log.Println(jsonData)
		mux := http.NewServeMux()
		mux.HandleFunc("/updateip", updateIP)
		mux.HandleFunc("/", displayIP)
		server := &http.Server{
			Addr:    addr,
			Handler: mux,
		}
		err = server.ListenAndServe()
		if nil != err {
			log.Fatal(err)
		}

	}

	log.Println("web end")
}

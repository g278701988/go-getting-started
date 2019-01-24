package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

var configFileName string
func getPublicIP2() string {

	resp, err := http.Get("http://206.189.91.42/")
	if err != nil {
		//log.Fatalln(err)
		return string("")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatalln(err)
		return string("")
	}
	return string(body)
}

func displayIP(responseWriter http.ResponseWriter, request *http.Request) {

	var jsonData Jsdata
	ModifyConfig(configFileName, &jsonData, true)
	//log.Println(jsonData)

	//ip2 := getPublicIP2()
	fmt.Fprintf(responseWriter, "ip:%s\n", jsonData.Value)

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

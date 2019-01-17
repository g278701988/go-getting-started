package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var configFileName string

func displayIP(responseWriter http.ResponseWriter, request *http.Request) {

	var jsonData Jsdata
	ModifyConfig(configFileName, &jsonData, true)
	//log.Println(jsonData)

	fmt.Fprintf(responseWriter, "ip:%s", jsonData.Value)

}
func updateIP(responseWriter http.ResponseWriter, request *http.Request) {

	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {

		fmt.Fprintf(responseWriter, "userip: %q is not IP:port", request.RemoteAddr)
	}
	var jsonData = Jsdata{Key: "ip", Value: ip}
	ModifyConfig(configFileName, &jsonData, false)
	//log.Println(jsonData)
	fmt.Fprintf(responseWriter, "%s!", "ok")

}
func main() {

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
			Addr:    "0.0.0.0:8080",
			Handler: mux,
		}
		err = server.ListenAndServe()
		if nil != err {
			log.Fatal(err)
		}

	}

	log.Println("web end")
}

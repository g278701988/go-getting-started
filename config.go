package main

import (
	"encoding/json"
	"log"
	"os"
)

//Jsdata for data
type Jsdata struct {
	Key   string
	Value string
}

//ModifyConfig  Value From fileName
//https://blog.golang.org/json-and-go
func ModifyConfig(fileName string, jsonData *Jsdata, isRead bool) {

	flag := os.O_CREATE | os.O_RDONLY
	if false == isRead {
		flag = os.O_CREATE | os.O_WRONLY
		err := os.Remove(fileName)
		if nil == err {
			log.Printf("remove %s ok \n", fileName)
		} else {
			log.Println(err)
		}

	}

	file, err := os.OpenFile(fileName, flag, 0755)
	if err == nil {

		defer file.Close()

		log.Printf("open file \"%s\" sucessed\n", fileName)

		if false == isRead {
			writeConfig(file, jsonData)
		} else {
			readConfig(file, jsonData)

		}

	} else {
		log.Fatal(err)
	}
}

// read request
func readConfig(file *os.File, jsonData *Jsdata) bool {
	ret := true
	dec := json.NewDecoder(file)

	if err := dec.Decode(&jsonData); nil != err {
		ret = false
	}

	// nCount, err := file.Read(data)
	// if nil == err && nCount <= 1024 {
	// 	err = json.Unmarshal(data[:nCount], &rdata)
	// 	if nil == err {
	// 		log.Println(rdata)
	// 	} else {
	// 		log.Fatal(err)
	// 	}

	// } else {
	// 	log.Fatal(err)
	// 	ret = false
	// }
	return ret
}

//only save the config  have changed.
//Each page  per config file
func writeConfig(file *os.File, jsonData *Jsdata) bool {
	ret := true
	data, err := json.Marshal(jsonData)
	if nil == err {
		_, err = file.Write(data)
		if nil == err {
			log.Println("write successed")
		} else {
			log.Fatal(err)
		}

	}
	return ret
}

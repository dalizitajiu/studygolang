package main

import (
	"log"
	"flag"
	"fmt"
	"encoding/json"
)

var data = flag.String("line","[1,2,3,4]","参数")

func main(){
	flag.Parse()
	log.Println("[接收到的参数]",*data)
	var testres = map[string]interface{}{}
	testres["line"] = []float64{12,13}
	resbyte,_:=json.Marshal(testres)
	fmt.Println(string(resbyte))
}
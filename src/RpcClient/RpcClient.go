package main

import (
	"net/rpc/jsonrpc"
	"log"
)

func main()  {
	rpcClient,err:=jsonrpc.Dial("tcp","127.0.0.1:9009")
	if err!=nil{
		log.Fatalln(err.Error())
	}
	var reply  = map[string]interface{}{}
	var args = map[string]interface{}{}
	args["line"]=[]int{4,5,16}
	err=rpcClient.Call("Rpc.Handle1",args,&reply)

	if err!=nil{
		log.Fatalln(err.Error())
	}
	log.Println(reply)

}

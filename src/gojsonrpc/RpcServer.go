package main

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

type RpcFunction func(*map[string]interface{},*map[string]interface{}) 
type Rpc struct {

}
func (this *Rpc)Handle1(args *map[string]interface{},reply *map[string]interface{})error{
	temp:=(*args)["line"].([]interface{})
	res:=make([]float64,0)
	for _,v:=range temp{
		res=append(res,v.(float64)*2)
	}
	*reply = map[string]interface{}{}
	(*reply)["res"] = res
	return nil
}
func main(){
	log.Println("JsonRpcServer")
	rpchandle:=new(Rpc)
	rpc.Register(rpchandle)
	addr, _ := net.ResolveTCPAddr("tcp", ":9009")
	ln, e := net.ListenTCP("tcp", addr)
	if e != nil {
		panic(e)
	}
	for {
		conn, e := ln.Accept()
		log.Println("[连接上了]")
		if e != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
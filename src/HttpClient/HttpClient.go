package HttpClient

import (
	"github.com/smallnest/goreq"
	"time"
	"fmt"
	"strings"
	"log"
	"encoding/json"
)

type Client struct {
	Prefix string
	Body []byte
	Timeout int
}

type Resp struct{
	StatusCode int
	RespBody []byte
}


func NewClient(prefix string,timeout int)*Client{
	res:=&Client{}
	res.Prefix = prefix
	res.Timeout = timeout
	return res
}
func (client *Client)Get(path string,query map[string]string)Resp{
	resp,body,err:=goreq.New().Timeout(3*time.Second).Get(client.Prefix+path+"?"+map2string(query)).End()
	if err!=nil{
		log.Fatalln(err)
		return Resp{resp.StatusCode,[]byte(body)}
	}
	return Resp{resp.StatusCode,[]byte(body)}
}

func (client *Client)PostForm(path string,body map[string]string)Resp{
	resp,body2,err:=goreq.New().Timeout(3*time.Second).ContentType("urlencoded").Post(client.Prefix+path).SendRawString(map2string(body)).End()
	if err!=nil{
		log.Fatalln(err)
		return Resp{resp.StatusCode,[]byte(body2)}
	}
	return Resp{resp.StatusCode,[]byte(body2)}
}

func (client *Client)PostJson(path string,body []byte)Resp{
	resp,body2,err:=goreq.New().Timeout(3*time.Second).ContentType("json").Post(client.Prefix+path).SendRawString(string(body)).End()
	if err!=nil{
		log.Fatalln(err)
		return Resp{resp.StatusCode,[]byte(body)}
	}
	return Resp{resp.StatusCode,[]byte(body2)}
}
func (client *Client)PostJsonObj(path string,obj interface{})Resp{
	return client.PostJson(path,obj2json(obj))
}
func obj2json(obj interface{})[]byte{
	res,err:= json.Marshal(obj)
	if err!=nil{
		log.Fatalln(err)
		return nil
	}
	return res
}

func map2string(query map[string]string)string{
	if query==nil{
		return ""
	}
	res:=make([]string,0)
	for k,v:=range query{
		res=append(res,fmt.Sprintf("%s=%s",k,v))
	}
	return strings.Join(res,"&")
}


package Job

import (
	"net/http"
	"context"
	"log"
	"io/ioutil"
	"github.com/tidwall/gjson"
)

const (
	Get = iota
	Post
	Patch
	Delete
	PostJson
)

type HttpResult struct {
	Status int
	Result gjson.Result
}
func NewHttpResult(status int,result gjson.Result)*HttpResult{
	res:=&HttpResult{}
	res.Status = status
	res.Result = result
	return res
}

type HttpJob struct {
	Job
	req *http.Request
	result *HttpResult
}
func NewHttpJob(ctx context.Context,cancel context.CancelFunc,req *http.Request)*HttpJob{
	res:=&HttpJob{}
	res.Job = *NewJob(ctx,cancel)
	res.req = req
	res.req.WithContext(ctx)
	return res
}
func (hjob *HttpJob)Start(){
	go func() {
		hjob.Wait()
	}()

	client:=http.DefaultClient
	resp,err:=client.Do(hjob.req)
	defer resp.Body.Close()
	hjob.Status = Running
	if err!=nil{
		hjob.Status = Failed
		log.Fatalln(err)
		hjob.Cancel()
		return
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		hjob.Cancel()
		hjob.Status=Failed
		return
	}
	hjob.result =NewHttpResult(resp.StatusCode,gjson.ParseBytes(body))
	hjob.Status = Complecated
	hjob.SetCanGoOn(true)
}
func(hjob *HttpJob)GetResult()*HttpResult{
	return hjob.result
}

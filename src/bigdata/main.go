package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sync"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"encoding/csv"
)

func fun() {
	fmt.Println("ok")
}

type Functional interface {
	Map(func(interface{}) interface{}) *Stream
	Filter(func(interface{}) interface{}) *Stream
	Reduce(func(interface{}, interface{}) interface{}) interface{}
}
type Stream struct {
	Iterator []interface{}
	Count    int
}

func ToSlice(arr interface{}) ([]interface{}, int) {
	v := reflect.ValueOf(arr)
	var l int
	var ret []interface{}
	if v.Kind() == reflect.Slice {
		l = v.Len()
		ret = make([]interface{}, l)
		for i := 0; i < l; i++ {
			ret[i] = v.Index(i).Interface()
		}
		return ret, l
	}

	res, ok := arr.(*list.List)
	if ok {
		ret = make([]interface{}, res.Len())
		count := 0
		for e := res.Front(); e != nil; e = e.Next() {
			ret[count] = e.Value
			count++
		}
		return ret, count
	}

	return nil, 0
}
func NewStream(sizable interface{}) *Stream {
	res := &Stream{}
	typename := reflect.TypeOf(sizable)

	typenamestr := typename.String()
	fmt.Println(typenamestr)
	itemslice, count := ToSlice(sizable)
	res.Count = count
	res.Iterator = itemslice
	log.Println(res)
	return res
}
func NewEmptyStream() *Stream {
	res := &Stream{}
	res.Count = 0
	res.Iterator = make([]interface{}, 0)
	return res
}

type InterfacesMapper func(item interface{}) interface{}
type InterfacesFilter func(item interface{}) bool
type InterfacesReducer func(a, b interface{}) interface{}

func (stream *Stream) Map(mapfunc InterfacesMapper) *Stream {
	log.Println("[Map]")
	res := NewEmptyStream()
	res.Iterator = make([]interface{}, stream.Count)
	res.Count = stream.Count
	var wg sync.WaitGroup
	wg.Add(stream.Count)
	for k, v := range stream.Iterator {
		copyk := k
		copyv := v
		go func() {
			// log.Println(k,`` v)
			(res.Iterator)[copyk] = mapfunc(copyv)
			wg.Done()
		}()
	}
	wg.Wait()
	return res
}
func (stream *Stream) Filter(filterfunc InterfacesFilter) *Stream {
	log.Println("[Filter]")
	res := NewEmptyStream()
	resarray := make([]interface{}, 0)
	res.Count = stream.Count
	for _, v := range stream.Iterator {
		if filterfunc(v) {
			resarray = append(resarray, v)
		}
	}
	res.Iterator = resarray
	return res
}

func (stream *Stream) Reduce(identity interface{}, reducefunc InterfacesReducer) interface{} {
	res := identity
	for _, v := range stream.Iterator {
		res = reducefunc(res, v)
	}
	return res
}

func (stream *Stream) Show() {
	log.Println("[Show]")
	for _, v := range stream.Iterator {
		fmt.Println(v)
	}
}
func (stream *Stream) ToList() []interface{} {
	log.Println("[ToList]")
	return stream.Iterator
}

func main() {
	// testlist2 := []string{"daliz", "hello3333", "world1234444"}
	testlist2 := list.New()
	testlist2.PushBack("dalizi")
	testlist2.PushBack("hello3333")
	testlist2.PushBack("world12345")

	stream := NewStream(testlist2)
	stream2 := stream.Map(func(a interface{}) interface{} {
		var temp interface{}
		temp = a.(string) + "llll"
		// panic("what is wrong")
		return temp
	}).Filter(func(a interface{}) bool {
		return len(a.(string)) > 9
	}).Reduce("", func(a interface{}, b interface{}) interface{} {
		var res interface{}
		res = a.(string) + b.(string)

		return res
	})
	log.Println(stream2)

	var temp interface{}
	testlist := list.New()
	temp = testlist
	_, ok := temp.(*list.List)
	if ok {
		log.Println("equal")
	}
	filepath := "E://jmeter.log"
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatalln("open file failed")
	}
	br := bufio.NewReader(file)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		log.Println(line)
	}
	list := arraylist.New()
	list.Add(1)
	list.Add(4)
	list.Add(2)
	log.Println(list.Contains(4))
	list.Sort(utils.IntComparator)
	log.Println(list)
	tempfile,err:=os.Open("C:\\\\Users\\Administrator\\Desktop\\gobigdata\\src\\testdata\\test1.csv")
	tempreader:=csv.NewReader(tempfile)
	tempreader.TrimLeadingSpace=true
	tempreader.Comment='#'

	res,err:=tempreader.ReadAll()
	log.Println(res[3][1])
	defer tempfile.Close()

	tempfile2,err:=os.Open("C:\\\\Users\\Administrator\\Desktop\\gobigdata\\src\\testdata\\test2.tsv")
	tempreader2:=csv.NewReader(tempfile2)
	//tempreader.TrimLeadingSpace=true
	//tempreader.Comment='#'
	tempreader.Comma= ';'

	res2,err:=tempreader2.ReadAll()
	if err!=nil{
		log.Panicln(err)
	}
	log.Println(res2)
	defer tempfile2.Close()


}

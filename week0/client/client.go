package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	h := &hello{
		endpoint: "http://localhost:8080/",
	}
	//msg, _ := h.SayHello("golang")
	//fmt.Println(msg)
	SetFuncField(h)
	msg,_ := h.Funcfield("golang")
	fmt.Println(msg)
}
type hello struct{
	endpoint string
	//只能改这个
	Funcfield func(name string)(string,error)
}
func SetFuncField(val interface{}){
	// 反射 reflection
	// t := reflect.TypeOf(va)

	v := reflect.ValueOf(val)  //这是指针的反射
	ele := v.Elem()   //拿到了指针指向的结构体
	t := ele.Type()   //拿到了指针指向的结构体的类型信息

	num := t.NumField()
	for i := 0; i < num; i++ {
		f := ele.Field(i)

		if f.CanSet(){
			fn := func(args []reflect.Value) (results []reflect.Value){
				name := args[0].Interface().(string)

				client := http.Client{}
				resp, err := client.Get("http://localhost:8080/"+name)
				if err != nil{
					print(err.Error())
					return []reflect.Value{reflect.ValueOf(""),reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil{
					print(err.Error())
					return []reflect.Value{reflect.ValueOf(""),reflect.ValueOf(err)}
				}

				fmt.Printf("%s\n",resp.Status)

				fmt.Printf(string(data))

				return []reflect.Value{reflect.ValueOf(string(data)),reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}

			f.Set(reflect.MakeFunc(f.Type(),fn))

			fmt.Println("aaa")
		}
	}
}

type HelloService interface {
	SayHello(name string) (string, error)
}
//改不了这个
func (h hello) SayHello(name string) (string, error) {
	return "", nil
}

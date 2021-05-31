package main

import (
	"errors"
	"fmt"
)

func main() {
	values := []interface{}{0,1,2,3,5}
	newValues, err:= Add(values,4,4);
	//ouput 0,1,2,3,4,5

	if err!= nil{
		fmt.Print(err.Error())
		return
	}

	for i, value := range newValues {
		fmt.Printf("index: %d, value: %d\n",i,value)
	}

	fmt.Println("")
	newValues = Delete(newValues,3)

	for i, value := range newValues {
		fmt.Printf("index: %d, value: %d\n",i,value)
	}

}


func Add(values []interface{}, val interface{},index int)([]interface{},error)  {
		res :=[]interface{}{}

	if index < 0 || index > len(values){
		return nil,errors.New("index illegal!")
	}
	for i := 0; i < index; i++ {
		res = append(res,values[i])
	}

	res = append(res,val)

	for i := index; i < len(values); i++ {
		res = append(res,values[i])
	}

	return res,nil
}

func Delete(values []interface{},index int) []interface{} {
	if index < 0 || index > len(values){
		return values
	}
	res := []interface{}{}

	for i := 0; i < len(values); i++ {
		if i == index {
			continue
		}
		res = append(res,values[i])
	}
	return res
}
package main

import (
	"fmt"
	"reflect"
)

func cstore(arr interface{}, ans *[]interface{}){ // function to store a given interface into our ans
	arr1:=reflect.ValueOf(arr)
	switch arr1.Kind(){
	case reflect.Slice:
		if arr1.Len() == 0 { // if empty slice break
			break
		} else {
			for i:=0;i<arr1.Len();i++{
				*ans=append(*ans,arr1.Index(i).Interface())
			}
			
		}
	default:
		*ans = append(*ans, arr1.Interface())

	}

}

func store(arr1 reflect.Value,arr2 reflect.Value, ans *[]interface{}){ // function to append two slice into our ans where are arr2 is smaller
	
		i:=0
		for i = 0; i < arr2.Len(); i++ {
			
			*ans = append(*ans, arr1.Index(i).Interface())
			*ans = append(*ans, arr2.Index(i).Interface())
		}

		for j:=i;j<arr1.Len();j++{
			*ans=append(*ans,arr1.Index(j).Interface())
		}

}

func merge(arr interface{}, nn interface{}) (interface{}, error) {

	ans := []interface{}{} //for storing answer
	if arr== nil && nn== nil{ // if both are nil return an empty array of interface
		return ans,nil
	}
	if arr== nil || nn==nil{ // if one of them is nil, the other one is sent to the cstore 
		if arr==nil{
			cstore(nn,&ans)
			}else{
				cstore(arr,&ans)
			}
		return ans,nil
	}
	sl:=false
	arr1 := reflect.ValueOf(arr) // value of first parameter
	switch arr1.Kind() {
	case reflect.Slice:
		if arr1.Len() == 0 { // if empty slice break
			break
		} else {
			sl=true // true indicating the first one is a slice
			
		}
	default:
		ans = append(ans, arr1.Interface()) // for other values simply append them
	}
	// doing the same for the second interface
	arr2 := reflect.ValueOf(nn)

	switch arr2.Kind() {
	case reflect.Slice: // if both are slice call store function epending on whose length is more
		if sl{
			if arr1.Len()>arr2.Len(){
				store(arr1,arr2,&ans)
			}else{
				store(arr2,arr1,&ans)
			}
			sl=false;// sl is set false to indicate that we have stored arr slice
		}else{
			cstore(nn,&ans) // if arr was not a slice store nn
		}
	default:
		ans = append(ans, arr2.Interface())
	}
	if sl{ // if not stored than store arr slice in ans
		cstore(arr,&ans)
	}

	return ans, nil

}

func main() {

	// arr1 := []int{1, 2, 3}
	// arr2:=[]interface{}{4,5,6}
	// n1:= "77660029"
	// n2:= []string{"1234567", "4354654746"}
	n3:=[]string{"77660029","1111"}
	n4:="a"

	a, _ := merge(n3, n4)

	fmt.Println(a)

}

package main

import (
	"fmt"
	"reflect"
)

//约定如果包含map，必须是map[string]string,如果是切片，必须是[]string
func ReadabilityMap(amap interface{}) interface{} {
	if amap == nil {
		return nil
	}
	v_amap := reflect.ValueOf(amap)
	if v_amap.Kind() == reflect.Ptr {
		v_amap = v_amap.Elem()
	}
	for i := 0; i < v_amap.NumField(); i++ {
		fmt.Println(v_amap.Field(i).Kind(), v_amap.Field(i))
		field := v_amap.Field(i)
		if field.Kind() == reflect.String {
			if field.String() == "" {
			}
		}
	}
	return nil
}

type test struct {
	a string
	b int
	c []string
	d map[string]string
}

func main() {
	t := test{
		a: "",
		b: 1,
		c: []string{"", "asd"},
		d: map[string]string{
			"first":  "asd",
			"second": "",
		},
	}
	ReadabilityMap(t)
	fmt.Println(t)
}

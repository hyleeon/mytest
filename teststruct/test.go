package main

import (
	"strconv"
	"reflect"
	"fmt"
)

// InnerSt ....
type InnerSt struct {
	ix int;
	str string;
}

//St is test struct
type St struct {
	key string	`required:"true"`
	value string
	ifll int
	c byte
	//
	ist  InnerSt `required:"true"` ;
}

//Copy to copy test
// type Copy interface {
// 	copy(st *St) St;
// }

func (st *St) copy() *St{
	var rst St
	rst.key = st.key
	rst.value = st.value
	return &rst;
}


func (st *St) copy4(st2 St) *St{
	var rst St
	rst.key = st.key
	rst.value = st.value
	return &rst;
}


func copy2(st *St) *St{
	var rst St
	rst.key = st.key
	rst.value = st.value
	return &rst;
}

func copy3(st St) *St{
	var rst St
	rst.key = st.key
	rst.value = st.value
	return &rst;
}



func main() {

	//var pst *St

	st := St{key:"abc", value:"hahaha"}

	//pst = &st

	rst := copy3(st)

	rst = copy2(&st)

	rst = st.copy()

	rst = st.copy4(st)

	fmt.Printf("rst: %s=%s\n", rst.key, rst.value);

	var v reflect.Value;
	v = reflect.ValueOf(st);
	fmt.Println("v:", v);
	fmt.Println("v.Kind:", v.Kind());

	var t reflect.Type;
	t = v.Type();

	fmt.Println("v.NumField(), t.NumField():", v.NumField(), t.NumField());

	for i :=0; i < t.NumField(); i++ {
		var f reflect.StructField;
		f = t.Field(i);
		fmt.Println(strconv.Itoa(i) + " f:", f.Tag);
	}
}






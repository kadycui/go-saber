package main

import (
	"reflect"
	"testing"
)

type data struct {
	N int
}

const AssignTimes = 100000000

func BenchmarkNativeAssign(t *testing.B) {
	v := data{N: 1}
	for i := 0; i < AssignTimes; i++ {
		v.N = 3
	}

}

func BenchmarkReflectAssign(t *testing.B) {
	v := data{N: 1}
	vv := reflect.ValueOf(&v).Elem()
	for i := 0; i < AssignTimes; i++ {
		vv.FieldByName("N").SetInt(3)
	}
}

func foo(v int) {

}

const CallTimes = 100000000

func BenchmarkNativeCall(t *testing.B) {
	for i := 0; i < CallTimes; i++ {
		foo(i)
	}
}

func BenchmarkReflectCall(t *testing.B) {
	v := reflect.ValueOf(foo)
	for i := 0; i < CallTimes; i++ {
		v.Call([]reflect.Value{reflect.ValueOf(2)})
	}
}

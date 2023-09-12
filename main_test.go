package main

import (
	"testing"
)

func BenchmarkLargeMapLength(b *testing.B) {
	myMap := make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		myMap[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = len(myMap)
	}
}

func BenchmarkLargeArrayLength(b *testing.B) {
	type MyStruct struct {
		Field1 int
		Field2 string
	}

	myStructArray := make([]MyStruct, 10000)
	for i := 0; i < 10000; i++ {
		myStructArray[i] = MyStruct{i, "Item"}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = len(myStructArray)
	}
}


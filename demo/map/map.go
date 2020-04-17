package main

import (
	"fmt"
	"github.com/chentaihan/container/hashmap"
)

func main() {
	tests := []struct {
		key   string
		value string
	}{
		{
			"key1",
			"value1",
		},
		{
			"key2",
			"value2",
		},
		{
			"key3",
			"value3",
		},
		{
			"key4",
			"value4",
		},
		{
			"key5",
			"value5",
		},
	}
	sm := hashmap.NewMapSync()
	for _, test := range tests {
		sm.Set(test.key, test.value)
	}
	for _, test := range tests {
		value, _ := sm.Get(test.key)
		if value.(string) != test.value {
			fmt.Println("equal ", test.key, value.(string), test.value)
		}
		if !sm.Exist(test.key) {
			fmt.Println("exist ", test.key, test.value)
		}
	}
	if sm.Len() != len(tests) {
		fmt.Println("len: ", sm.Len(), len(tests))
	}
	if sm.Exist("asdfghjtre") {
		fmt.Println("exist ", "asdfghjtre")
	}
	data, _ := sm.Marshal()
	fmt.Println(string(data))
	fmt.Println("success")

	dataString := `{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5","key6":"value6"}`
	err := sm.Unmarshal([]byte(dataString))
	if err != nil {
		fmt.Println(err)
	}
	for _, test := range tests {
		if !sm.Exist(test.key) {
			fmt.Println("exist ", test.key, test.value)
		}
	}
	if !sm.Exist("key6") {
		fmt.Println("key6 not exist ")
	}
	value6, _ := sm.Get("key6")
	if value6 != "value6" {
		fmt.Println("key6 value error ", value6)
	}
	sm.Clear()
	if sm.Len() != 0 {
		fmt.Println("clear ", sm.Len())
	}
}

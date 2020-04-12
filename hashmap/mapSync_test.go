package hashmap

import "testing"

func TestSyncMap_Set(t *testing.T) {
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
	sm := NewSyncMap()
	for _, test := range tests {
		sm.Set(test.key, test.value)
	}
	for _, test := range tests {
		value, _ := sm.Get(test.key)
		if value.(string) != test.value {
			t.Fatal("equal ", test.key, value.(string), test.value)
		}
		if !sm.Exist(test.key) {
			t.Fatal("exist ", test.key, test.value)
		}
	}
	if sm.Len() != len(tests) {
		t.Fatal("len: ", sm.Len(), len(tests))
	}
	if sm.Exist("asdfghjtre") {
		t.Fatal("exist ", "asdfghjtre")
	}
	data, _ := sm.Marshal()
	t.Log(string(data))
	t.Log("success")

	dataString := `{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5","key6":"value6"}`
	err := sm.Unmarshal([]byte(dataString))
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		if !sm.Exist(test.key) {
			t.Fatal("exist ", test.key, test.value)
		}
	}
	if !sm.Exist("key6") {
		t.Fatal("key6 not exist ")
	}
	value6, _ := sm.Get("key6")
	if value6 != "value6" {
		t.Fatal("key6 value error ", value6)
	}
	sm.Clear()
	if sm.Len() != 0 {
		t.Fatal("clear ", sm.Len())
	}

	var sm1 *MapSync
	if sm1.Len() != 0 {
		t.Fatal("clear ", sm1.Len())
	}
}

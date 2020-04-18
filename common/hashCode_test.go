package common

import "testing"

func TestEntity_GetHashCode(t *testing.T) {
	tests := []string{
		"1234567890",
		"1234567891",
		"1234567892",
		"12",
		"123",
		"qwer",
		"qwert",
		"qwerty",
		"qwertyu",
		"qwertyui",
		"qwertyuio",
		"qwertyuiop",
		"qwertyuiop[",
		"qwertyuiop[]",
		"qwertyuiop[]a",
		"qwertyuiop[]as",
		"qwertyuiop[]asd",
		"qwertyuiop[]asdf",
		"qwertyuiop[]asdfg",
		"qwertyuiop[]asdfgh",
		"qwertyuiop[]asdfghj",
		"qwertyuiop[]asdfghjk",
		"qwertyuiop[]asdfghjkl",
		"qwertyuiop[]asdfghjkl;",
		"qwertyuiop[]asdfghjkl;'",
		"qwertyuiop[]asdfghjkl;'z",
		"21",
		"1",
		"0",
		"00",
		"",
		" ",
		"  ",
		"asdfghjk",
		"asdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfge",
		"asdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfgeasdfghrewsdsdfgewsdfgwsdfge",
		"123456tresdfgh543wsgt543wsdft",
		"123456tresdfgh543wsgt543wsdfh",
		"123456tresdfgh543wsgt543wsdfi",
		"123456tresdfgh543wsgt543wsdfi0",
	}
	for _,str := range  tests {
		hashCode := GetHashCode([]byte(str))
		t.Log(hashCode)
	}
}


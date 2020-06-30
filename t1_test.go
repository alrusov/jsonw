package jsonw

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/alrusov/misc"
)

//----------------------------------------------------------------------------------------------------------------------------//

type (
	testStructType struct {
		Field1 uint    `json:"f1"`
		Field2 int     `json:"f2"`
		Field3 float64 `json:"f3"`
		Field4 string  `json:"f4"`
		Struct testSubStructType
		Array  []int `json:"a"`
		Map    misc.InterfaceMap
	}

	testSubStructType struct {
		Field1 uint    `json:"f1"`
		Field2 int     `json:"f2"`
		Field3 float64 `json:"f3"`
		Field4 string  `json:"f4"`
	}
)

var (
	testStruct = &testStructType{
		Field1: 10,
		Field2: -10,
		Field3: 10.101,
		Field4: `"zzz"/\'В чащах Юга жил бы цитрус?`,
		Struct: testSubStructType{
			Field1: 20,
			Field2: -20,
			Field3: 20.202,
			Field4: `"ZZZ"/\'Да, но фальшивый экземпляр`,
		},
		Array: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		Map:   misc.InterfaceMap{"v1": 1., "v2": "sss", "v3": false},
	}

	testJSON = []byte(`{"f1":10,"f2":-10,"f3":10.101,"f4":"\"zzz\"/\\'В чащах Юга жил бы цитрус?","Struct":{"f1":20,"f2":-20,"f3":20.202,"f4":"\"ZZZ\"/\\'Да, но фальшивый экземпляр"},"a":[9,8,7,6,5,4,3,2,1],"Map":{"v1":1,"v2":"sss","v3":false}}`)
)

//----------------------------------------------------------------------------------------------------------------------------//

func TestStdMarshal(t *testing.T) {
	tMarshal(t, true)
}

func TestAltMarshal(t *testing.T) {
	tMarshal(t, false)
}

func tMarshal(t *testing.T, useStd bool) {
	UseStd(useStd)
	j, err := Marshal(testStruct)
	if err != nil {
		t.Fatal(err)
		return
	}

	if !bytes.Equal(j, testJSON) {
		t.Errorf(`got:\n%s\nexpected:\n%s\n`, j, testJSON)
	}
}

//----------------------------------------------------------------------------------------------------------------------------//

func TestStdUnmarshal(t *testing.T) {
	tUnmarshal(t, true)
}

func TestAltUnmarshal(t *testing.T) {
	tUnmarshal(t, false)
}

func tUnmarshal(t *testing.T, useStd bool) {
	UseStd(useStd)

	var d testStructType

	err := Unmarshal(testJSON, &d)
	if err != nil {
		t.Fatal(err)
		return
	}

	if !reflect.DeepEqual(&d, testStruct) {
		t.Errorf("got:\n%#v\nexpected:\n%#v\n", &d, testStruct)
	}
}

//----------------------------------------------------------------------------------------------------------------------------//

func BenchmarkStdJSONmarshal(b *testing.B) {
	bMarshal(b, true)
}

func BenchmarkAltJSONmarshal(b *testing.B) {
	bMarshal(b, false)
}

func bMarshal(b *testing.B, useStd bool) {
	UseStd(useStd)

	bp := misc.InterfaceMap{}
	for i := 0; i < 100; i++ {
		bp[fmt.Sprintf("p_int_%d", i)] = i
		bp[fmt.Sprintf("p_str_%d", i)] = `"zzz"/\'В чащах Юга жил бы цитрус?`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Marshal(bp)
		if err != nil {
			b.Fatal(err)
			break
		}
	}
	b.StopTimer()
}

//----------------------------------------------------------------------------------------------------------------------------//

func BenchmarkStdJSONunmarshal(b *testing.B) {
	bUnmarshal(b, true)
}

func BenchmarkAltJSONunmarshal(b *testing.B) {
	bUnmarshal(b, false)
}

func bUnmarshal(b *testing.B, useStd bool) {
	UseStd(useStd)

	bp := misc.InterfaceMap{}
	for i := 0; i < 100; i++ {
		bp[fmt.Sprintf("p_int_%d", i)] = i
		bp[fmt.Sprintf("p_str_%d", i)] = `"zzz"/\'В чащах Юга жил бы цитрус?`
	}

	bb, err := Marshal(bp)
	if err != nil {
		b.Fatal(err)
		return
	}

	var d misc.InterfaceMap

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := Unmarshal(bb, &d)
		if err != nil {
			b.Fatal(err)
			break
		}
	}
	b.StopTimer()
}

//----------------------------------------------------------------------------------------------------------------------------//

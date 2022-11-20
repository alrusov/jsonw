package jsonw

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

//----------------------------------------------------------------------------------------------------------------------------//

var (
	useStdLib = true
)

//----------------------------------------------------------------------------------------------------------------------------//

// UseStd --
func UseStd(useStd bool) {
	useStdLib = useStd
}

//----------------------------------------------------------------------------------------------------------------------------//

// Marshal --
func Marshal(data any) ([]byte, error) {
	if useStdLib {
		return json.Marshal(data)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(data)
}

// Unmarshal --
func Unmarshal(data []byte, obj any) error {
	if useStdLib {
		return json.Unmarshal(data, obj)
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, obj)
}

//----------------------------------------------------------------------------------------------------------------------------//

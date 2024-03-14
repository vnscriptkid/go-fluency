package main

import (
	"encoding/json"
	"fmt"
)

func CaseHappy() {
	jsonStr := `{"age": 6}`

	type Obj struct {
		Age int `json:"age"`
	}

	var obj Obj

	err := json.Unmarshal([]byte(jsonStr), &obj)

	if err != nil {
		panic(err)
	}

	fmt.Printf("obj: %+v", obj)
}

func CaseWeird1() {
	// Using unmatched type between jsonStr and struct
	// jsonStr: number
	// Obj struct: string
	// !!! panic: json: cannot unmarshal number into Go struct field Obj.age of type string
	jsonStr := `{"age": 6}`

	type Obj struct {
		Age string
	}

	var obj Obj

	err := json.Unmarshal([]byte(jsonStr), &obj)

	if err != nil {
		panic(err)
	}

	fmt.Printf("obj: %+v", obj)
}

func CaseWeird2() {
	// Using unmatched type between jsonStr and struct
	// jsonStr: string
	// Obj struct: int
	// !!! panic: json: cannot unmarshal string into Go struct field Obj.age of type int
	jsonStr := `{"age": "6"}`

	type Obj struct {
		Age int `json:"age"`
	}

	var obj Obj

	err := json.Unmarshal([]byte(jsonStr), &obj)

	if err != nil {
		panic(err)
	}

	fmt.Printf("obj: %+v", obj)
}

func CaseWeird3() {
	// Using mis-leading tag
	// => obj: {Age:0}
	jsonStr := `{"age": 6}`

	type Obj struct {
		Age int `json:"ageeeeeeeeee"`
	}

	var obj Obj

	err := json.Unmarshal([]byte(jsonStr), &obj)

	if err != nil {
		panic(err)
	}

	fmt.Printf("obj: %+v", obj)
}

func CaseWeird4() {
	// Missing ref &
	// !!! panic: json: Unmarshal(non-pointer main.Obj)
	jsonStr := `{"age": 6}`

	type Obj struct {
		Age int `json:"age"`
	}

	var obj Obj

	err := json.Unmarshal([]byte(jsonStr), obj)

	if err != nil {
		panic(err)
	}

	fmt.Printf("obj: %+v", obj)
}

func main() {
	// CaseHappy()
	// CaseWeird1()
	// CaseWeird2()
	// CaseWeird3()
	CaseWeird4()

}

package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page  int
	Fruit []string
}

type response2 struct {
	Page  int      `json:"page"`
	Fruit []string `json:"fruit"`
}

func main() {

	//Json encode
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	//Json encode
	intB, _ := json.Marshal(25)
	fmt.Println(string(intB))

	slcB := []string{"apple", "peach", "mango"}
	slc, _ := json.Marshal(slcB)
	fmt.Println(string(slc))

	mapB := map[string]int{"one": 1, "two": 2, "three": 3}
	mapM, _ := json.Marshal(mapB)
	fmt.Println(string(mapM))

	resID := &response1{
		Page:  1,
		Fruit: []string{"apple", "peach", "banana"}}

	res1B, _ := json.Marshal(resID)
	fmt.Println(string(res1B))

	resId2 := &response2{
		Page:  2,
		Fruit: []string{"mango", "sugar cane", "watermelan"}}
	res2B, _ := json.Marshal(resId2)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	//create variable to hold value
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

// 只有 可导出 的字段才会被 JSON 编码/解码。
// 必须以大写字母开头的字段才是可导出的。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println("bolB: ", string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println("intB: ", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("fltB: ", string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println("strB: ", string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)

	fmt.Println("slcB: ", string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("mapB: ", string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("res1B: ", string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("res2B: ", string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat: ", dat)

	num := dat["num"].(float64)
	fmt.Println("num: ", num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println("str1: ", str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

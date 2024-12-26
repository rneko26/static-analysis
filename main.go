package main

import "fmt"

func main() {

	foo := map[string]interface{}{
		"channel": "GOPAY",
		"code":    "200",
	}

	result := foo["not_exist"].(map[string]interface{})["RequestData"]

	arr := [2]string{"foo", "bar"}

	fmt.Println(result)

	fmt.Println(arr[3])

}

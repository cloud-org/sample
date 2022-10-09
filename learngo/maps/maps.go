package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "not bad",
	}

	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map

	fmt.Println(m2)

	var m3 map[string]int // m3 == nil

	fmt.Println(m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	if courseName, ok := m["couasrt"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

package lesson8

import (
	"fmt"
)

func Entry() {
	fmt.Println("\nMaps:")
	map1 := make(map[string]string) // map1 := make(map[string]string, size)
	map2 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	map1["key1"] = "value1"

	fmt.Println(map1)
	fmt.Println(map2)
	mutation(map2)
	fmt.Println(map2)
	delete(map2, "key2")
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Println(key, value)
	}
}

func mutation(m map[string]string) {
	m["key1"] = "new_value_1"
}

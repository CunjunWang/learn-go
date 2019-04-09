package main

import "fmt"

/*
m := map[string]string {
	"name": "cunjunwang",
	"course": "golang",
}
*/

// map定义方式:
//  map[KeyType]ValueType
// 复合map: map[K1]map[K2]V

func main() {

	fmt.Println("Creating Maps")

	map1 := map[string]string{
		"name":   "cunjunwang",
		"course": "golang",
	}

	map2 := make(map[string]int) // m2 == empty map

	var map3 map[string]int // m3 == nil

	fmt.Println("map1 = ", map1)
	fmt.Println("map2 = ", map2)
	fmt.Println("map3 = ", map3)

	fmt.Println("Traversing Maps")

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for k := range map1 {
		fmt.Println(k)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

	// key在map中是无序的, 是一个hashMap

	fmt.Println("Get values from map")
	courseName, ok := map1["course"]
	fmt.Println(courseName, ok)
	causeName, ok := map1["cause"] // 取不存在的值, 拿到zero value = "" for string
	fmt.Println(causeName, ok)
	// 用ok判断key在map中是否存在
	if causeName, ok := map1["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key [cause] not in the map")
	}

	fmt.Println("Delete values from map")
	if name, ok1 := map1["name"]; ok1 {
		fmt.Printf("delete key [name] value [%s] from map\n", name)
		delete(map1, "name")
	} else {
		fmt.Println("key [name] not in the map")
	}

	fmt.Println("name = ", map1["name"])

	// map的key
	// map使用哈希表，必须可以比较相等
	// 类似于java中必须要重写equals和hashcode方法
	// 除了slice, map, function的内建类型都可以作为key
	// Struct不包含上述字段类型的话也可以作为key

}

package base

import (
	"fmt"
	"sort"
)

func MapT() {
	var maps map[string]string
	maps = map[string]string{}
	fmt.Println(maps)

	map2 := make(map[string]string)

	map2["aa"] = "aa"
	fmt.Println(map2)

	map3 := make(map[string]map[string]string)
	map4 := make(map[string]string)
	map4["b"] = "ccc"
	map3["a"] = map4
	d, ok := map3["a"]["b"]
	if !ok {
		fmt.Println("没有拿到数据！")
		return
	}
	fmt.Println(d)

	for key, val := range map2 {
		fmt.Println(key)
		fmt.Println(val)
	}

	slimap := make([]map[string]string, 10)

	for i, v := range slimap {
		v = make(map[string]string)
		slimap[i] = v
		fmt.Println(slimap[i])
	}
	fmt.Println(slimap)

	map5 := make(map[int]string)
	map5[2] = "aaa"
	map5[3] = "bbb"
	map5[1] = "ccc"
	map5[4] = "ddd"

	slit := make([]int, len(map5))

	i := 0
	for key, _ := range map5 {
		slit[i] = key
		i++
	}
	sort.Ints(slit)

	fmt.Println(slit)

}

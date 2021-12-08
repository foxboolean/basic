package main

func main() {
	// 指定容量为 2
	m := make(map[string]string, 2)
	// 不指定容量
	m1 := make(map[string]string)
	m2 := map[string]string {
		"Name":"Tom",
	}

	m2["Age"] = "12"
	val := m2["Name"]
	println(val)

	m1["Name"] = "Lisa"
	sex, ok := m["sex"]
	if ok {
		println(sex)
	} else {
		println("key not found")
	}

	for key, val := range m2 {
		println(key + ":" + val)
	}
}

package main

import "strings"

func main() {
	println(equationsPossible([]string{"a==b", "e==c", "b==c", "a!=e"}))
}
func equationsPossible(equations []string) bool {
	uf := UF{}
	uf.parents = map[string]string{}
	uf.size = map[string]int{}
	uf.count = 0

	var noEqualIndex []int //不相等方程式的索引
	for i := 0; i < len(equations); i++ {
		if strings.Contains(equations[i], "!=") {
			noEqualIndex = append(noEqualIndex, i)
			continue
		}
		alphabets := strings.Split(equations[i], "==")
		uf.addElement(alphabets[0])
		uf.addElement(alphabets[1])
		uf.union(alphabets[0], alphabets[1]) //将相等的连通
	}

	for _, i := range noEqualIndex {
		alphabets := strings.Split(equations[i], "!=") //不相等的方程式
		uf.addElement(alphabets[0])
		uf.addElement(alphabets[1])
		if alphabets[0] == alphabets[1] { //相等字母
			return false
		}
		if uf.connected(alphabets[0], alphabets[1]) { //连通过，之前肯定是相等的，则矛盾
			return false
		}
	}
	return true
}

type UF struct {
	parents map[string]string
	size    map[string]int
	count   int
}

func (uf UF) union(p string, q string) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	if uf.size[rootP] > uf.size[rootQ] {
		uf.parents[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	} else {
		uf.parents[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	}
	uf.count--
}

func (uf UF) connected(p string, q string) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}

func (uf UF) find(p string) string {
	for uf.parents[p] != p {
		// 路径压缩
		uf.parents[p] = uf.parents[uf.parents[p]]
		p = uf.parents[p]
	}
	return p
}

func (uf UF) counts() int {
	return uf.count
}

func (uf UF) addElement(c string) {
	if _, ok := uf.parents[c]; !ok {
		uf.parents[c] = c
	}
}

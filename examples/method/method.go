package method

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

func trace(s string) string {
	fmt.Printf("entering :%s\n", s)
	return s
}

func un(s string) {
	fmt.Printf("leaving :%s\n", s)
}

func Follow() {
	un(trace("follow"))
}

func DeferMethod() {
	for i := 0; i < 2; i++ {
		fmt.Printf("loop i: %d\n", i)
		defer fmt.Printf("defer perform loop %d\n", i)
	}
	defer fmt.Printf("defer perform\n")
	for i := 3; i < 5; i++ {
		fmt.Printf("loop i: %d\n", i)
		defer fmt.Printf("defer perform loop %d\n", i)
	}
}

// Contents 内容返回文件的内容作字符串
// 内容返回文件的内容作为字符串。
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()  // 我们结束后就关闭了f

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append稍后讨论。
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err  // 如果我们回到这里，f就关闭了。
		}
	}
	return string(result), nil // 如果我们回到这里，f就关闭了。
}

func NextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {
	}
	x := 0
	for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}
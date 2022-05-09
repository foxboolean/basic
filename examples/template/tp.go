package main

import (
	"os"
	"strings"
	"text/template"
)

// . 标识渲染变量
const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	// 注册自定义函数
	funcMap := template.FuncMap{"title": strings.Title}
	// 生成模板对象
	tpl := template.New("try_template")
	// 解析模板
	tpl, _ = tpl.Funcs(funcMap).Parse(templateText)
	// 加载数据
	data := map[string]string{
		"Name1": "try",
		"Name2": "go",
		"Name3": "template",
	}
	// 将结果输出到 os.Stdout
	_ = tpl.Execute(os.Stdout, data)
}

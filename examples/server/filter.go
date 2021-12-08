package main

import (
	"fmt"
	"time"
)

// Filter 定义方法类型结构体
type Filter func(c *Context)

// FilterBuilder 构造出一个 Filter 结构体
type FilterBuilder func(next Filter) Filter

var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("spend times is %d ns", end - start)
	}
}
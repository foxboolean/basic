package main

import (
	"fmt"
	"syscall"
)

func main() {
	tz, ok := syscall.Getenv("TZ")
	fmt.Println(tz, ok)
	//switch {
	//case !ok:
	//	z, err := loadLocation("localtime", []string{"/etc/"})
	//	if err == nil {
	//		localLoc = *z
	//		localLoc.name = "Local"
	//		return
	//	}
	//case tz != "" && tz != "UTC":
	//	if z, err := loadLocation(tz, zoneSources); err == nil {
	//		localLoc = *z
	//		return
	//	}
	//}
}

package main

import (
	"fmt"
	"os"
	"syscall"
)

/**
 * 如果想要变更目录的权限，就必须执行syscall.Umask(0)，因为
 * 系统件文件夹的权限是 0777 & ~mask
 * 所以，我们将umask设置为0，最终的权限即为用户自定的任意权限
 */
func main() {
	mask := syscall.Umask(0)
	fmt.Println("umask: ", mask)
	defer syscall.Umask(mask)

	err := os.Mkdir("/tmp/test_umask/", 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Create Tmp Directory")
}

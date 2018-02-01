package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

/**
 * 如果缺少rand.Seed函数，那么每次执行产生的随机数序列都是相同的
 * 即使包含rand.Seed函数，也有可能产生相同序列的随机数
 *     这是因为rand.Seed底层实现是取模运算（2^32-1)，而传入的seed参数是64位，因此还是有一定几率生成相同的seed
 * 这就是伪随机数的悲惨命运
 *
 * 真正要产生随机数，还得用crypto/rand包下面的随机数生成函数
 *     注意：crypto/rand产生随机数的方式要比math/rand慢上一个数量级
 */

func main() {
	// test incorrect rand gen
	fmt.Println("-- incorrect --")
	randE()
	// test correct rand gen
	fmt.Println("-- correct --")
	randC()

	// real rand
	fmt.Println("-- real --")
	randR()
}

func randR() {
	var x uint32
	binary.Read(crand.Reader, binary.BigEndian, &x)
	fmt.Println(x)
}

func randC() {
	rand.Seed(time.Now().Unix())
	randE()
}

func randE() {
	for i := 0; i < 10; i++ {
		fmt.Println(getRand())
	}
}

func getRand() int {
	return rand.Intn(10)
}

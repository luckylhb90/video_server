package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int

	/*<-是对chan类型来说的。chan类型类似于一个数组。
	当<- chan 的时候是对chan中的数据读取；
	相反 chan <- value 是对chan赋值。*/
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Panicln("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New Connection coming: %d", c)
}

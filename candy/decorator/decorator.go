// @Title  decorator
// @Description
// @Author  EwdAger
// @Update  2021/7/19 10:34

package main

import (
	"fmt"
	"time"
)

func hello(name string) string {
	return fmt.Sprintf("Hello my name is %s", name)
}

// 定义一个 type 可以极大的提升可读性
type helloFunc func(name string) string

func timeIt(fun helloFunc) helloFunc {
	return func(name string) string {
		timeStart := time.Now()

		res := fun(name)

		timeElapsed := time.Since(timeStart)
		fmt.Println("用时：", timeElapsed)
		return res
	}
}

func ping(pong string) func(fun helloFunc) helloFunc {
	fmt.Println(pong)
	wrapper := func(fun helloFunc) helloFunc {
		return func(name string) string {
			timeStart := time.Now()

			res := fun(name)

			timeElapsed := time.Since(timeStart)
			fmt.Println("用时：", timeElapsed)
			return res
		}
	}

	return wrapper
}

func main() {

	// 不带参数的装饰器
	res := timeIt(hello)("Tom")
	fmt.Println(res)

	// 带参数的装饰器
	res2 := ping("pong")(hello)("Tom")
	fmt.Println(res2)
}

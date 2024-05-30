package main

import "fmt"

func main() {
	fmt.Println("Hello, World!" + "12345")
	fmt.Println("菜鸟教程：runoob.com")
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	//fmt.Printf(url,stockcode,enddate)
	var a, b = 9, 2
	//var c = true
	var c = false
	var d = "hello"
	if c {
		fmt.Println(a, b, c, d)
	}
	for a <= 13 {
		fmt.Println("Hello, World!" + "123")
		a++
	}
	fmt.Println(a, b)
	ret := max1(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
	//a1, b1 := swap("Google", "Runoob")
	var a1, b1 = swap("Google", "Runoob")
	fmt.Println(a1, b1)

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	//var g = 20   /* 声明实际变量 */
	var g int
	g = 10
	var ip *int /* 声明指针变量 */

	ip = &g /* 指针变量的存储地址 */

	fmt.Printf("g 变量的地址是: %x\n", &g)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

}

/* 函数返回两个数的最大值 */
func max1(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func swap(x, y string) (string, string) {
	return y, x
}

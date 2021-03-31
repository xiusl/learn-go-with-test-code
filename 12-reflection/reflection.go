package _2_reflection

func Walk(x interface{}, fn func(input string)) {
	fn("write something")
}


/*NOTE
Go 运行我们使用 interface{} 来作为任意类型

除非真的必要否则不要使用反射

如果想实现函数多态性，不推荐使用反射，而是通过接口实现（不是 interface{} 类型）

*/


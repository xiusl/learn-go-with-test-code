package _2_reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), fn)
		}
	}
}


/*NOTE
使用反射包的 ValueOf 函数，查看 x 的属性，返回给定变量的 Value

我们对传入的值做了乐观的断定
field := val.Field(0) 获取第一字段，可能根本没有字段而引起 panic
field.String() 调用 String()，以字符串的形式返回底层值，但我们并不清除它是否一定是字符串
*/

/*NOTE
Go 运行我们使用 interface{} 来作为任意类型

除非真的必要否则不要使用反射

如果想实现函数多态性，不推荐使用反射，而是通过接口实现（不是 interface{} 类型）

*/

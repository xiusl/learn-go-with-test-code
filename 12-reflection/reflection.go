package _2_reflection

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}


func WalkV2(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		WalkV2(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			WalkV2(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			WalkV2(v.Interface(), fn)
		}
	}

	/*
	bad 优化，getField maybe nil
	for i := 0; i < numberOfValues; i++ {
		WalkV2(getField(i).Interface(), fn)
	}
	*/
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

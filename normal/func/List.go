package main

// 接口定义
// 方法名字大写开头是公有方法，小写开头是私有方法
type List interface {
	Add(idx int, val any)
	Append(val any)
	Delete(index int)
}

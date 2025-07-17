// 1114. Print in Order
// LeetCode 上的 1114. Print in Order 是一道多執行緒同步（Concurrency）的設計題，要求你設計 Foo 類別，確保三個方法 first(), second(), third() 在多執行緒下仍按順序呼叫。

type Foo struct {
	firstOnce  sync.Once
	secondOnce sync.Once
}

func Constructor() Foo {
	return Foo{}
}

func (f *Foo) first(printFirst func()) {
	f.firstOnce.Do(printFirst)
}

func (f *Foo) second(printSecond func()) {
	f.firstOnce.Do(func(){}) // 等待 first 完
	f.secondOnce.Do(printSecond)
}

func (f *Foo) third(printThird func()) {
	f.secondOnce.Do(func(){}) // 等待 second 完
	printThird()
}
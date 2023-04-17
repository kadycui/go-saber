package utils

import "fmt"

// 非侵入式接口

/*
不需要显式地创建一个类去实现一个接口。

◇ 去掉了繁杂的继承体系，Go语言的标准库再也不需要绘制类库的继承树图。在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法、每个方法是何含义就足够了。
◇ 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理。接口由使用方按需定义，而不用事前规划。
◇ 不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦合。接口由使用方按自身需求来定义，使用方无须关心是否有其他模块定义过类似的接口。
	总的来说，非入侵式的接口设计更简洁、灵活，更注重实用性。下面我们从代码层面来更深
	入地理解。
*/

type MobilPhoner interface {
	Call() error
	Video() error
	Game() error
}

type Apple interface {
	Call() error
	Video() error
}
type HuaWei interface {
	Call() error
	Game() error
}

type Phone struct {
	Name string
}

func (p *Phone) Video() error {
	fmt.Println(p.Name, "Start Vidio")
	return nil
}

func (p *Phone) Call() error {
	fmt.Println(p.Name, "Start Call")
	return nil
}

func (p *Phone) Game() error {
	fmt.Println(p.Name, "Start Game")
	return nil
}

func Demo12() {
	// var apple Apple = &Phone{"apple"}
	// var huawei HuaWei = &Phone{"huewei"}

	var apple = new(Phone)
	var huawei = new(Phone)

	apple.Call()
	huawei.Game()

}

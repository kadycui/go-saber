package utils

import "fmt"


// 接口可以通过嵌套大接口包含小接口



//定义接口
type A interface {
    test1()
}
//定义接口
type B interface {
    test2()
}
//定义接口
type C interface {
    A
    B
    test3()
}

//2.实现类
type Cat struct { //如果想实现接口从，那不止要实现接口c的方法，还要实现接口A，B中方法

}

//连接实现类来实现接口
func (c Cat) test1(){
    fmt.Println("test1()...")
}

func (c Cat) test2(){
    fmt.Println("test2()...")
}

func (c Cat) test3(){
    fmt.Println("test3()...")
}



func Demo6(){
    var cat Cat = Cat{}
    cat.test1()//test1()...
    cat.test2() //test2()...
    cat.test3() //test3()...

    fmt.Println("----------------")
    var a1 A = cat
    a1.test1()  //test1()...

    var b1 B = cat
    b1.test2()  //test2()...

    var c1 C = cat
    c1.test1() //test1()...
    c1.test2() //test2()...
    c1.test3() //test3()...
}

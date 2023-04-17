package utils

import "fmt"

/*
值传递：将变量的一个副本传递给函数，函数中不管如何操作该变量副本，都不会改变原变量的值。
引用传递：将变量的内存地址传递给函数，函数中操作变量时会找到保存在该地址的变量，对其进行操作，会改变原变量的值。
*/

func passByValue(numPara int) { // 值传递函数
	fmt.Printf("passByValue函数中变量numPara地址为：%p\n", &numPara)
	numPara = 100

}

func passByReference(numPara *int) { // 引用传递函数
	fmt.Printf("passByReference函数中指针变量numPara地址为：%p\n", &numPara)
	fmt.Printf("passByReference函数中指针变量numPara指向的地址为：%p\n", numPara)
	*numPara = 100
}

func Demo3() {
	fmt.Println("==========================")
	num := 1
	fmt.Printf("main函数中变量num地址为：%p\n", &num)
	passByValue(num)
	fmt.Printf("num变量值为：%d\n", num)
	passByReference(&num)
	fmt.Printf("num变量值为：%d\n", num)

	/*
		num变量在传递到passByReference函数后，其原本的值发生了变
		化。原因是其传递的是num变量的地址，该地址经过拷贝后传递给passByReference函数，但是其指
		向的值为num变量的地址（0xc042066080），因此在passByReference函数中的“*numPara = 100”会
		影响原变量的值。
	*/

}

func passByReference1(mapNumReference map[int]int) {
	fmt.Printf("passByReference函数中变量mapNumReference地址为：%p\n", mapNumReference)
	fmt.Printf("passByReference函数中变量mapNumReference所属指针地址为：%p\n", &mapNumReference)
	mapNumReference[1] = 100
}

func Demo4() {
	// 使用map变量传递的是原变量的指针
	fmt.Println("==========================")
	mapNum := map[int]int{1: 10}
	fmt.Printf("main函数中变量mapNum地址为：%p\n", mapNum)
	fmt.Printf("main函数中变量mapNum所属指针的地址为：%p\n", &mapNum)
	fmt.Printf("mapNum变量值为：%d\n", mapNum)
	passByReference1(mapNum)
	fmt.Printf("mapNum变量值为：%d\n", mapNum)

	/*
		通过map类型变量参数传递的例子我们可以发现，函数进行值传递后拷贝的是map类型变量指
		针，但拷贝后的指针地址指向的还是map的地址，从而导致在函数passByReference中对map的操作
		会影响原变量。
	*/
}

func Demo5() {
	//将4、5、6追加到sourceSlice切片中
	sourceSlice := []int{1, 2, 3}
	sourceSlice = append(sourceSlice, 4, 5, 6)
	fmt.Println("sourceSlice:", sourceSlice)
	//将sourceSlice切片中的元素复制到targetSlice切片中
	targetSlice := make([]int, 3)
	num := copy(targetSlice, sourceSlice)
	fmt.Println("复制成功的元素个数为：", num)
	fmt.Println("targetSlice:", targetSlice)
	fmt.Println("targetSlice切片长度为:", len(targetSlice))
	fmt.Println("targetSlice切片容量为:", cap(targetSlice))
}

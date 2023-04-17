package utils

import (
	"fmt"
	"os"
	"strconv"
)

// 对于文件的创建与打开，使用的是标准库os中的OpenFile
// func OpenFile(name string, flag int, perm FileMode) (file ＊File, err error)

/*
位掩码参数flag用于指定文件的访问模式，可用的值在os中定义为常量（以下值并非所有操作系统都可用）：
const (
 O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
 O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
 O_RDWR int = syscall.O_RDWR // 读写模式打开文件
 O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
 O_CREATE int = syscall.O_CREAT // 如果不存在将创建一个新文件
 O_EXCL int = syscall.O_EXCL // 和O_CREATE配合使用，文件必须不
 存在
 O_SYNC int = syscall.O_SYNC // 打开文件用于同步I/O
 O_TRUNC int = syscall.O_TRUNC // 如果可能，打开时清空文件
)


其中，O_RDONLY、O_WRONLY、O_RDWR应该只指定一个，剩下的通过“|”操作符来指定。
该函数内部会给flags加上syscall.O_CLOEXEC，在fork子进程时会关闭通过OpenFile打开的文件，即子进程不会重用该文件描述符。

*/

/*
const (
 // 单字符是被 String 方法用于格式化的属性缩写
 ModeDir FileMode = 1 << (32 - 1 - iota) // d: 目录
 ModeAppend // a: 只能写入，且只能写入到末尾
 ModeExclusive // l: 用于执行
 ModeTemporary // T: 临时文件（非备份文件）
 ModeSymlink // L: 符号链接（不是快捷方式文件）
 ModeDevice // D: 设备
 ModeNamedPipe // p: 命名管道（FIFO）
 ModeSocket // S: Unix域socket
 ModeSetuid // u: 表示文件具有其创建者用户id权限
 ModeSetgid // g: 表示文件具有其创建者组id的权限
 ModeCharDevice // c: 字符设备，需已设置ModeDevice
 ModeSticky // t: 只有root/创建者能删除/移动文件

 // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
 ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket |
ModeDevice
 ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)

*/

func Demo5() {
	//以读写方式打开文件，如果不存在，则创建
	file, err := os.OpenFile("chatper_10/utils/1.txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	file.Close()
}

// 读取文件可以使用os库中的Read接口
// func (f ＊File) Read(b []byte) (n int, err error)

func ReadFile(path string) {
	// 读取文件内容
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	// 创建byte的slice用于接收文件读取数据
	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容:")

	// 循环读取
	for {
		// Read 函数会改变文件当前偏移量
		len, _ := file.Read(buf)
		// 读取字节数为0时跳出循环
		if len == 0 {
			break
		}
		fmt.Println(string(buf))
	}
	file.Close()
}

// 当遇到特别大的文件，并且只需要读取文件最后部分的内容时，
//Read接口就不能满足我们的 需要了，这时可以使用另外一个文件读取接口ReadAt
// func (f ＊File) ReadAt(b []byte, off int64) (n int, err error)

func ReadFile2(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 1024)
	fmt.Println("以下是文件内容:")

	_, _ = file.ReadAt(buf, 45)
	fmt.Println(string(buf))
	file.Close()
}

// 文件写入
// func (f ＊File) Write(b []byte) (n int, err error)

func WriteFile() {
	file, err := os.Create("chatper_10/utils/2.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := "这是要写入的数据\r\n"
	for i := 0; i < 3; i++ {
		file.Write([]byte(data))
	}
	file.Close()
}

// 使用WriteAt可以指定从文件的什么位置开始写
// func (f ＊File) WriteAt(b []byte, off int64) (n int, err error)

func WriteFile2() {
	file, err := os.Create("chatper_10/utils/3.txt")
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 3; i++ {
		// 写入偏移量
		ix := i * 64
		file.WriteAt([]byte("我是写入的数据"+strconv.Itoa(i)+"\r\n"), int64(ix))
	}
	file.Close()

}

// 删除文件
// 删除文件使用os库中的Remove和RemoveAll接口，与删除目录的接口一致

func DeleteFile() {
	// 删除文件
	err := os.Remove("chatper_10/utils/2.txt")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("删除文件成功!")
	}

	// 删除目录下所有文件
	err2 := os.RemoveAll("D:\\vvv\\")
	if err2 != nil {
		fmt.Println(err2)

	} else {
		fmt.Println("删除文件成功!")
	}

}

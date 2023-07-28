package main

import (
	"gostart/feature/channel_operate/utils"
)

func main() {

	// utils.ChUse()
	// utils.ChannelWg()

	// utils.AssertChannel()
	// utils.ChannelRange()

	// utils.EmptyInterfaceChan()

	// utils.ConcurrentSync()

	// utils.CloseChannel()

	// utils.ReadWriteChan()


	utils.FindPrimeNum()

}

/*
关闭channel后，无法向channel 再发送数据(引发 panic 错误后导致接收立即返回零值)；
关闭channel后，可以继续从channel接收数据；
对于nil channel，无论收发都会被阻塞。
*/

package main

import (
	"io"
	"log"
	"net"
	"time"

	"wzmiiiiii.cn/zinx/znet"
)

func main() {
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("connect server err:", err)
		return
	}
	defer conn.Close()

	var cnt uint32 = 1

	for {
		// 向服务器发送数据
		dp := znet.NewDataPack()
		binaryMsg, _ := dp.Pack(znet.NewMessage(cnt, []byte("ZinxV0.5 client Test msg")))

		if _, err := conn.Write(binaryMsg); err != nil {
			log.Println("Client send msg fail,msgID:", cnt)
		}
		cnt++

		// 接受服务器发送过来的数据
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, headData)
		if err != nil {
			log.Println("read head error")
			return
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			log.Println(err)
			return
		}
		if msg.GetMsgLen() > 0 {
			// msg中有数据.
			data := make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(conn, data)
			if err != nil {
				log.Println(err)
				return
			}
			msg.SetData(data)
		}
		log.Println("recv server data:")
		log.Println(msg.GetMsgId())
		log.Println(string(msg.GetData()))

		time.Sleep(2 * time.Second)
	}
}

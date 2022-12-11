package znet_test

import (
	"io"
	"log"
	"net"
	"testing"

	"wzmiiiiii.cn/zinx/znet"
)

func TestDataPack(t *testing.T) {
	listener, _ := net.Listen("tcp", ":8080")

	defer listener.Close()

	go func() {
		for {
			conn, _ := listener.Accept()

			go func(conn net.Conn) {
				defer conn.Close()
				dp := znet.NewDataPack()

				for {
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

					log.Println(msg.GetMsgId())
					log.Println(msg.GetMsgLen())
					log.Println(string(msg.GetData()))
				}
			}(conn)
		}
	}()

	//	模拟客户端
	conn, _ := net.Dial("tcp", ":8080")

	dp := znet.NewDataPack()

	msg1 := &znet.Message{
		Id:   1,
		Len:  4,
		Data: []byte{'z', 'i', 'n', 'x'},
	}

	msg2 := &znet.Message{
		Id:   2,
		Len:  5,
		Data: []byte{'h', 'e', 'l', 'l', 'o'},
	}
	// 模拟粘包过程,封装两个msg一起发送.
	data1, _ := dp.Pack(msg1)
	data2, _ := dp.Pack(msg2)
	data1 = append(data1, data2...)
	_, _ = conn.Write(data1)
	select {}
}

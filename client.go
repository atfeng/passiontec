package passiontec

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	http.Client
	Logined bool
}

func New() *Client {
	cli := &Client{}
	cli.Jar, _ = cookiejar.New(nil)

	go cli.Heartbeat()
	return cli
}

func (cli *Client) Heartbeat() {
	c := time.Tick(1 * time.Minute)
	for range c {
		u, err := cli.GetCurrentUserInfo()
		if err != nil {
			log.Println("心跳检测失败:", err)
			continue
		}

		log.Printf("当前登陆用户信息", u)
	}
}

type BaseResponse struct {
	Code    int
	Message string
}

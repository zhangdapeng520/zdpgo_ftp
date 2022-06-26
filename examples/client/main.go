package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ftp"
	"github.com/zhangdapeng520/zdpgo_log"
)

/*
@Time : 2022/6/7 11:53
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	f := zdpgo_ftp.NewWithConfig(&zdpgo_ftp.Config{
		Debug:    true,
		Host:     "localhost",
		Port:     2122,
		Username: "admin",
		Password: "admin",
	}, zdpgo_log.Tmp)
	client, err := f.GetClient()
	if err != nil {
		panic(err)
	}

	err = client.Upload("README.md", "README.md")
	err = client.Download("README.md", "README1.md")

	toBytes, err := client.DownloadToBytes("README.md")
	fmt.Println(string(toBytes))
}

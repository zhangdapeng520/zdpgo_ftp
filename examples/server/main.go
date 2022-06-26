package main

import (
	"github.com/zhangdapeng520/zdpgo_ftp"
	"github.com/zhangdapeng520/zdpgo_log"
)

func main() {
	f := zdpgo_ftp.NewWithConfig(&zdpgo_ftp.Config{
		Debug:    true,
		Host:     "localhost",
		Port:     2122,
		Username: "admin",
		Password: "admin",
	}, zdpgo_log.Tmp)
	s := f.GetServer()
	err := s.Run()
	if err != nil {
		panic(err)
	}
}

package main

import "github.com/zhangdapeng520/zdpgo_ftp"

func main() {
	f := zdpgo_ftp.NewWithConfig(&zdpgo_ftp.Config{
		Debug:    true,
		Host:     "localhost",
		Port:     2122,
		Username: "admin",
		Password: "admin",
	})
	s := f.GetServer()
	err := s.Run()
	if err != nil {
		panic(err)
	}
}

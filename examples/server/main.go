// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/zhangdapeng520/zdpgo_ftp/ftp"
	"github.com/zhangdapeng520/zdpgo_ftp/server/core"
	"github.com/zhangdapeng520/zdpgo_ftp/server/driver/file"
)

func main() {
	err := os.MkdirAll("./testdata", os.ModePerm)
	if err != nil {
		panic(err)
	}

	var perm = core.NewSimplePerm("test", "test")
	opt := &core.ServerOpts{
		Name: "test ftpd",
		Factory: &file.DriverFactory{
			RootPath: "./testdata",
			Perm:     perm,
		},
		Port: 2122,
		Auth: &core.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Logger: new(core.DiscardLogger),
	}
	s := core.NewServer(opt)

	go func() {
		err = s.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	// Give server 0.5 seconds to get to the listening state
	timeout := time.NewTimer(time.Millisecond * 500)

	f, err := ftp.Connect("localhost:2122")
	if err != nil && len(timeout.C) == 0 { // Retry errors
	}
	f.Login("admin", "admin")

	var content = `test`
	f.Stor("test.txt", strings.NewReader(content))
	names, err := f.NameList("/")
	fmt.Println(names, err)
	bs, err := ioutil.ReadFile("./testdata/test.txt")
	fmt.Println(content, string(bs))
	entries, err := f.List("/")
	fmt.Println(len(entries))
	curDir, err := f.CurrentDir()
	fmt.Println(curDir)
	size, err := f.FileSize("/test.txt")
	fmt.Println(size)
	r, err := f.RetrFrom("/test.txt", 2)
	buf, err := ioutil.ReadAll(r)
	r.Close()
	fmt.Println(string(buf))
	err = f.Rename("/test.txt", "/test.go")
	err = f.MakeDir("/src")
	err = f.Delete("/test.go")
	err = f.ChangeDir("/src")
	curDir, err = f.CurrentDir()
	fmt.Println(curDir)
	f.Stor("test.txt", strings.NewReader(content))
	r, err = f.Retr("/src/test.txt")
	buf, err = ioutil.ReadAll(r)
	r.Close()
	fmt.Println(string(buf))
	err = f.RemoveDir("/src")
	err = f.Quit()

	for {
		time.Sleep(time.Hour * 24)
	}
}

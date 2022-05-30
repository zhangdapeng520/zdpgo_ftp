package zdpgo_ftp

import (
	"github.com/zhangdapeng520/zdpgo_test/assert"
	"testing"
)

/*
@Time : 2022/5/30 9:50
@Author : 张大鹏
@File : server_test.go
@Software: Goland2021.3.1
@Description:
*/

func TestServer_Run(t *testing.T) {
	f := getFtp()
	s := f.GetServer()
	err := s.Run()
	assert.NoError(t, err)
}

package zdpgo_ftp

import (
	"github.com/zhangdapeng520/zdpgo_test/assert"
	"testing"
)

/*
@Time : 2022/5/28 17:31
@Author : 张大鹏
@File : ftp_test.go
@Software: Goland2021.3.1
@Description:
*/

func getFtp() *Ftp {
	return NewWithConfig(&Config{
		Debug:    true,
		Host:     "localhost",
		Port:     2122,
		Username: "admin",
		Password: "admin",
	})
}

func TestFtp_IsHealth(t *testing.T) {
	f := getFtp()
	assert.Equal(t, true, f.IsHealth())
	assert.NotEqual(t, false, f.IsHealth())
}

func TestFtp_GetClient(t *testing.T) {
	f := getFtp()
	client, err := f.GetClient()
	assert.Equal(t, nil, err)
	assert.NotEqual(t, client, nil)
}

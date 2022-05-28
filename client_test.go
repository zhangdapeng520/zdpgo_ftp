package zdpgo_ftp

import (
	"github.com/zhangdapeng520/zdpgo_test/assert"
	"testing"
)

/*
@Time : 2022/5/28 17:39
@Author : 张大鹏
@File : client_test.go
@Software: Goland2021.3.1
@Description:
*/

func getClient() *Client {
	f := getFtp()
	if f.IsHealth() {
		client, err := f.GetClient()
		if err != nil {
			panic(err)
		}
		return client
	}
	return nil
}

func TestClient_Upload(t *testing.T) {
	c := getClient()
	assert.NotEqual(t, c, nil)
	err := c.Upload("README.md", "README.md")
	assert.NoError(t, err)
}

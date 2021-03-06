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

func TestClient_UploadAndCheckMd5(t *testing.T) {
	f := NewWithConfig(&Config{
		Debug:    true,
		Host:     "localhost",
		Port:     2122,
		Username: "admin",
		Password: "admin",
	})
	client, err := f.GetClient()
	assert.NoError(t, err)

	flag := client.UploadAndCheckMd5("README.md", "README1.md")
	assert.Equal(t, flag, true)
}

func TestClient_Download(t *testing.T) {
	c := getClient()
	assert.NotEqual(t, c, nil)
	err := c.Download("README.md", "README1.md")
	assert.NoError(t, err)
}

func TestClient_DownloadToBytes(t *testing.T) {
	c := getClient()
	assert.NotEqual(t, c, nil)
	toBytes, err := c.DownloadToBytes("README.md")
	assert.NoError(t, err)
	assert.NotEqual(t, len(toBytes), 0)
}

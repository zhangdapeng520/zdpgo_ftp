package zdpgo_ftp

import (
	"github.com/zhangdapeng520/zdpgo_ftp/ftp"
	"github.com/zhangdapeng520/zdpgo_log"
	"io"
	"os"
	"path/filepath"
)

/*
@Time : 2022/5/28 17:27
@Author : 张大鹏
@File : client.go
@Software: Goland2021.3.1
@Description:
*/

type Client struct {
	Conn *ftp.ServerConn
	Log  *zdpgo_log.Log
}

// Close 关闭客户端
func (c *Client) Close() error {
	return c.Conn.Quit()
}

// Upload 上传文件
func (c *Client) Upload(srcPath, destPath string) error {
	// 切割目标路径
	destDir, destFileName := filepath.Split(destPath)

	// 获取当前目录，比对
	currentDir, err := c.CurrentDir()
	if err != nil {
		c.Log.Error("获取当前工作目录失败", "error", err)
		return err
	}
	if destDir != "" && currentDir != destDir {
		err = c.MakeDir(destDir) // 不存在则创建目录
		if err != nil {
			c.Log.Error("创建目标目录失败", "error", err, "currentDir", currentDir, "destDir", destDir)
			return err
		}

		// 切换工作目录
		err = c.ChangeDir(destDir)
		if err != nil {
			c.Log.Error("切换工作目录失败", "error", err, "destDir", destDir)
			return err
		}
	}

	// 打开文件，存储在工作目录
	file, err := os.Open(srcPath)
	if err != nil {
		c.Log.Error("打开本地文件失败", "error", err, "srcPath", srcPath)
		return err
	}
	defer file.Close()
	err = c.Store(destFileName, file)
	if err != nil {
		c.Log.Error("存储文件失败", "error", err, "destFileName", destFileName)
	}

	// 返回
	return nil
}

// ChangeDir 切换工作目录
func (c *Client) ChangeDir(dirPath string) error {
	err := c.Conn.ChangeDir(dirPath)
	if err != nil {
		c.Log.Error("切换工作目录失败", "error", err, "dirPath", dirPath)
	}
	return err
}

// CurrentDir 获取当前目录
func (c *Client) CurrentDir() (string, error) {
	currentDir, err := c.Conn.CurrentDir()
	if err != nil {
		c.Log.Error("获取当前目录失败", "error", err)
		return "", err
	}
	return currentDir, nil
}

// MakeDir 创建目录
func (c *Client) MakeDir(dirPath string) error {
	err := c.Conn.MakeDir(dirPath)
	if err != nil {
		c.Log.Error("创建目录失败", "error", err)
		return err
	}
	return nil
}

// Store 存储文件
func (c *Client) Store(fileName string, fileContent io.Reader) error {
	err := c.Conn.Stor(fileName, fileContent)
	if err != nil {
		c.Log.Error("存储文件失败", "error", err, "fileName", fileName)
		return err
	}
	return nil
}

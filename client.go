package zdpgo_ftp

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ftp/ftp"
	"github.com/zhangdapeng520/zdpgo_log"
	"io"
	"io/ioutil"
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

// UploadAndCheckMd5 上传文件并校验MD5值
func (c *Client) UploadAndCheckMd5(srcPath, destPath string) bool {
	// 切割目标路径
	destDir, destFileName := filepath.Split(destPath)

	// 获取当前目录，比对
	currentDir, err := c.CurrentDir()
	if err != nil {
		c.Log.Error("获取当前工作目录失败", "error", err)
		return false
	}
	if destDir != "" && currentDir != destDir {
		err = c.MakeDir(destDir) // 不存在则创建目录
		if err != nil {
			c.Log.Error("创建目标目录失败", "error", err, "currentDir", currentDir, "destDir", destDir)
			return false
		}

		// 切换工作目录
		err = c.ChangeDir(destDir)
		if err != nil {
			c.Log.Error("切换工作目录失败", "error", err, "destDir", destDir)
			return false
		}
	}

	// 打开文件，存储在工作目录
	file, err := os.Open(srcPath)
	if err != nil {
		c.Log.Error("打开本地文件失败", "error", err, "srcPath", srcPath)
		return false
	}
	defer file.Close()

	// 获取md5
	srcFileData, err := ioutil.ReadFile(srcPath)
	if err != nil {
		c.Log.Error("读取源文件数据失败", "error", err)
		return false
	}
	srcFileMd5 := c.GetMd5(srcFileData)

	// 存储文件
	err = c.Store(destFileName, file)
	if err != nil {
		c.Log.Error("存储文件失败", "error", err, "destFileName", destFileName)
		return false
	}

	// 下载文件
	destFileBytes, err := c.DownloadToBytes(destFileName)
	if err != nil {
		c.Log.Error("下载目标文件失败", "error", err, "destFileName", destFileName)
		return false
	}

	// 比较
	flag := srcFileMd5 == c.GetMd5(destFileBytes)

	// 返回
	return flag
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

// Download 下载文件
func (c *Client) Download(srcPath, destPath string) error {
	fileBytes, err := c.Retr(srcPath)
	if err != nil {
		c.Log.Error("获取文件内容失败", "error", err)
		return err
	}

	dirPath, _ := filepath.Split(destPath)
	_ = os.MkdirAll(dirPath, os.ModePerm)

	err = ioutil.WriteFile(destPath, fileBytes, os.ModePerm)
	if err != nil {
		c.Log.Error("保存文件失败", "error", err)
		return err
	}

	return nil
}

// Retr 获取文件内容
func (c *Client) Retr(filePath string) ([]byte, error) {
	if c.Conn == nil {
		return nil, errors.New("连接对象为空，无法连接FTP服务")
	}

	response, err := c.Conn.Retr(filePath)
	if err != nil {
		c.Log.Error("获取文件失败", "error", err, "filePath", filePath)
		return nil, err
	}
	defer response.Close()

	fileBytes, err := ioutil.ReadAll(response)
	if err != nil {
		c.Log.Error("读取文件失败", "error", err)
		return nil, err
	}

	return fileBytes, nil
}

// DownloadToBytes 下载文件并转换我bytes数组
func (c *Client) DownloadToBytes(filePath string) ([]byte, error) {
	return c.Retr(filePath)
}

// GetMd5 获取数据的MD5值
func (c *Client) GetMd5(data []byte) string {
	has := md5.Sum(data)
	result := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return result
}

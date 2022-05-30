package zdpgo_ftp

/*
@Time : 2022/5/28 17:05
@Author : 张大鹏
@File : ftp.go
@Software: Goland2021.3.1
@Description: ftp操作核心方法
*/

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ftp/ftp"
	"github.com/zhangdapeng520/zdpgo_log"
	"time"
)

type Ftp struct {
	Config *Config
	Log    *zdpgo_log.Log
}

func New() *Ftp {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Ftp {
	f := &Ftp{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_ftp.log"
	}
	f.Log = zdpgo_log.NewWithDebug(config.Debug, config.LogFilePath)

	// 配置
	if config.WorkDir == "" {
		config.WorkDir = ".zdpgo_ftp_files"
	}
	if config.Port == 0 {
		config.Port = 34333
	}
	if config.Username == "" {
		config.Username = "zhangdapeng"
	}
	if config.Password == "" {
		config.Password = "zhangdapeng"
	}
	f.Config = config

	// 返回
	return f
}

func (f *Ftp) IsHealth() bool {
	//连接远程服务器
	addr := fmt.Sprintf("%s:%d", f.Config.Host, f.Config.Port)
	client, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return false
	}
	defer client.Quit()

	//登陆
	err = client.Login(f.Config.Username, f.Config.Password)
	return err == nil
}

// GetClient 获取客户端
func (f *Ftp) GetClient() (*Client, error) {
	//连接远程服务器
	addr := fmt.Sprintf("%s:%d", f.Config.Host, f.Config.Port)
	client, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		f.Log.Error("连接FTP服务器失败", "error", err, "config", f.Config)
		return nil, err
	}

	//登陆
	err = client.Login(f.Config.Username, f.Config.Password)
	if err != nil {
		f.Log.Error("登录FTP服务器失败", "error", err, "config", f.Config)
		return nil, err
	}

	// 创建客户端对象
	result := &Client{
		Conn: client,
		Log:  f.Log,
	}

	// 返回
	return result, nil
}

// GetServer 获取服务端对象
func (f *Ftp) GetServer() *Server {
	return &Server{
		Config: f.Config,
		Log:    f.Log,
	}
}

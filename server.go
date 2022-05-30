package zdpgo_ftp

import (
	"github.com/zhangdapeng520/zdpgo_ftp/server/core"
	"github.com/zhangdapeng520/zdpgo_ftp/server/driver/file"
	"github.com/zhangdapeng520/zdpgo_log"
	"os"
)

/*
@Time : 2022/5/30 9:36
@Author : 张大鹏
@File : server.go
@Software: Goland2021.3.1
@Description:
*/

type Server struct {
	Config *Config
	Log    *zdpgo_log.Log
}

// Run 运行服务
func (s *Server) Run() error {
	// 创建工作目录
	err := os.MkdirAll(s.Config.WorkDir, os.ModePerm)
	if err != nil {
		s.Log.Error("创建工作目录失败", "error", err, "dirPath", s.Config.WorkDir)
		return err
	}

	// 创建权限
	var perm = core.NewSimplePerm("zdpgo_ftp", "zdpgo_ftp")
	opt := &core.ServerOpts{
		Name: "zdpgo_ftp",
		Factory: &file.DriverFactory{
			RootPath: s.Config.WorkDir,
			Perm:     perm,
		},
		Port: s.Config.Port,
		Auth: &core.SimpleAuth{
			Name:     s.Config.Username,
			Password: s.Config.Password,
		},
		Logger: new(core.DiscardLogger),
	}

	// 创建服务
	server := core.NewServer(opt)

	// 启动服务
	err = server.ListenAndServe()
	if err != nil {
		s.Log.Error("启动服务失败", "error", err)
		return err
	}

	return nil
}

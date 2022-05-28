package server

import (
	"github.com/zhangdapeng520/zdpgo_ftp/server/core"
	"github.com/zhangdapeng520/zdpgo_ftp/server/driver/file"
	"github.com/zhangdapeng520/zdpgo_ftp/server/driver/minio"
)

// Backwards compatible types for the server code.
//
// New code should import github.com/zhangdapeng520/zdpgo_ftp/server/core
type (
	Auth                  = core.Auth
	Command               = core.Command
	Conn                  = core.Conn
	DataSocket            = core.DataSocket
	DiscardLogger         = core.DiscardLogger
	Driver                = core.Driver
	DriverFactory         = core.DriverFactory
	FileInfo              = core.FileInfo
	Logger                = core.Logger
	MultipleDriver        = core.MultipleDriver
	MultipleDriverFactory = core.MultipleDriverFactory
	Notifier              = core.Notifier
	NullNotifier          = core.NullNotifier
	Perm                  = core.Perm
	Server                = core.Server
	ServerOpts            = core.ServerOpts
	SimpleAuth            = core.SimpleAuth
	SimplePerm            = core.SimplePerm
	StdLogger             = core.StdLogger
)

// Backwards compatible functions and variables for the server code.
//
// New code should import github.com/zhangdapeng520/zdpgo_ftp/server/core
var (
	ErrServerClosed = core.ErrServerClosed
	NewServer       = core.NewServer
	NewSimplePerm   = core.NewSimplePerm
	Version         = core.Version
)

// Backwards compatible types for the file driver code.
//
// New code should import github.com/zhangdapeng520/zdpgo_ftp/server/driver/file
type (
	FileDriver        = file.Driver
	FileDriverFactory = file.DriverFactory
)

// Backwards compatible types for the minio driver code.
//
// New code should import github.com/zhangdapeng520/zdpgo_ftp/server/driver/minio
type (
	MinioDriver        = minio.Driver
	MinioDriverFactory = minio.DriverFactory
)

// Backwards compatible functions for the minio driver code.
//
// New code should import github.com/zhangdapeng520/zdpgo_ftp/server/driver/minio
var NewMinioDriverFactory = minio.NewDriverFactory

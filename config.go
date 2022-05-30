package zdpgo_ftp

/*
@Time : 2022/5/28 17:06
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description:
*/

type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"`
	WorkDir     string `yaml:"root_dir" json:"root_dir"`
	Host        string `yaml:"host" json:"host"`
	Port        int    `yaml:"port" json:"port"`
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
}

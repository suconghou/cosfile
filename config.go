package main

var Config struct {
	Version   string
	Name      string
	AppID     string
	SecretID  string
	SecretKey string
	Bucket    string
	ApiUrl    string
}

func init() {
	Config.Version = "1.0"
	Config.Name = "cosfile"
	Config.AppID = "1251010988"
	Config.SecretID = "AKIDd19fjBLrtOIPmYR7KHrmDnSqLikNKHGT"
	Config.SecretKey = "Zf0nJIoiIB9ak0qsbgIt1mtigUzYbYEn"
	Config.Bucket = "filekeeper"
	Config.ApiUrl = "http://filekeeper-1251010988.costj.myqcloud.com/"
}

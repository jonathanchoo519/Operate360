package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"golang.org/x/sys/windows/registry"
	"time"
)

func main() {
	var filePath string
	//访问注册表查看360安装目录是否存在
	k,err := registry.OpenKey(registry.LOCAL_MACHINE,`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\360sd.exe`,registry.QUERY_VALUE)
	if err != nil{
		fmt.Println("未检测到安装360杀毒")
	}else {
		s,_,err := k.GetStringValue("Path")
		filePath = s
		if err != nil{
			panic(err)
		}
		fmt.Println("检测到已安装360杀毒")
		fmt.Printf("安装目录为：%s\n",s)
	}
	defer k.Close()


	configFile,err := goconfig.LoadConfigFile(filePath+"\\setting.ini")
	var currentDate string

	if err!=nil {
		panic(err)
	}else {
		value,err := configFile.GetValue("engupdate","1")
		if err!=nil {
			panic(err)
		}else {
			currentDate = string([]rune(value)[:10])
			fmt.Println("本地病毒库日期:",currentDate)
		}
	}


	newDate,_ := SpiderPage()
	fmt.Println("最新的病毒库日期：",newDate)

	currentTime, _ := time.Parse("2006-01-02 15:04:05",currentDate)
	newTime,_ := time.Parse("2006-01-02 15:04:05",newDate)

	if currentTime.Before(newTime) {
		fmt.Println("执行下载更新...")
		DownLoadAndExec(newDate)

	}else {
		fmt.Println("已是最新病毒库版本")
	}

}







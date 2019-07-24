package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func DownLoadAndExec(data string){
	sysBit := 32 << (^uint(0) >> 63)

	ch := make(chan bool)
	datapath := "360Update_"+data+".exe"
	go func() {

		ch <- true

		/*
			360杀毒离线病毒库
			32位：http://down.360safe.com/offline/360sd-upd.exe
			64位：http://down.360safe.com/offline/360sd-upd-x64.exe
		*/
		var (
			url32 = "http://down.360safe.com/offline/360sd-upd.exe"
			url64 = "http://down.360safe.com/offline/360sd-upd-x64.exe"
		)

		if sysBit == 64 {
			res, err := http.Get(url64)
			if err != nil {
				panic(err)
			}

			f, err := os.Create(datapath)
			if err != nil {
				panic(err)
			}
			io.Copy(f, res.Body)
		}else {
			res, err := http.Get(url32)
			if err != nil {
				panic(err)
			}

			f, err := os.Create(datapath)
			if err != nil {
				panic(err)
			}
			io.Copy(f, res.Body)
		}

	}()

	<-ch

	cmd := exec.Command("cmd.exe", "/c", "start "+datapath)
	cmd.Run()

}

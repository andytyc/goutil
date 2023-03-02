package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/andytyc/goutil/app/media/video"
)

var FlagFileName string

func init() {
	flag.StringVar(&FlagFileName, "f", "", "本地文件路径")
	flag.Parse()

	log.Println("FlagFileName :", FlagFileName)
}

func main() {
	funcGetVideoInfo := func(localFilename string) {
		tagmsg := "funcGetVideoInfo"

		ret, videoFileInfo, err := video.GetMediaFileInfo(localFilename)
		if err != nil {
			log.Panicln(tagmsg, "err", err)
			return
		}
		log.Println("file :", localFilename, "run command ret :", ret)

		if videoFileInfo.Format.Size == "" {
			err = fmt.Errorf("videoFileInfo.Format.Size is empty :%s", localFilename)
			log.Panicln(tagmsg, "err", err)
			return
		}

		fileSize, err := strconv.ParseInt(videoFileInfo.Format.Size, 10, 64)
		if err != nil {
			err = fmt.Errorf("videoFileInfo.Format.Size strconv.ParseInt failed :%s:%s", localFilename, videoFileInfo.Format.Size)
			log.Panicln(tagmsg, "err", err)
			return
		}
		log.Println(tagmsg, "文件大小", fileSize)

		totalFrame, err := video.GetVideoTotalFrame(videoFileInfo)
		if err != nil {
			log.Panicln(tagmsg, "err", err)
			return
		}
		log.Println(tagmsg, "文件总帧数", totalFrame)
	}

	funcGetVideoInfo(FlagFileName)
}

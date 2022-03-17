package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"web-dl/api"
)

const (
	PPT_SUFFIX     = "-ppt"
	TEACHER_SUFFIX = "-teacher"
	STUDENT_SUFFIX = "-student"
)

var (
	helpFlag      bool
	pptFlag       bool
	teacherFlag   bool
	studentFlag   bool
	filePathFlag  string
	outputDirFlag string
)

func init() {
	flag.BoolVar(&helpFlag, "h", false, "使用说明")
	flag.BoolVar(&pptFlag, "p", false, "是否下载ppt录播")
	flag.BoolVar(&teacherFlag, "t", false, "是否下载教师端录播")
	flag.BoolVar(&studentFlag, "s", false, "是否下载学生端录像")
	flag.StringVar(&filePathFlag, "i", "", "json文件路径")
	flag.StringVar(&outputDirFlag, "o", "", "视频输出目录")
}

func main() {
	fmt.Println("学在西电录播下载器")
	flag.Parse()
	if helpFlag {
		flag.Usage()
		return
	}
	file, err := os.Open(filePathFlag)
	if err != nil {
		panic(err)
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	xdlb := api.Xdlb{}
	err = json.Unmarshal(fileBytes, &xdlb)
	if err != nil {
		panic(err)
	}
	if outputDirFlag == "" {
		outputDirFlag = filePathFlag[strings.LastIndexAny(filePathFlag, `/\`)+1 : strings.LastIndexByte(filePathFlag, '.')]
	}
	if pptFlag {
		api.M3u8dl(xdlb.VideoPath.PptVideo, outputDirFlag+PPT_SUFFIX)
	}
	if teacherFlag {
		api.M3u8dl(xdlb.VideoPath.TeacherTrack, outputDirFlag+TEACHER_SUFFIX)
	}
	if studentFlag {
		api.M3u8dl(xdlb.VideoPath.StudentFull, outputDirFlag+STUDENT_SUFFIX)
	}
}

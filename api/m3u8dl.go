package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const DL_PATH = "dl/"
const TSREGEXP = `.*\.ts`

func M3u8dl(m3u8Url string, name string) {
	fmt.Printf("start download %v\n", name)
	// dir
	vedioDir := DL_PATH + name
	_, err := os.Stat(vedioDir)
	if err != nil {
		os.MkdirAll(vedioDir, os.ModePerm)
	}
	// m3u8
	urlBase := m3u8Url[0:strings.LastIndexByte(m3u8Url, '/')]
	resp, err := http.Get(m3u8Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(bodyBytes)
	m3u8File, err := os.Create(vedioDir + "/0.m3u8")
	if err != nil {
		panic(err)
	}
	defer m3u8File.Close()
	io.Copy(m3u8File, strings.NewReader(body))
	// tslist
	tslistFileName := vedioDir + "/tslist.txt"
	tsListFile, err := os.Create(tslistFileName)
	if err != nil {
		panic(err)
	}
	// ts
	tsregexp := regexp.MustCompile(TSREGEXP)
	tsNames := tsregexp.FindAllString(body, -1)
	for i, tsName := range tsNames {
		io.Copy(tsListFile, strings.NewReader("file '"+tsName+"'\n"))
		filePath := DL_PATH + name + "/" + tsName
		_, err := os.Stat(filePath)
		if err != nil {
			tsUrl := urlBase + "/" + tsName
			resp, err := http.Get(tsUrl)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			tsFile, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
			defer tsFile.Close()
			written, err := io.Copy(tsFile, resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\r%v/%v, size: %vBytes", i, len(tsNames), written)
		} else {
			fmt.Printf("\r%v skip", i)
		}
	}
	fmt.Println()
	fmt.Println("Done.")
	fmt.Println("start concat ts")
	outputFileName := vedioDir + ".mp4"
	_, err = os.Stat(outputFileName)
	if err != nil {
		concatCmd := exec.Command("ffmpeg", "-f", "concat", "-i", tslistFileName, "-c", "copy", outputFileName)
		out, err := concatCmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(out))
			panic(err)
		}
	}
	fmt.Println("Done.")
}

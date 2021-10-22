// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/9/13 10:20
// @Update  姓名（需要改）  2021/9/13 10:20
package sample_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadTagsFromFile(fileName string){
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开错误！", err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("读取文件错误！", err)
			}
		}
		fmt.Println(line)
	}
}
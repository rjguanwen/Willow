// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/9/13 10:47
// @Update  姓名（需要改）  2021/9/13 10:47
package sample_1

import "testing"

func TestReadTagsFromFile(t *testing.T) {
	fileName := "../data/tags-json-sample.txt"
	ReadTagsFromFile(fileName)
}
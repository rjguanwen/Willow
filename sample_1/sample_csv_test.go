// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/10/22 15:25
// @Update  姓名（需要改）  2021/10/22 15:25
package sample_1

import "testing"

func TestReadCsv(t *testing.T) {
	filename := "../data/tags-all.csv"
	ReadCsv(filename)
}
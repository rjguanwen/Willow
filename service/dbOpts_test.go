// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2021/9/15 17:45
// @Update  姓名（需要改）  2021/9/15 17:45
package service

import (
	"log"
	"testing"
)

func TestQueryMembersByTags(t *testing.T) {
	SetTagsFile("../data/tags-json-sample.txt")
	// 初始化
	TagsDb = InitDB()
	log.Println(QueryMembersByTags("dp_dc", "1"))
}
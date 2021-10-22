// @Title  dbOpts
// @Description  数据库相关操作
// @Author  zhengbin  2021/9/15 15:48
package service

import (
	log "github.com/cihub/seelog"
	"github.com/rjguanwen/willow/public"
	"strconv"
)

// @title    QueryMembersByTags
// @description   通过标签查询成员记录
// @auth      zhengbin     2021/9/15 15:48
// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"
func QueryMembersByTags(index string, args string) ([]*public.Tags, error){
	log.Info("dbOpts - QueryMembersByTags() - begin...")
	// 创建只读事务
	txn := TagsDb.Txn(false)
	defer txn.Abort()

	// 将传入的 string 类型参数转换为 int
	tv, err := strconv.Atoi(args)

	// 根据索引查询
	it, err := txn.Get("memberTags", index, tv)
	if err != nil {
		//panic(err)
		return nil, err
	}

	rtnSlice := make([]*public.Tags, 10)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		member := obj.(*public.Tags)
		rtnSlice = append(rtnSlice, member)
	}
	log.Info("dbOpts - QueryMembersByTags() - end!")
	return rtnSlice, nil
}
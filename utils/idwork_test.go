// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  ligx   2020/7/23 8:53
// @Update  姓名（需要改）  2020/7/23 8:53
package utils

import (
	"fmt"
	"sync"
	"testing"
)

type Id struct{
	id map[int64]int
	rm sync.RWMutex
}

var idwork *Worker
var ids *Id
var e map[int64]int

// @title    函数名称
// @description   函数的详细描述
// @auth      ligx             2020/7/23 8:53
// @param     输入参数名        参数类型         "解释"
// @return    返回参数名        参数类型         "解释"
func TestWorker_GetId(t *testing.T) {

	var wg sync.WaitGroup
	idwork, _ = NewWorker(int64(100))
	ids=NewConnManger()
	//fmt.Printf("bit:[%d]\n",(-1 ^ (-1 << 10)))
	for i:=0;i<30;i++{
		fmt.Printf("start i:[%d]\n",i)
		wg.Add(1)
		go func() {
			CreateId(&wg)
		}()
	}
	wg.Wait()
	fmt.Printf("ids[%d]\n",ids.id)
	for j:=range ids.id{
		if ids.id[j]>1{
			e[j]=ids.id[j]
		}
	}
	fmt.Printf("e[%d]\n",e)
}

func NewConnManger() *Id {
	cm := &Id{
		id: make(map[int64]int),
	}
	return cm
}

func CreateId(wg *sync.WaitGroup) {

	for i := 0; i < 100000; i++ {
		id := idwork.GetId()
		ids.rm.Lock()
		i := ids.id[id]
		ids.id[id] = i + 1
		ids.rm.Unlock()
	}
	wg.Done()
}

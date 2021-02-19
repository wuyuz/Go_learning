package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"114.215.84.163:2379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 创建连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	ctx1,concalFunc1 := context.WithCancel(context.TODO())

	// 模拟etcd中的kv变化
	go func(ctx context.Context) {
		for{
			select {
			case <-ctx.Done():
				fmt.Println("协程关闭")
				goto DONE
			default:
				kv.Put(context.TODO(), "/cron/jobs/job7", "I'm job 7")
				kv.Delete(context.TODO(), "/cron/jobs/job7")
				time.Sleep(1 * time.Second)
			}
		}
		DONE:
	}(ctx1)

	// 获取当前的值，并监听后续的变化
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}

	// 现在是有值的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值：", string(getResp.Kvs[0].Value))
	}
	// 当前etcd的集群事务的ID，单增
	watchStartRevision = getResp.Header.Revision + 1

	// 创建一个监听器
	watcher = clientv3.NewWatcher(client)
	// 启动监听
	fmt.Println("从该版本开始监听：", watchStartRevision)

	ctx, cancelFunc := context.WithCancel(ctx1)
	// 五秒后定时取消，相当于5秒后不在监听
	time.AfterFunc(5*time.Second, func() {
		concalFunc1()
		time.Sleep(1*time.Second)
		cancelFunc()
	})

	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, )
			case mvccpb.DELETE:
				fmt.Println("删除了,Revision:",event.Kv.ModRevision)
			}
		}
	}

}

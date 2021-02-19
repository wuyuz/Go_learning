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
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
		kvpair  *mvccpb.KeyValue
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

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil { // Get方法中有很多参数配置
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs) // [key:"/cron/jobs/job1" create_revision:4 mod_revision:13 version:7 value:"wang" ] version表示被修改的次数
	}

	//kv.Put(context.TODO(),"/cron/jobs/job2","{xxx}")

	// 读取/cron/jobs/为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		// [key:"/cron/jobs/job1" create_revision:4 mod_revision:13 version:7 value:"wang"  key:"/cron/jobs/job2" create_revision:14 mod_revision:14 version:1 value:"{xxx}" ]
		fmt.Println(getResp.Kvs)
	}

	// 删除,当然也可以指定到"/cron/jobs/"来删除所有子级目录
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job2", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		// 被删除之前的val
		if len(delResp.PrevKvs) != 0 {
			for _, kvpair = range delResp.PrevKvs {
				fmt.Println("删除了：", string(kvpair.Key))
			}
		}
	}
}

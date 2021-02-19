package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
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

	// 申请一个lease租约
	lease = clientv3.NewLease(client)

	// 申请一个10秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约ID
	leaseId = leaseGrantResp.ID

	// 5秒后自动取消自动续租，也就是说续租5秒，停止了还有10秒，所以一共15秒生命期
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)

	// 自动续租,不断的刷新租约id
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 消费续约,处理续约应答
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:  // 自动续租取消后，接收到的keepResp为nil
				if keepRespChan == nil || keepResp == nil {
					fmt.Println("租约已经失效")
					goto END
				} else  {
					fmt.Println("收到自动续租应答", keepResp)
				}
			}
			time.Sleep(100*time.Millisecond)
		}
		END:
	}()

	// 获取kv对象
	kv = clientv3.NewKV(client)

	// put一个kv，让它和租约相关联起来，从而实现10秒过期
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("写入成功：", putResp.Header.Revision)
	}

	// 定时看一下key过期了吗
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期")
			break
		}
		fmt.Println("kv没过期", getResp.Kvs)
		time.Sleep(1 * time.Second)
	}

}

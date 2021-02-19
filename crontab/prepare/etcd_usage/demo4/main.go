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
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
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

	// 上锁（创建租约、自动续租、拿着租约抢占一个key）
	lease = clientv3.NewLease(client)
	//申请租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}
	// 拿到租约的ID
	leaseId = leaseGrantResp.ID

	// 准备一个用于取消续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())
	defer cancelFunc()                          // 确保函数退出后，自动续租停止，比如服务宕机
	defer lease.Revoke(context.TODO(), leaseId) // 立即实现锁和租约

	// 自动续租10秒
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约的应答展示协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil || keepResp == nil {
					fmt.Println("租约失效")
					goto END
				} else {
					fmt.Println("收到自动续租应答：", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 抢锁 if 不存在key， then 设置它 else 强锁失败
	kv = clientv3.NewKV(client)
	// 创建事务
	txn = kv.Txn(context.TODO())
	// 定义事务
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).Then(
		clientv3.OpPut("/cron/lock/job9", "xxxx", clientv3.WithLease(leaseId))).Else(
		clientv3.OpGet("/cron/lock/job9")) // 抢锁失败后打印下

	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	} // 提交事务

	// 判读是否抢锁成功
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 处理业务
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)

	// 释放锁（取消自动续租、释放租约立即）上面的两个defer会把租约释放掉，关联的kv会被删除

}

package tests

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
	"unioj-Judger/service"
)

func TestEtcd(t *testing.T) {
	service.InitEtcd(5 * time.Second)
}

func TestEtcdPut(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"39.105.152.206:20000", "39.105.152.206:20002", "39.105.152.206:20004"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed: ", err)
		return
	}
	fmt.Println("connect to etcd succeeded")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	_, err = cli.Put(ctx, "/logagent/192.168.4.147/colletc_config", `[{"topic":"web_log","path":"/home/et/Desktop/tmp/logs/web_log.log"},{"topic":"redis_log","path":"/home/et/Desktop/tmp/logs/redis_log.log"},{"topic":"mysql_log","path":"/home/et/Desktop/tmp/logs/mysql_log.log"}]`)
	cancel()
	if err != nil {
		fmt.Println("put failed  ", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	res, err := cli.Get(ctx, "/logagent/colletc_config", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed,err:", err)
		return
	}
	kvs := res.Kvs
	val := string(kvs[0].Value)
	log.Println("result:", val, val == "dsb")
}

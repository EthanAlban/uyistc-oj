package service

// 使用etcd作为判题器集群的配置中心  加入watcher监控配置项目的改变

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"go.etcd.io/etcd/clientv3"
	"time"
	"unioj-Judger/conf"
)

var (
	cli *clientv3.Client
)

type EtcdConf struct {
	KafkaHost       string `json:"kafka_host"`
	JudgerTaskTopic string `json:"judger_task_topic"`
	SqlHost         string `json:"sql_host"`
	SqlUser         string `json:"sql_user"`
	SqlPassword     string `json:"sql_password"`
}

// InitEtcd Init 初始化etcd，连接etcd，设置连接超时时间
func InitEtcd(timeout time.Duration) (err error) {
	etcdHost := conf.CFG.Section("server").Key("host").String()
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdHost},
		DialTimeout: timeout,
	})
	if err != nil {
		logger.Error("连接etcd失败... ", err)
		return
	}
	logger.Debug("etcd连接成功...")
	return
}

// GetConf 加载etcd中的配置文件
func GetConf(key string) (logEntryConf *EtcdConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		logger.Error("从etcd拉取key=%v失败...  err:%v\n", key, err)
		return
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s %s \n",ev.Key,ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			logger.Error("unmarshal etcd value %v falied err:%v\n", ev.Value, err)
			return
		}
	}
	return
}

func WatchConf(key string, newConfChan chan<- []*EtcdConf) {
	ch := cli.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("配置项目更新 type:%v  key:%v  value:%v \n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			//通知taillog.tskMgr
			//先判断操作的类型
			var newConf []*EtcdConf
			if evt.Type != clientv3.EventTypeDelete {
				//如果是删除操作,手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("json unmarshal failed,err:", err)
					continue
				}
			}
			fmt.Printf("get new conf%v\n", newConf)
			newConfChan <- newConf
		}
	}
}

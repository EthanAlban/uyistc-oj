package tests

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"testing"
	"time"
)

//集群地址
var address = []string{"39.105.152.206:9092"}

//创建kafka任务发布者
func TestKafkaProducter(t *testing.T) {
	//配置发布者
	config := sarama.NewConfig()
	//确认返回，记得一定要写，因为本次例子我用的是同步发布者
	config.Producer.Return.Successes = true
	//设置超时时间 这个超时时间一旦过期，新的订阅者在这个超时时间后才创建的，就不能订阅到消息了
	config.Producer.Timeout = 5 * time.Second
	//连接发布者，并创建发布者实例
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	//程序退出时释放资源
	defer p.Close()
	//设置一个逻辑上的分区名，叫安彦飞
	topic := "anyanfei"
	//这个是发布的内容
	srcValue := "sync: this is a message. index=%d"
	//发布者循环发送0-9的消息内容
	for i := 0; i < 100; i++ {
		value := fmt.Sprintf(srcValue, i)
		//创建发布者消息体
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		//发送消息并返回消息所在的物理分区和偏移量
		partition, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", partition, offset)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

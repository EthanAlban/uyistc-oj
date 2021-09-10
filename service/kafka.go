package service

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/go-ini/ini"
	uuid "github.com/satori/go.uuid"
	"github.com/wonderivan/logger"
	"strings"
	"sync"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/service/structs"
)

var DeathKafka = make(chan bool)

// NewKafkaConsumer 开始创建kafka订阅者
func NewKafkaConsumer(kafka_host string) {
	/**
	  group:
	   设置订阅者群 如果多个订阅者group一样，则随机挑一个进行消费，当然也可以设置轮训，在设置里面修改；
	   若多个订阅者的group不同，则一旦发布者发布消息，所有订阅者都会订阅到同样的消息；
	  topics:
	   逻辑分区必须与发布者相同，还是用安彦飞，不然找不到内容咯
	   当然订阅者是可以订阅多个逻辑分区的，只不过因为演示方便我写了一个，你可以用英文逗号分割在这里写多个
	*/
	cfg, err := ini.Load("../conf/conf.ini")
	if err != nil {
		logger.Error("Fail to read file: ", err)
	}
	//topics := cfg.Section("client").Key("kafka_topic").String()
	topics := "judger_task" + uuid.NewV4().String()
	logger.Debug("新的topic: ", topics)
	cfg.Section("client").Key("kafka_topic").SetValue(topics)
	err = cfg.SaveTo("../conf/conf.ini")
	if err != nil {
		logger.Error(err)
	}
	conf.CFG = cfg
	var (
		group = "Consumer1"
		//topics = "judger_task"
	)
	logger.Info("kafka消费者启动中...")
	//配置订阅者
	config := sarama.NewConfig()
	//config.Consumer.Group.Session.Timeout = 2 * time.Second
	//config.Consumer.Group.Heartbeat.Interval  = 1 * time.Second
	//配置偏移量
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Net.WriteTimeout = 10 * time.Second
	//开始创建订阅者
	consumer := Consumer{
		ready: make(chan bool),
	}
	//创建一个上下文对象，实际项目中也一定不要设置超时（当然，按你项目需求，我是没见过有项目需求要多少时间后取消订阅的）
	ctx, _ := context.WithCancel(context.Background())
	//创建订阅者群，集群地址发布者代码里已定义
	client, err := sarama.NewConsumerGroup([]string{kafka_host}, group, config)
	if err != nil {
		logger.Fatal("创建kafka消费者失败，", err)
	}

	//创建同步组
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// 重试次数
		defaultRetries, _ := conf.CFG.Section("client").Key("kafka_retry").Int()
		count := defaultRetries
		defer wg.Done()
		for {
			/**
			  官方说：`订阅者`应该在无限循环内调用
			  当`发布者`发生变化时
			  需要重新创建`订阅者`会话以获得新的声明
			  所以这里把订阅者放在了循环体内
			*/
			err := client.Consume(ctx, strings.Split(topics, ","), &consumer)
			if err != nil {
				logger.Error("kafka消费者错误 重试次数：", count, "  err:", err)
				count--
				time.Sleep(time.Second * 2)
				if count == 0 {
					return
				}
			} else if count < defaultRetries {
				logger.Debug("kafka恢复，重置重试次数...")
				count = defaultRetries
			}

			// 检查上下文是否被取消，收到取消信号应当立刻在本协程中取消循环
			if ctx.Err() != nil {
				logger.Error("ctx.Err() != nil")
				return
			}

			//获取订阅者准备就绪信号
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // 获取到了订阅者准备就绪信号后打印下面的话
	logger.Info("kafka消费者启动成功...,topic: " + topics)
	//golang优雅退出的信号通道创建
	//sigterm := make(chan os.Signal, 1)
	////golang优雅退出的信号获取
	//signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	////创建选择器，如果不是上下文取消或者用户ctrl+c这种系统级退出，则就不向下执行了
	//select {
	//case <-ctx.Done():
	//	logger.Debug("context被取消了...")
	//case <-sigterm:
	//	logger.Debug("接收到终止信号...")
	//}
	//logger.Debug("开始清退kafka环境...")
	//////取消上下文
	//cancel()
	wg.Wait()
	//关闭客户端
	if err = client.Close(); err != nil {
		logger.Error("关闭kafka消费者出现错误 ", err)
	}
	DeathKafka <- true
	return
}

// Consumer 重写订阅者，并重写订阅者的所有方法
type Consumer struct {
	ready chan bool
}

// Setup Setup方法在新会话开始时运行的，然后才使用声明
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup 一旦所有的订阅者协程都退出，Cleaup方法将在会话结束时运行
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim 订阅者在会话中消费消息，并标记当前消息已经被消费。
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		//log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		var submission structs.Submission
		err := json.Unmarshal([]byte(string(message.Value)), &submission)
		if err != nil {
			logger.Error(err)
		}
		go JudgeUserCode(submission)
		session.MarkMessage(message, "")
	}
	return nil
}

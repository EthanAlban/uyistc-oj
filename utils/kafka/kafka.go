package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/wonderivan/logger"
	"unioj/models"
)

func KafkaHealthCheck(host string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "healthcheck"
	msg.Value = sarama.StringEncoder("this is a healthcheck log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		logger.Error("[5] ERROR 初始化kafka失败...,error:" + err.Error())
		logger.Error("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	_, _, err = client.SendMessage(msg)
	if err != nil {
		logger.Error("[5] ERROR 初始化kafka失败...,error:", err)
		return
	}
	logger.Debug("[5] INFO 初始化kafka成功...")
}

func SendToKafka(host, topic string, task models.Submission) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	bytes, _ := json.Marshal(task)
	msg.Value = sarama.StringEncoder(string(bytes))
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		logger.Error("ERROR kafka生产者异常...,error:", err)
		return err
	}
	defer client.Close()
	// 发送消息
	_, _, err = client.SendMessage(msg)
	if err != nil {
		logger.Error("ERROR kafka发送消息失败,err:", err)
		return err
	}
	logger.Debug("消息发送成功")
	return err
}

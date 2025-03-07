package kafka

import (
	"fmt"
	kk "go_server/pkg/util/kafka"
	"go_server/pkg/util/log"

	"github.com/spf13/viper"
)

var Writer *kk.KafkaWriter
var Reader *kk.KafkaReader

func init() {
	Writer = &kk.KafkaWriter{}
	Reader = &kk.KafkaReader{}
}

/*
初始化kafka加载配置

	param:
		readKafka bool 是消费者还是生产者
	return:
		error 错误信息
*/
func WriterSetup() (err error) {
	topic := viper.GetString("kafka.topic")
	brokers := viper.GetStringSlice("kafka.brokers")
	if len(brokers) == 0 {
		return fmt.Errorf("kafka.brokers不能为空")
	}
	err = kk.CreateKafkaTopic(topic, brokers[0], 1, len(brokers))
	if err != nil {
		return fmt.Errorf("[Kafka Connect Url:%v]:%v", brokers, err)
	}
	err = Writer.Init(brokers, topic, &kk.RoundRobin{})
	if err != nil {
		return err
	}
	log.Info("kafka消息队列服务已连接,Brokers:[%+v]", brokers)
	return nil
}

func ReaderSetup() (err error) {
	topic := viper.GetString("kafka.topic")
	brokers := viper.GetStringSlice("kafka.brokers")
	if len(brokers) == 0 {
		return fmt.Errorf("kafka.brokers不能为空")
	}
	err = Reader.Init(brokers, topic,
		viper.GetInt("kafka.minBytes"),
		viper.GetInt("kafka.maxBytes"),
		viper.GetInt("kafka.partition"),
		viper.GetString("kafka.group"))
	if err != nil {
		return err
	}
	log.Info("kafka消息队列服务已连接,Brokers:[%+v]", brokers)
	return nil
}

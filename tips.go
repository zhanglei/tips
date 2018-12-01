package tips

import (
	"context"
	"fmt"

	"github.com/shafreeck/tips/store/pubsub"
)

var (
	ErrNotFound = "%s can not found"
)

type Tips struct {
	ps *pubsub.Pubsub
}
type Topic struct {
	pubsub.Topic
}
type Subscription struct {
	pubsub.Subscription
}

func NewTips(path string) (tips Pubsub, err error) {
	ps, err := pubsub.Open(path)
	if err != nil {
		return nil, err
	}
	return &Tips{
		ps: ps,
	}, nil
}

//创建一个topic
func (ti *Tips) CreateTopic(cxt context.Context, topic string) error {
	txn, err := ti.ps.Begin()
	if err != nil {
		return err
	}
	if _, err = txn.CreateTopic(topic); err != nil {
		return err
	}
	if err = txn.Commit(cxt); err != nil {
		return err
	}
	return nil

}

//查看当前topic订阅信息
func (ti *Tips) Topic(cxt context.Context, name string) (*Topic, error) {
	txn, err := ti.ps.Begin()
	if err != nil {
		return nil, err
	}
	//查看当前topic是否存在
	t, err := txn.GetTopic(name)
	if err == pubsub.ErrNotFound {
		return nil, fmt.Errorf(ErrNotFound, "topic")
	}

	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	topic.Topic = *t
=======
	//如果存在则返回topic信息
	topic := &Topic{Topic: *t}
>>>>>>> a899340c7635d36613f2d5b7a6bed82b6ef47a12

	if err = txn.Commit(cxt); err != nil {
		return nil, err
	}
	return topic, nil
}

//销毁一个topic
func (ti *Tips) Destroy(cxt context.Context, topic string) error {
	txn, err := ti.ps.Begin()
	if err != nil {
		return err
	}
	if err = txn.DeleteTopic(topic); err != nil {
		return err
	}
	if err = txn.Commit(cxt); err != nil {
		return err
	}
	return nil
}

//Publish 消息下发 支持批量下发,返回下发成功的msgids
//msgids 返回的序列和下发消息序列保持一直
func (ti *Tips) Publish(cxt context.Context, msg []string, topic string) ([]string, error) {
	//获取当前topic
	txn, err := ti.ps.Begin()
	if err != nil {
		return nil, err
	}
	//查看当前topic是否存在
	t, err := txn.GetTopic(topic)
	//如果当前的topic不存在，那么返回错误
	if err == pubsub.ErrNotFound {
		return nil, fmt.Errorf(ErrNotFound, "topic")
	}

	if err != nil {
		return nil, err
	}
	//将传递进来的msg转化成Append需要的格式
	message := make([]*pubsub.Message, len(msg))
	for i := range msg {
		message[i] = &pubsub.Message{
			Payload: []byte(msg[i]),
		}
	}
	//如果当前的topic存在 则调用Append接口将消息存储到对应的topic下
	// f func(topic *pubsub.Topic, messages ...*pubsub.Message) ([]pubsub.MessageID, error)i
	messageID, err := txn.Append(t, message...)
	if err != nil {
		return nil, err
	}
	if err = txn.Commit(cxt); err != nil {
		return nil, err
	}
	MessageID := make([]string, len(messageID))
	for i := range messageID {
		MessageID[i] = messageID[i].String()
	}

	return MessageID, nil
}

func (ti *Tips) Ack(cxt context.Context, msgids []string) (err error) {
	return nil
}

//Subscribe 创建topic 和 subscription 订阅关系
func (ti *Tips) Subscribe(cxt context.Context, subName string, topic string) (*Subscription, error) {
	txn, err := ti.ps.Begin()
	if err != nil {
		return nil, err
	}
	//查看当前topic是否存在
	t, err := txn.GetTopic(topic)
	//如果当前的topic不存在，那么返回错误
	if err != nil {
		return nil, err
	}
	//func (txn *Transaction) CreateSubscription(t *Topic, name string) (*Subscription, error)
	s, err := txn.CreateSubscription(t, subName)
	if err != nil {
		return nil, err
	}
	if err = txn.Commit(cxt); err != nil {
		return nil, err
	}
	sub := &Subscription{}
	sub.Subscription = *s
	return sub, nil
}

//Unsubscribe 指定topic 和 subscription 订阅关系
func (ti *Tips) Unsubscribe(cxt context.Context, subName string, topic string) error {
	txn, err := ti.ps.Begin()
	if err != nil {
		return err
	}
	//查看当前topic是否存在
	t, err := txn.GetTopic(topic)
	//如果当前的topic不存在，那么返回错误
	if err != nil {
		return err
	}
	if err := txn.DeleteSubscription(t, subName); err != nil {
		return err
	}
	return nil
}

//Subscription 查询当前subscription的信息
//func (ti *Tips) Subscription(cxt context.Context, subName string) (string, error) {
//Pull 拉取消息
func (ti *Tips) Pull(cxt context.Context, subName string, topic string, index, limit int64, ack bool) ([]string, int64, error) {
	txn, err := ti.ps.Begin()
	if err != nil {
		return nil, 0, err
	}

	return nil, 0, nil
}

func (ti *Tips) CreateSnapshots(cxt context.Context, name string, subName string) (int, error) {
	return 0, nil
}
func (ti *Tips) DeleteSnapshots(cxt context.Context, name string, subName string) error {
	return nil
}
func (ti *Tips) Seek(cxt context.Context, name string) (int64, error) {
	return 0, nil
}

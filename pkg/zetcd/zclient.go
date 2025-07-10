// 利用go-zero对etcd的连接配置，实现put、get、delete等方法的调用，并支持在rpc启动时创建pub对象并发布内容
package zetcd

import (
	"errors"
	"time"

	"context"

	"github.com/zeromicro/go-zero/core/discov"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

const (
	_ = iota
	indexOfId
)

const timeToLive int64 = 10

// TimeToLive is seconds to live in etcd.
var (
	TimeToLive = timeToLive
)

// ZEtcdConf is the config item with the given key on etcd.
type ZEtcdConf struct {
	EtcdHosts          []string
	ID                 int64  `json:",optional"`
	User               string `json:",optional"`
	Pass               string `json:",optional"`
	CertFile           string `json:",optional"`
	CertKeyFile        string `json:",optional=CertFile"`
	CACertFile         string `json:",optional=CertFile"`
	InsecureSkipVerify bool   `json:",optional"`
}

type ZEtcdClient struct {
	Client     *clientv3.Client
	Publisher  *Publisher
	Subscriber *discov.Subscriber
}

type EtcdClientInterface interface {
	ActiveConnection() *grpc.ClientConn
	Close() error
	Ctx() context.Context
	Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error)
	Grant(ctx context.Context, ttl int64) (*clientv3.LeaseGrantResponse, error)
	KeepAlive(ctx context.Context, id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error)
	Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error)
	Revoke(ctx context.Context, id clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error)
	Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan
}

// NewEtcdClient 创建EtcdClient并初始化etcd连接
func NewZEtcdClient(cfg ZEtcdConf) (*ZEtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.EtcdHosts,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return nil, err
	}
	publisher := NewPublisher(cli)
	if publisher == nil {
		return nil, errors.New("failed to create publisher")
	}
	zEtcdClient := &ZEtcdClient{
		Client:     cli,
		Publisher:  publisher,
		Subscriber: nil,
	}

	return zEtcdClient, nil
}

func (e *ZEtcdClient) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return e.Client.Put(e.Client.Ctx(), key, val, opts...)
}

func (e *ZEtcdClient) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return e.Client.Get(e.Client.Ctx(), key, opts...)
}

func (e *ZEtcdClient) Publish(key, val string) error {
	if e.Publisher == nil {
		return errors.New("publisher is not initialized")
	}
	publisher := e.Publisher
	publisher.key = key
	publisher.value = val
	return publisher.KeepAlive()
}

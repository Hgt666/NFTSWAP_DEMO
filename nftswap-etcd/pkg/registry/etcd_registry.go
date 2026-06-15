package registry

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// EtcdConfig etcd 连接配置
type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
	Timeout   int      `yaml:"timeout"`
}

// Service 服务元信息
type Service struct {
	Name string // 服务名
	Addr string // 服务地址 ip:port
	Port int    `yaml:"port"`
}

// EtcdRegistry etcd 注册器
type EtcdRegistry struct {
	cli    *clientv3.Client
	config EtcdConfig
}

// NewEtcdRegistry 创建注册器
func NewEtcdRegistry(cfg EtcdConfig) (*EtcdRegistry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.Timeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &EtcdRegistry{
		cli:    cli,
		config: cfg,
	}, nil
}

// Register 服务注册 + 保活（租约）
// path 注册路径规范：/service/服务名/实例唯一标识
func (r *EtcdRegistry) Register(svc Service, ttl int64) error {
	ctx := context.Background()
	// 创建租约
	leaseResp, err := r.cli.Grant(ctx, ttl)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("/service/%s/%s", svc.Name, svc.Addr)
	// 注册 key-value 并绑定租约
	_, err = r.cli.Put(ctx, key, svc.Addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	// 开启续租（心跳保活）
	go r.keepAlive(ctx, leaseResp.ID, key)
	log.Printf("服务注册成功: %s -> %s", svc.Name, svc.Addr)
	return nil
}

// 租约续租
func (r *EtcdRegistry) keepAlive(ctx context.Context, leaseID clientv3.LeaseID, key string) {
	ch, err := r.cli.KeepAlive(ctx, leaseID)
	if err != nil {
		log.Printf("租约保活失败: %v", err)
		return
	}
	for range ch {
	}
	// 续租中断，删除失效节点
	_, _ = r.cli.Delete(ctx, key)
	log.Printf("服务节点下线: %s", key)
}

// Discover 服务发现：根据服务名获取所有实例地址
func (r *EtcdRegistry) Discover(svcName string) ([]string, error) {
	ctx := context.Background()
	prefix := fmt.Sprintf("/service/%s/", svcName)
	resp, err := r.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var addrs []string
	for _, kv := range resp.Kvs {
		addrs = append(addrs, string(kv.Value))
	}
	return addrs, nil
}

// Close 关闭客户端
func (r *EtcdRegistry) Close() error {
	return r.cli.Close()
}
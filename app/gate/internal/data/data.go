package data

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/vulcan-frame/vulcan-gate/app/gate/internal/conf"
	"github.com/vulcan-frame/vulcan-pkg-tool/data/cache"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	Rdb redis.Cmdable
}

func NewData(c *conf.Data) (d *Data, cleanup func(), err error) {
	var rdb redis.Cmdable

	if c.Redis.Cluster {
		rdb, cleanup, err = cache.NewRedisCluster(&redis.ClusterOptions{
			Addrs:        []string{c.Redis.Addr},
			Password:     c.Redis.Password,
			DialTimeout:  c.Redis.DialTimeout.AsDuration(),
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		})
	} else {
		rdb, cleanup, err = cache.NewRedis(&redis.Options{
			Addr:         c.Redis.Addr,
			Password:     c.Redis.Password,
			DialTimeout:  c.Redis.DialTimeout.AsDuration(),
			WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
			ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		})
	}
	if err != nil {
		return
	}

	d = &Data{
		Rdb: rdb,
	}
	return
}

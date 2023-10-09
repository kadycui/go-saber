package utils

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:         "172.20.166.56:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	return &RedisClient{client: client}

}

// 关闭redis连接
func (r *RedisClient) Close() error {
	err := r.client.Close()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(r.client.Context(), key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.client.Context(), key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Ping
func (r *RedisClient) Ping() (string, error) {
	pong, err := r.client.Ping(r.client.Context()).Result()
	if err != nil {
		return "", err
	}

	return pong, nil
}

// 清空当前数据库
func (r *RedisClient) FlushDB() error {
	_, err := r.client.FlushDB(r.client.Context()).Result()
	if err != nil {
		return err
	}

	return nil
}

// 清空所有数据库
func (r *RedisClient) FlushAll() error {
	_, err := r.client.FlushAll(r.client.Context()).Result()
	if err != nil {
		return err
	}

	return nil
}


// 通过Key删除
func (r *RedisClient) Del(key ...string) (int64, error) {
	delCount, err := r.client.Del(r.client.Context(), key...).Result()
	if err != nil {
		return 0, err
	}

	return delCount, nil
}


func (r *RedisClient) LPush(key string, values ...interface{}) (int64, error) {
	lpushCount, err := r.client.LPush(r.client.Context(), key, values...).Result()
	if err != nil {
		return 0, err
	}

	return lpushCount, nil
}


// 在列表中的 pivot 之前插入值。
func (r *RedisClient) LInsertBefore(key, pivot, value string) (int64, error) {
	val, err := r.client.LInsert(r.client.Context(), key, "BEFORE", pivot, value).Result()
	if err != nil {
		return 0, err
	}

	return val, nil
}

// LInsertAfter在列表中的 pivot 之后插入值。
func (r *RedisClient) LInsertAfter(key, pivot, value string) (int64, error) {
	val, err := r.client.LInsert(r.client.Context(), key, "AFTER", pivot, value).Result()
	if err != nil {
		return 0, err
	}

	return val, nil
}

// LSet使用新值设置索引处的元素。
func (r *RedisClient) LSet(key string, index int64, value interface{}) (string, error) {
	val, err := r.client.LSet(r.client.Context(), key, index, value).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// LLen返回存储在key的列表的长度。
func (r *RedisClient) LLen(key string) (int64, error) {
	val, err := r.client.LLen(r.client.Context(), key).Result()
	if err != nil {
		return 0, err
	}

	return val, nil
}

// LRange从存储在key的列表中返回指定范围的元素。
func (r *RedisClient) LRange(key string, start, stop int64) ([]string, error) {
	val, err := r.client.LRange(r.client.Context(), key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

// LIndex返回存储在key的列表中索引处的元素。
func (r *RedisClient) LIndex(key string, index int64) (string, error) {
	val, err := r.client.LIndex(r.client.Context(), key, index).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// LRem从存储在key的列表中删除第一次出现等于value的元素的计数。
func (r *RedisClient) LRem(key string, count int64, value string) (int64, error) {
	val, err := r.client.LRem(r.client.Context(), key, count, value).Result()
	if err != nil {
		return 0, err
	}

	return val, nil
}


// SAdd adds the specified members to the set stored at key.
func (r *RedisClient) SAdd(key string, members ...interface{}) (int64, error) {
	addCount, err := r.client.SAdd(r.client.Context(), key, members...).Result()
	if err != nil {
		return 0, err
	}

	return addCount, nil
}

// SPop 从存储在指定键的集合中删除并返回一个或多个随机元素。
func (r *RedisClient) SPop(key string, count ...int64) ([]string, error) {
	var poppedValues []string
	var err error

	if len(count) > 0 {
		poppedValues, err = r.client.SPopN(r.client.Context(), key, count[0]).Result()
	} else {
		poppedValue, err := r.client.SPop(r.client.Context(), key).Result()
		if err == nil {
			poppedValues = append(poppedValues, poppedValue)
		}
	}

	if err != nil {
		return nil, err
	}

	return poppedValues, nil
}

// SRem方法从存储在key中的集合中移除指定的成员。
func (r *RedisClient) SRem(key string, members ...interface{}) (int64, error) {
	remCount, err := r.client.SRem(r.client.Context(), key, members...).Result()
	if err != nil {
		return 0, err
	}

	return remCount, nil
}

// SMembers方法返回存储在key中的集合的所有成员。
func (r *RedisClient) SMembers(key string) ([]string, error) {
	members, err := r.client.SMembers(r.client.Context(), key).Result()
	if err != nil {
		return nil, err
	}

	return members, nil
}

// SIsMember方法返回成员是否是存储在key中的集合的成员。
func (r *RedisClient) SIsMember(key string, member interface{}) (bool, error) {
	isMember, err := r.client.SIsMember(r.client.Context(), key, member).Result()
	if err != nil {
		return false, err
	}

	return isMember, nil
}

// SCard方法返回存储在key中的集合的基数（元素数量）。
func (r *RedisClient) SCard(key string) (int64, error) {
	cardinality, err := r.client.SCard(r.client.Context(), key).Result()
	if err != nil {
		return 0, err
	}

	return cardinality, nil
}


// ZAdd方法将一个或多个成员以及它们的分值添加到存储在key中的有序集合中。
func (r *RedisClient) ZAdd(key string, members ...*redis.Z) (int64, error) {
    // 调用Redis客户端的ZAdd方法将成员添加到有序集合中
    zAddCmd := r.client.ZAdd(r.client.Context(), key, members...)
    return zAddCmd.Result()   // 返回ZAdd命令的执行结果
}

// ZIncrBy方法按increment参数增加存储在key中的有序集合中成员member的分数。
func (r *RedisClient) ZIncrBy(key string, increment float64, member string) (float64, error) {
	zIncrByCmd := r.client.ZIncrBy(r.client.Context(), key, increment, member)
	return zIncrByCmd.Result()
}

// ZRange方法返回存储在key中的有序集合中指定范围的元素。
func (r *RedisClient) ZRange(key string, start, stop int64) ([]string, error) {
	zRangeCmd := r.client.ZRange(r.client.Context(), key, start, stop)
	return zRangeCmd.Result()
}

// ZRevRange方法以倒序（从最高分到最低分）返回存储在key中的有序集合中指定范围的元素。
func (r *RedisClient) ZRevRange(key string, start, stop int64) ([]string, error) {
	zRevRangeCmd := r.client.ZRevRange(r.client.Context(), key, start, stop)
	return zRevRangeCmd.Result()
}

// ZRangeByScore方法返回在key中分数介于min和max（包含）之间的所有元素。
func (r *RedisClient) ZRangeByScore(key string, min, max string) ([]string, error) {
	zRangeByScoreCmd := r.client.ZRangeByScore(r.client.Context(), key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	})
	return zRangeByScoreCmd.Result()
}

// ZRevRangeByScore方法以倒序（从最高分到最低分）返回在key中分数介于max和min（包含）之间的所有元素。
func (r *RedisClient) ZRevRangeByScore(key string, max, min string) ([]string, error) {
	zRevRangeByScoreCmd := r.client.ZRevRangeByScore(r.client.Context(), key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	})
	return zRevRangeByScoreCmd.Result()
}

// ZCard方法返回存储在key中的有序集合的基数（元素数量）。
func (r *RedisClient) ZCard(key string) (int64, error) {
	zCardCmd := r.client.ZCard(r.client.Context(), key)
	return zCardCmd.Result()
}

// ZCount方法返回在key中分数介于min和max（包含）之间的有序集合的元素数量。
func (r *RedisClient) ZCount(key, min, max string) (int64, error) {
	zCountCmd := r.client.ZCount(r.client.Context(), key, min, max)
	return zCountCmd.Result()
}

// ZScore方法返回在key中有序集合中成员的分数。
func (r *RedisClient) ZScore(key, member string) (float64, error) {
	zScoreCmd := r.client.ZScore(r.client.Context(), key, member)
	return zScoreCmd.Result()
}

// ZRank方法返回存储在key中的有序集合中成员的排名（从低到高排序），即成员在有序集合中的位置。
func (r *RedisClient) ZRank(key, member string) (int64, error) {
	zRankCmd := r.client.ZRank(r.client.Context(), key, member)
	return zRankCmd.Result()
}

// ZRevRank方法返回存储在key中的有序集合中成员的排名（从高到低排序），即成员在有序集合中的位置。
func (r *RedisClient) ZRevRank(key, member string) (int64, error) {
	zRevRankCmd := r.client.ZRevRank(r.client.Context(), key, member)
	return zRevRankCmd.Result()
}

// ZRem方法从存储在key中的有序集合中移除一个或多个成员。
func (r *RedisClient) ZRem(key string, members ...interface{}) (int64, error) {
	zRemCmd := r.client.ZRem(r.client.Context(), key, members...)
	return zRemCmd.Result()
}

// ZRemRangeByRank方法从存储在key中的有序集合中移除排名在start和stop之间（包括start和stop）的所有元素。
func (r *RedisClient) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	zRemRangeByRankCmd := r.client.ZRemRangeByRank(r.client.Context(), key, start, stop)
	return zRemRangeByRankCmd.Result()
}

// ZRemRangeByScore方法从存储在key中的有序集合中移除分数在min和max之间（包括min和max）的所有元素。
func (r *RedisClient) ZRemRangeByScore(key, min, max string) (int64, error) {
	zRemRangeByScoreCmd := r.client.ZRemRangeByScore(r.client.Context(), key, min, max)
	return zRemRangeByScoreCmd.Result()
}



func (r *RedisClient) HSet(key, field string, value interface{}) error {
	return r.client.HSet(r.client.Context(), key, field, value).Err()
}

func (r *RedisClient) HMSet(key string, values map[string]interface{}) error {
	return r.client.HMSet(r.client.Context(), key, values).Err()
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	return r.client.HGet(r.client.Context(), key, field).Result()
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	return r.client.HGetAll(r.client.Context(), key).Result()
}

func (r *RedisClient) HDel(key string, fields ...string) error {
	return r.client.HDel(r.client.Context(), key, fields...).Err()
}

func (r *RedisClient) HExists(key, field string) (bool, error) {
	return r.client.HExists(r.client.Context(), key, field).Result()
}

func (r *RedisClient) HLen(key string) (int64, error) {
	return r.client.HLen(r.client.Context(), key).Result()
}
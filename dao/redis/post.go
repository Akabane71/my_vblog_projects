package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func GetPostListInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFormKey(key, p.Page, p.Size)

}

func getIDsFormKey(key string, page, pageSize int64) ([]string, error) {
	start := (page - 1) * pageSize
	end := start + pageSize - 1
	return rdb.ZRevRange(key, start, end).Result()
}

// 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteDate(ids []string) (data []int64, err error) {

	// 访问redis的操作过于频繁了，当dis过大的时候
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPre + id)
	//	// 查找key中分数是1的元素的数量---> 统计每篇帖子赞成票的数量
	//	V1 := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, V1)
	//}

	// 使用pipline一次发送多条命令，减少RTT
	pipline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPre + id)
		pipline.ZCount(key, "1", "1")
	}
	cmders, err := pipline.Exec()
	if err != nil {
		return nil, err
	}

	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按照社区id查询
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// 使用 zinterstore 把分区的帖子set和帖子分数的zset生成一个新的zset，
	//  针对新的zset按之前的逻辑取数据

	// 社区的key
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存key减少zinterstore执行的次数
	key := p.Order + strconv.Itoa(int(p.CommunityID))

	if rdb.Exists(p.Order).Val() < 1 {
		// 不存在,需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, p.Order) // zinterstore 计算
		pipeline.Expire(key, 60*time.Second) // 超时
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	// 存在的话直接根据key查询ids
	return getIDsFormKey(key, p.Page, p.Size)
}

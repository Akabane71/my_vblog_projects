package redis

import (
	"bluebell/models"
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 60 * 60 // 一个星期的秒数
	scorePerVote     = 432              // 每一票多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票的限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 2和3 需要放到一个pipeline操作中
	// 2. 更新分数
	// 先查询当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPre+postID), postID).Val()

	// 投过然后还投,就直接返回一个报错
	if value == ov {
		return ErrVoteRepeated
	}

	var op float64
	// 计算方向
	if value > ov {
		op = 1
	} else {
		op = 0
	}

	diff := math.Abs(ov - value)                                                               // 计算两次投票的差值
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostTimeZSet), op*diff*scorePerVote, postID).Result() // 原子性操作的增加
	if ErrVoteTimeExpire != nil {
		return err
	}

	pipeline := rdb.Pipeline()

	// 3. 记录用户为该帖子投票的数据
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPre+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPre+postID), redis.Z{
			Score:  value,  // 赞成票 还是 反对票
			Member: userID, // 哪个用户
		})
	}
	_, err = pipeline.Exec()
	return err
}

func CreatePost(postID, communityID int64) error {
	// 使用原子性操作
	pipeline := rdb.TxPipeline()

	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)

	// 执行
	_, err := pipeline.Exec()

	return err
}

// 按照
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1. 根据用户请求中携带的order参数确定要查询的redis的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else {
		key = getRedisKey(KeyPostTimeZSet)
	}

	// 2. 确认查询索引起终点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	// 3. ZREVRANGE 查询
	return rdb.ZRevRange(key, start, end).Result()

}

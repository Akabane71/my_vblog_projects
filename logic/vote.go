package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 基于用户投票的相关算法,推荐阅读
// https://ruanyifeng.com/blog/algorithm/

// 投票功能:

// 本项目使用简化版本的投票分数
// 投一票就加432分   86400/200 ---> 需要200张赞成才可以给你的帖子续一天 --->  例子来自于<redis实战>

/* 投票的几种情况:

direction = 1 时,有两种情况:
	1. 之前没有投过票,现在投了赞成票 --> 更新分数和投票记录  差值的绝对值:1 	+432
	2. 之前投反对票,现在改投赞成票 --> 更新分数和投票记录  差值的绝对值:2 + 	+432*2

direction = 0 时,也有两种情况:
	1. 之前投过票,现在要取消投票 --> 更新分数和投票记录  差值的绝对值:1 	    -432
	2. 之前投反对票，现在要取消投票 --> 更新分数和投票记录  差值的绝对值:1 		+432

direction = -1 时,有两种情况:
	1. 之前没有投过票，现在投反对票 --> 更新分数和投票记录   差值的绝对值:1 	-432
	2. 之前投赞成票，现在改投反对票 --> 更新分数和投票记录   差值的绝对值:2 	-432*2

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期不允许在投票了.
	1. 到期之后将redis中保存的额赞成票和反对票数存储到mysql中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamsVoteData) error {
	zap.L().Debug(
		"VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(
		strconv.Itoa(int(userID)),
		p.PostID,
		float64(p.Direction))
}

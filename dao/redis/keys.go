package redis

// redis key

// redis key 尽量使用命名空间的方式,方便查询和拆分

const (
	KeyPrefix           = "bluebell"            // 项目的Key前缀
	KeyPostTimeZSet     = "bluebell:post:time"  // ZSet; 帖子以发帖时间为分数
	KeyPostScoreZSet    = "bluebell:post:score" // ZSet; 帖子及投票分数
	KeyPostVotedZSetPre = "post:voted:"         // ZSet; 记录用户及投票的类型;参数是post的id
	KeyCommunitySetPF   = "community:"          // Set;保存每个分区下的帖子的id
)

// 给Redis的key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

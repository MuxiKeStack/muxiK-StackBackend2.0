package redis_store

type Store interface {
	// 保存验证码
	Set(key string, value string, expire int) error

	// 获取验证码
	Get(key string, clear bool) string

	// 检查验证码
	Verify(id, answer string, clear bool) bool
}

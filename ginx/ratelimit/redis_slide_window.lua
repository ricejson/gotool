-- 限流对象
local key = KEYS[1]
-- 窗口值
local window = tonumber(ARGV[1])
-- 限流阈值
local threshold = tonumber(ARGV[2])
-- 当前时间
local now = tonumber(ARGV[3])
-- 移除窗口外的元素
redis.call('ZREMRANGEBYSCORE', key, '-INF', now - window)
-- 计算当前窗口的元素个数
local cnt = redis.call('ZCOUNT', key, now - window, now)
if cnt >= threshold then
    -- 执行限流
    return "true"
else
    -- 设置当前请求
    redis.call('ZADD', key, now, now)
    -- 防止key永不过期
    redis.call('PEXPIRE', key, window)
    return "false"
end

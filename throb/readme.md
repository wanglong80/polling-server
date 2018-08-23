# 建议通过设定 crontab 每天定时清除缓存数据
`redis-cli keys 'imx_*' | xargs redis-cli del`
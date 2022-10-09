# golang short me

golang 实现短地址服务

```bash
# ronething @ ashings-macbook-pro in ~/Documents/app/redis/local/redis-4.0.0 [1:51:46]
$ curl -X POST http://127.0.0.1:8888/api/shorten \
-H 'Content-Type: application/json' \
-d '{
"url": "https://blog.ronething.cn",
"expiration_in_minutes": 1
}'
{"short_link":"2"}%
# ronething @ ashings-macbook-pro in ~/Documents/app/redis/local/redis-4.0.0 [1:55:12]
$ curl -X GET http://127.0.0.1:8888/api/info\?shortLink\=2
{"create_at":"2020-03-27 01:55:12.589573 +0800 CST m=+501.744003023","expiration_in_minutes":1,"url":"https://blog.ronething.cn"}%
# ronething @ ashings-macbook-pro in ~/Documents/app/redis/local/redis-4.0.0 [1:55:21]
$ curl -X GET http://127.0.0.1:8888/2
<a href="https://blog.ronething.cn">Temporary Redirect</a>.

```


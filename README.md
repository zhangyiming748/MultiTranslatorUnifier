# MultiTranslatorUnifier

聚合翻译服务

# 使用的引擎

- [x] Bing
- [x] Google
- [x] Deeplx
- [x] LinuxDo
- [ ] 百度翻译开放平台

# ToDo

- [ ] 引入redis 单次运行时作为缓存
- [ ] 引入mysql 单次运行完成持久化

# 逻辑


if 提供proxy 查询 translateshell google引擎
if 



# 接口文档

### 心跳测试

```shell
curl --location --request GET 'http://127.0.0.1:8192/api/v1/translate?user=zen'
```

### 执行翻译

```shell
curl --location --request POST 'http://127.0.0.1:8192/api/v1/translate' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "src":"",
	    "proxy":"",
	    "linuxdokey":""
	}'
```
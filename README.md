


### 使用说明

此工具用来上传文件到对象存储，目前支持阿里云、aws、金山云、华为云。

**说明：** 主要用于文件上传至对象存储
- 支持配置文件及命令行（命令行会覆盖同名配置块）
- 支持子目录
- 支持文件及目录
- 支持多类型对象存储

```bash
#local env: go 1.17.1
```

### usage demo 

```bash
./upload2remote -t OSS -s ./test_dir_or_filename -d oss://gourds_111/ ----AccessKeyID xxx --AccessKeySecret xxx --Region xxx --Endpoint ks3-cn-beijing.ksyun.co
```

### Todo

- [x] 基本集成
- [ ] 一次认证多次上传
- [ ] 并发上传支持
- [ ] 添加统计信息（耗时、计数）
- [ ] 内容未改变跳过上传


### support & Document

- [x] Aliyun: https://help.aliyun.com/document_detail/88426.html
- [ ] 微软云
- [x] 华为云：https://sdkcenter.developer.huaweicloud.com/?product=obs
- [x] aws: https://docs.aws.amazon.com/sdk-for-go/api/service/s3/
- [x] 金山云 ks3 https://console.huaweicloud.com/console/?region=cn-east-3#/obs/manage/jws2-ops-backup-data/object/list




```

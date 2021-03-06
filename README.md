


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
./upload2remote -t OSS -s ./test_dir_or_filename -d oss://gourds_111/ --AccessKeyID xxx --AccessKeySecret xxx --Region xxx --Endpoint ks3-cn-beijing.ksyun.co
```

### Todo

- [x] 基本集成
- [x] 一次认证多次上传
- [x] 统计耗时
- [x] 统计数量
- [x] 并发上传支持


### support & Document

- [x] [阿里云 oss](https://help.aliyun.com/document_detail/88426.html)
- [x] [华为云 obs](https://sdkcenter.developer.huaweicloud.com/?product=obs)
- [x] [亚马逊 s3](https://docs.aws.amazon.com/sdk-for-go/api/service/s3/)
- [x] [金山云 ks3](https://console.huaweicloud.com/console/?region=cn-east-3#/obs/manage/jws2-ops-backup-data/object/list)
- [x] [微软云 blob](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/storage#Blob.CreateBlockBlobFromReader)



```

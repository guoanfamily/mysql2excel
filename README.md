# mysql2excel
mysql生成excel数据库文档

生成表目录，表明细内容，可根据实际需求自行修改调整

每张表为一个sheet页

表目录包含超链接，可快速跳转至对应表明细内容
## 使用说明
1. 安装程序，运行：
```shell script
go get github.com/guoanfamily/mysql2excel
```
2. 创建config.yaml
填写数据库连接串，生成文档名称，数据库名
```yaml
connectstr : "user:password@tcp(localhost:3306)"
file : "dbdoc.xlsx"
dbname : "dbname"
```
3. 在config.yaml同目录中执行命令
```shell script
mysql2execl
```
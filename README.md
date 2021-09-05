# 运行注意事项
- 在goland下开发由于相对文件路径的问题，建议先将run/run.go编译成可执行文件进行执行
- 启动之前为了方便后期的docker打包，需要在/etc/hosts文件中添加
  127.0.0.1 unioj-beego  这样一条解析，对应于，/server/conf.ini中的host字段，为了方便本地开发也可以字节将host字段写为ip:port的形式
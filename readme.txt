配置文件服务项目

当启动alphabetweb服务时，设置参数 ` -config_url "http://xxxx/yyy" ` ，那么会访问sysconfig_server中的远程配置文件信息，动态装载配置信息。
启动命令：
go run alphabetweb/src/alphabet/cmd/abserver.go -start /Users/vikingbays/golang/AlphabetwebProject/sysconfig_service

1、配置文件存储结构

存储的目录结构如下：
    store
      - [projectid]+
         - [aliaskey]+
            - [appname]+

例如：
    store                                   ## 存储根目录
      - octopus                             ## 定义的一个projectid，与project对应
         - base                             ## 一组 sysconfig别名
            - sample_octopus_frontend       ## 针对 project中的appname对应
            - sample_octopus_service        ## 针对 project中的appname对应


2、配置文件web界面查看方式：
  http://localhost:10060/sysconfigservice/sysconfig/filelist


3、配置文件服务的访问方式：
  http://localhost:10060/sysconfigservice/sysconfig/downzip/octopus/base_remote/

  octopus  是 定义的一个projectid，与project对应
  base_remote  一组 sysconfig别名

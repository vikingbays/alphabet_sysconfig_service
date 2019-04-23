# Vikingbays • AlphabetWeb : alphabet_sysconfig_service

alphabet_sysconfig_service 作为AlphabetWeb的统一配置中心。

统一配置中心解决了本地化配置管理有其局限性，例如在微服务体系下，可能因为程序频繁更新，会使配置文件不一致概率增大，存在冲突。因此需要有一个统一配置中心，保障配置信息高度一致性。

当前AlphabetWeb项目提供了一个简易的统一配置中心，对应的项目是：alphabet_sysconfig_service 。

## 1. 统一配置中心：alphabet_sysconfig_service

### 1.1. 用途说明
Sysconfig 项目配置项文件可以存储到 统一配置中心（sysconfig_service）中，使 AlphabetWeb项目 和 Sysconfig 项目配置项文件 分离。 AlphabetWeb项目 在启动时候，可以连接统一配置中心的服务，下载Sysconfig信息。

### 1.2. 安装和启动

下载[sysconfig_service]

启动：
```
AlphabetwebProject="/Users/vikingbays/golang/AlphabetwebProject/alphabetweb"

sysconfig_service="/Users/vikingbays/golang/AlphabetwebProject/alphabet_sysconfig_service"

go run $AlphabetwebProject/src/alphabet/cmd/abserver.go -start "$sysconfig_service"
```

访问，可以查看当前Sysconfig配置项：
```
http://localhost:10060/sysconfigservice/sysconfig/filelist
```

sysconfig配置项下载地址：
```
http://localhost:10060/sysconfigservice/sysconfig/downzip/{projectid}/{aliaskey}/{appname}
```
其中：{projectid} 是项目别名，{aliaskey}是关键字，{appname}是项目的应用包名。


### 1.3. 配置Sysconfig规则
统一配置中心（sysconfig_service）的Sysconfig配置文件默认存储在： sysconfig_service/store 。

假设当前AlphabetWeb项目 sample_octopus_frontend_service , 他有两个应用包 sample_octopus_frontend 和 sample_octopus_service ，具体目录结构是：
```
  sample_octopus_frontend_service
      ...
      src
          sample_octopus_frontend        # 应用包
              ...
          sample_octopus_service         # 应用包
              ...
```

那么在统一配置中心（sysconfig_service）配置此项目的Sysconfig
```
  sysconfig_service
      store
          ...
          octopus                               # 项目别名，可以和sample_octopus_frontend_service不一致
              ...
              dev_remote                        # key
                  sample_octopus_frontend       # 与 sample_octopus_frontend_service的应用包名称一致
                      dbsconfig.toml
                      logconfig_http.toml
                      logconfig.toml
                      ms_client_config_01.toml
                      webconfig.toml
                  sample_octopus_service        # 与 sample_octopus_frontend_service的应用包名称一致
                      dbsconfig.toml
                      logconfig_http.toml
                      logconfig.toml
                      ms_server_config.toml
                      webconfig.toml                    

```

sysconfig_service/store/octopus 是项目别名，与AlphabetWeb项目 sample_octopus_frontend_service 对应。

sysconfig_service/store/octopus/dev_remote 是关键字，可以定义多个，例如:test_remote，release_remote 。

sysconfig_service/store/octopus/dev_remote/sample_octopus_frontend 对应了 sample_octopus_frontend_service/src/sample_octopus_frontend ，名称必须一致 。

### 1.4. 服务访问说明

##### 1.4.1. 查看Sysconfig
可以通过界面查看有哪些服务存在。
访问地址是：
```
http://localhost:10060/sysconfigservice/sysconfig/filelist
```
![sysconfig_filelist](../../frame/resource/mkimages/develop_001.png "" "" "vikings-img-Class1")

##### 1.4.1. 获取Sysconfig

sysconfig配置项下载服务，访问地址是一个Restful风格：
```
http://localhost:10060/sysconfigservice/sysconfig/downzip/{projectid}/{aliaskey}/{appname}
```
其中：{projectid} 是项目别名，{aliaskey}是关键字，{appname}是项目的应用包名。

假设1.3中的sample_octopus_frontend_service的配置项来看，他们的对应关系是：

地址中的关键字{{class="table table-bordered vikings-table-Simple-max200"}}  |  实例值  |  sysconfig_service的路径   |  sample_octopus_frontend_service对应信息
------|------|-------|------
{projectid} | octopus |  store/octopus | sample_octopus_frontend_service
{aliaskey} | dev_remote |  store/octopus/dev_remote | src/../sysconfig.dev_remote
{appname} | sample_octopus_frontend |  store/octopus/dev_remote/sample_octopus_frontend  | src/sample_octopus_frontend

## 2. 远程获取配置

AlphabetWeb项目调用统一配置中心（sysconfig_service）的Sysconfig配置项，只需要在启动时加上 -config_url 参数。

基于上面例子的启动命令是：
```
go run /Users/vikingbays/golang/AlphabetwebProject/alphabetweb/src/alphabet/cmd/abserver.go
       -start "/Users/vikingbays/golang/AlphabetwebProject/sample_octopus_frontend_service"  
       -config_url "http://localhost:10060/sysconfigservice/sysconfig/downzip/octopus/dev_remote/"

```

说明： -config_url对应的参数值，只需要指向：`http://localhost:10060/sysconfigservice/sysconfig/downzip/{projectid}/{aliaskey}` 。启动时，可以根据应用包拼装获取其sysconfig配置项。

启动完成后，下载的sysconfig配置项，存储在应用包下，sysconfig配置项的目录是 `sysconfig.`{aliaskey}  的组合。 例如：
```
sysconfig.dev_remote
```

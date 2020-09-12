# gin-demo
用gin框架搭建的一个项目结构，方便快速开发项目。

### 特点

- 集成gorm，用户mysql存储层操作

- 集成go redis，用户操作缓存

- 集成uber/zap, zap是一个高效的日志组件

  项目中日志进行了分解:

  - request日志  

    记录http请求的request和response结果

  - init日志

    记录服务启动的日志

  - mysql日志 

    记录启动项目时，迁移数据库执行的sql情况日志

  - err日志

    记录http请求产生的错误日志，方便快速定位错误问题

  - app日志

    记录开发过程中，我们记录的业务日志，这是我们最常用的日志，日志当中记录了traceId，方便快速根据请求查看当前请求的日志流,日志格式如下,包括了traceId、file、keywords，以及我们记录的重要信息，提现在data里面:

    ```she
    {"level":"info","time":"2020-09-12 10:56:07.432","keywords":"create User","traceId":"1hops10syiy0tycvpgnpkab9x31","file":"controller/home.go:37","data":{"ID":5,"CreatedAt":"2020-09-12T10:56:07.4261674+08:00","UpdatedAt":"2020-09-12T10:56:07.4261674+08:00","DeletedAt":null,"Username":"范兄弟","Password":"3333","Status":1,"CreateAt":"2020-09-12T10:56:07.4251674+08:00"}}
    {"level":"warn","time":"2020-09-12 10:56:07.432","keywords":"redis","traceId":"1hops10syiy0tycvpgnpkab9x31","file":"controller/home.go:38","data":{"message":"specified duration is 5ns, but minimal supported value is 1ms - truncating to 1ms"}}
    ```

- 集成jwt，一种流行的web身份认证方式，减轻服务端压力，将用户登录验证信息存储在可以端

- 集成gopkg.in/ini，用户解析我们的配置项

### 目录结构

- config -配置
- constant -常量
- controller -业务控制器
- logs -日志目录
- middleware -中间件目录
- model -表模型目录
- router -路由配置目录
- service -服务存放目录，包括db、log、redis
- util -其它工具目录
- main.go -程序入库文件

### 其它

- 日志按天进行分割
- 根目录的.env.local进行本地配置,.env作为线上配置，避免误将本地配置提交到git仓库
- 业务日志可以根据traceId查看当次请求的所有日志，方便定位问题，分析逻辑行为
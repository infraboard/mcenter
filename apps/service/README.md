# 服务管理

服务管理服务, 管理着服务完整的生命周期

+ 服务创建
    + 代码仓库创建, 代码模版
+ 服务上线
    + 构建
        + 构建资源申请
        + 服务构建(构建变量管理, 默认变量和用户自定义变量)
        + 制品保存
    + 发布
        + 部署资源申请
        + 上线申请
        + 部署配置(敏感信息管理)
        + 服务部署(部署变量管理, 默认变量和用户自定义变量)
        + 服务验证
        + 部署完成
        + 发版记录
+ 服务下线
    + 代码归档, 资源回收释放


## 服务事件


+ SCM类型事件:
    + [Gitlab Web Hooks](https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#webhooks)
    + [Github Web Hooks]()
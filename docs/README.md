# 使用文档

创建数据库用户:
```sh
# docker exec -it mongo mongo
> use mcenter
switched to db mcenter
> db.createUser({user: "mcenter", pwd: "123456", roles: [{ role: "dbOwner", db: "mcenter" }]})
```
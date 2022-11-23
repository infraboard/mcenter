# LDAP 认证

## 环境准备

采用Docker安装LDAP开发环境:
```
# 使用docker安装LDAP
docker run -p 389:389 -p 636:636 --name my-openldap-container --detach osixia/openldap:1.5.0
# 添加用户
docker exec my-openldap-container ldapsearch -x -H ldap://localhost -b dc=example,dc=org -D "cn=admin,dc=example,dc=org" -w admin
```

更详细安装请参考: [docker-openldap](https://github.com/osixia/docker-openldap)
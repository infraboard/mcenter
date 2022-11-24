# LDAP 认证

## 环境准备

执行下面的脚本安装LDAP和Web管理
```sh
#!/bin/bash -e
docker run -p 389:389 -p 636:636 --name ldap-service --hostname ldap-service --detach osixia/openldap:1.5.0
docker run -p 6443:443 --name phpldapadmin-service --hostname phpldapadmin-service --link ldap-service:ldap-host --env PHPLDAPADMIN_LDAP_HOSTS=ldap-host --detach osixia/phpldapadmin:0.9.0

echo "Go to: https://localhost:6443/"
echo "Login DN: cn=admin,dc=example,dc=org"
echo "Password: admin"
```


执行下面命令测试LDAP搜索功能
```
docker exec ldap-service ldapsearch -x -H ldap://localhost -b dc=example,dc=org -D "cn=admin,dc=example,dc=org" -w admin
```

更详细安装请参考: [docker-openldap](https://github.com/osixia/docker-openldap)


## 创建用户

1. 首先创建一个Grou: cn=dev,dc=example,dc=org

|  Attribute   | New Value  |
|  ----  | ----  |
| Group  | test |
| GID Number  | 500 |
| objectClass  | posixGroup |

2. 在dev组创建一个用户: cn=old fish,cn=dev,dc=example,dc=org

|  Attribute   | New Value  |
|  ----  | ----  |
| Given Name  | old |
| Last name  | fish |
| Common Name  | old fish |
| User ID  | oldfish |
| Email  | oldfish@devcloud.io |
| Password | Password |
| objectClass | inetOrgPerson|
		

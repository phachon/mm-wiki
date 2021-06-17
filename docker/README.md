# [mm-wiki](https://github.com/chaiyd/docker/tree/master/mm-wiki)

## docker
- https://github.com/chaiyd/docker/tree/master/mm-wiki
- 若挂载配置文件，则下列变量不生效
  ```
  HTTP_PORT=8081
  DB_HOST=mysql
  DB_PORT=3306
  DB_NAME=mm_wiki
  DB_USER=mm_wiki
  DB_PASS=ukC2ZkcG_ZTeb
  COOKIE=2160000
  ```
- 数据库准备
  - 导入docs/databases/data.sql和docs/databases/table.sql（注：需取消注释data.sql中第一条管理用户插入语句）
  ```
  -- 系统用户（root）, password：123456, 自动安装不需要插入，手动安装时需要插入该数据
  -- --------------------------------------
  INSERT INTO `mw_user` (`user_id`, `username`, `password`, `given_name`, `email`,  `mobile`, `role_id`, `is_delete`, `create_time`, `update_time`)
  VALUES ('1', 'root', 'e10adc3949ba59abbe56e057f20f883e', 'root', 'root@123456.com', '1102222', '1', '0', unix_timestamp(now()), unix_timestamp(now()));
  ```  
    
- 目录结构
  | /app/ | mm-wiki/ | mm-wiki |
  | ----- | -------- | ------- |
  |       |          | conf/   |
  |       |          | data/   |
  

## mm-wiki官方文档
- https://github.com/phachon/mm-wiki.git


## openldap
```
{
    "basedn": "cn=dev,dc=mmwiki,dc=com",
    "bind_username": "cn=readonly,dc=mmwiki,dc=com",
    "bind_password": "mmwiki", 
	"account_pattern": "(&(objectClass=inetOrgPerson)(uid=%s))",// ldap search pattern; 非必填可以为空，默认值为(&(objectClass=User)(userPrincipalName=%s))
	"given_name_key": "cn",  // ldap 查询用户名对应的 key，非必填可以为空，默认为 cn
    "email_key": "mail",
    "mobile_key": "mobile", // ldap 查询手机号对应的 key，非必填可以为空
    "phone_key": "telephoneNumber", // ldap 查询电话对应的 key，非必填可以为空
    "department_key": "department", // ldap 查询部门对应的 key，非必填可以为空
    "position_key": "Position", // ldap 查询职位对应的 key，非必填可以为空
    "location_key": "physicalDeliveryOfficeName", // ldap 查询位置对应的 key，非必填可以为空
    "im_key": "im" // ldap 查询 im 信息对应的 key，非必填可以为空
}

```
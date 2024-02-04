
# 拉取 mysql 5.7
docker pull mysql:5.7

# 启动容器
docker run -itd --name mysql-kitter \
  -p 3306:3636 \
  -e MYSQL_ROOT_PASSWORD=kitter_pass \
  -d mysql:5.7

# 创建挂载目录
mkdir -p /usr/local/docker_v/etc

# 复制文件到宿主机
docker cp mysql-kitter:/etc/mysql /usr/local/docker_v/etc/mysql/

# 停止容器
docker stop mysql-kitter
# 删除容器
docker rm mysql-kitter

# 启动容器带挂载目录
docker run -itd --name mysql-kitter \
  -p 3636:3306 \
  -e MYSQL_ROOT_PASSWORD=kitter_pass \
  -v /usr/local/docker_v/etc/mysql:/etc/mysql \
  -d mysql:5.7

docker exec -it mysql-kitter bash

mysql -u root -pkitter_pass
grant all privileges on kms.* to 'root'@'%' identified by 'kitter_pass';
flush privileges;
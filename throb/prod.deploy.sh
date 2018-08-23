#!/bin/sh
######################################################################
## 本地自动化发布配置，下面你需要修改相关参数
######################################################################

# 服务器地址
host=139.196.162.43

# 服务器用户
user=chuxin

# 项目目录
project=./dist

# 部署到指定的目录
path=/data0/chuxin/jxdpm-polling/production

# 保留历史版本的数量
keep_releases=5

# 是否启用压缩传输（TODO: 本地tar压缩后scp，服务器tar解压）
trans_zip=false

######################################################################
## 在这里添加你的自定义命令，可以用来定制构建化工作
######################################################################

rm -rf $project/*
mkdir $project/conf
GOOS=linux GOARCH=amd64 go build
mv throb ./$project/
cp conf/app.prod.conf ./$project/conf/app.conf

######################################################################
## 以下是自动化发布脚本，尽量不要碰我 ~
######################################################################

# SSH 连接地址
server="${user}@${host}"

# 当前时间戳
currentTime=`date "+%Y%m%d%H%M%S"`

# 发布远程的绝对路径
deploy_to="${server}:${path}/releases/$currentTime"

# 远程创建发布目录
ssh $server "mkdir -p ${path}/releases/$currentTime"
scp -r $project/* $deploy_to

# 远程创建软链接到发布目录
# 往前数第 keep_releases 个版本的 release 号
# 获取当前 releases 目录下的所有版本
# 循环删除最早保留版本之前的所有目录
remote_drop_old_releases="
rm -rf ${path}/current && ln -s ${path}/releases/$currentTime ${path}/current
&& oldest_release=\`ls -t ${path}/releases | head -${keep_releases} | tail -1\`
&& release_dirs=\`ls ${path}/releases\`
&& for release_dir in \$release_dirs;
    do if [ \$release_dir -lt \$oldest_release ];then
        rm -rf ${path}/releases/\$release_dir;
    fi;
done
"
# 删除旧的 release
ssh $server $remote_drop_old_releases

# 杀掉并重启进程
ssh $server "sudo /data0/scripts/jxdpm-polling-kill.sh"

#ssh $server "cd /data0/chuxin/jxdpm-polling/production/current && nohup ./throb --tries=1 --sleep=0 &"



# ps aux | grep bigscreen | grep -v grep | cut -c 9-15 | xargs kill -9
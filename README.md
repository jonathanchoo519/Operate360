# Operate360
1.通过注册表判断是否安装360

2.读取本地病毒库日期

3.爬取官方病毒库日期与本地对比，如果发布新版本根据机器位数爬取新的病毒库增量安装包并安装

# 存在问题
由于360杀毒5.0版本的问题，导致必须在第一次安装后手动在客户端执行一次更新病毒库才能获取到病毒库日期信息
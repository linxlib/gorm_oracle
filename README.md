# GORM Oracle driver

## Description
Gorm原生go的驱动

早期版本为基于 [github.com/wdrabbit/gorm-oracle](https://github.com/wdrabbit/gorm-oracle), 升级 https://github.com/sijms/go-ora 到 v2 版本 (支持更新版本的oracle数据库)

使用方法参考原仓库

1. 不支持First之类的方法，因为oracle需要使用rownum进行分页
2. oracle://user:password@ip:port/service

- Update(2023-1-13): 原仓库也已经升级了依赖，可以用回原仓库
- Update(2023-1-13): 升级依赖的gorm到最新版本, 并修复了gorm该版本中limit类型从int变为*int的问题




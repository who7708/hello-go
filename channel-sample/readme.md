# 完全使用Go语言

Docker

k8s

Caddy

CockroachDB


# 部分使用

MongoDB/Couchbase

Dropbox

Uber

Google


# 语言发展趋势

https://octoverse.github.com

TIOBE

# Go 历史

2009 开源

2012 1.0 发布

2015 1.5 发布,自编译,重写垃圾回收器,更好的并发

目前 1.19



# 设计初衷

像C/C++ 那样的性能,可以做系统开发

没有繁琐的类型系统,有简单统一化的模块依赖管理,编译速度快

像java那样拥有垃圾回收

像python那样简单易学,援助灵活的类型,支持函数编程,异步IO

有编译器进行静态类型检查

加入并发编程

# Go语言的归类

类型检查: 编译时

运行环境: 编译成机器代码直接运行

编程范式: 面向接口,函数式编程,并发编程

# GO语言的并发编程

采用CSP(Communication Sequential Process)模型

不需要锁,不需要callbak,底层有,但是编写时不需要关心

并发编程vs并行计算


# 归并排序

将数据分为左右两半,分别归并排序,再把两个有序数据归并


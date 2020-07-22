## 框架简介

之前使用"云捷Go"进行开发，后发现前后端分离更适合Go，所以根据若依和云捷Go修改为了前后端分离的后台系统
该框架依然十分精简易上手，出错概率低。同时支持移动客户端访问。



## 核心技术及组件
> web服务框架    GoFrame        v1.13.1
>
> 若依前端组件   RuoYi-Vue           v2.2.0
>
> 权限组件       casbin/v2     v2.8.6
>
> 导出excel文件  tealeg/xlsx    v1.0.5   
>
> api文档生成    gf-swagger/swagger    v1.0.3
>
> 图形验证码     base64Captcha  v1.2.2  
>
> 服务器监控     gopsutil       v2.19.12+incompatible   

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  通知公告：系统通知公告信息发布维护。
9.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10.  登录日志：系统登录日志记录查询包含登录异常。
11.  在线用户：当前系统中活跃用户状态监控。
12.  定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13.  代码生成：前后端代码的生成（Go、vue、js、sql） 。
14.  系统接口：根据业务代码自动生成相关的api接口文档。
15.  服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
16.  在线构建器：拖动表单元素生成相应的HTML代码。
17.  案例演示：常用的前端组件整合演示。

## 使用方法
```
1. git clone git@github.com:1307super/gea.git
2. 导入document/sql/gea.sql
3. 修改配置文件
4. go run main.go 启动
```


## 感谢(排名不分先后)
> gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 
>
> RuoYi框架 [https://github.com/yangzongzhuan/RuoYi](https://github.com/yangzongzhuan/RuoYi)
>
> 云捷go [https://github.com/zeroing-zy1/yjgo](https://github.com/zeroing-zy1/yjgo)
>
> tealeg [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
>
> swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)
>
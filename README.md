## 项目简介
lottery-annual宗旨是打造一个简单的抽奖系统，通过简单的配置，快速开始一个抽奖活动。本人是一名JAVA开发人员，目前正在学习Golang，学以致用，以此记录自己的学习成果，有很多不足之处，欢迎交流指正，在此不胜感激！！！



## 部分截图

- 签到页面

  ![image-20241029210134627](.\doc\img\image-20241029210134627.png)

- 抽奖页面

  ![image-20241029210801812](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20241029210801812.png)

- 活动设置相关

  ![image-20241029210900684](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20241029210900684.png)

- 登录页面

  ![image-20241029211257937](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20241029211257937.png)

- 注册页面

  ![image-20241029211323691](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20241029211323691.png)

- 活动管理相关

  ![image-20241029211442715](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20241029211442715.png)



## 技术选型

本着方便开发，快速学习，项目太小，前后端不分离！！！

- ##### 前端相关

  抽奖首页：3D照片球改造，three.js

  后台管理：jquery、bootstrap相关如bootstrap-table、bootstrap-validate、bootstrap-datepicker、layer等等，

  在此，感谢<font color=red> 笔下光年 </font>提供的模板 <a> https://gitee.com/yinqi/Light-Year-Admin-Template</a>

- ##### 后端相关（go+gorm+redis）

  语言：Golang

  缓存Redis：redigo

  数据库：gorm

  安全：dgrijalva

  感谢后端 <font color=red>张奇峰 </font>的go框架：<a><font color=red> https://gitee.com/daitougege/GinSkeleton</font></a>

  

## 快速开始

1. 拷贝项目 git  clone  https://github.com/mingdw/lottery_annual.git

2. 执行脚本

   找到doc目录的lottery-annual.sql

3. 修改数据库配置

   找到resource 的-dev，修改数据库配置

4. 启动

   cd 到项目根目录，执行命令 go run lotteryStartMain.go



## 里程碑

```mermaid

gantt
        dateFormat  YYYY-MM-DD
        title 开发进度情况
        section 已完成
        后端框架搭建               :done,   des1, 2024-08-01,2024-08-08
        jwt安全校验               :done, des2 , 2024-08-08,2024-08-10
        使用redis缓存token               :done,  des3, 2024-08-10, 2d
        活动管理               :         done, des4,after des3, 5d
        登陆页面与功能  :         done, des5,after des4, 2d
        注册页面与功能  :         done, des6,after des5, 2d
        
        section 进行中
        抽奖页面3d球特效开发               :active,    des1, 2024-09-06,2024-09-20
        抽奖切换开发               :active,  des2, 2024-09-20, 5d
        抽奖配置页面开发               :active         des3, after des2, 5d
        活动管理奖项管理               :active         des4, after des3, 5d
        活动管理人员管理               :active         des5, after des4, 5d
        
        section 未开始
        后台管理个人资料               :    des1, 2024-11-15,2024-11-20
        3d旋转球调整特效               :  des2, after des1, 5d
        其它优化               :         des3, after des2, 10d


```

## 项目简介
lottery-annual宗旨是打造一个简单的抽奖系统，通过简单的配置，快速开始一个抽奖活动。本人是一名JAVA开发人员，目前正在学习Golang，学以致用，以此记录自己的学习成果，有很多不足之处，欢迎交流指正，在此不胜感激！！！



## 部分截图

- 签到页面

  ![image-20241029210134627](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029210134627.png)

- 抽奖页面

  ![image-20241029210801812](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029210801812.png)

- 活动设置相关

  ![image-20241029210900684](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029210900684.png)

- 登录页面

  ![image-20241029211257937](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029211257937.png)

- 注册页面

  ![image-20241029211323691](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029211323691.png)

- 活动管理相关

  ![image-20241029211442715](https://github.com/mingdw/lottery_annual/blob/master/doc/img/image-20241029211442715.png)

## 项目特点

- 简单快捷。  Golang前后端不分离，开箱即用，直接go run即可，要做的仅仅更改一下数据库配置

- 适合新手。本人也是新手，怎么简单方便怎么来，页面不多，想怎么改怎么改

- 操作简单，功能单一。功能简单，后台配置活动，前端负责抽奖，其它一概没有

  ...



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
  目前框架基本结构已搭建差不多了，许多细节和功能任然是待开发状态，由于非全职开发和对于前后端以及Golang的不熟悉，导致进度较慢，但是会一直更新下去

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



## 关于开源

目前本人并非全职开发，也不是专业的前端和Golang开发人员，也只是一个初学者，目前还有好多功能没有实现，同时对于已有的一些功能实现的也并不是太完美，希望有大佬路过指点，不胜感激。

站在巨人的肩膀上，前端主要借鉴 **[笔下光年](https://gitee.com/yinqi)/[Light Year Admin](https://gitee.com/yinqi/Light-Year-Admin-Template)**，后端主要框架借鉴 <a href="https://gitee.com/daitougege/GinSkeleton"><strong>张奇峰/GinSkeleton</strong></a> ，感谢你们的开源奉献！！！

本项目没有啥版权，喜欢就拿去玩去，别干违法犯罪的事即可（哈哈，貌似也没那功能）

未来，会一直持续更新下去，欢迎交流，欢迎指正！！！




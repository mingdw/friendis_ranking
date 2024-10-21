--系统用户表
CREATE TABLE `sys_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userName` varchar(32) DEFAULT NULL COMMENT '用户真实姓名',
  `nickName` varchar(64) DEFAULT NULL COMMENT '用户昵称',
  `account` varchar(64) DEFAULT NULL COMMENT '账号',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `phone` varchar(32) DEFAULT NULL COMMENT '电话',
  `job` varchar(64) DEFAULT NULL COMMENT '职业',
  `level` int(11) DEFAULT NULL COMMENT '级别',
  `birthDay` date DEFAULT NULL COMMENT '出生日期',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `registerTime` datetime DEFAULT NULL COMMENT '注册日期',
  `lastLoginTime` datetime DEFAULT NULL COMMENT '最后登录日期',
  `createTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) DEFAULT NULL COMMENT '修改人',
  `isDelete` int(11) DEFAULT NULL COMMENT '是否删除',
  `idCard` varchar(24) DEFAULT NULL COMMENT '身份证',
  `sex` int(4) DEFAULT NULL COMMENT '性别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

--活动表
CREATE TABLE `sys_activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `code` varchar(32) DEFAULT NULL COMMENT '活动编码',
  `title` varchar(64) DEFAULT NULL COMMENT '活动标题',
  `ac_desc` varchar(64) DEFAULT NULL COMMENT '活动描述',
  `startTime` timestamp NULL DEFAULT NULL COMMENT '开始时间',
  `endTime` timestamp NULL DEFAULT NULL COMMENT '结束时间',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '活动状态',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) DEFAULT NULL COMMENT '修改人',
  `isDelete` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4;

--奖项表
CREATE TABLE `sys_prize` (
  `id` int(11) NOT NULL,
  `ac_id` int(11) NOT NULL COMMENT '活动id',
  `name` varchar(255) NOT NULL COMMENT '奖项名称',
  `item_code` int(11) NOT NULL COMMENT '奖项编号',
  `prize_num` int(11) DEFAULT NULL COMMENT '中奖人数',
  `is_repeat` tinyint(4) DEFAULT NULL COMMENT '是否可重复抽取',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) DEFAULT NULL COMMENT '修改人',
  `isDelete` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--奖品表
CREATE TABLE `sys_prize_item` (
  `id` int(11) NOT NULL,
  `prize_id` int(11) NOT NULL COMMENT '活动id',
  `item_name` varchar(255) NOT NULL COMMENT '奖品名称',
  `item_code` int(11) NOT NULL COMMENT '奖项编号',
  `item_count` int(11) DEFAULT NULL COMMENT '奖品数量',
  `page_url` tinyint(4) DEFAULT NULL COMMENT '商品图片链接',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) DEFAULT NULL COMMENT '修改人',
  `isDelete` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--中奖记录表
CREATE TABLE `sys_activity_record` (
  `id` int(11) NOT NULL,
  `ac_id` int(11) NOT NULL COMMENT '活动id',
  `prize_id` int(11) NOT NULL COMMENT '奖项id',
  `prize_item_id` int(11) NOT NULL COMMENT '奖品id',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态（是否发放等等）',
  `win_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '中奖时间',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) DEFAULT NULL COMMENT '修改人',
  `isDelete` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


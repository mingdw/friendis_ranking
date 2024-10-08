CREATE TABLE `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `userName` varchar(32) DEFAULT NULL COMMENT '用户真实姓名',
  `nickName` varchar(64) DEFAULT NULL COMMENT '用户昵称',
  `account` varchar(64) DEFAULT NULL COMMENT '账号',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `idCard` varchar(64) DEFAULT NULL,
  `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
  `age` int DEFAULT NULL COMMENT '年龄',
  `phone` varchar(32) DEFAULT NULL COMMENT '电话',
  `job` varchar(64) DEFAULT NULL COMMENT '职业',
  `level` int DEFAULT NULL COMMENT '级别',
  `birthDay` date DEFAULT NULL COMMENT '出生日期',
  `status` int DEFAULT NULL COMMENT '状态',
  `registerTime` datetime DEFAULT NULL COMMENT '注册日期',
  `lastLoginTime` datetime DEFAULT NULL COMMENT '最后登录日期',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '修改人',
  `isDelete` int DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=109 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `sys_activity` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `code` varchar(32) DEFAULT NULL COMMENT '活动编码',
  `title` varchar(64) DEFAULT NULL COMMENT '活动标题',
  `desc` varchar(64) DEFAULT NULL COMMENT '活动描述',
  `startTime`  date NOT NULL COMMENT '开始时间',
  `endTime`  date NOT NULL COMMENT '结束时间',
  `status` int Not NULL DEFAULT 0 COMMENT '活动状态',
  `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `creator` varchar(64) DEFAULT NULL COMMENT '创建人',
  `editTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改日期',
  `editor` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '修改人',
  `isDelete` int DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
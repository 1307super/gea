/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : gea

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 01/03/2021 08:19:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `ptype` varchar(10) DEFAULT NULL,
  `v0` varchar(256) DEFAULT NULL,
  `v1` varchar(256) DEFAULT NULL,
  `v2` varchar(256) DEFAULT NULL,
  `v3` varchar(256) DEFAULT NULL,
  `v4` varchar(256) DEFAULT NULL,
  `v5` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('g', 'admin', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/online', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/server', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/build', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/operlog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/logininfor', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/import', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/user/resetPwd', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/menu', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dept', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/post/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/type/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/config/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/operlog/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/operlog', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/operlog/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/logininfor/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/logininfor', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/logininfor/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/online/batchLogout', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/online', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job/changeStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen/preview', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen/batchGenCode', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/tool/gen/db/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/job/run', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/jobLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/jobLog', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/jobLog/clean', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/dataScope', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/role/changeStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/system/dict/data/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/operlog/clean', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/monitor/logininfor/clean', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '测试测试2', 'pop', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/online', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/server', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/online/batchLogout', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/online', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job/changeStatus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job/export', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/job/run', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/jobLog', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/jobLog', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'pop', '/monitor/jobLog/clean', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'test008', 'admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'test007', 'admin', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table` (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) DEFAULT NULL COMMENT '其它生成选项',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COMMENT='代码生成业务表';

-- ----------------------------
-- Records of gen_table
-- ----------------------------
BEGIN;
INSERT INTO `gen_table` VALUES (33, 'sys_dept', '部门表', 'dept', 'crud', 'gea', 'system', 'dept', '部门', 'GEA', '{\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\"}', 'admin', '2020-02-16 20:58:20', 'admin', '2020-07-18 13:00:03', '');
INSERT INTO `gen_table` VALUES (34, 'sys_user_online', '在线用户记录', 'user_online', 'crud', 'yj-app', 'module', 'online', '在线用户记录', 'yunjie', '', 'admin', '2020-02-17 14:03:51', '', NULL, '');
INSERT INTO `gen_table` VALUES (35, 'sys_job', '定时任务调度表', 'job', 'crud', 'yj-app', 'module', 'job', '定时任务调度', 'yunjie', '', 'admin', '2020-02-18 15:44:13', '', NULL, '');
INSERT INTO `gen_table` VALUES (36, 'sys_job_log', '定时任务调度日志表', 'job_log', 'crud', 'yj-app', 'module', 'log', '定时任务调度日志', 'yunjie', '', 'admin', '2020-02-18 15:44:13', '', NULL, '');
INSERT INTO `gen_table` VALUES (37, 'econtract_user', '电子合同用户', 'econtracuser', 'crud', 'econtract', 'econtract', 'user', '电子合同用户', '1307', '', 'admin', '2020-04-16 16:35:17', 'admin', '2020-04-16 16:39:38', '');
INSERT INTO `gen_table` VALUES (38, 'econtract', '电子合同合同', 'econtract', 'crud', 'econtract', 'econtract', 'docment', '电子合同合同', '1307', '', 'admin', '2020-04-20 14:53:32', 'admin', '2020-04-20 14:56:17', '电子合同文档');
INSERT INTO `gen_table` VALUES (40, 'econtract_auth_temp', '电子合同用户id关联金格人脸认证业务编码', 'econtract_auth_temp', 'crud', 'econtract', 'econtract', 'auth_temp', '电子合同用户id关联金格人脸认证业务编码', 'yunjie', '{\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\"}', 'admin', '2020-06-22 21:28:02', 'admin', '2020-06-22 21:51:43', '');
INSERT INTO `gen_table` VALUES (41, 'casbin_rule', '', 'casbin_rule', 'crud', 'gea', 'module', 'rule', '', '1307', '', 'admin', '2020-07-20 23:56:27', '', '2020-07-20 23:41:34', '');
INSERT INTO `gen_table` VALUES (43, 't_tq', '测试表', 'tq', 'crud', 'gea', 'admin', 'tq', '测试1', '1307', '{\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\"}', 'admin', '2021-02-01 15:27:56', 'admin', '2021-02-01 16:33:14', '');
COMMIT;

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_table_column`;
CREATE TABLE `gen_table_column` (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) DEFAULT '' COMMENT '字典类型',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`column_id`)
) ENGINE=InnoDB AUTO_INCREMENT=442 DEFAULT CHARSET=utf8 COMMENT='代码生成业务表字段';

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------
BEGIN;
INSERT INTO `gen_table_column` VALUES (355, 33, 'dept_id', '部门id', 'bigint(20)', 'int64', 'DeptId', 'deptId', '1', '1', '0', '1', '', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (356, 33, 'parent_id', '父部门id', 'bigint(20)', 'int64', 'ParentId', 'parentId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (357, 33, 'ancestors', '祖级列表', 'varchar(50)', 'string', 'Ancestors', 'ancestors', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (358, 33, 'dept_name', '部门名称', 'varchar(30)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (359, 33, 'order_num', '显示顺序', 'int(4)', 'int', 'OrderNum', 'orderNum', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (360, 33, 'leader', '负责人', 'varchar(20)', 'string', 'Leader', 'leader', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (361, 33, 'phone', '联系电话', 'varchar(11)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (362, 33, 'email', '邮箱', 'varchar(50)', 'string', 'Email', 'email', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (363, 33, 'status', '部门状态（0正常 1停用）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (364, 33, 'del_flag', '删除标志（0代表存在 2代表删除）', 'char(1)', 'string', 'DelFlag', 'delFlag', '0', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (365, 33, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (366, 33, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (367, 33, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (368, 33, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 14, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (369, 34, 'sessionId', '用户会话id', 'varchar(50)', 'string', 'SessionId', 'sessionId', '1', '0', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (370, 34, 'login_name', '登录账号', 'varchar(50)', 'string', 'LoginName', 'loginName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (371, 34, 'dept_name', '部门名称', 'varchar(50)', 'string', 'DeptName', 'deptName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (372, 34, 'ipaddr', '登录IP地址', 'varchar(50)', 'string', 'Ipaddr', 'ipaddr', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (373, 34, 'login_location', '登录地点', 'varchar(255)', 'string', 'LoginLocation', 'loginLocation', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (374, 34, 'browser', '浏览器类型', 'varchar(50)', 'string', 'Browser', 'browser', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (375, 34, 'os', '操作系统', 'varchar(50)', 'string', 'Os', 'os', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (376, 34, 'status', '在线状态on_line在线off_line离线', 'varchar(10)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (377, 34, 'start_timestamp', 'session创建时间', 'datetime', 'Time', 'StartTimestamp', 'startTimestamp', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datatime', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (378, 34, 'last_access_time', 'session最后访问时间', 'datetime', 'Time', 'LastAccessTime', 'lastAccessTime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'datatime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (379, 34, 'expire_time', '超时时间，单位为分钟', 'int(5)', 'int', 'ExpireTime', 'expireTime', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (380, 35, 'job_id', '任务ID', 'bigint(20)', 'int64', 'JobId', 'jobId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (381, 35, 'job_name', '任务名称', 'varchar(64)', 'string', 'JobName', 'jobName', '1', '0', '1', '1', '0', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (382, 35, 'job_group', '任务组名', 'varchar(64)', 'string', 'JobGroup', 'jobGroup', '1', '0', '0', '1', '0', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (383, 35, 'invoke_target', '调用目标字符串', 'varchar(500)', 'string', 'InvokeTarget', 'invokeTarget', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (384, 35, 'cron_expression', 'cron执行表达式', 'varchar(255)', 'string', 'CronExpression', 'cronExpression', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (385, 35, 'misfire_policy', '计划执行错误策略（1立即执行 2执行一次 3放弃执行）', 'varchar(20)', 'string', 'MisfirePolicy', 'misfirePolicy', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (386, 35, 'concurrent', '是否并发执行（0允许 1禁止）', 'char(1)', 'string', 'Concurrent', 'concurrent', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (387, 35, 'status', '状态（0正常 1暂停）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (388, 35, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (389, 35, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (390, 35, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (391, 35, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (392, 35, 'remark', '备注信息', 'varchar(500)', 'string', 'Remark', 'remark', '0', '0', '0', '1', '1', '1', '0', 'EQ', 'input', '', 13, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (393, 36, 'job_log_id', '任务日志ID', 'bigint(20)', 'int64', 'JobLogId', 'jobLogId', '1', '1', '0', '1', '0', '1', '1', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (394, 36, 'job_name', '任务名称', 'varchar(64)', 'string', 'JobName', 'jobName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (395, 36, 'job_group', '任务组名', 'varchar(64)', 'string', 'JobGroup', 'jobGroup', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (396, 36, 'invoke_target', '调用目标字符串', 'varchar(500)', 'string', 'InvokeTarget', 'invokeTarget', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (397, 36, 'job_message', '日志信息', 'varchar(500)', 'string', 'JobMessage', 'jobMessage', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (398, 36, 'status', '执行状态（0正常 1失败）', 'char(1)', 'string', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (399, 36, 'exception_info', '异常信息', 'varchar(2000)', 'string', 'ExceptionInfo', 'exceptionInfo', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (400, 36, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datatime', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (401, 37, 'id', 'ID', 'int(11) unsigned', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (402, 37, 'name', '姓名', 'char(30)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (403, 37, 'idcard', '身份证号码', 'varchar(30)', 'string', 'Idcard', 'idcard', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (404, 37, 'phone', '手机号码', 'char(20)', 'string', 'Phone', 'phone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (405, 37, 'face_auth', '实名认证', 'tinyint(1) unsigned', 'int', 'FaceAuth', 'faceAuth', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', 'econtract_face_auth', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (406, 37, 'sso_id', '单点id', 'varchar(64)', 'string', 'SsoId', 'ssoId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (407, 37, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (408, 37, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (409, 37, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (410, 37, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (411, 38, 'id', 'ID', 'int(11) unsigned', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (412, 38, 'contractName', '合同名称', 'varchar(255)', 'string', 'ContractName', 'contractName', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (413, 38, 'contractCode', '合同编号', 'varchar(255)', 'string', 'ContractCode', 'contractCode', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (414, 38, 'filePath', '文件地址', 'varchar(255)', 'string', 'FilePath', 'filePath', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (415, 38, 'signPhone', '签署方手机号', 'varchar(255)', 'string', 'SignPhone', 'signPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (416, 38, 'signName', '签署方姓名', 'varchar(255)', 'string', 'SignName', 'signName', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (417, 38, 'cntId', '金格合同id', 'varchar(255)', 'string', 'CntId', 'cntId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (418, 38, 'status', '状态', 'tinyint(1)', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', '', 8, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (419, 38, 'create_by', '创建者', 'varchar(64)', 'string', 'CreateBy', 'createBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 9, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (420, 38, 'create_time', '创建时间', 'datetime', 'Time', 'CreateTime', 'createTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 10, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (421, 38, 'update_by', '更新者', 'varchar(64)', 'string', 'UpdateBy', 'updateBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 11, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (422, 38, 'update_time', '更新时间', 'datetime', 'Time', 'UpdateTime', 'updateTime', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 12, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (426, 40, 'id', '', 'int(11) unsigned', 'int64', 'Id', 'id', '1', '1', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (427, 40, 'user_id', '用户id', 'varchar(64)', 'string', 'UserId', 'userId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (428, 40, 'biz_id', '业务编码', 'varchar(64)', 'string', 'BizId', 'bizId', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (429, 41, 'ptype', '', 'varchar(10)', 'string', 'Ptype', 'ptype', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'select', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (430, 41, 'v0', '', 'varchar(256)', 'string', 'V0', 'v0', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (431, 41, 'v1', '', 'varchar(256)', 'string', 'V1', 'v1', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (432, 41, 'v2', '', 'varchar(256)', 'string', 'V2', 'v2', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (433, 41, 'v3', '', 'varchar(256)', 'string', 'V3', 'v3', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (434, 41, 'v4', '', 'varchar(256)', 'string', 'V4', 'v4', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (435, 41, 'v5', '', 'varchar(256)', 'string', 'V5', 'v5', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (439, 43, 'id', 'ID', 'int(11)', 'int64', 'Id', 'id', '1', '0', '0', '1', '0', '0', '0', 'EQ', 'input', '', 1, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (440, 43, 'test_name', '测试名称', 'varchar(255)', 'string', 'TestName', 'testName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, 'admin', NULL, '', NULL);
INSERT INTO `gen_table_column` VALUES (441, 43, 'test_phone', '测试手机号', 'varchar(255)', 'string', 'TestPhone', 'testPhone', '0', '0', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, 'admin', NULL, '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 'admin', '2018-03-16 11:33:00', '', '2020-02-12 15:32:15', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 'admin', '2018-03-16 11:33:00', '', '2020-02-05 10:46:28', '深黑主题theme-dark，浅色主题theme-light，深蓝主题theme-blue');
INSERT INTO `sys_config` VALUES (4, '静态资源网盘存储', 'sys.resource.url', 'http://127.0.0.1:8199/', 'Y', 'admin', '2020-02-18 20:10:33', '', '2020-02-19 10:36:22', 'public目录下的静态资源存储到OSS/COS等网盘，如果不需要动静分离设为null，如果需要设置OSS/COS等网盘网址即可');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(50) DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` char(1) DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8 COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'gea', 0, 'admin', '', '110@qq.com', '0', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2020-06-21 00:21:53');
INSERT INTO `sys_dept` VALUES (110, 100, '0,100', '测试部门', 1, '1307', '', '', '0', '0', 'admin', '2019-12-02 17:07:02', 'admin', '2020-07-18 11:34:19');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '性别男');
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '性别女');
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '显示菜单');
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '隐藏菜单');
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '默认分组');
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统分组');
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认是');
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统默认否');
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知');
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '公告');
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '关闭状态');
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '新增操作');
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '修改操作');
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '删除操作');
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '授权操作');
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导出操作');
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '导入操作');
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '强退操作');
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '生成操作');
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '清空操作');
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '正常状态');
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '停用状态');
INSERT INTO `sys_dict_data` VALUES (29, 0, '免费用户', '0', 'zjuser_type', '0', 'admin', '2019-12-02 16:56:16', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 1, '付费用户', '1', 'zjuser_type', '0', 'admin', '2019-12-02 16:56:40', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '微信用户', '0', 'zxuser_type', '0', 'admin', '2019-12-02 17:14:32', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 1, 'QQ用户', '1', 'zxuser_type', '0', 'admin', '2019-12-02 17:14:55', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (33, 2, '抖音用户', '2', 'zxuser_type', '0', 'admin', '2019-12-02 17:15:21', 'admin', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (34, 0, '未认证', '0', 'econtract_face_auth', '0', 'admin', '2020-04-16 16:37:51', '', NULL, '');
INSERT INTO `sys_dict_data` VALUES (35, 1, '已认证', '1', 'econtract_face_auth', '0', 'admin', '2020-04-16 16:38:07', '', NULL, '');
INSERT INTO `sys_dict_data` VALUES (36, 1, '在线', 'on_line', 'sys_online_status', '0', 'admin', '2020-07-16 21:00:14', '', NULL, '在线');
INSERT INTO `sys_dict_data` VALUES (37, 2, '离线', 'off_line', 'sys_online_status', '0', 'admin', '2020-07-16 21:00:32', '', NULL, '离线');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`),
  UNIQUE KEY `dict_type` (`dict_type`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '用户性别列表');
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '菜单状态列表');
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统开关列表');
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务状态列表');
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '任务分组列表');
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '系统是否列表');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知类型列表');
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '通知状态列表');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '操作类型列表');
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2018-03-16 11:33:00', 'admin', '2018-03-16 11:33:00', '登录状态列表');
INSERT INTO `sys_dict_type` VALUES (11, '专家用户类别', 'zjuser_type', '0', 'admin', '2019-12-02 16:55:42', 'admin', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (12, '咨询用户类别', 'zxuser_type', '0', 'admin', '2019-12-02 17:14:07', 'admin', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (15, '在线用户状态', 'sys_online_status', '0', 'admin', '2020-07-16 20:59:54', '', NULL, '用户在线状态');
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(255) DEFAULT NULL COMMENT '任务参数',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` char(1) DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` VALUES (9, 'test2', 'param1|param1', 'DEFAULT', 'test2', '* * * * * *', '1', '1', '1', 'admin', '2020-04-22 14:30:13', '', NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log` (
  `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) DEFAULT NULL COMMENT '日志信息',
  `status` char(1) DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) DEFAULT '' COMMENT '异常信息',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='定时任务调度日志表';

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor` (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` char(1) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`)
) ENGINE=InnoDB AUTO_INCREMENT=323 DEFAULT CHARSET=utf8 COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
BEGIN;
INSERT INTO `sys_logininfor` VALUES (264, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-18 13:10:00');
INSERT INTO `sys_logininfor` VALUES (265, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '账号或密码不正确', '2020-07-18 13:13:05');
INSERT INTO `sys_logininfor` VALUES (266, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-18 13:13:12');
INSERT INTO `sys_logininfor` VALUES (267, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 09:03:59');
INSERT INTO `sys_logininfor` VALUES (268, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 20:53:45');
INSERT INTO `sys_logininfor` VALUES (269, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 23:20:52');
INSERT INTO `sys_logininfor` VALUES (270, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 23:33:00');
INSERT INTO `sys_logininfor` VALUES (271, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 23:36:33');
INSERT INTO `sys_logininfor` VALUES (272, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 23:39:04');
INSERT INTO `sys_logininfor` VALUES (273, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-20 23:41:55');
INSERT INTO `sys_logininfor` VALUES (274, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 09:00:50');
INSERT INTO `sys_logininfor` VALUES (275, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 10:59:27');
INSERT INTO `sys_logininfor` VALUES (276, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 11:00:34');
INSERT INTO `sys_logininfor` VALUES (277, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 11:02:11');
INSERT INTO `sys_logininfor` VALUES (278, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 11:40:36');
INSERT INTO `sys_logininfor` VALUES (279, 'testtest', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2020-07-21 13:34:33');
INSERT INTO `sys_logininfor` VALUES (280, 'testtest', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2020-07-21 13:34:40');
INSERT INTO `sys_logininfor` VALUES (281, '测试测试2', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2020-07-21 13:35:38');
INSERT INTO `sys_logininfor` VALUES (282, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 13:35:51');
INSERT INTO `sys_logininfor` VALUES (283, '测试测试2', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 13:36:40');
INSERT INTO `sys_logininfor` VALUES (284, '测试测试2', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 13:39:44');
INSERT INTO `sys_logininfor` VALUES (285, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 14:43:00');
INSERT INTO `sys_logininfor` VALUES (286, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 14:44:51');
INSERT INTO `sys_logininfor` VALUES (287, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 14:52:31');
INSERT INTO `sys_logininfor` VALUES (288, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 15:08:39');
INSERT INTO `sys_logininfor` VALUES (289, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2020-07-21 15:18:29');
INSERT INTO `sys_logininfor` VALUES (290, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:14:18');
INSERT INTO `sys_logininfor` VALUES (291, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:15:40');
INSERT INTO `sys_logininfor` VALUES (292, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:17:04');
INSERT INTO `sys_logininfor` VALUES (293, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:18:28');
INSERT INTO `sys_logininfor` VALUES (294, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:25:27');
INSERT INTO `sys_logininfor` VALUES (295, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 10:35:21');
INSERT INTO `sys_logininfor` VALUES (296, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 15:32:26');
INSERT INTO `sys_logininfor` VALUES (297, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 16:33:54');
INSERT INTO `sys_logininfor` VALUES (298, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 16:35:52');
INSERT INTO `sys_logininfor` VALUES (299, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 16:38:21');
INSERT INTO `sys_logininfor` VALUES (300, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-22 16:46:01');
INSERT INTO `sys_logininfor` VALUES (301, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 09:21:39');
INSERT INTO `sys_logininfor` VALUES (302, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 09:22:23');
INSERT INTO `sys_logininfor` VALUES (303, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 09:22:40');
INSERT INTO `sys_logininfor` VALUES (304, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 09:23:37');
INSERT INTO `sys_logininfor` VALUES (305, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 09:57:46');
INSERT INTO `sys_logininfor` VALUES (306, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 10:03:33');
INSERT INTO `sys_logininfor` VALUES (307, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 10:08:05');
INSERT INTO `sys_logininfor` VALUES (308, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 10:10:22');
INSERT INTO `sys_logininfor` VALUES (309, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 10:14:59');
INSERT INTO `sys_logininfor` VALUES (310, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-23 10:23:03');
INSERT INTO `sys_logininfor` VALUES (311, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-25 08:16:53');
INSERT INTO `sys_logininfor` VALUES (312, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-25 08:28:08');
INSERT INTO `sys_logininfor` VALUES (313, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:22:07');
INSERT INTO `sys_logininfor` VALUES (314, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:23:08');
INSERT INTO `sys_logininfor` VALUES (315, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:26:49');
INSERT INTO `sys_logininfor` VALUES (316, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:27:17');
INSERT INTO `sys_logininfor` VALUES (317, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:28:34');
INSERT INTO `sys_logininfor` VALUES (318, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '1', '账号或密码不正确', '2021-01-25 09:29:20');
INSERT INTO `sys_logininfor` VALUES (319, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-25 09:39:55');
INSERT INTO `sys_logininfor` VALUES (320, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-25 09:42:09');
INSERT INTO `sys_logininfor` VALUES (321, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-01-25 09:51:10');
INSERT INTO `sys_logininfor` VALUES (322, 'admin', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', '0', '登陆成功', '2021-02-09 10:38:00');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `order_num` int(4) DEFAULT '0' COMMENT '显示顺序',
  `path` varchar(200) DEFAULT '#' COMMENT '请求地址',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `is_frame` tinyint(1) unsigned DEFAULT '1' COMMENT '打开方式（1页签 2新窗口）',
  `menu_type` char(1) DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` tinyint(1) unsigned DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `status` tinyint(1) unsigned DEFAULT '0' COMMENT '菜单状态（0正常 1停用）',
  `perms` varchar(100) DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT '' COMMENT '备注',
  `url` varchar(255) DEFAULT NULL COMMENT '接口地址',
  `method` char(10) DEFAULT NULL COMMENT '请求方法',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1077 DEFAULT CHARSET=utf8 COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` VALUES (1, '系统管理', 0, 1, 'system', '', 1, 'M', 0, 0, '', 'system', 'admin', '2021-01-25 08:28:49', 'admin', '2018-03-16 11:33:00', '', '', '');
INSERT INTO `sys_menu` VALUES (2, '系统监控', 0, 2, 'monitor', NULL, 1, 'M', 0, 0, '', 'monitor', 'admin', '2018-03-16 11:33:00', 'ry', '2018-03-16 11:33:00', '系统监控目录', NULL, NULL);
INSERT INTO `sys_menu` VALUES (3, '系统工具', 0, 3, 'tool', NULL, 1, 'M', 0, 0, '', 'tool', 'admin', '2018-03-16 11:33:00', 'ry', '2018-03-16 11:33:00', '系统工具目录', NULL, NULL);
INSERT INTO `sys_menu` VALUES (4, '若依官网', 0, 4, 'http://www.truckgogo.com2', '', 0, 'M', 0, 0, '', 'guide', '', '2021-01-25 09:56:46', 'admin', '2018-03-16 11:33:00', '', '', '');
INSERT INTO `sys_menu` VALUES (100, '用户管理', 1, 1, 'user', 'system/user/index', 0, 'C', 0, 0, '/system/user', 'user', '', '2020-07-20 23:39:10', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/user', 'GET');
INSERT INTO `sys_menu` VALUES (101, '角色管理', 1, 2, 'role', 'system/role/index', 0, 'C', 0, 0, '/system/role', 'peoples', '', '2020-07-20 23:42:13', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/role', 'GET');
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 1, 3, 'menu', 'system/menu/index', 0, 'C', 0, 0, '/system/menu', 'tree-table', '', '2020-07-20 23:42:23', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/menu', 'GET');
INSERT INTO `sys_menu` VALUES (103, '部门管理', 1, 4, 'dept', 'system/dept/index', 0, 'C', 0, 0, '/system/dept', 'tree', '', '2020-07-20 23:42:35', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/dept', 'GET');
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 1, 5, 'post', 'system/post/index', 0, 'C', 0, 0, '/system/post', 'post', '', '2020-07-20 23:42:45', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/post', 'GET');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 1, 6, 'dict', 'system/dict/index', 0, 'C', 0, 0, '/system/dict/type', 'dict', '', '2020-07-21 11:52:58', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/dict', 'GET');
INSERT INTO `sys_menu` VALUES (106, '参数设置', 1, 7, 'config', 'system/config/index', 0, 'C', 0, 0, '/system/config', 'edit', '', '2020-07-20 23:43:03', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/config', 'GET');
INSERT INTO `sys_menu` VALUES (108, '日志管理', 1, 9, 'log', 'system/log/index', 1, 'M', 0, 0, '', 'log', 'admin', '2018-03-16 11:33:00', 'ry', '2018-03-16 11:33:00', '日志管理菜单', NULL, NULL);
INSERT INTO `sys_menu` VALUES (109, '在线用户', 2, 1, 'online', 'monitor/online/index', 0, 'C', 0, 0, '/monitor/online', 'online', '', '2020-06-29 10:04:18', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/online', 'GET');
INSERT INTO `sys_menu` VALUES (110, '定时任务', 2, 2, 'job', 'monitor/job/index', 0, 'C', 0, 0, '/monitor/job', 'job', '', '2020-06-29 10:05:17', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/job', 'GET');
INSERT INTO `sys_menu` VALUES (111, '数据监控', 2, 3, 'druid', 'monitor/druid/index', 1, 'C', 1, 1, '/monitor/druid', 'druid', '', '2020-07-20 23:54:31', 'ry', '2018-03-16 11:33:00', '', '', 'GET');
INSERT INTO `sys_menu` VALUES (112, '服务监控', 2, 4, 'server', 'monitor/server/index', 0, 'C', 0, 0, '/monitor/server', 'server', '', '2020-06-29 10:06:33', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/server', 'GET');
INSERT INTO `sys_menu` VALUES (113, '表单构建', 3, 1, 'build', 'tool/build/index', 0, 'C', 0, 0, '/tool/build', 'build', '', '2020-07-20 23:54:41', 'ry', '2018-03-16 11:33:00', '', '', 'GET');
INSERT INTO `sys_menu` VALUES (114, '代码生成', 3, 2, 'gen', 'tool/gen/index', 0, 'C', 0, 0, '/tool/gen', 'code', '', '2020-06-29 10:07:41', 'ry', '2018-03-16 11:33:00', '', 'GET:/tool/gen', 'GET');
INSERT INTO `sys_menu` VALUES (115, '系统接口', 3, 3, 'swagger', '', 0, 'C', 1, 1, '/tool/swagger', 'swagger', '', '2020-06-23 08:44:22', 'ry', '2018-03-16 11:33:00', '', NULL, NULL);
INSERT INTO `sys_menu` VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', 0, 'C', 0, 0, '/monitor/operlog', 'form', '', '2020-06-29 10:01:51', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/operlog', 'GET');
INSERT INTO `sys_menu` VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', 0, 'C', 0, 0, '/monitor/logininfor', 'logininfor', '', '2020-06-29 10:02:06', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/logininfor', 'GET');
INSERT INTO `sys_menu` VALUES (1001, '用户查询', 100, 1, '', '', 0, 'F', 0, 0, '/system/user/info', '', '', '2020-07-20 23:43:33', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/user/info', 'GET');
INSERT INTO `sys_menu` VALUES (1002, '用户新增', 100, 2, '', '', 0, 'F', 0, 0, '/system/user', '', '', '2020-07-20 23:43:37', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/user', 'POST');
INSERT INTO `sys_menu` VALUES (1003, '用户修改', 100, 3, '', '', 0, 'F', 0, 0, '/system/user', '', '', '2020-07-20 23:43:41', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/user', 'PUT');
INSERT INTO `sys_menu` VALUES (1004, '用户删除', 100, 4, '', '', 0, 'F', 0, 0, '/system/user', '', '', '2020-07-20 23:43:53', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/user', 'DELETE');
INSERT INTO `sys_menu` VALUES (1005, '用户导出', 100, 5, '', '', 0, 'F', 0, 0, '/system/user/export', '', '', '2020-07-20 23:48:43', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/user/export', 'GET');
INSERT INTO `sys_menu` VALUES (1006, '用户导入', 100, 6, '', '', 0, 'F', 0, 0, '/system/user/import', '', '', '2020-07-20 23:48:50', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/user/import', 'GET');
INSERT INTO `sys_menu` VALUES (1007, '重置密码', 100, 7, '', '', 0, 'F', 0, 0, '/system/user/resetPwd', '', '', '2020-07-20 23:48:55', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/user/resetPwd', 'PUT');
INSERT INTO `sys_menu` VALUES (1008, '角色查询', 101, 1, '', '', 0, 'F', 0, 0, '/system/role/info', '', '', '2020-07-20 23:49:18', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/role/info', 'GET');
INSERT INTO `sys_menu` VALUES (1009, '角色新增', 101, 2, '', '', 0, 'F', 0, 0, '/system/role', '', '', '2020-07-20 23:49:14', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/role', 'POST');
INSERT INTO `sys_menu` VALUES (1010, '角色修改', 101, 3, '', '', 0, 'F', 0, 0, '/system/role', '', '', '2020-07-20 23:49:27', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/role', 'PUT');
INSERT INTO `sys_menu` VALUES (1011, '角色删除', 101, 4, '', '', 0, 'F', 0, 0, '/system/role', '', '', '2020-07-20 23:49:33', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/role', 'DELETE');
INSERT INTO `sys_menu` VALUES (1012, '角色导出', 101, 5, '', '', 0, 'F', 0, 0, '/system/role/export', '', '', '2020-07-20 23:49:38', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/role/export', 'GET');
INSERT INTO `sys_menu` VALUES (1013, '菜单查询', 102, 1, '', '', 0, 'F', 0, 0, '/system/menu/info', '', '', '2020-07-20 23:49:51', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/menu/info', 'GET');
INSERT INTO `sys_menu` VALUES (1014, '菜单新增', 102, 2, '', '', 0, 'F', 0, 0, '/system/menu', '', '', '2020-07-20 23:50:20', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/menu', 'POST');
INSERT INTO `sys_menu` VALUES (1015, '菜单修改', 102, 3, '', '', 0, 'F', 0, 0, '/system/menu', '', '', '2020-07-20 23:50:25', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/menu', 'PUT');
INSERT INTO `sys_menu` VALUES (1016, '菜单删除', 102, 4, '', '', 0, 'F', 0, 0, '/system/menu', '', '', '2020-07-20 23:50:31', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/menu', 'DELETE');
INSERT INTO `sys_menu` VALUES (1017, '部门查询', 103, 1, '', '', 0, 'F', 0, 0, '/system/dept/info', '', '', '2020-07-20 23:50:45', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/dept/info', 'GET');
INSERT INTO `sys_menu` VALUES (1018, '部门新增', 103, 2, '', '', 0, 'F', 0, 0, '/system/dept', '', '', '2020-07-20 23:50:51', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/dept', 'POST');
INSERT INTO `sys_menu` VALUES (1019, '部门修改', 103, 3, '', '', 0, 'F', 0, 0, '/system/dept', '', '', '2020-07-20 23:50:55', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/dept', 'PUT');
INSERT INTO `sys_menu` VALUES (1020, '部门删除', 103, 4, '', '', 0, 'F', 0, 0, '/system/dept', '', '', '2020-07-20 23:51:00', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/dept', 'DELETE');
INSERT INTO `sys_menu` VALUES (1021, '岗位查询', 104, 1, '', '', 0, 'F', 0, 0, '/system/post/info', '', '', '2020-07-20 23:51:08', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/post/info', 'GET');
INSERT INTO `sys_menu` VALUES (1022, '岗位新增', 104, 2, '', '', 0, 'F', 0, 0, '/system/post', '', '', '2020-07-20 23:51:14', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/post', 'POST');
INSERT INTO `sys_menu` VALUES (1023, '岗位修改', 104, 3, '', '', 0, 'F', 0, 0, '/system/post', '', '', '2020-07-20 23:51:19', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/post', 'PUT');
INSERT INTO `sys_menu` VALUES (1024, '岗位删除', 104, 4, '', '', 0, 'F', 0, 0, '/system/post', '', '', '2020-07-20 23:51:25', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/post', 'DELETE');
INSERT INTO `sys_menu` VALUES (1025, '岗位导出', 104, 5, '', '', 0, 'F', 0, 0, '/system/post/export', '', '', '2020-07-20 23:51:30', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/post/export', 'GET');
INSERT INTO `sys_menu` VALUES (1026, '字典查询', 105, 1, '', '', 0, 'F', 0, 0, '/system/dict/type/info', '', '', '2020-07-21 11:53:08', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/dict/info', 'GET');
INSERT INTO `sys_menu` VALUES (1027, '字典新增', 105, 2, '', '', 0, 'F', 0, 0, '/system/dict/type', '', '', '2020-07-21 11:53:14', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/dict', 'POST');
INSERT INTO `sys_menu` VALUES (1028, '字典修改', 105, 3, '', '', 0, 'F', 0, 0, '/system/dict/type', '', '', '2020-07-21 11:53:20', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/dict', 'PUT');
INSERT INTO `sys_menu` VALUES (1029, '字典删除', 105, 4, '', '', 0, 'F', 0, 0, '/system/dict/type', '', '', '2020-07-21 11:53:24', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/dict', 'DELETE');
INSERT INTO `sys_menu` VALUES (1030, '字典导出', 105, 5, '', '', 0, 'F', 0, 0, '/system/dict/type/export', '', '', '2020-07-21 11:53:28', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/dict/export', 'GET');
INSERT INTO `sys_menu` VALUES (1031, '参数查询', 106, 1, '', '', 0, 'F', 0, 0, '/system/config/info', '', '', '2020-07-20 23:52:07', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/config/info', 'GET');
INSERT INTO `sys_menu` VALUES (1032, '参数新增', 106, 2, '', '', 0, 'F', 0, 0, '/system/config', '', '', '2020-07-20 23:52:11', 'ry', '2018-03-16 11:33:00', '', 'POST:/system/config', 'POST');
INSERT INTO `sys_menu` VALUES (1033, '参数修改', 106, 3, '', '', 0, 'F', 0, 0, '/system/config', '', '', '2020-07-20 23:52:17', 'ry', '2018-03-16 11:33:00', '', 'PUT:/system/config', 'PUT');
INSERT INTO `sys_menu` VALUES (1034, '参数删除', 106, 4, '', '', 0, 'F', 0, 0, '/system/config', '', '', '2020-07-20 23:52:23', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/system/config', 'DELETE');
INSERT INTO `sys_menu` VALUES (1035, '参数导出', 106, 5, '', '', 0, 'F', 0, 0, '/system/config/export', '', '', '2020-07-20 23:52:28', 'ry', '2018-03-16 11:33:00', '', 'GET:/system/config/export', 'GET');
INSERT INTO `sys_menu` VALUES (1040, '操作查询', 500, 1, '', '', 0, 'F', 0, 0, '/monitor/operlog/info', '', '', '2020-07-20 23:52:46', 'ry', '2018-03-16 11:33:00', '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1041, '操作删除', 500, 2, '', '', 0, 'F', 0, 0, '/monitor/operlog', '', '', '2020-06-29 10:03:10', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/monitor/operlog', 'DELETE');
INSERT INTO `sys_menu` VALUES (1042, '日志导出', 500, 4, '', '', 0, 'F', 0, 0, '/monitor/operlog/export', '', '', '2020-07-20 23:52:58', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/operlog/export', 'GET');
INSERT INTO `sys_menu` VALUES (1043, '登录查询', 501, 1, '', '', 0, 'F', 0, 0, '/monitor/logininfor/info', '', '', '2020-07-20 23:53:11', 'ry', '2018-03-16 11:33:00', '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1044, '登录删除', 501, 2, '', '', 0, 'F', 0, 0, '/monitor/logininfor', '', '', '2020-07-20 23:53:18', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/monitor/logininfor', 'DELETE');
INSERT INTO `sys_menu` VALUES (1045, '日志导出', 501, 3, '', '', 0, 'F', 0, 0, '/monitor/logininfor/export', '', '', '2020-07-20 23:53:23', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/logininfor/export', 'GET');
INSERT INTO `sys_menu` VALUES (1047, '批量强退', 109, 2, '', '', 0, 'F', 0, 0, '/monitor/online/batchLogout', '', '', '2020-07-20 23:53:41', 'ry', '2018-03-16 11:33:00', '', 'PUT:/monitor/online', 'PUT');
INSERT INTO `sys_menu` VALUES (1048, '单条强退', 109, 3, '', '', 0, 'F', 0, 0, '/monitor/online', '', '', '2020-07-21 11:58:16', 'ry', '2018-03-16 11:33:00', '', 'PUT:/monitor/online', 'DELETE');
INSERT INTO `sys_menu` VALUES (1049, '任务查询', 110, 1, '', '', 0, 'F', 0, 0, '/monitor/job/info', '', '', '2020-07-20 23:53:57', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/job/info', 'GET');
INSERT INTO `sys_menu` VALUES (1050, '任务新增', 110, 2, '', '', 0, 'F', 0, 0, '/monitor/job', '', '', '2020-07-20 23:54:01', 'ry', '2018-03-16 11:33:00', '', 'POST:/monitor/job', 'POST');
INSERT INTO `sys_menu` VALUES (1051, '任务修改', 110, 3, '', '', 0, 'F', 0, 0, '/monitor/job', '', '', '2020-07-20 23:54:10', 'ry', '2018-03-16 11:33:00', '', 'PUT:/monitor/job', 'PUT');
INSERT INTO `sys_menu` VALUES (1052, '任务删除', 110, 4, '', '', 0, 'F', 0, 0, '/monitor/job', '', '', '2020-07-20 23:54:15', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/monitor/job', 'DELETE');
INSERT INTO `sys_menu` VALUES (1053, '状态修改', 110, 5, '', '', 0, 'F', 0, 0, '/monitor/job/changeStatus', '', '', '2020-07-20 23:54:20', 'ry', '2018-03-16 11:33:00', '', 'PUT:/monitor/job/changeStatus', 'PUT');
INSERT INTO `sys_menu` VALUES (1054, '任务导出', 110, 7, '', '', 0, 'F', 0, 0, '/monitor/job/export', '', '', '2020-07-20 23:54:26', 'ry', '2018-03-16 11:33:00', '', 'GET:/monitor/job/export', 'GET');
INSERT INTO `sys_menu` VALUES (1055, '生成查询', 114, 1, '', '', 0, 'F', 0, 0, '/tool/gen/info', '', '', '2020-06-29 10:07:55', 'ry', '2018-03-16 11:33:00', '', 'GET:/tool/gen/info', 'GET');
INSERT INTO `sys_menu` VALUES (1056, '生成修改', 114, 2, '', '', 0, 'F', 0, 0, '/tool/gen', '', '', '2020-07-20 23:55:56', 'ry', '2018-03-16 11:33:00', '', 'PUT:/tool/gen', 'PUT');
INSERT INTO `sys_menu` VALUES (1057, '生成删除', 114, 3, '', '', 0, 'F', 0, 0, '/tool/gen', '', '', '2020-06-29 10:08:09', 'ry', '2018-03-16 11:33:00', '', 'DELETE:/tool/gen', 'DELETE');
INSERT INTO `sys_menu` VALUES (1058, '导入代码', 114, 2, '', '', 0, 'F', 0, 0, '/tool/gen', '', '', '2020-07-20 23:57:08', 'ry', '2018-03-16 11:33:00', '', 'POST:/tool/gen/importTable', 'POST');
INSERT INTO `sys_menu` VALUES (1059, '预览代码', 114, 4, '', '', 0, 'F', 0, 0, '/tool/gen/preview', '', '', '2020-07-20 23:57:22', 'ry', '2018-03-16 11:33:00', '', 'GET:/tool/gen/preview', 'GET');
INSERT INTO `sys_menu` VALUES (1060, '生成代码', 114, 5, '', '', 0, 'F', 0, 0, '/tool/gen/batchGenCode', '', '', '2020-07-20 23:57:38', 'ry', '2018-03-16 11:33:00', '', 'GET:/tool/gen/batchGenCode', 'GET');
INSERT INTO `sys_menu` VALUES (1061, '表查询', 114, 6, '', '', 1, 'F', 0, 0, '/tool/gen/db/list', '', 'admin', '2020-07-20 23:59:35', '', NULL, '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1064, '执行一次', 110, 8, '', '', 1, 'F', 0, 0, '/monitor/job/run', '', 'admin', '2020-07-21 11:43:11', '', NULL, '', '', 'PUT');
INSERT INTO `sys_menu` VALUES (1065, '任务日志', 110, 0, '', '', 1, 'F', 0, 0, '/monitor/jobLog', '', 'admin', '2020-07-21 11:44:24', '', NULL, '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1066, '任务日志删除', 110, 0, '', '', 1, 'F', 0, 0, '/monitor/jobLog', '', 'admin', '2020-07-21 11:44:41', '', NULL, '', '', 'DELETE');
INSERT INTO `sys_menu` VALUES (1067, '任务日志清空', 110, 0, '', '', 1, 'F', 0, 0, '/monitor/jobLog/clean', '', 'admin', '2020-07-21 11:44:55', '', NULL, '', '', 'DELETE');
INSERT INTO `sys_menu` VALUES (1068, '数据权限', 101, 0, '', '', 1, 'F', 0, 0, '/system/role/dataScope', '', 'admin', '2020-07-21 11:48:48', '', NULL, '', '', 'PUT');
INSERT INTO `sys_menu` VALUES (1069, '修改状态', 101, 0, '', '', 1, 'F', 0, 0, '/system/role/changeStatus', '', 'admin', '2020-07-21 11:49:00', '', NULL, '', '', 'PUT');
INSERT INTO `sys_menu` VALUES (1070, '字典数据查询', 105, 0, '', '', 1, 'F', 0, 0, '/system/dict/data', '', 'admin', '2020-07-21 11:54:51', '', NULL, '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1071, '字典数据新增', 105, 0, '', '', 1, 'F', 0, 0, '/system/dict/data', '', 'admin', '2020-07-21 11:55:00', '', NULL, '', '', 'POST');
INSERT INTO `sys_menu` VALUES (1072, '字典数据修改', 105, 0, '', '', 1, 'F', 0, 0, '/system/dict/data', '', 'admin', '2020-07-21 11:55:13', '', NULL, '', '', 'PUT');
INSERT INTO `sys_menu` VALUES (1073, '字典数据删除', 105, 0, '', '', 1, 'F', 0, 0, '/system/dict/data', '', 'admin', '2020-07-21 11:55:22', '', NULL, '', '', 'DELETE');
INSERT INTO `sys_menu` VALUES (1074, '字典数据详情', 105, 0, '', '', 1, 'F', 0, 0, '/system/dict/data/info', '', 'admin', '2020-07-21 11:55:44', '', NULL, '', '', 'GET');
INSERT INTO `sys_menu` VALUES (1075, '日志清空', 500, 0, '', '', 1, 'F', 0, 0, '/monitor/operlog/clean', '', 'admin', '2020-07-21 11:56:41', '', NULL, '', '', 'DELETE');
INSERT INTO `sys_menu` VALUES (1076, '日志清空', 501, 0, '', '', 1, 'F', 0, 0, '/monitor/logininfor/clean', '', 'admin', '2020-07-21 11:57:03', '', NULL, '', '', 'DELETE');
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice` (
  `notice_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) NOT NULL COMMENT '公告标题',
  `notice_type` char(1) NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` varchar(2000) DEFAULT NULL COMMENT '公告内容',
  `status` char(1) DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知公告表';

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
  `oper_param` text COMMENT '请求参数',
  `json_result` text COMMENT '返回参数',
  `status` int(1) DEFAULT '0' COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1883 DEFAULT CHARSET=utf8 COMMENT='操作日志记录';

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_oper_log` VALUES (1712, '操作日志管理', 9, '/monitor/operlog/clean', 'DELETE', 1, 'admin', '信息部', '/monitor/operlog/clean', '127.0.0.1', '内网IP', '{\"jwtUid\":\"1\"}', '{\"code\":500,\"msg\":\"清空失败\",\"data\":\"\",\"otype\":9,\"module\":\"操作日志管理\"}', 1, '', '2020-07-18 11:38:04');
INSERT INTO `sys_oper_log` VALUES (1713, '登录日志管理', 3, '/monitor/logininfor?ids=263', 'DELETE', 1, 'admin', '信息部', '/monitor/logininfor?ids=263', '127.0.0.1', '内网IP', '{\"ids\":\"263\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"登录日志管理\"}', 0, '', '2020-07-18 11:38:22');
INSERT INTO `sys_oper_log` VALUES (1714, '登录日志管理', 9, '/monitor/logininfor/clean', 'DELETE', 1, 'admin', '信息部', '/monitor/logininfor/clean', '127.0.0.1', '内网IP', '{\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":9,\"module\":\"登录日志管理\"}', 0, '', '2020-07-18 11:38:27');
INSERT INTO `sys_oper_log` VALUES (1715, '定时任务管理', 2, '/monitor/job/changeStatus', 'PUT', 1, 'admin', '信息部', '/monitor/job/changeStatus', '127.0.0.1', '内网IP', '{\"jobId\":10,\"jwtUid\":\"1\",\"status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"定时任务管理\"}', 0, '', '2020-07-18 11:38:38');
INSERT INTO `sys_oper_log` VALUES (1716, '定时任务管理', 2, '/monitor/job/changeStatus', 'PUT', 1, 'admin', '信息部', '/monitor/job/changeStatus', '127.0.0.1', '内网IP', '{\"jobId\":10,\"jwtUid\":\"1\",\"status\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"定时任务管理\"}', 0, '', '2020-07-18 11:38:53');
INSERT INTO `sys_oper_log` VALUES (1717, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"deptIds\":[],\"jwtUid\":\"1\",\"menuIds\":\"\",\"role_key\":\"admin\",\"role_name\":\"ces\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"菜单不能为空\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2020-07-20 15:05:24');
INSERT INTO `sys_oper_log` VALUES (1718, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"deptIds\":[],\"jwtUid\":\"1\",\"menuIds\":\"1,100,1001,1002,1003,1004,1005,1006,1007,101,1008,1009,1010,1011,1012,102,1013,1014,1015,1016,103,1017,1018,1019,1020,104,1021,1022,1023,1024,1025,105,1026,1027,1028,1029,1030,106,1031,1032,1033,1034,1035,107,1036,1037,1038,1039,108,500,1040,1041,1042,501,1043,1044,1045\",\"role_key\":\"admin\",\"role_name\":\"ces\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"角色权限已存在\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2020-07-20 15:05:27');
INSERT INTO `sys_oper_log` VALUES (1719, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:33:40');
INSERT INTO `sys_oper_log` VALUES (1720, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1001,\"menu_name\":\"用户查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:34:47');
INSERT INTO `sys_oper_log` VALUES (1721, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1002,\"menu_name\":\"用户新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:34:57');
INSERT INTO `sys_oper_log` VALUES (1722, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1003,\"menu_name\":\"用户修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:35:12');
INSERT INTO `sys_oper_log` VALUES (1723, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:35:51');
INSERT INTO `sys_oper_log` VALUES (1724, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:36:50');
INSERT INTO `sys_oper_log` VALUES (1725, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:37:03');
INSERT INTO `sys_oper_log` VALUES (1726, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:37:40');
INSERT INTO `sys_oper_log` VALUES (1727, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/user/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"user\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":100,\"menu_name\":\"用户管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"user\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:39:10');
INSERT INTO `sys_oper_log` VALUES (1728, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/role/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"peoples\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":101,\"menu_name\":\"角色管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":2,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"role\",\"perms\":\"/system/role\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/role\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:42:13');
INSERT INTO `sys_oper_log` VALUES (1729, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/menu/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"tree-table\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":102,\"menu_name\":\"菜单管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":3,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"menu\",\"perms\":\"/system/menu\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/menu\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:42:23');
INSERT INTO `sys_oper_log` VALUES (1730, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/dept/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"tree\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":103,\"menu_name\":\"部门管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":4,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"dept\",\"perms\":\"/system/dept\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dept\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:42:35');
INSERT INTO `sys_oper_log` VALUES (1731, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/post/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"post\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":104,\"menu_name\":\"岗位管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"post\",\"perms\":\"/system/post\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/post\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:42:45');
INSERT INTO `sys_oper_log` VALUES (1732, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/dict/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"dict\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":105,\"menu_name\":\"字典管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":6,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"dict\",\"perms\":\"/system/dict\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:42:55');
INSERT INTO `sys_oper_log` VALUES (1733, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/config/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"edit\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":106,\"menu_name\":\"参数设置\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":7,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"config\",\"perms\":\"/system/config\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/config\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:03');
INSERT INTO `sys_oper_log` VALUES (1734, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/notice/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"message\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":107,\"menu_name\":\"通知公告\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":8,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"notice\",\"perms\":\"/system/notice\",\"remark\":\"通知公告菜单\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:12');
INSERT INTO `sys_oper_log` VALUES (1735, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1001,\"menu_name\":\"用户查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:33');
INSERT INTO `sys_oper_log` VALUES (1736, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1002,\"menu_name\":\"用户新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:37');
INSERT INTO `sys_oper_log` VALUES (1737, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1003,\"menu_name\":\"用户修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:41');
INSERT INTO `sys_oper_log` VALUES (1738, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1004,\"menu_name\":\"用户删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"system:user:remove\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:44');
INSERT INTO `sys_oper_log` VALUES (1739, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1004,\"menu_name\":\"用户删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/user\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:43:53');
INSERT INTO `sys_oper_log` VALUES (1740, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1005,\"menu_name\":\"用户导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:48:43');
INSERT INTO `sys_oper_log` VALUES (1741, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1006,\"menu_name\":\"用户导入\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":6,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user/import\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/user/import\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:48:50');
INSERT INTO `sys_oper_log` VALUES (1742, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1007,\"menu_name\":\"重置密码\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":7,\"parentName\":\"用户管理\",\"parent_id\":100,\"path\":\"\",\"perms\":\"/system/user/resetPwd\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/user/resetPwd\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:48:55');
INSERT INTO `sys_oper_log` VALUES (1743, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1008,\"menu_name\":\"角色查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role/query\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/role/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:08');
INSERT INTO `sys_oper_log` VALUES (1744, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1009,\"menu_name\":\"角色新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/role\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:14');
INSERT INTO `sys_oper_log` VALUES (1745, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1008,\"menu_name\":\"角色查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/role/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:18');
INSERT INTO `sys_oper_log` VALUES (1746, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1010,\"menu_name\":\"角色修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/role\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:27');
INSERT INTO `sys_oper_log` VALUES (1747, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1011,\"menu_name\":\"角色删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/role\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:33');
INSERT INTO `sys_oper_log` VALUES (1748, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1012,\"menu_name\":\"角色导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"角色管理\",\"parent_id\":101,\"path\":\"\",\"perms\":\"/system/role/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/role/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:38');
INSERT INTO `sys_oper_log` VALUES (1749, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1013,\"menu_name\":\"菜单查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"菜单管理\",\"parent_id\":102,\"path\":\"\",\"perms\":\"/system/menu/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/menu/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:49:51');
INSERT INTO `sys_oper_log` VALUES (1750, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1014,\"menu_name\":\"菜单新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"菜单管理\",\"parent_id\":102,\"path\":\"\",\"perms\":\"/system/menu\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/menu\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:20');
INSERT INTO `sys_oper_log` VALUES (1751, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1015,\"menu_name\":\"菜单修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"菜单管理\",\"parent_id\":102,\"path\":\"\",\"perms\":\"/system/menu\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/menu\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:25');
INSERT INTO `sys_oper_log` VALUES (1752, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1016,\"menu_name\":\"菜单删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"菜单管理\",\"parent_id\":102,\"path\":\"\",\"perms\":\"/system/menu\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/menu\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:31');
INSERT INTO `sys_oper_log` VALUES (1753, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1017,\"menu_name\":\"部门查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"部门管理\",\"parent_id\":103,\"path\":\"\",\"perms\":\"/system/dept/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dept/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:45');
INSERT INTO `sys_oper_log` VALUES (1754, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1018,\"menu_name\":\"部门新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"部门管理\",\"parent_id\":103,\"path\":\"\",\"perms\":\"/system/dept\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/dept\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:51');
INSERT INTO `sys_oper_log` VALUES (1755, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1019,\"menu_name\":\"部门修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"部门管理\",\"parent_id\":103,\"path\":\"\",\"perms\":\"/system/dept\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/dept\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:50:55');
INSERT INTO `sys_oper_log` VALUES (1756, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1020,\"menu_name\":\"部门删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"部门管理\",\"parent_id\":103,\"path\":\"\",\"perms\":\"/system/dept\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/dept\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:00');
INSERT INTO `sys_oper_log` VALUES (1757, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1021,\"menu_name\":\"岗位查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"岗位管理\",\"parent_id\":104,\"path\":\"\",\"perms\":\"/system/post/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/post/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:08');
INSERT INTO `sys_oper_log` VALUES (1758, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1022,\"menu_name\":\"岗位新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"岗位管理\",\"parent_id\":104,\"path\":\"\",\"perms\":\"/system/post\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/post\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:14');
INSERT INTO `sys_oper_log` VALUES (1759, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1023,\"menu_name\":\"岗位修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"岗位管理\",\"parent_id\":104,\"path\":\"\",\"perms\":\"/system/post\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/post\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:19');
INSERT INTO `sys_oper_log` VALUES (1760, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1024,\"menu_name\":\"岗位删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"岗位管理\",\"parent_id\":104,\"path\":\"\",\"perms\":\"/system/post\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/post\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:25');
INSERT INTO `sys_oper_log` VALUES (1761, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1025,\"menu_name\":\"岗位导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"岗位管理\",\"parent_id\":104,\"path\":\"\",\"perms\":\"/system/post/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/post/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:30');
INSERT INTO `sys_oper_log` VALUES (1762, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1026,\"menu_name\":\"字典查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:37');
INSERT INTO `sys_oper_log` VALUES (1763, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1027,\"menu_name\":\"字典新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:42');
INSERT INTO `sys_oper_log` VALUES (1764, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1028,\"menu_name\":\"字典修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:48');
INSERT INTO `sys_oper_log` VALUES (1765, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1029,\"menu_name\":\"字典删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:53');
INSERT INTO `sys_oper_log` VALUES (1766, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1030,\"menu_name\":\"字典导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:51:57');
INSERT INTO `sys_oper_log` VALUES (1767, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1031,\"menu_name\":\"参数查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"参数设置\",\"parent_id\":106,\"path\":\"\",\"perms\":\"/system/config/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/config/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:07');
INSERT INTO `sys_oper_log` VALUES (1768, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1032,\"menu_name\":\"参数新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"参数设置\",\"parent_id\":106,\"path\":\"\",\"perms\":\"/system/config\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/config\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:11');
INSERT INTO `sys_oper_log` VALUES (1769, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1033,\"menu_name\":\"参数修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"参数设置\",\"parent_id\":106,\"path\":\"\",\"perms\":\"/system/config\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/config\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:17');
INSERT INTO `sys_oper_log` VALUES (1770, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1034,\"menu_name\":\"参数删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"参数设置\",\"parent_id\":106,\"path\":\"\",\"perms\":\"/system/config\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/config\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:23');
INSERT INTO `sys_oper_log` VALUES (1771, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1035,\"menu_name\":\"参数导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"参数设置\",\"parent_id\":106,\"path\":\"\",\"perms\":\"/system/config/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/config/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:28');
INSERT INTO `sys_oper_log` VALUES (1772, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"#\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1040,\"menu_name\":\"操作查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"操作日志\",\"parent_id\":500,\"path\":\"#\",\"perms\":\"/monitor/operlog/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:46');
INSERT INTO `sys_oper_log` VALUES (1773, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1042,\"menu_name\":\"日志导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":4,\"parentName\":\"操作日志\",\"parent_id\":500,\"path\":\"\",\"perms\":\"/monitor/operlog/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/monitor/operlog/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:52:58');
INSERT INTO `sys_oper_log` VALUES (1774, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"#\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1043,\"menu_name\":\"登录查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"登录日志\",\"parent_id\":501,\"path\":\"#\",\"perms\":\"/monitor/logininfor/query\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:08');
INSERT INTO `sys_oper_log` VALUES (1775, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1043,\"menu_name\":\"登录查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"登录日志\",\"parent_id\":501,\"path\":\"\",\"perms\":\"/monitor/logininfor/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:12');
INSERT INTO `sys_oper_log` VALUES (1776, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1044,\"menu_name\":\"登录删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":2,\"parentName\":\"登录日志\",\"parent_id\":501,\"path\":\"\",\"perms\":\"/monitor/logininfor\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/monitor/logininfor\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:18');
INSERT INTO `sys_oper_log` VALUES (1777, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1045,\"menu_name\":\"日志导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":3,\"parentName\":\"登录日志\",\"parent_id\":501,\"path\":\"\",\"perms\":\"/monitor/logininfor/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/monitor/logininfor/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:23');
INSERT INTO `sys_oper_log` VALUES (1778, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1046,\"menu_name\":\"在线查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"在线用户\",\"parent_id\":109,\"path\":\"\",\"perms\":\"/monitor/online/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/monitor/online/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:35');
INSERT INTO `sys_oper_log` VALUES (1779, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1047,\"menu_name\":\"批量强退\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":2,\"parentName\":\"在线用户\",\"parent_id\":109,\"path\":\"\",\"perms\":\"/monitor/online/batchLogout\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/monitor/online\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:41');
INSERT INTO `sys_oper_log` VALUES (1780, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1048,\"menu_name\":\"单条强退\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"在线用户\",\"parent_id\":109,\"path\":\"\",\"perms\":\"/monitor/online/forceLogout\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/monitor/online\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:47');
INSERT INTO `sys_oper_log` VALUES (1781, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1049,\"menu_name\":\"任务查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/monitor/job/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:53:57');
INSERT INTO `sys_oper_log` VALUES (1782, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1050,\"menu_name\":\"任务新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/monitor/job\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:01');
INSERT INTO `sys_oper_log` VALUES (1783, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1051,\"menu_name\":\"任务修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/monitor/job\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:10');
INSERT INTO `sys_oper_log` VALUES (1784, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1052,\"menu_name\":\"任务删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/monitor/job\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:15');
INSERT INTO `sys_oper_log` VALUES (1785, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1053,\"menu_name\":\"状态修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":5,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job/changeStatus\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/monitor/job/changeStatus\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:20');
INSERT INTO `sys_oper_log` VALUES (1786, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1054,\"menu_name\":\"任务导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":7,\"parentName\":\"定时任务\",\"parent_id\":110,\"path\":\"\",\"perms\":\"/monitor/job/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/monitor/job/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:26');
INSERT INTO `sys_oper_log` VALUES (1787, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"monitor/druid/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"druid\",\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":111,\"menu_name\":\"数据监控\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":3,\"parentName\":\"系统监控\",\"parent_id\":2,\"path\":\"druid\",\"perms\":\"/monitor/druid\",\"remark\":\"数据监控菜单\",\"status\":\"1\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:31');
INSERT INTO `sys_oper_log` VALUES (1788, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"tool/build/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"build\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":113,\"menu_name\":\"表单构建\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"系统工具\",\"parent_id\":3,\"path\":\"build\",\"perms\":\"/tool/build\",\"remark\":\"表单构建菜单\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:54:41');
INSERT INTO `sys_oper_log` VALUES (1789, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1060,\"menu_name\":\"生成代码\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"代码生成\",\"parent_id\":114,\"path\":\"\",\"perms\":\"/tool/gen\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/tool/gen/batchGenCode\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:55:27');
INSERT INTO `sys_oper_log` VALUES (1790, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1056,\"menu_name\":\"生成修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":2,\"parentName\":\"代码生成\",\"parent_id\":114,\"path\":\"\",\"perms\":\"/tool/gen\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/tool/gen\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:55:56');
INSERT INTO `sys_oper_log` VALUES (1791, '代码生成管理', 1, '/tool/gen?tables=casbin_rule', 'POST', 1, 'admin', '信息部', '/tool/gen?tables=casbin_rule', '127.0.0.1', '内网IP', '{\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"tables\":\"casbin_rule\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2020-07-20 23:56:27');
INSERT INTO `sys_oper_log` VALUES (1792, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1058,\"menu_name\":\"导入代码\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"代码生成\",\"parent_id\":114,\"path\":\"\",\"perms\":\"/tool/gen\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/tool/gen/importTable\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:57:08');
INSERT INTO `sys_oper_log` VALUES (1793, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1059,\"menu_name\":\"预览代码\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":4,\"parentName\":\"代码生成\",\"parent_id\":114,\"path\":\"\",\"perms\":\"/tool/gen/preview\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/tool/gen/preview\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:57:22');
INSERT INTO `sys_oper_log` VALUES (1794, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1060,\"menu_name\":\"生成代码\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"代码生成\",\"parent_id\":114,\"path\":\"\",\"perms\":\"/tool/gen/batchGenCode\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/tool/gen/batchGenCode\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:57:38');
INSERT INTO `sys_oper_log` VALUES (1795, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"表查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":6,\"parent_id\":114,\"perms\":\"/tool/gen/db/list\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-20 23:59:35');
INSERT INTO `sys_oper_log` VALUES (1796, '菜单管理', 3, '/system/menu?id=1039', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1039', '127.0.0.1', '内网IP', '{\"id\":\"1039\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 09:01:46');
INSERT INTO `sys_oper_log` VALUES (1797, '菜单管理', 3, '/system/menu?id=1038', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1038', '127.0.0.1', '内网IP', '{\"id\":\"1038\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 09:01:48');
INSERT INTO `sys_oper_log` VALUES (1798, '菜单管理', 3, '/system/menu?id=1037', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1037', '127.0.0.1', '内网IP', '{\"id\":\"1037\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 09:01:50');
INSERT INTO `sys_oper_log` VALUES (1799, '菜单管理', 3, '/system/menu?id=1036', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1036', '127.0.0.1', '内网IP', '{\"id\":\"1036\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 09:01:51');
INSERT INTO `sys_oper_log` VALUES (1800, '菜单管理', 3, '/system/menu?id=107', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=107', '127.0.0.1', '内网IP', '{\"id\":\"107\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 09:01:54');
INSERT INTO `sys_oper_log` VALUES (1801, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"用户信息（权限）\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":8,\"parent_id\":100,\"perms\":\"/getInfo\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:02:49');
INSERT INTO `sys_oper_log` VALUES (1802, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"权限路由\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parent_id\":102,\"perms\":\"/getRouters\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:03:15');
INSERT INTO `sys_oper_log` VALUES (1803, '菜单管理', 3, '/system/menu?id=1062', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1062', '127.0.0.1', '内网IP', '{\"id\":\"1062\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:33:12');
INSERT INTO `sys_oper_log` VALUES (1804, '菜单管理', 3, '/system/menu?id=1063', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1063', '127.0.0.1', '内网IP', '{\"id\":\"1063\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:33:17');
INSERT INTO `sys_oper_log` VALUES (1805, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"执行一次\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":8,\"parent_id\":110,\"perms\":\"/monitor/job/run\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:43:11');
INSERT INTO `sys_oper_log` VALUES (1806, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"任务日志\",\"menu_type\":\"F\",\"method\":\"GET\",\"parent_id\":110,\"perms\":\"/monitor/jobLog\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:44:24');
INSERT INTO `sys_oper_log` VALUES (1807, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"任务日志删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"parent_id\":110,\"perms\":\"/monitor/jobLog\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:44:41');
INSERT INTO `sys_oper_log` VALUES (1808, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"任务日志清空\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"parent_id\":110,\"perms\":\"/monitor/jobLog/clean\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:44:55');
INSERT INTO `sys_oper_log` VALUES (1809, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"数据权限\",\"menu_type\":\"F\",\"method\":\"PUT\",\"parent_id\":101,\"perms\":\"/system/role/dataScope\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:48:48');
INSERT INTO `sys_oper_log` VALUES (1810, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"修改状态\",\"menu_type\":\"F\",\"method\":\"PUT\",\"parent_id\":101,\"perms\":\"/system/role/changeStatus\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:49:00');
INSERT INTO `sys_oper_log` VALUES (1811, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"system/dict/index\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"dict\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":105,\"menu_name\":\"字典管理\",\"menu_type\":\"C\",\"method\":\"GET\",\"order_num\":6,\"parentName\":\"系统管理\",\"parent_id\":1,\"path\":\"dict\",\"perms\":\"/system/dict/type\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:52:58');
INSERT INTO `sys_oper_log` VALUES (1812, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1026,\"menu_name\":\"字典查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":1,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/type/info\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict/info\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:53:08');
INSERT INTO `sys_oper_log` VALUES (1813, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1027,\"menu_name\":\"字典新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"order_num\":2,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/type\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"POST:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:53:14');
INSERT INTO `sys_oper_log` VALUES (1814, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1028,\"menu_name\":\"字典修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"order_num\":3,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/type\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:53:20');
INSERT INTO `sys_oper_log` VALUES (1815, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1029,\"menu_name\":\"字典删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":4,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/type\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"DELETE:/system/dict\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:53:24');
INSERT INTO `sys_oper_log` VALUES (1816, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1030,\"menu_name\":\"字典导出\",\"menu_type\":\"F\",\"method\":\"GET\",\"order_num\":5,\"parentName\":\"字典管理\",\"parent_id\":105,\"path\":\"\",\"perms\":\"/system/dict/type/export\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"GET:/system/dict/export\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:53:28');
INSERT INTO `sys_oper_log` VALUES (1817, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"字典数据查询\",\"menu_type\":\"F\",\"method\":\"GET\",\"parent_id\":105,\"perms\":\"/system/dict/data\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:54:51');
INSERT INTO `sys_oper_log` VALUES (1818, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"字典数据新增\",\"menu_type\":\"F\",\"method\":\"POST\",\"parent_id\":105,\"perms\":\"/system/dict/data\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:55:00');
INSERT INTO `sys_oper_log` VALUES (1819, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"字典数据修改\",\"menu_type\":\"F\",\"method\":\"PUT\",\"parent_id\":105,\"perms\":\"/system/dict/data\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:55:13');
INSERT INTO `sys_oper_log` VALUES (1820, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"字典数据删除\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"parent_id\":105,\"perms\":\"/system/dict/data\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:55:22');
INSERT INTO `sys_oper_log` VALUES (1821, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"字典数据详情\",\"menu_type\":\"F\",\"method\":\"GET\",\"parent_id\":105,\"perms\":\"/system/dict/data/info\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:55:44');
INSERT INTO `sys_oper_log` VALUES (1822, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"日志清空\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"parent_id\":500,\"perms\":\"/monitor/operlog/clean\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:56:41');
INSERT INTO `sys_oper_log` VALUES (1823, '菜单管理', 1, '/system/menu', 'POST', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"is_frame\":\"1\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_name\":\"日志清空\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"parent_id\":501,\"perms\":\"/monitor/logininfor/clean\",\"status\":\"0\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:57:03');
INSERT INTO `sys_oper_log` VALUES (1824, '菜单管理', 3, '/system/menu?id=1046', 'DELETE', 1, 'admin', '信息部', '/system/menu?id=1046', '127.0.0.1', '内网IP', '{\"id\":\"1046\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:58:08');
INSERT INTO `sys_oper_log` VALUES (1825, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '信息部', '/system/menu', '127.0.0.1', '内网IP', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"\",\"is_frame\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menu_id\":1048,\"menu_name\":\"单条强退\",\"menu_type\":\"F\",\"method\":\"DELETE\",\"order_num\":3,\"parentName\":\"在线用户\",\"parent_id\":109,\"path\":\"\",\"perms\":\"/monitor/online\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"PUT:/monitor/online\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2020-07-21 11:58:16');
INSERT INTO `sys_oper_log` VALUES (1826, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2018-03-16 11:33:00\",\"dataScope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"1,100,1001,1002,1003,1004,1005,1006,1007,101,1008,1009,1010,1011,1012,1068,1069,102,1013,1014,1015,1016,103,1017,1018,1019,1020,104,1021,1022,1023,1024,1025,105,1026,1027,1028,1029,1030,1070,1071,1072,1073,1074,106,1031,1032,1033,1034,1035,108,500,1040,1041,1042,1075,501,1043,1044,1045,1076,2,109,1047,1048,110,1049,1050,1051,1052,1053,1054,1064,1065,1066,1067,111,112,3,113,114,1055,1056,1057,1058,1059,1060,1061,115,4\",\"remark\":\"管理员\",\"role_id\":1,\"role_key\":\"admin\",\"role_name\":\"管理员\",\"role_sort\":1,\"status\":\"0\",\"update_by\":\"\",\"update_time\":\"2020-05-08 16:44:11\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 12:05:42');
INSERT INTO `sys_oper_log` VALUES (1827, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2018-03-16 11:33:00\",\"dataScope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"1,100,1001,1002,1003,1004,1005,1006,1007,101,1008,1009,1010,1011,1012,1068,1069,102,1013,1014,1015,1016,103,1017,1018,1019,1020,104,1021,1022,1023,1024,1025,105,1026,1027,1028,1029,1030,1070,1071,1072,1073,1074,106,1031,1032,1033,1034,1035,108,500,1040,1041,1042,1075,501,1043,1044,1045,1076,2,109,1047,1048,110,1049,1050,1051,1052,1053,1054,1064,1065,1066,1067,111,112,3,113,114,1055,1056,1057,1058,1059,1060,1061,115,4\",\"remark\":\"管理员\",\"role_id\":1,\"role_key\":\"admin\",\"role_name\":\"管理员\",\"role_sort\":1,\"status\":\"0\",\"update_by\":\"\",\"update_time\":\"2020-07-21 12:05:42\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 12:06:26');
INSERT INTO `sys_oper_log` VALUES (1828, '用户管理', 2, '/system/user', 'PUT', 1, 'admin', '信息部', '/system/user', '127.0.0.1', '内网IP', '{\"avatar\":\"\",\"create_by\":\"admin\",\"create_time\":\"2020-07-17 14:22:00\",\"del_flag\":\"0\",\"dept\":null,\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"login_date\":null,\"login_ip\":\"\",\"login_name\":\"测试测试\",\"password\":\"\",\"phonenumber\":\"18983690296\",\"postIds\":\"1\",\"remark\":\"测试\",\"roleIds\":\"1\",\"roles\":null,\"salt\":\"e3B7rF\",\"sex\":\"0\",\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2020-07-17 21:45:14\",\"user_id\":6,\"user_name\":\"test\",\"user_type\":\"\"}', '{\"code\":500,\"msg\":\"用户名称长度为5到30位\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 1, '', '2020-07-21 13:34:08');
INSERT INTO `sys_oper_log` VALUES (1829, '用户管理', 2, '/system/user', 'PUT', 1, 'admin', '信息部', '/system/user', '127.0.0.1', '内网IP', '{\"avatar\":\"\",\"create_by\":\"admin\",\"create_time\":\"2020-07-17 14:22:00\",\"del_flag\":\"0\",\"dept\":null,\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"login_date\":null,\"login_ip\":\"\",\"login_name\":\"测试测试\",\"password\":\"\",\"phonenumber\":\"18983690296\",\"postIds\":\"1\",\"remark\":\"测试\",\"roleIds\":\"1\",\"roles\":null,\"salt\":\"e3B7rF\",\"sex\":\"0\",\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2020-07-17 21:45:14\",\"user_id\":6,\"user_name\":\"testtest\",\"user_type\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2020-07-21 13:34:17');
INSERT INTO `sys_oper_log` VALUES (1830, '用户管理', 2, '/system/user/resetPwd', 'PUT', 1, 'admin', '信息部', '/system/user/resetPwd', '127.0.0.1', '内网IP', '{\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"password\":\"123456\",\"userId\":6}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2020-07-21 13:34:21');
INSERT INTO `sys_oper_log` VALUES (1831, '用户管理', 2, '/system/user/resetPwd', 'PUT', 1, 'admin', '信息部', '/system/user/resetPwd', '127.0.0.1', '内网IP', '{\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"password\":\"123456\",\"userId\":6}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2020-07-21 13:36:28');
INSERT INTO `sys_oper_log` VALUES (1832, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"deptIds\":[],\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"4\",\"role_key\":\"test\",\"role_name\":\"测试角色\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2020-07-21 14:43:11');
INSERT INTO `sys_oper_log` VALUES (1833, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"deptIds\":[],\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"4\",\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:45:17');
INSERT INTO `sys_oper_log` VALUES (1834, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2020-07-21 14:45:17\",\"data_scope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"4\",\"remark\":\"\",\"role_id\":6,\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"\",\"update_time\":null}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:45:41');
INSERT INTO `sys_oper_log` VALUES (1835, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2020-07-21 14:45:17\",\"data_scope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"114,1055,1056,1057,1058,1059,1060,1061,3\",\"remark\":\"\",\"role_id\":6,\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"\",\"update_time\":\"2020-07-21 14:45:41\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:46:50');
INSERT INTO `sys_oper_log` VALUES (1836, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2020-07-21 14:45:17\",\"data_scope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"4\",\"remark\":\"\",\"role_id\":6,\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"\",\"update_time\":\"2020-07-21 14:46:50\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:47:00');
INSERT INTO `sys_oper_log` VALUES (1837, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '信息部', '/system/role', '127.0.0.1', '内网IP', '{\"create_by\":\"admin\",\"create_time\":\"2020-07-21 14:45:17\",\"data_scope\":\"1\",\"del_flag\":\"0\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"menuIds\":\"114,1055,1056,1057,1058,1059,1060,1061,3\",\"remark\":\"\",\"role_id\":6,\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"\",\"update_time\":\"2020-07-21 14:47:00\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:52:38');
INSERT INTO `sys_oper_log` VALUES (1838, '角色管理', 3, '/system/role?ids=6', 'DELETE', 1, 'admin', '信息部', '/system/role?ids=6', '127.0.0.1', '内网IP', '{\"ids\":\"6\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2020-07-21 14:52:45');
INSERT INTO `sys_oper_log` VALUES (1839, '用户管理', 2, '/system/user', 'PUT', 1, 'admin', '信息部', '/system/user', '127.0.0.1', '内网IP', '{\"avatar\":\"\",\"create_by\":\"admin\",\"create_time\":\"2020-07-17 14:22:00\",\"del_flag\":\"0\",\"dept\":null,\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"login_date\":null,\"login_ip\":\"\",\"login_name\":\"测试测试2\",\"password\":\"\",\"phonenumber\":\"18983690296\",\"postIds\":\"1\",\"remark\":\"测试\",\"roleIds\":\"\",\"roles\":null,\"salt\":\"jmUs7p\",\"sex\":\"0\",\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2020-07-21 13:34:17\",\"user_id\":6,\"user_name\":\"testtest\",\"user_type\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2020-07-21 15:08:49');
INSERT INTO `sys_oper_log` VALUES (1840, '用户管理', 2, '/system/user', 'PUT', 1, 'admin', '信息部', '/system/user', '127.0.0.1', '内网IP', '{\"avatar\":\"\",\"create_by\":\"admin\",\"create_time\":\"2020-07-17 14:22:00\",\"del_flag\":\"0\",\"dept\":null,\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"jwtLoginName\":\"admin\",\"jwtUid\":\"1\",\"login_date\":null,\"login_ip\":\"\",\"login_name\":\"测试测试2\",\"password\":\"\",\"phonenumber\":\"18983690296\",\"postIds\":\"1\",\"remark\":\"测试\",\"roleIds\":\"3\",\"roles\":null,\"salt\":\"jmUs7p\",\"sex\":\"0\",\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2020-07-21 15:08:49\",\"user_id\":6,\"user_name\":\"testtest\",\"user_type\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2020-07-21 15:19:08');
INSERT INTO `sys_oper_log` VALUES (1841, '', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网1\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 0, '', '2021-01-25 09:08:15');
INSERT INTO `sys_oper_log` VALUES (1842, '', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 0, '', '2021-01-25 09:10:20');
INSERT INTO `sys_oper_log` VALUES (1843, '', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网1\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 0, '', '2021-01-25 09:10:50');
INSERT INTO `sys_oper_log` VALUES (1844, '', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 0, '', '2021-01-25 09:12:04');
INSERT INTO `sys_oper_log` VALUES (1845, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网1\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2021-01-25 09:16:45');
INSERT INTO `sys_oper_log` VALUES (1846, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2021-01-25 09:17:01');
INSERT INTO `sys_oper_log` VALUES (1847, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网1\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2021-01-25 09:18:26');
INSERT INTO `sys_oper_log` VALUES (1848, '菜单管理', 2, '/system/menu', 'PUT', 1, 'admin', '测试部门', '/system/menu', '172.17.0.1', '', '{\"component\":\"\",\"create_by\":\"\",\"create_time\":null,\"icon\":\"guide\",\"is_frame\":\"0\",\"menu_id\":4,\"menu_name\":\"若依官网\",\"menu_type\":\"M\",\"method\":\"\",\"order_num\":4,\"parentName\":\"\",\"parent_id\":0,\"path\":\"http://www.truckgogo.com2\",\"perms\":\"\",\"remark\":\"\",\"status\":\"0\",\"update_by\":\"\",\"update_time\":null,\"url\":\"\",\"visible\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"菜单管理\"}', 0, '', '2021-01-25 09:56:46');
INSERT INTO `sys_oper_log` VALUES (1849, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"data_scope\":\"2\",\"del_flag\":\"0\",\"menuIds\":\"2,109,1047,1048,110,1049,1050,1051,1052,1053,1054,1064,1065,1066,1067,111,112,4\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2020-07-18 11:31:41\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:38:53');
INSERT INTO `sys_oper_log` VALUES (1850, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"1\",\"data_scope\":\"2\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:38:52\"}', '{\"code\":500,\"msg\":\"保存数据失败\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:39:52');
INSERT INTO `sys_oper_log` VALUES (1851, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"1\",\"data_scope\":\"2\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:38:52\"}', '{\"code\":500,\"msg\":\"保存数据失败\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:40:14');
INSERT INTO `sys_oper_log` VALUES (1852, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"1\",\"data_scope\":\"2\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:38:52\"}', '{\"code\":500,\"msg\":\"保存数据失败\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:40:43');
INSERT INTO `sys_oper_log` VALUES (1853, '', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"1\",\"data_scope\":\"2\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:38:52\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":4,\"module\":\"\"}', 0, '', '2021-01-25 13:42:08');
INSERT INTO `sys_oper_log` VALUES (1854, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"4\",\"data_scope\":\"1\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:42:08\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:42:26');
INSERT INTO `sys_oper_log` VALUES (1855, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"5\",\"data_scope\":\"4\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:42:26\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:43:03');
INSERT INTO `sys_oper_log` VALUES (1856, '角色管理', 4, '/system/role/dataScope', 'PUT', 1, 'admin', '测试部门', '/system/role/dataScope', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"5\",\"data_scope\":\"5\",\"del_flag\":\"0\",\"deptIds\":\"100,110\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:43:03\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":4,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:44:18');
INSERT INTO `sys_oper_log` VALUES (1857, '', 2, '/system/role/changeStatus', 'PUT', 1, 'admin', '测试部门', '/system/role/changeStatus', '172.17.0.1', '', '{\"roleId\":1,\"status\":\"1\"}', '{\"code\":500,\"msg\":\"不能停用超级管理员\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 1, '', '2021-01-25 13:48:29');
INSERT INTO `sys_oper_log` VALUES (1858, '', 2, '/system/role/changeStatus', 'PUT', 1, 'admin', '测试部门', '/system/role/changeStatus', '172.17.0.1', '', '{\"roleId\":3,\"status\":\"1\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"\"}', 0, '', '2021-01-25 13:48:30');
INSERT INTO `sys_oper_log` VALUES (1859, '角色管理', 2, '/system/role/changeStatus', 'PUT', 1, 'admin', '测试部门', '/system/role/changeStatus', '172.17.0.1', '', '{\"roleId\":3,\"status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:48:34');
INSERT INTO `sys_oper_log` VALUES (1860, '角色管理', 2, '/system/role', 'PUT', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"create_by\":\"admin\",\"create_time\":\"2020-04-16 14:34:27\",\"dataScope\":\"5\",\"del_flag\":\"0\",\"menuIds\":\"2,109,1047,1048,110,1049,1050,1051,1052,1053,1054,1064,1065,1066,1067,111,112\",\"remark\":\"普通员工\",\"role_id\":3,\"role_key\":\"pop\",\"role_name\":\"员工\",\"role_sort\":0,\"status\":\"0\",\"update_by\":\"admin\",\"update_time\":\"2021-01-25 13:44:18\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:51:36');
INSERT INTO `sys_oper_log` VALUES (1861, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:51:46');
INSERT INTO `sys_oper_log` VALUES (1862, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"role_key\":\"tesftlkls\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:52:10');
INSERT INTO `sys_oper_log` VALUES (1863, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"role_key\":\"tesftlkls\",\"role_name\":\"测试\",\"role_sort\":1,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:52:13');
INSERT INTO `sys_oper_log` VALUES (1864, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"role_key\":\"tesftlkls\",\"role_name\":\"测试\",\"role_sort\":1,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:52:22');
INSERT INTO `sys_oper_log` VALUES (1865, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"role_key\":\"tesftlkls\",\"role_name\":\"测试\",\"role_sort\":1,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:54:11');
INSERT INTO `sys_oper_log` VALUES (1866, '角色管理', 3, '/system/role?ids=4', 'DELETE', 1, 'admin', '测试部门', '/system/role?ids=4', '172.17.0.1', '', '{\"ids\":\"4\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 13:55:23');
INSERT INTO `sys_oper_log` VALUES (1867, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"remark\":\"123\",\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":500,\"msg\":\"添加失败\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 1, '', '2021-01-25 13:55:32');
INSERT INTO `sys_oper_log` VALUES (1868, '角色管理', 1, '/system/role', 'POST', 1, 'admin', '测试部门', '/system/role', '172.17.0.1', '', '{\"deptIds\":[],\"menuIds\":\"4\",\"remark\":\"123\",\"role_key\":\"test\",\"role_name\":\"测试\",\"role_sort\":0,\"status\":\"0\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 14:07:26');
INSERT INTO `sys_oper_log` VALUES (1869, '角色管理', 3, '/system/role?ids=8', 'DELETE', 1, 'admin', '测试部门', '/system/role?ids=8', '172.17.0.1', '', '{\"ids\":\"8\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"角色管理\"}', 0, '', '2021-01-25 14:07:30');
INSERT INTO `sys_oper_log` VALUES (1870, '用户管理', 1, '/system/user', 'POST', 1, 'admin', '测试部门', '/system/user', '172.17.0.1', '', '{\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"login_name\":\"test007\",\"password\":\"123456\",\"phonenumber\":\"18983690292\",\"postIds\":\"1\",\"roleIds\":\"1\",\"sex\":\"0\",\"status\":\"0\",\"user_name\":\"测试\"}', '{\"code\":500,\"msg\":\"Error 1062: Duplicate entry \'2-1\' for key \'PRIMARY\', INSERT INTO `sys_user_role`(`user_id`,`role_id`) VALUES(2,1) \\n\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 1, '', '2021-01-25 14:46:59');
INSERT INTO `sys_oper_log` VALUES (1871, '用户管理', 1, '/system/user', 'POST', 1, 'admin', '测试部门', '/system/user', '172.17.0.1', '', '{\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"login_name\":\"test007\",\"password\":\"123456\",\"phonenumber\":\"18983690292\",\"postIds\":\"1\",\"roleIds\":\"1\",\"sex\":\"0\",\"status\":\"0\",\"user_name\":\"测试\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-01-25 14:48:20');
INSERT INTO `sys_oper_log` VALUES (1872, '用户管理', 1, '/system/user', 'POST', 1, 'admin', '测试部门', '/system/user', '172.17.0.1', '', '{\"dept_id\":110,\"email\":\"xiahaowen126@gmail.com\",\"login_name\":\"test008\",\"password\":\"123456\",\"phonenumber\":\"18983690293\",\"postIds\":\"1\",\"roleIds\":\"1\",\"sex\":\"0\",\"status\":\"0\",\"user_name\":\"测试2\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-01-25 14:50:25');
INSERT INTO `sys_oper_log` VALUES (1873, '用户管理', 2, '/system/user', 'PUT', 1, 'admin', '测试部门', '/system/user', '172.17.0.1', '', '{\"avatar\":\"\",\"create_by\":\"admin\",\"create_time\":\"2021-01-25 14:48:20\",\"del_flag\":\"0\",\"dept\":null,\"dept_id\":110,\"email\":\"xiahaowen127@gmail.com\",\"login_date\":null,\"login_ip\":\"\",\"login_name\":\"test007\",\"password\":\"\",\"phonenumber\":\"18983690292\",\"postIds\":\"1\",\"remark\":\"\",\"roleIds\":\"1\",\"roles\":null,\"salt\":\"IaR6wG\",\"sex\":\"0\",\"status\":\"\",\"update_by\":\"\",\"update_time\":null,\"user_id\":3,\"user_name\":\"1234\",\"user_type\":\"\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-01-25 14:51:18');
INSERT INTO `sys_oper_log` VALUES (1874, '用户管理', 2, '/system/user/changeStatus', 'PUT', 1, 'admin', '测试部门', '/system/user/changeStatus', '172.17.0.1', '', '{\"status\":\"0\",\"userId\":3}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-01-25 14:51:27');
INSERT INTO `sys_oper_log` VALUES (1875, '代码生成管理', 1, '/tool/gen?tables=test_test', 'POST', 1, 'admin', '测试部门', '/tool/gen?tables=test_test', '172.17.0.1', '', '{\"tables\":\"test_test\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 14:08:06');
INSERT INTO `sys_oper_log` VALUES (1876, '代码生成管理', 2, '/tool/gen', 'PUT', 1, 'admin', '测试部门', '/tool/gen', '172.17.0.1', '', '{\"business_name\":\"test\",\"class_name\":\"testest\",\"columns\":\"[{\\\"column_id\\\":436,\\\"table_id\\\":42,\\\"column_name\\\":\\\"id\\\",\\\"column_comment\\\":\\\"ID\\\",\\\"column_type\\\":\\\"int(11)\\\",\\\"go_type\\\":\\\"int64\\\",\\\"go_field\\\":\\\"Id\\\",\\\"html_field\\\":\\\"id\\\",\\\"is_pk\\\":\\\"1\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"0\\\",\\\"is_list\\\":\\\"0\\\",\\\"is_query\\\":\\\"0\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":1,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":437,\\\"table_id\\\":42,\\\"column_name\\\":\\\"test_name\\\",\\\"column_comment\\\":\\\"测试名称\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestName\\\",\\\"html_field\\\":\\\"testName\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"1\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"LIKE\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":2,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":438,\\\"table_id\\\":42,\\\"column_name\\\":\\\"test_phone\\\",\\\"column_comment\\\":\\\"测试手机号\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestPhone\\\",\\\"html_field\\\":\\\"testPhone\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":3,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null}]\",\"create_by\":\"admin\",\"create_time\":\"2021-02-01 14:08:03\",\"function_author\":\"1307\",\"function_name\":\"测试\",\"module_name\":\"system\",\"options\":\"\",\"package_name\":\"gea\",\"params\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"pkColumn\":{\"column_comment\":\"\",\"column_id\":0,\"column_name\":\"\",\"column_type\":\"\",\"create_by\":\"\",\"create_time\":null,\"dict_type\":\"\",\"go_field\":\"\",\"go_type\":\"\",\"html_field\":\"\",\"html_type\":\"\",\"is_edit\":\"\",\"is_increment\":\"\",\"is_insert\":\"\",\"is_list\":\"\",\"is_pk\":\"\",\"is_query\":\"\",\"is_required\":\"\",\"query_type\":\"\",\"sort\":0,\"table_id\":0,\"update_by\":\"\",\"update_time\":null},\"remark\":\"\",\"table_comment\":\"测试表\",\"table_id\":42,\"table_name\":\"test_test\",\"tpl_category\":\"crud\",\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\",\"update_by\":\"\",\"update_time\":null}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 14:09:42');
INSERT INTO `sys_oper_log` VALUES (1877, '代码生成管理', 1, '/tool/gen?tables=t_tq', 'POST', 1, 'admin', '测试部门', '/tool/gen?tables=t_tq', '172.17.0.1', '', '{\"tables\":\"t_tq\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 15:27:59');
INSERT INTO `sys_oper_log` VALUES (1878, '代码生成管理', 2, '/tool/gen', 'PUT', 1, 'admin', '测试部门', '/tool/gen', '172.17.0.1', '', '{\"business_name\":\"tq\",\"class_name\":\"tq\",\"columns\":\"[{\\\"column_id\\\":439,\\\"table_id\\\":43,\\\"column_name\\\":\\\"id\\\",\\\"column_comment\\\":\\\"ID\\\",\\\"column_type\\\":\\\"int(11)\\\",\\\"go_type\\\":\\\"int64\\\",\\\"go_field\\\":\\\"Id\\\",\\\"html_field\\\":\\\"id\\\",\\\"is_pk\\\":\\\"1\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"0\\\",\\\"is_list\\\":\\\"0\\\",\\\"is_query\\\":\\\"0\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":1,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":440,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_name\\\",\\\"column_comment\\\":\\\"测试名称\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestName\\\",\\\"html_field\\\":\\\"testName\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"1\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"LIKE\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":2,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":441,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_phone\\\",\\\"column_comment\\\":\\\"测试手机号\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestPhone\\\",\\\"html_field\\\":\\\"testPhone\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":3,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null}]\",\"create_by\":\"admin\",\"create_time\":\"2021-02-01 15:27:56\",\"function_author\":\"1307\",\"function_name\":\"测试\",\"module_name\":\"system\",\"options\":\"\",\"package_name\":\"gea\",\"params\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"pkColumn\":{\"column_comment\":\"\",\"column_id\":0,\"column_name\":\"\",\"column_type\":\"\",\"create_by\":\"\",\"create_time\":null,\"dict_type\":\"\",\"go_field\":\"\",\"go_type\":\"\",\"html_field\":\"\",\"html_type\":\"\",\"is_edit\":\"\",\"is_increment\":\"\",\"is_insert\":\"\",\"is_list\":\"\",\"is_pk\":\"\",\"is_query\":\"\",\"is_required\":\"\",\"query_type\":\"\",\"sort\":0,\"table_id\":0,\"update_by\":\"\",\"update_time\":null},\"remark\":\"\",\"table_comment\":\"测试表\",\"table_id\":43,\"table_name\":\"t_tq\",\"tpl_category\":\"crud\",\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\",\"update_by\":\"\",\"update_time\":null}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 15:30:36');
INSERT INTO `sys_oper_log` VALUES (1879, '代码生成管理', 2, '/tool/gen', 'PUT', 1, 'admin', '测试部门', '/tool/gen', '172.17.0.1', '', '{\"business_name\":\"tq\",\"class_name\":\"tq\",\"columns\":\"[{\\\"column_id\\\":439,\\\"table_id\\\":43,\\\"column_name\\\":\\\"id\\\",\\\"column_comment\\\":\\\"ID\\\",\\\"column_type\\\":\\\"int(11)\\\",\\\"go_type\\\":\\\"int64\\\",\\\"go_field\\\":\\\"Id\\\",\\\"html_field\\\":\\\"id\\\",\\\"is_pk\\\":\\\"1\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"0\\\",\\\"is_list\\\":\\\"0\\\",\\\"is_query\\\":\\\"0\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":1,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":440,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_name\\\",\\\"column_comment\\\":\\\"测试名称\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestName\\\",\\\"html_field\\\":\\\"testName\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"1\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"LIKE\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":2,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":441,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_phone\\\",\\\"column_comment\\\":\\\"测试手机号\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestPhone\\\",\\\"html_field\\\":\\\"testPhone\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":3,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null}]\",\"create_by\":\"admin\",\"create_time\":\"2021-02-01 15:27:56\",\"function_author\":\"1307\",\"function_name\":\"测试\",\"module_name\":\"admin\",\"options\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"package_name\":\"gea\",\"params\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"pkColumn\":{\"column_comment\":\"\",\"column_id\":0,\"column_name\":\"\",\"column_type\":\"\",\"create_by\":\"\",\"create_time\":null,\"dict_type\":\"\",\"go_field\":\"\",\"go_type\":\"\",\"html_field\":\"\",\"html_type\":\"\",\"is_edit\":\"\",\"is_increment\":\"\",\"is_insert\":\"\",\"is_list\":\"\",\"is_pk\":\"\",\"is_query\":\"\",\"is_required\":\"\",\"query_type\":\"\",\"sort\":0,\"table_id\":0,\"update_by\":\"\",\"update_time\":null},\"remark\":\"\",\"table_comment\":\"测试表\",\"table_id\":43,\"table_name\":\"t_tq\",\"tpl_category\":\"crud\",\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\",\"update_by\":\"admin\",\"update_time\":\"2021-02-01 15:30:36\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 15:41:01');
INSERT INTO `sys_oper_log` VALUES (1880, '代码生成管理', 2, '/tool/gen', 'PUT', 1, 'admin', '测试部门', '/tool/gen', '172.17.0.1', '', '{\"business_name\":\"tq\",\"class_name\":\"tq\",\"columns\":\"[{\\\"column_id\\\":439,\\\"table_id\\\":43,\\\"column_name\\\":\\\"id\\\",\\\"column_comment\\\":\\\"ID\\\",\\\"column_type\\\":\\\"int(11)\\\",\\\"go_type\\\":\\\"int64\\\",\\\"go_field\\\":\\\"Id\\\",\\\"html_field\\\":\\\"id\\\",\\\"is_pk\\\":\\\"1\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"0\\\",\\\"is_list\\\":\\\"0\\\",\\\"is_query\\\":\\\"0\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":1,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":440,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_name\\\",\\\"column_comment\\\":\\\"测试名称\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestName\\\",\\\"html_field\\\":\\\"testName\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"1\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"LIKE\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":2,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null},{\\\"column_id\\\":441,\\\"table_id\\\":43,\\\"column_name\\\":\\\"test_phone\\\",\\\"column_comment\\\":\\\"测试手机号\\\",\\\"column_type\\\":\\\"varchar(255)\\\",\\\"go_type\\\":\\\"string\\\",\\\"go_field\\\":\\\"TestPhone\\\",\\\"html_field\\\":\\\"testPhone\\\",\\\"is_pk\\\":\\\"0\\\",\\\"is_increment\\\":\\\"0\\\",\\\"is_required\\\":\\\"0\\\",\\\"is_insert\\\":\\\"1\\\",\\\"is_edit\\\":\\\"1\\\",\\\"is_list\\\":\\\"1\\\",\\\"is_query\\\":\\\"1\\\",\\\"query_type\\\":\\\"EQ\\\",\\\"html_type\\\":\\\"input\\\",\\\"dict_type\\\":\\\"\\\",\\\"sort\\\":3,\\\"create_by\\\":\\\"admin\\\",\\\"create_time\\\":null,\\\"update_by\\\":\\\"\\\",\\\"update_time\\\":null}]\",\"create_by\":\"admin\",\"create_time\":\"2021-02-01 15:27:56\",\"function_author\":\"1307\",\"function_name\":\"测试1\",\"module_name\":\"admin\",\"options\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"package_name\":\"gea\",\"params\":\"{\\\"treeCode\\\":\\\"\\\",\\\"treeName\\\":\\\"\\\",\\\"treeParentCode\\\":\\\"\\\"}\",\"pkColumn\":{\"column_comment\":\"\",\"column_id\":0,\"column_name\":\"\",\"column_type\":\"\",\"create_by\":\"\",\"create_time\":null,\"dict_type\":\"\",\"go_field\":\"\",\"go_type\":\"\",\"html_field\":\"\",\"html_type\":\"\",\"is_edit\":\"\",\"is_increment\":\"\",\"is_insert\":\"\",\"is_list\":\"\",\"is_pk\":\"\",\"is_query\":\"\",\"is_required\":\"\",\"query_type\":\"\",\"sort\":0,\"table_id\":0,\"update_by\":\"\",\"update_time\":null},\"remark\":\"\",\"table_comment\":\"测试表\",\"table_id\":43,\"table_name\":\"t_tq\",\"tpl_category\":\"crud\",\"treeCode\":\"\",\"treeName\":\"\",\"treeParentCode\":\"\",\"update_by\":\"admin\",\"update_time\":\"2021-02-01 15:40:58\"}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"代码生成管理\"}', 0, '', '2021-02-01 16:33:14');
INSERT INTO `sys_oper_log` VALUES (1881, '用户管理', 2, '/api/system/user/changeStatus', 'PUT', 1, 'admin', '测试部门', '/api/system/user/changeStatus', '172.17.0.1', '', '{\"status\":\"1\",\"userId\":4}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-02-09 10:43:31');
INSERT INTO `sys_oper_log` VALUES (1882, '用户管理', 2, '/api/system/user/changeStatus', 'PUT', 1, 'admin', '测试部门', '/api/system/user/changeStatus', '172.17.0.1', '', '{\"status\":\"0\",\"userId\":4}', '{\"code\":0,\"msg\":\"操作成功\",\"data\":\"\",\"otype\":0,\"module\":\"用户管理\"}', 0, '', '2021-02-09 10:43:39');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, '0', 'admin', '2018-03-16 11:33:00', '', '2020-02-04 19:36:13', '4223434');
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理2', 2, '0', 'admin', '2018-03-16 11:33:00', '', '2020-07-18 11:36:18', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `status` char(1) NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` VALUES (1, '管理员', 'admin', 1, '1', '0', '0', 'admin', '2018-03-16 11:33:00', '', '2020-07-21 12:06:26', '管理员');
INSERT INTO `sys_role` VALUES (3, '员工', 'pop', 0, '5', '0', '0', 'admin', '2020-04-16 14:34:27', 'admin', '2021-01-25 13:51:36', '普通员工');
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` VALUES (3, 100);
INSERT INTO `sys_role_dept` VALUES (3, 110);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` VALUES (1, 1);
INSERT INTO `sys_role_menu` VALUES (1, 2);
INSERT INTO `sys_role_menu` VALUES (1, 3);
INSERT INTO `sys_role_menu` VALUES (1, 4);
INSERT INTO `sys_role_menu` VALUES (1, 100);
INSERT INTO `sys_role_menu` VALUES (1, 101);
INSERT INTO `sys_role_menu` VALUES (1, 102);
INSERT INTO `sys_role_menu` VALUES (1, 103);
INSERT INTO `sys_role_menu` VALUES (1, 104);
INSERT INTO `sys_role_menu` VALUES (1, 105);
INSERT INTO `sys_role_menu` VALUES (1, 106);
INSERT INTO `sys_role_menu` VALUES (1, 108);
INSERT INTO `sys_role_menu` VALUES (1, 109);
INSERT INTO `sys_role_menu` VALUES (1, 110);
INSERT INTO `sys_role_menu` VALUES (1, 111);
INSERT INTO `sys_role_menu` VALUES (1, 112);
INSERT INTO `sys_role_menu` VALUES (1, 113);
INSERT INTO `sys_role_menu` VALUES (1, 114);
INSERT INTO `sys_role_menu` VALUES (1, 115);
INSERT INTO `sys_role_menu` VALUES (1, 500);
INSERT INTO `sys_role_menu` VALUES (1, 501);
INSERT INTO `sys_role_menu` VALUES (1, 1001);
INSERT INTO `sys_role_menu` VALUES (1, 1002);
INSERT INTO `sys_role_menu` VALUES (1, 1003);
INSERT INTO `sys_role_menu` VALUES (1, 1004);
INSERT INTO `sys_role_menu` VALUES (1, 1005);
INSERT INTO `sys_role_menu` VALUES (1, 1006);
INSERT INTO `sys_role_menu` VALUES (1, 1007);
INSERT INTO `sys_role_menu` VALUES (1, 1008);
INSERT INTO `sys_role_menu` VALUES (1, 1009);
INSERT INTO `sys_role_menu` VALUES (1, 1010);
INSERT INTO `sys_role_menu` VALUES (1, 1011);
INSERT INTO `sys_role_menu` VALUES (1, 1012);
INSERT INTO `sys_role_menu` VALUES (1, 1013);
INSERT INTO `sys_role_menu` VALUES (1, 1014);
INSERT INTO `sys_role_menu` VALUES (1, 1015);
INSERT INTO `sys_role_menu` VALUES (1, 1016);
INSERT INTO `sys_role_menu` VALUES (1, 1017);
INSERT INTO `sys_role_menu` VALUES (1, 1018);
INSERT INTO `sys_role_menu` VALUES (1, 1019);
INSERT INTO `sys_role_menu` VALUES (1, 1020);
INSERT INTO `sys_role_menu` VALUES (1, 1021);
INSERT INTO `sys_role_menu` VALUES (1, 1022);
INSERT INTO `sys_role_menu` VALUES (1, 1023);
INSERT INTO `sys_role_menu` VALUES (1, 1024);
INSERT INTO `sys_role_menu` VALUES (1, 1025);
INSERT INTO `sys_role_menu` VALUES (1, 1026);
INSERT INTO `sys_role_menu` VALUES (1, 1027);
INSERT INTO `sys_role_menu` VALUES (1, 1028);
INSERT INTO `sys_role_menu` VALUES (1, 1029);
INSERT INTO `sys_role_menu` VALUES (1, 1030);
INSERT INTO `sys_role_menu` VALUES (1, 1031);
INSERT INTO `sys_role_menu` VALUES (1, 1032);
INSERT INTO `sys_role_menu` VALUES (1, 1033);
INSERT INTO `sys_role_menu` VALUES (1, 1034);
INSERT INTO `sys_role_menu` VALUES (1, 1035);
INSERT INTO `sys_role_menu` VALUES (1, 1040);
INSERT INTO `sys_role_menu` VALUES (1, 1041);
INSERT INTO `sys_role_menu` VALUES (1, 1042);
INSERT INTO `sys_role_menu` VALUES (1, 1043);
INSERT INTO `sys_role_menu` VALUES (1, 1044);
INSERT INTO `sys_role_menu` VALUES (1, 1045);
INSERT INTO `sys_role_menu` VALUES (1, 1047);
INSERT INTO `sys_role_menu` VALUES (1, 1048);
INSERT INTO `sys_role_menu` VALUES (1, 1049);
INSERT INTO `sys_role_menu` VALUES (1, 1050);
INSERT INTO `sys_role_menu` VALUES (1, 1051);
INSERT INTO `sys_role_menu` VALUES (1, 1052);
INSERT INTO `sys_role_menu` VALUES (1, 1053);
INSERT INTO `sys_role_menu` VALUES (1, 1054);
INSERT INTO `sys_role_menu` VALUES (1, 1055);
INSERT INTO `sys_role_menu` VALUES (1, 1056);
INSERT INTO `sys_role_menu` VALUES (1, 1057);
INSERT INTO `sys_role_menu` VALUES (1, 1058);
INSERT INTO `sys_role_menu` VALUES (1, 1059);
INSERT INTO `sys_role_menu` VALUES (1, 1060);
INSERT INTO `sys_role_menu` VALUES (1, 1061);
INSERT INTO `sys_role_menu` VALUES (1, 1064);
INSERT INTO `sys_role_menu` VALUES (1, 1065);
INSERT INTO `sys_role_menu` VALUES (1, 1066);
INSERT INTO `sys_role_menu` VALUES (1, 1067);
INSERT INTO `sys_role_menu` VALUES (1, 1068);
INSERT INTO `sys_role_menu` VALUES (1, 1069);
INSERT INTO `sys_role_menu` VALUES (1, 1070);
INSERT INTO `sys_role_menu` VALUES (1, 1071);
INSERT INTO `sys_role_menu` VALUES (1, 1072);
INSERT INTO `sys_role_menu` VALUES (1, 1073);
INSERT INTO `sys_role_menu` VALUES (1, 1074);
INSERT INTO `sys_role_menu` VALUES (1, 1075);
INSERT INTO `sys_role_menu` VALUES (1, 1076);
INSERT INTO `sys_role_menu` VALUES (3, 2);
INSERT INTO `sys_role_menu` VALUES (3, 109);
INSERT INTO `sys_role_menu` VALUES (3, 110);
INSERT INTO `sys_role_menu` VALUES (3, 111);
INSERT INTO `sys_role_menu` VALUES (3, 112);
INSERT INTO `sys_role_menu` VALUES (3, 1047);
INSERT INTO `sys_role_menu` VALUES (3, 1048);
INSERT INTO `sys_role_menu` VALUES (3, 1049);
INSERT INTO `sys_role_menu` VALUES (3, 1050);
INSERT INTO `sys_role_menu` VALUES (3, 1051);
INSERT INTO `sys_role_menu` VALUES (3, 1052);
INSERT INTO `sys_role_menu` VALUES (3, 1053);
INSERT INTO `sys_role_menu` VALUES (3, 1054);
INSERT INTO `sys_role_menu` VALUES (3, 1064);
INSERT INTO `sys_role_menu` VALUES (3, 1065);
INSERT INTO `sys_role_menu` VALUES (3, 1066);
INSERT INTO `sys_role_menu` VALUES (3, 1067);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `login_name` varchar(30) NOT NULL COMMENT '登录账号',
  `user_name` varchar(30) NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` varchar(11) DEFAULT '' COMMENT '手机号码',
  `sex` char(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) DEFAULT '' COMMENT '头像路径',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  `salt` varchar(20) DEFAULT '' COMMENT '盐加密',
  `status` char(1) DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(50) DEFAULT '' COMMENT '最后登陆IP',
  `login_date` datetime DEFAULT NULL COMMENT '最后登陆时间',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 110, 'admin', '超级管理员', '00', '1307@qq.com', '123', '0', '/upload/avatar/1/c3o6k3bwvquw6qiwrz', '9eb3cbe1c7c81a8d5982f295361ea4b2', 'Sp6Jqx', '0', '0', '127.0.0.1', '2020-01-13 13:20:40', 'admin', '2018-03-16 11:33:00', 'admin', '2020-01-27 08:53:05', '管理员');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online` (
  `token` varchar(255) NOT NULL DEFAULT '' COMMENT '用户会话token',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` varchar(10) DEFAULT '' COMMENT '在线状态on_line在线off_line离线',
  `start_timestamp` datetime DEFAULT NULL COMMENT '创建时间',
  `last_access_time` datetime DEFAULT NULL COMMENT '最后访问时间',
  `expire_time` int(5) DEFAULT '0' COMMENT '超时时间，单位为分钟',
  PRIMARY KEY (`token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='在线用户记录';

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_online` VALUES ('GJWT1', 'admin', '', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', 'on_line', '2020-07-20 09:03:59', '2020-07-20 09:03:59', 3600);
INSERT INTO `sys_user_online` VALUES ('GJWTadmin', 'admin', '', '172.17.0.1', '', 'Chrome', 'Intel Mac OS X 10_15_1', 'on_line', '2021-02-09 10:38:00', '2021-02-09 10:38:00', 864000);
INSERT INTO `sys_user_online` VALUES ('GJWT测试测试2', '测试测试2', '', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_1', 'on_line', '2020-07-21 13:39:44', '2020-07-21 13:39:44', 864000);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (1, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

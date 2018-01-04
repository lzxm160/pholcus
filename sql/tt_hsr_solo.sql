/*
Navicat MySQL Data Transfer

Source Server         : hsr
Source Server Version : 50634
Source Host           : rm-bp1lqwyzn178u1cqao.mysql.rds.aliyuncs.com:3306
Source Database       : wnt_91pool

Target Server Type    : MYSQL
Target Server Version : 50634
File Encoding         : 65001

Date: 2018-01-04 15:57:17
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tt_hsr_solo
-- ----------------------------
DROP TABLE IF EXISTS `tt_hsr_solo`;
CREATE TABLE `tt_hsr_solo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pool` varchar(50) NOT NULL,
  `addr` varchar(50) NOT NULL,
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL,
  `hashrate` int(11) NOT NULL COMMENT 'mb',
  `hit` double NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1655 DEFAULT CHARSET=utf8;

/*
Navicat MySQL Data Transfer

Source Server         : ali.banixc.com_3306
Source Server Version : 50549
Source Host           : ali.banixc.com:3306
Source Database       : paper

Target Server Type    : MYSQL
Target Server Version : 50549
File Encoding         : 65001

Date: 2016-09-14 02:51:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for date
-- ----------------------------
DROP TABLE IF EXISTS `date`;
CREATE TABLE `date` (
  `id` int(11) NOT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of date
-- ----------------------------
INSERT INTO `date` VALUES ('0', '2016-09-11');
INSERT INTO `date` VALUES ('1', '2016-09-28');

-- ----------------------------
-- Table structure for document
-- ----------------------------
DROP TABLE IF EXISTS `document`;
CREATE TABLE `document` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `path` varchar(255) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of document
-- ----------------------------
INSERT INTO `document` VALUES ('43', 'D:DocumentsREADME1.pdf', '/static/upload/2016/09/', '2016-09-14 02:38:33');
INSERT INTO `document` VALUES ('44', 'README2.doc', '/static/upload/2016/09/', '2016-09-14 02:39:40');
INSERT INTO `document` VALUES ('45', 'README3.doc', '/static/upload/2016/09/', '2016-09-14 02:40:13');
INSERT INTO `document` VALUES ('46', 'Paper.pdf', '/static/upload/2016/09/', '2016-09-14 02:40:29');
INSERT INTO `document` VALUES ('47', 'Paper-T.pdf', '/static/upload/2016/09/', '2016-09-14 02:40:59');

-- ----------------------------
-- Table structure for major
-- ----------------------------
DROP TABLE IF EXISTS `major`;
CREATE TABLE `major` (
  `id` int(11) NOT NULL,
  `major_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of major
-- ----------------------------
INSERT INTO `major` VALUES ('1', '空气动力学');
INSERT INTO `major` VALUES ('2', '化学材料');
INSERT INTO `major` VALUES ('3', '软件架构');

-- ----------------------------
-- Table structure for selected
-- ----------------------------
DROP TABLE IF EXISTS `selected`;
CREATE TABLE `selected` (
  `user_id` varchar(10) NOT NULL,
  `subject_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`subject_id`),
  KEY `subject_id` (`subject_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `selected_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `selected_ibfk_2` FOREIGN KEY (`subject_id`) REFERENCES `subject` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of selected
-- ----------------------------
INSERT INTO `selected` VALUES ('2012303001', '22');
INSERT INTO `selected` VALUES ('2012303004', '22');

-- ----------------------------
-- Table structure for status
-- ----------------------------
DROP TABLE IF EXISTS `status`;
CREATE TABLE `status` (
  `id` int(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of status
-- ----------------------------
INSERT INTO `status` VALUES ('10', '已发布 待初审');
INSERT INTO `status` VALUES ('11', '初审通过 待学生选择');
INSERT INTO `status` VALUES ('12', '一审未通过 请修改');
INSERT INTO `status` VALUES ('15', '已有学生选择 待锁定');
INSERT INTO `status` VALUES ('16', '已锁定 等待上传计划书');
INSERT INTO `status` VALUES ('20', '计划书已上传 待审核');
INSERT INTO `status` VALUES ('21', '计划书通过 待上传开题报告');
INSERT INTO `status` VALUES ('22', '计划书审核未通过 待修改');
INSERT INTO `status` VALUES ('30', '开题报告已上传 待审核');
INSERT INTO `status` VALUES ('31', '开题报告通过 待上传中期报告');
INSERT INTO `status` VALUES ('32', '开题报告未通过 待修改');
INSERT INTO `status` VALUES ('40', '中期报告已上传 待审核');
INSERT INTO `status` VALUES ('41', '中期报告已通过 待上传论文');
INSERT INTO `status` VALUES ('42', '中期报告未通过 待修改');
INSERT INTO `status` VALUES ('50', '论文已上传 待修改');
INSERT INTO `status` VALUES ('51', '论文已通过 待上传论文译文');
INSERT INTO `status` VALUES ('52', '论文未通过 待审核');
INSERT INTO `status` VALUES ('60', '论文译文已上传 待审核');
INSERT INTO `status` VALUES ('61', '论文译文已通过 待答辩');
INSERT INTO `status` VALUES ('62', '论文译文未通过 待修改');
INSERT INTO `status` VALUES ('70', '已答辩');
INSERT INTO `status` VALUES ('71', '答辩未通过');

-- ----------------------------
-- Table structure for subject
-- ----------------------------
DROP TABLE IF EXISTS `subject`;
CREATE TABLE `subject` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `presentation` varchar(255) NOT NULL,
  `sender_id` char(10) DEFAULT NULL,
  `status_id` int(11) NOT NULL DEFAULT '10',
  `task_id` int(11) DEFAULT NULL,
  `primary_id` int(11) DEFAULT NULL,
  `interim_id` int(11) DEFAULT NULL,
  `paper_id` int(11) DEFAULT NULL,
  `translate_id` int(11) DEFAULT NULL,
  `student_id` char(10) DEFAULT NULL,
  `paper_grade` decimal(11,2) DEFAULT NULL,
  `defence_grade` decimal(11,2) DEFAULT NULL,
  `comment` varchar(255) DEFAULT NULL,
  `time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `subject_ibfk_6` (`translate_id`),
  KEY `paper_id` (`paper_id`),
  KEY `primary_id` (`primary_id`),
  KEY `task_id` (`task_id`),
  KEY `interim_id` (`interim_id`),
  KEY `sender_id` (`sender_id`),
  KEY `subject_ibfk_7` (`status_id`),
  KEY `user_id` (`student_id`) USING BTREE,
  CONSTRAINT `subject_ibfk_1` FOREIGN KEY (`paper_id`) REFERENCES `document` (`id`),
  CONSTRAINT `subject_ibfk_2` FOREIGN KEY (`primary_id`) REFERENCES `document` (`id`),
  CONSTRAINT `subject_ibfk_3` FOREIGN KEY (`task_id`) REFERENCES `document` (`id`),
  CONSTRAINT `subject_ibfk_4` FOREIGN KEY (`interim_id`) REFERENCES `document` (`id`),
  CONSTRAINT `subject_ibfk_6` FOREIGN KEY (`translate_id`) REFERENCES `document` (`id`),
  CONSTRAINT `subject_ibfk_7` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`),
  CONSTRAINT `subject_ibfk_8` FOREIGN KEY (`sender_id`) REFERENCES `user` (`user_id`),
  CONSTRAINT `subject_ibfk_9` FOREIGN KEY (`student_id`) REFERENCES `user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of subject
-- ----------------------------
INSERT INTO `subject` VALUES ('22', '关于飞机机翼的研究', '这是一个关于飞机机翼的研究。', 'J201210001', '15', null, null, null, null, null, null, '0.00', '0.00', '', '2016-09-14 02:29:39');
INSERT INTO `subject` VALUES ('23', '1,4—环己烷二甲醇改性聚对苯二甲酸乙二酯结晶行为的研究', '合成1,4-环己烷二甲醇(CHDM)共聚改性的聚对苯二甲酸 乙二酯(PET),用示差扫描量热(DSC)和傅立叶变换红外(FT-IR)方法研究了该体系的等温结晶动力学和结晶度,结果表明,共聚物组成和结晶温度 均会影响材料的结晶行为,这将有利于控制和改善PET的加工性.', 'J201210002', '70', '43', '44', '45', '46', '47', '2012303002', '87.50', '88.60', '', '2016-09-14 02:41:44');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` char(10) NOT NULL,
  `password` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `email` varchar(255) DEFAULT '',
  `gender` tinyint(4) NOT NULL COMMENT '0为男 非0为1',
  `phone` varchar(40) DEFAULT NULL,
  `type` int(11) NOT NULL DEFAULT '1',
  `major_id` int(11) DEFAULT NULL,
  `max_student` int(11) DEFAULT '0',
  `now_student` int(11) DEFAULT '0',
  PRIMARY KEY (`user_id`),
  KEY `major_id` (`major_id`),
  KEY `type` (`type`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`major_id`) REFERENCES `major` (`id`),
  CONSTRAINT `user_ibfk_2` FOREIGN KEY (`type`) REFERENCES `user_type` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('2012303001', '123456', '学生A', 'zhangsan@exp.com', '0', '123456789', '1', '1', '0', '0');
INSERT INTO `user` VALUES ('2012303002', '123456', '学生B', '', '0', null, '1', '2', '0', '0');
INSERT INTO `user` VALUES ('2012303003', '123456', '学生C', '', '0', null, '1', '1', '0', '0');
INSERT INTO `user` VALUES ('2012303004', '123456', '学生D', '', '0', null, '1', '1', '0', '0');
INSERT INTO `user` VALUES ('2012303005', '123456', '学生E', '', '0', null, '1', '1', '0', '0');
INSERT INTO `user` VALUES ('J201210001', 'aaaaaa', '张老师', '', '0', null, '2', '1', '0', '0');
INSERT INTO `user` VALUES ('J201210002', 'bbbbbb', '王老师', '', '0', null, '2', '2', '0', '0');
INSERT INTO `user` VALUES ('J201210003', 'cccccc', '李主任', '', '0', null, '3', '2', '0', '0');
INSERT INTO `user` VALUES ('J201210004', 'dddddd', '马主任', '', '0', null, '3', '1', '0', '0');
INSERT INTO `user` VALUES ('J201210005', 'eeeeee', '方主任', '', '0', null, '5', null, '0', '0');
INSERT INTO `user` VALUES ('J201210006', 'ffffff', '钟秘书', '', '0', null, '4', null, '0', '0');

-- ----------------------------
-- Table structure for user_type
-- ----------------------------
DROP TABLE IF EXISTS `user_type`;
CREATE TABLE `user_type` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_type
-- ----------------------------
INSERT INTO `user_type` VALUES ('1', 'Stuent');
INSERT INTO `user_type` VALUES ('2', 'Teacher');
INSERT INTO `user_type` VALUES ('3', 'Leader');
INSERT INTO `user_type` VALUES ('4', 'Sercetary');
INSERT INTO `user_type` VALUES ('5', 'Admin');

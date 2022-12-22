/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80018
 Source Host           : localhost:3306
 Source Schema         : xbd_iptv

 Target Server Type    : MySQL
 Target Server Version : 80018
 File Encoding         : 65001

 Date: 19/01/2022 23:02:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `id` int(11) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字典名称',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------

-- ----------------------------
-- Table structure for tv_channel
-- ----------------------------
DROP TABLE IF EXISTS `tv_channel`;
CREATE TABLE `tv_channel`  (
  `id` bigint(11) NOT NULL,
  `tv_id` bigint(64) NOT NULL COMMENT '节目Id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '频道',
  `plot` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '情节',
  `outline` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '简介',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_tv_id`(`tv_id`) USING BTREE,
  CONSTRAINT `fk_tv_id` FOREIGN KEY (`tv_id`) REFERENCES `tv_show` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tv_channel
-- ----------------------------

-- ----------------------------
-- Table structure for tv_route
-- ----------------------------
DROP TABLE IF EXISTS `tv_route`;
CREATE TABLE `tv_route`  (
  `id` bigint(11) NOT NULL,
  `channel_id` bigint(11) NULL DEFAULT NULL,
  `uri` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `grade` int(255) NULL DEFAULT NULL COMMENT '线路品质',
  `dpi` int(255) NULL DEFAULT NULL COMMENT '分辨率',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `kf_channel_id`(`channel_id`) USING BTREE,
  CONSTRAINT `kf_channel_id` FOREIGN KEY (`channel_id`) REFERENCES `tv_channel` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tv_route
-- ----------------------------

-- ----------------------------
-- Table structure for tv_show
-- ----------------------------
DROP TABLE IF EXISTS `tv_show`;
CREATE TABLE `tv_show`  (
  `id` bigint(64) NOT NULL,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分类：TVLIVE:MOVE:TVSHOW:SEASON',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `poster` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '海报',
  `fanart` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '横板海报',
  `country` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '国家',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '省份',
  `plot` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '剧情',
  `outline` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '简介',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tv_show
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;

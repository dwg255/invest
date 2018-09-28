/*
 Navicat Premium Data Transfer

 Source Server         : baidu
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 182.61.24.31:3306
 Source Schema         : game

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 23/09/2018 17:49:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for game_invest_base
-- ----------------------------
DROP TABLE IF EXISTS `game_invest_base`;
CREATE TABLE `game_invest_base`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_times_id` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `periods` int(11) NOT NULL,
  `game_pool` int(11) NOT NULL COMMENT '奖池',
  `stake_detail` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `game_result` int(11) NOT NULL COMMENT '结果',
  `start_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 164 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '投资大亨押注表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for game_invest_user_stake
-- ----------------------------
DROP TABLE IF EXISTS `game_invest_user_stake`;
CREATE TABLE `game_invest_user_stake`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `game_times_id` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一标识',
  `periods` int(11) NOT NULL COMMENT '期数',
  `room_id` int(11) NOT NULL COMMENT '房间id',
  `room_type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '房间类型 1 普通场 2 搞几场',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_all_stake` int(11) NOT NULL COMMENT '本场用户总押注',
  `get_gold` int(11) NOT NULL COMMENT '赢得金币',
  `stake_detail` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '押注详情',
  `game_result` int(11) NOT NULL COMMENT '本场开奖',
  `game_pool` int(11) NOT NULL COMMENT '奖池',
  `last_stake_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '投资大亨' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

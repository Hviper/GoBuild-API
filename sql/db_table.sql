/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 18/09/2021 17:06:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for db_table
-- ----------------------------
DROP TABLE IF EXISTS `db_table`;
CREATE TABLE `db_table`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` char(18) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of db_table
-- ----------------------------
INSERT INTO `db_table` VALUES (1, 'dillonl', '14789');
INSERT INTO `db_table` VALUES (2, 'asdjsavjla', 'Y7Kuj8MUk3ifTPs');
INSERT INTO `db_table` VALUES (3, 'Leader', '6xQhZZtZWuYpAzY');
INSERT INTO `db_table` VALUES (4, 'adasdasda', 'dasdasdsadsad');
INSERT INTO `db_table` VALUES (5, 'admin520', 'admin520');
INSERT INTO `db_table` VALUES (6, 'undefineds', 'ssssssss');
INSERT INTO `db_table` VALUES (7, '123456789', '987654321');
INSERT INTO `db_table` VALUES (8, '135135135', 'asdfgh');
INSERT INTO `db_table` VALUES (9, 'admin', 'admin');
INSERT INTO `db_table` VALUES (10, '爱迪生', '123456789');

SET FOREIGN_KEY_CHECKS = 1;

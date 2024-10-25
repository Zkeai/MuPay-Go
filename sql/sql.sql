/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80300 (8.3.0)
 Source Host           : localhost:3306
 Source Schema         : yuka

 Target Server Type    : MySQL
 Target Server Version : 80300 (8.3.0)
 File Encoding         : 65001

 Date: 10/09/2024 13:58:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for yu_business
-- ----------------------------
DROP TABLE IF EXISTS `yu_business`;
CREATE TABLE `yu_business` (
                               `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                               `user_id` int unsigned NOT NULL COMMENT '用户id',
                               `shop_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '店铺名称',
                               `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器标题',
                               `notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '店铺公告',
                               `service_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '网页客服链接',
                               `subdomain` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '子域名',
                               `topdomain` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '顶级域名',
                               `master_display` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '主站显示：0=否，1=是',
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of yu_business
-- ----------------------------
BEGIN;
INSERT INTO `yu_business` (`id`, `user_id`, `shop_name`, `title`, `notice`, `service_url`, `subdomain`, `topdomain`, `master_display`, `create_time`) VALUES (2, 1001, '木鱼店铺', '木鱼店铺', '本程序为开源程序，使用者造成的一切法律后果与作者无关。', NULL, NULL, 'localhost:2900', 0, '2024-09-05 11:13:24');
COMMIT;

-- ----------------------------
-- Table structure for yu_category
-- ----------------------------
DROP TABLE IF EXISTS `yu_category`;
CREATE TABLE `yu_category` (
                               `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                               `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品分类名称',
                               `sort` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `owner` int unsigned NOT NULL DEFAULT '0' COMMENT '所属会员：0=系统，其他等于会员UID',
                               `icon` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '分类图标',
                               `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态：0=停用，1=启用',
                               `hide` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '隐藏：1=隐藏，0=不隐藏',
                               `user_level_config` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '会员配置',
                               PRIMARY KEY (`id`,`hide`),
                               UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of yu_category
-- ----------------------------
BEGIN;
INSERT INTO `yu_category` (`id`, `name`, `sort`, `create_time`, `owner`, `icon`, `status`, `hide`, `user_level_config`) VALUES (3, 'DEMO', 0, '2024-09-05 11:14:46', 1001, 'icon-shangpinguanli', 0, 0, NULL);
INSERT INTO `yu_category` (`id`, `name`, `sort`, `create_time`, `owner`, `icon`, `status`, `hide`, `user_level_config`) VALUES (4, 'DEMO1', 0, '2024-09-05 11:14:51', 1001, 'icon-shangpinguanli', 0, 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for yu_commodity
-- ----------------------------
DROP TABLE IF EXISTS `yu_commodity`;
CREATE TABLE `yu_commodity` (
                                `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                `category_id` int unsigned NOT NULL COMMENT '商品分类ID',
                                `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品名称',
                                `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '商品说明',
                                `cover` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '商品封面图片',
                                `factory_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '成本单价',
                                `price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品单价(未登录)',
                                `user_price` decimal(10,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '商品单价(会员价)',
                                `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '状态：0=下架，1=上架',
                                `owner` int unsigned NOT NULL DEFAULT '0' COMMENT '所属会员：0=系统，其他等于会员UID',
                                `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `api_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'API对接：0=关闭，1=启用',
                                `code` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品代码(API对接)',
                                `delivery_way` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '发货方式：0=自动发货，1=手动发货/插件发货',
                                `delivery_auto_mode` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '自动发卡模式：0=旧卡先发，1=随机发卡，2=新卡先发',
                                `delivery_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手动发货显示信息',
                                `contact_type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '联系方式：0=任意，1=手机，2=邮箱，3=QQ',
                                `password_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '订单密码：0=关闭，1=启用',
                                `sort` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
                                `coupon` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '优惠卷：0=关闭，1=启用',
                                `shared_id` int unsigned DEFAULT NULL COMMENT '共享平台ID',
                                `shared_code` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '共享平台-商品代码',
                                `shared_premium` float(10,2) unsigned DEFAULT '0.00' COMMENT '商品加价',
                                `shared_premium_type` tinyint unsigned DEFAULT '0' COMMENT '加价模式',
                                `seckill_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '商品秒杀：0=关闭，1=开启',
                                `seckill_start_time` datetime DEFAULT NULL COMMENT '秒杀开始时间',
                                `seckill_end_time` datetime DEFAULT NULL COMMENT '秒杀结束时间',
                                `draft_status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '指定卡密购买：0=关闭，1=启用',
                                `draft_premium` decimal(10,2) unsigned DEFAULT '0.00' COMMENT '指定卡密购买时溢价',
                                `inventory_hidden` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '隐藏库存：0=否，1=是',
                                `leave_message` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '发货留言',
                                `recommend` tinyint unsigned DEFAULT '0' COMMENT '推荐商品：0=否，1=是',
                                `send_email` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '发送邮件：0=否，1=是',
                                `only_user` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '限制登录购买：0=否，1=是',
                                `purchase_count` int unsigned NOT NULL DEFAULT '0' COMMENT '限制购买数量：0=无限制',
                                `widget` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '控件',
                                `level_price` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci COMMENT '会员等级-定制价格',
                                `level_disable` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '禁用会员等级折扣，0=关闭，1=启用',
                                `minimum` int unsigned NOT NULL DEFAULT '0' COMMENT '最低购买数量，0=无限制',
                                `maximum` int unsigned NOT NULL DEFAULT '0' COMMENT '最大购买数量，0=无限制',
                                `shared_sync` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '同步平台价格：0=关，1=开',
                                `config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '配置文件',
                                `hide` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '隐藏：1=隐藏，0=不隐藏',
                                `inventory_sync` tinyint NOT NULL DEFAULT '0' COMMENT '同步库存数量: 0=关，1=开',
                                PRIMARY KEY (`id`,`name`) USING BTREE,
                                UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of yu_commodity
-- ----------------------------
BEGIN;
INSERT INTO `yu_commodity` (`id`, `category_id`, `name`, `description`, `cover`, `factory_price`, `price`, `user_price`, `status`, `owner`, `create_time`, `api_status`, `code`, `delivery_way`, `delivery_auto_mode`, `delivery_message`, `contact_type`, `password_status`, `sort`, `coupon`, `shared_id`, `shared_code`, `shared_premium`, `shared_premium_type`, `seckill_status`, `seckill_start_time`, `seckill_end_time`, `draft_status`, `draft_premium`, `inventory_hidden`, `leave_message`, `recommend`, `send_email`, `only_user`, `purchase_count`, `widget`, `level_price`, `level_disable`, `minimum`, `maximum`, `shared_sync`, `config`, `hide`, `inventory_sync`) VALUES (1, 3, '商品名称', '商品描述', '商品封面链接', 0.00, 0.00, 1.00, 0, 0, '2024-09-06 23:06:13', 0, 'bb6ac94671284d09b33586f52d02d5c5', 0, 0, '', 0, 0, 0, 0, 0, NULL, 0.00, 0, 0, NULL, NULL, 0, NULL, 0, NULL, 0, 0, 0, 0, NULL, NULL, 0, 0, 0, 0, NULL, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for yu_pay
-- ----------------------------
DROP TABLE IF EXISTS `yu_pay`;
CREATE TABLE `yu_pay` (
                          `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
                          `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
                          `code` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付代码',
                          `commodity` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '前台状态',
                          `recharge` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '充值状态',
                          `handle` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '支付平台',
                          `sort` smallint unsigned NOT NULL DEFAULT '0' COMMENT '排序',
                          `equipment` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '设备：0=通用 1=手机 2=电脑',
                          `cost` decimal(10,3) unsigned DEFAULT '0.000' COMMENT '手续费',
                          `cost_type` tinyint DEFAULT NULL COMMENT '手续费模式：0=单笔固定，1=百分比',
                          `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `unique_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of yu_pay
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for yu_user
-- ----------------------------
DROP TABLE IF EXISTS `yu_user`;
CREATE TABLE `yu_user` (
                           `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
                           `wallet` varchar(42) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'eth 钱包地址',
                           `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '账号状态',
                           `type` int NOT NULL DEFAULT '0' COMMENT '账号类型 0-用户 1-商户 2-管理员',
                           `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
                           `login_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '登录ip',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `unique_wallet` (`wallet`)
) ENGINE=InnoDB AUTO_INCREMENT=1003 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';

-- ----------------------------
-- Records of yu_user
-- ----------------------------
BEGIN;
INSERT INTO `yu_user` (`id`, `wallet`, `status`, `type`, `create_time`, `login_time`, `login_ip`) VALUES (1002, '0xf5f77d4e368f6ee1115515db3a537bbf98888888', 1, 0, '2024-09-05 11:13:24', NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

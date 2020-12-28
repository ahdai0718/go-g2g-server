-- MySQL dump 10.13  Distrib 8.0.20, for macos10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: dev_g2g_server_base
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `announcement_setting`
--

DROP TABLE IF EXISTS `announcement_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `announcement_setting` (
  `sn` int NOT NULL AUTO_INCREMENT,
  `game_type` int NOT NULL,
  `announcement_title` varchar(40) NOT NULL,
  `announcement_status` tinyint NOT NULL,
  `news_ticker_status` tinyint NOT NULL,
  `announcement_content` varchar(40) NOT NULL,
  `created_at` bigint DEFAULT NULL,
  PRIMARY KEY (`sn`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `announcement_setting`
--

LOCK TABLES `announcement_setting` WRITE;
/*!40000 ALTER TABLE `announcement_setting` DISABLE KEYS */;
INSERT INTO `announcement_setting` VALUES (1,1001,'test',1,1,'testtest',1591088489),(2,1001,'test',0,0,'testtesttesttest',1591088666),(3,1,'test',1,1,'testtest',1591090546),(6,1,'test',0,1,'testtesttest',1591090804);
/*!40000 ALTER TABLE `announcement_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `currency`
--

DROP TABLE IF EXISTS `currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `currency` (
  `name` varchar(16) NOT NULL,
  `multiply_to_credit` double unsigned DEFAULT '1',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `currency`
--

LOCK TABLES `currency` WRITE;
/*!40000 ALTER TABLE `currency` DISABLE KEYS */;
INSERT INTO `currency` VALUES ('TWD',10000);
/*!40000 ALTER TABLE `currency` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `game_server_info`
--

DROP TABLE IF EXISTS `game_server_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game_server_info` (
  `sn` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `host` varchar(128) NOT NULL DEFAULT '0.0.0.0',
  `port` int NOT NULL DEFAULT '0',
  `protocol` varchar(8) NOT NULL DEFAULT 'http',
  PRIMARY KEY (`sn`),
  UNIQUE KEY `sn_UNIQUE` (`sn`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game_server_info`
--

LOCK TABLES `game_server_info` WRITE;
/*!40000 ALTER TABLE `game_server_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `game_server_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `platform_provider`
--

DROP TABLE IF EXISTS `platform_provider`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `platform_provider` (
  `name` varchar(32) NOT NULL,
  `factory_name` varchar(32) NOT NULL DEFAULT 'default',
  `aes_key` varchar(64) NOT NULL,
  `aes_iv` varchar(64) NOT NULL,
  `auth_type` int NOT NULL DEFAULT '1' COMMENT '1: JWT 2: OAuth',
  `auth_id` varchar(64) NOT NULL,
  `auth_secret` varchar(64) NOT NULL,
  `auth_grant_type` varchar(256) NOT NULL DEFAULT 'client_credentials',
  `auth_scope` varchar(256) NOT NULL DEFAULT 'bets',
  `api_url_base` varchar(256) NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `platform_provider`
--

LOCK TABLES `platform_provider` WRITE;
/*!40000 ALTER TABLE `platform_provider` DISABLE KEYS */;
INSERT INTO `platform_provider` VALUES ('CYKJ','default','12345678','12345678',2,'1','3E8ADFA056924F3AB4AF3CFFF3E3C27D','client_credentials','bets','http://localhost/api'),('default','default','12345678','12345678',1,'default','12345678','client_credentials','bets','http://0.0.0.0:3000/api');
/*!40000 ALTER TABLE `platform_provider` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `server_info`
--

DROP TABLE IF EXISTS `server_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `server_info` (
  `name` varchar(32) NOT NULL,
  `group` varchar(32) NOT NULL DEFAULT 'group001',
  `type` int NOT NULL DEFAULT '1',
  `host` varchar(128) NOT NULL DEFAULT '0.0.0.0',
  `port` int NOT NULL DEFAULT '0',
  `protocol` varchar(8) NOT NULL DEFAULT 'ws',
  `server_info_route_path` varchar(32) NOT NULL DEFAULT '/server/info',
  `websocket_protocol` varchar(32) NOT NULL DEFAULT 'ws',
  `websocket_route_path` varchar(32) NOT NULL DEFAULT '/server/ws',
  `host_for_client` varchar(128) NOT NULL DEFAULT '0.0.0.0',
  `port_for_client` int NOT NULL DEFAULT '0',
  `websocket_protocol_for_client` varchar(32) NOT NULL DEFAULT 'ws',
  `websocket_route_path_for_client` varchar(32) NOT NULL DEFAULT 'ws',
  `public_ip_address` varchar(16) NOT NULL DEFAULT '0.0.0.0',
  `is_offline` tinyint NOT NULL DEFAULT '0',
  `game_type` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `server_info`
--

LOCK TABLES `server_info` WRITE;
/*!40000 ALTER TABLE `server_info` DISABLE KEYS */;
INSERT INTO `server_info` VALUES ('fluentd001','',3,'0.0.0.0',8888,'http','/server/info','ws','/server/ws','0.0.0.0',0,'ws','ws','0.0.0.0',0,0),('GameExample','',2,'0.0.0.0',3011,'http','/server/info','ws','/server/ws','0.0.0.0',3011,'ws','ws','0.0.0.0',0,1001),('GameExample2','',2,'0.0.0.0',3021,'http','/server/info','ws','/server/ws','0.0.0.0',3021,'ws','ws','0.0.0.0',0,1002),('GameExample','',2,'0.0.0.0',3001,'http','/server/info','ws','/server/ws','0.0.0.0',3001,'ws','ws','0.0.0.0',0,1),('Gateway001','',1,'0.0.0.0',3000,'http','/server/info','ws','/server/ws','0.0.0.0',3000,'ws','ws','0.0.0.0',0,0),('Offline00101','groupOffline',2,'0.0.0.0',0,'http','/server/info','ws','/server/ws','0.0.0.0',0,'ws','ws','0.0.0.0',1,0);
/*!40000 ALTER TABLE `server_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_log`
--

DROP TABLE IF EXISTS `transaction_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction_log` (
  `sn` int unsigned NOT NULL AUTO_INCREMENT,
  `type` int unsigned DEFAULT NULL,
  `player_id` varchar(64) NOT NULL,
  `game_type` varchar(64) NOT NULL,
  `game_id` varchar(64) NOT NULL,
  `round_id` varchar(64) NOT NULL,
  `timestamp` bigint unsigned DEFAULT NULL,
  `id_lock` varchar(64) DEFAULT NULL,
  `id_cancel_lock` varchar(64) DEFAULT NULL,
  `id_bet` varchar(64) DEFAULT NULL,
  `id_cancel_bet` varchar(64) DEFAULT NULL,
  `id_settle` varchar(64) DEFAULT NULL,
  `id_unlock` varchar(64) DEFAULT NULL,
  `amount` decimal(10,6) unsigned DEFAULT NULL,
  `credit` bigint unsigned DEFAULT NULL,
  `jackpot` bigint unsigned DEFAULT NULL,
  `currency` decimal(10,6) unsigned DEFAULT NULL,
  `currency_code` varchar(8) DEFAULT NULL,
  `system_fee` bigint unsigned DEFAULT NULL,
  `system_fee_jackpot` bigint unsigned DEFAULT NULL,
  `percentage_win_to_system_fee` int unsigned DEFAULT NULL,
  `percentage_system_fee_to_jackpot` int unsigned DEFAULT NULL,
  `percentage_win_jackpot` int unsigned DEFAULT NULL,
  `platform` varchar(64) DEFAULT NULL,
  `platform_player_id` varchar(64) DEFAULT NULL,
  `platform_player_name` varchar(64) DEFAULT NULL,
  `platform_player_display_name` varchar(64) DEFAULT NULL,
  `platform_player_balance` decimal(16,6) unsigned DEFAULT NULL,
  `raw_data` text,
  PRIMARY KEY (`sn`),
  KEY `player_id` (`player_id`),
  KEY `game_type` (`game_type`),
  KEY `game_id` (`game_id`),
  KEY `round_id` (`round_id`),
  KEY `platform` (`platform`),
  KEY `platform_player_id` (`platform_player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_log`
--

LOCK TABLES `transaction_log` WRITE;
/*!40000 ALTER TABLE `transaction_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'dev_g2g_server_base'
--

--
-- Dumping routines for database 'dev_g2g_server_base'
--
/*!50003 DROP PROCEDURE IF EXISTS `Currency_SP_GetAll` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `Currency_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `currency`;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `Currency_SP_GetByName` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `Currency_SP_GetByName`(
	IN `_name` VARCHAR(16)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `currency` WHERE `name`=_name;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `PlatformProvider_SP_GetAll` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `PlatformProvider_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `platform_provider`;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `PlatformProvider_SP_GetByName` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `PlatformProvider_SP_GetByName`(
	IN `_name` VARCHAR(64)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `platform_provider` WHERE `name`=_name;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `ServerInfo_SP_GetAll` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `ServerInfo_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `server_info`;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `ServerInfo_SP_GetByName` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `ServerInfo_SP_GetByName`(
	IN `_name` VARCHAR(64)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `server_info` WHERE `name`=_name;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `Table_SP_CreateAllMonth` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `Table_SP_CreateAllMonth`()
    SQL SECURITY INVOKER
BEGIN

	DECLARE now DATE;
    DECLARE current_day INT(10);
    DECLARE current_month INT(10);
    DECLARE current_year INT(10);
    DECLARE last_day INT(10);
    DECLARE `current_date` DATE;
    DECLARE ymd VARCHAR(16);
    DECLARE next_month DATE;

	SET now = NOW();
	SET current_day = 1;
	SET current_month = MONTH(now);
	SET current_year = YEAR(now);
	SET last_day = DAY(LAST_DAY(now));

	WHILE current_day <= last_day  DO
		SET `current_date` = CONCAT(current_year, '-', current_month, '-', current_day);
		SELECT DATE_FORMAT(DATE(`current_date`), "%Y%m%d") INTO ymd;
		CALL `Table_SP_CreateTransactionLog`(ymd);
		SET current_day = current_day + 1;
	END WHILE;

	SET next_month = DATE_ADD(DATE(now), INTERVAL 1 MONTH);
	SET current_day = 1;
    SET current_month = MONTH(next_month);
    SET current_year = YEAR(next_month);
    SET last_day = DAY(LAST_DAY(next_month));

	WHILE current_day <= last_day  DO
		SET `current_date` = CONCAT(current_year, '-', current_month, '-', current_day);
		SELECT DATE_FORMAT(DATE(`current_date`), "%Y%m%d") INTO ymd;
        CALL `Table_SP_CreateTransactionLog`(ymd);
		SET current_day = current_day + 1;
	END WHILE;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `Table_SP_CreateTransactionLog` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `Table_SP_CreateTransactionLog`(
	IN _ymd VARCHAR(11)
)
    SQL SECURITY INVOKER
BEGIN

	DECLARE `table_name` VARCHAR(64);

    SET `table_name` = CONCAT('`transaction_log_', _ymd, '`');
    SET @sql_statement = CONCAT('CREATE TABLE IF NOT EXISTS ', `table_name`, ' LIKE `transaction_log`;');

    PREPARE stmt from @sql_statement;
	EXECUTE stmt;
    DEALLOCATE PREPARE stmt;

END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `Transaction_SP_AddLog` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`admin`@`%` PROCEDURE `Transaction_SP_AddLog`(
	IN _type INT,
	IN _player_id VARCHAR(64),
	IN _game_type VARCHAR(64),
	IN _game_id VARCHAR(64),
	IN _round_id VARCHAR(64),
	IN _timestamp BIGINT UNSIGNED,
	IN _id_lock VARCHAR(64),
	IN _id_cancel_lock VARCHAR(64),
	IN _id_bet VARCHAR(64),
	IN _id_cancel_bet VARCHAR(64),
	IN _id_settle VARCHAR(64),
	IN _id_unlock VARCHAR(64),
	IN _amount decimal(10,6) unsigned,
	IN _credit BIGINT UNSIGNED,
	IN _jackpot BIGINT UNSIGNED,
	IN _currency decimal(10,6) unsigned,
	IN _currency_code VARCHAR(8),
	IN _system_fee BIGINT UNSIGNED,
	IN _system_fee_jackpot BIGINT UNSIGNED,
	IN _percentage_win_to_system_fee INT UNSIGNED,
	IN _percentage_system_fee_to_jackpot INT UNSIGNED,
	IN _percentage_win_jackpot INT UNSIGNED,
	IN _platform VARCHAR(64),
	IN _platform_player_id VARCHAR(64),
	IN _platform_player_name VARCHAR(64),
	IN _platform_player_display_name VARCHAR(64),
	IN _platform_player_balance decimal(16,6) unsigned,
	IN _raw_data TEXT
)
BEGIN

	DECLARE `table_name` VARCHAR(128);
	DECLARE today VARCHAR(16);

	SELECT DATE_FORMAT(NOW(), "%Y%m%d") INTO today;
	SET `table_name` = CONCAT('`transaction_log_', today, '`');

	SET @type=_type;
	SET @player_id=_player_id;
	SET @game_type=_game_type;
	SET @game_id=_game_id;
	SET @round_id=_round_id;
	SET @timestamp=_timestamp;
	SET @id_lock=_id_lock;
	SET @id_cancel_lock=_id_cancel_lock;
	SET @id_bet=_id_bet;
	SET @id_cancel_bet=_id_cancel_bet;
	SET @id_settle=_id_settle;
	SET @id_unlock=_id_unlock;
	SET @amount=_amount;
	SET @credit=_credit;
	SET @jackpot=_jackpot;
	SET @currency=_currency;
	SET @currency_code=_currency_code;
	SET @system_fee=_system_fee;
	SET @system_fee_jackpot=_system_fee_jackpot;
	SET @percentage_win_to_system_fee=_percentage_win_to_system_fee;
	SET @percentage_system_fee_to_jackpot=_percentage_system_fee_to_jackpot;
	SET @percentage_win_jackpot=_percentage_win_jackpot;
	SET @platform=_platform;
	SET @platform_player_id=_platform_player_id;
	SET @platform_player_name=_platform_player_name;
	SET @platform_player_display_name=_platform_player_display_name;
	SET @platform_player_balance=_platform_player_balance;
	SET @raw_data=_raw_data;

	SET @sql_statement = CONCAT('INSERT INTO ', `table_name`, ' VALUES(NULL,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);');

	PREPARE stmt FROM @sql_statement;
	EXECUTE stmt USING 
		@type,
		@player_id,
		@game_type,
		@game_id,
		@round_id,
		@timestamp,
		@id_lock,
		@id_cancel_lock,
		@id_bet,
		@id_cancel_bet,
		@id_settle,
		@id_unlock,
		@amount,
		@credit,
		@jackpot,
		@currency,
		@currency_code,
		@system_fee,
		@system_fee_jackpot,
		@percentage_win_to_system_fee,
		@percentage_system_fee_to_jackpot,
		@percentage_win_jackpot,
		@platform,
		@platform_player_id,
		@platform_player_name,
		@platform_player_display_name,
		@platform_player_balance,
		@raw_data;
    DEALLOCATE PREPARE stmt;

END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-17 16:39:43

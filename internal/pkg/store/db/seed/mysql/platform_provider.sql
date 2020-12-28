--
DROP TABLE IF EXISTS `platform_provider`;

CREATE TABLE `platform_provider` (
  `name` varchar(32) NOT NULL,
  `factory_name` varchar(32) NOT NULL DEFAULT 'default',
  `aes_key` varchar(64) NOT NULL,
  `aes_iv` varchar(64) NOT NULL,
  `auth_type` int(11) NOT NULL DEFAULT '1' COMMENT '1: JWT 2: OAuth',
  `auth_id` varchar(64) NOT NULL,
  `auth_secret` varchar(64) NOT NULL,
  `auth_grant_type` varchar(256) NOT NULL DEFAULT 'client_credentials',
  `auth_scope` varchar(256) NOT NULL DEFAULT 'bets',
  `api_url_base` varchar(256) NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `Platform_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `platform`;
END ;;
DELIMITER ;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `Platform_SP_GetByName`(
	IN `_name` VARCHAR(64)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `platform` WHERE `name`=_name;
END ;;
DELIMITER ;

INSERT INTO `platform_provider` VALUES ('default','default','12345678','12345678',1,'default','12345678','client_credentials','bets','http://0.0.0.0/');
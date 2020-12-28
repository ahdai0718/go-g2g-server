--
DROP TABLE IF EXISTS `server_info`;

CREATE TABLE `server_info` (
  `name` varchar(32) NOT NULL,
  `group` varchar(32) DEFAULT 'group001',
  `type` int(11) NOT NULL,
  `host` varchar(128) DEFAULT NULL,
  `port` int(11) DEFAULT NULL,
  `protocol` varchar(8) DEFAULT 'http',
  `server_info_route_path` varchar(32) DEFAULT '/server/info',
  `websocket_protocol` varchar(32) DEFAULT 'ws',
  `websocket_route_path` varchar(32) DEFAULT 'ws',
  `host_for_client` varchar(128) DEFAULT NULL,
  `port_for_client` int(11) DEFAULT NULL,
  `websocket_protocol_for_client` varchar(32) DEFAULT 'ws',
  `websocket_route_path_for_client` varchar(32) DEFAULT 'ws',
  `public_ip_address` varchar(16) DEFAULT NULL,
  `is_offline` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `ServerInfo_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `server_info`;
END ;;
DELIMITER ;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `ServerInfo_SP_GetByName`(
	IN `_name` VARCHAR(64)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `server_info` WHERE `name`=_name;
END ;;
DELIMITER ;

--
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('fluentd001','',3,'0.0.0.0',8888,'http','/server/info','ws','/server/ws','0.0.0.0',0,'ws','ws','0.0.0.0',0,0);
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('GameExample','',2,'0.0.0.0',3011,'http','/server/info','ws','/server/ws','0.0.0.0',3021,'ws','ws','0.0.0.0',0,1001);
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('GameExample','',2,'0.0.0.0',3021,'http','/server/info','ws','/server/ws','0.0.0.0',3031,'ws','ws','0.0.0.0',0,1002);
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('GameExample','',2,'0.0.0.0',3001,'http','/server/info','ws','/server/ws','0.0.0.0',3011,'ws','ws','0.0.0.0',0,1);
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('Gateway001','',1,'0.0.0.0',3000,'http','/server/info','ws','/server/ws','0.0.0.0',3000,'ws','ws','0.0.0.0',0,0);
INSERT INTO `server_info` (`name`,`group`,`type`,`host`,`port`,`protocol`,`server_info_route_path`,`websocket_protocol`,`websocket_route_path`,`host_for_client`,`port_for_client`,`websocket_protocol_for_client`,`websocket_route_path_for_client`,`public_ip_address`,`is_offline`,`game_type`) VALUES ('Offline00101','groupOffline',2,'0.0.0.0',0,'http','/server/info','ws','/server/ws','0.0.0.0',0,'ws','ws','0.0.0.0',1,0);
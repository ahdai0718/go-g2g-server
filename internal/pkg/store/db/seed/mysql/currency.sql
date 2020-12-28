--
DROP TABLE IF EXISTS `currency`;

CREATE TABLE `currency` (
  `name` varchar(16) NOT NULL,
  `multiply_to_credit` double unsigned DEFAULT '1',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `Currency_SP_GetAll`()
    SQL SECURITY INVOKER
BEGIN
  SELECT * FROM `currency`;
END ;;
DELIMITER ;

--
DELIMITER ;;
CREATE DEFINER=CURRENT_USER PROCEDURE `Currency_SP_GetByName`(
	IN `_name` VARCHAR(16)
)
    SQL SECURITY INVOKER
BEGIN
	SELECT * FROM `currency` WHERE `name`=_name;
END ;;
DELIMITER ;
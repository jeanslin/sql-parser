# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.16-1~exp1)
# Database: leonids_db_for_tests
# Generation Time: 2018-02-12 15:22:21 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table sometesttable
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sometesttable`;

CREATE TABLE `sometesttable` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `Field_1` int(11) NOT NULL DEFAULT '0',
  `Field_211111112` varchar(11) NOT NULL DEFAULT '',
  `Field_3` int(11) NOT NULL DEFAULT '0',
  `Field_4` int(11) NOT NULL DEFAULT '0',
  `Field_5` int(11) NOT NULL DEFAULT '0',
  `Field_6` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `sometesttable` WRITE;
/*!40000 ALTER TABLE `sometesttable` DISABLE KEYS */;

INSERT INTO `sometesttable` (`id`, `Field_1`, `Field_211111112`, `Field_3`, `Field_4`, `Field_5`, `Field_6`)
VALUES
	(1,0,'0',0,0,0,0),
	(2,0,'something',0,0,0,0),
	(3,0,'something',0,0,0,0),
	(4,0,'something',0,0,0,0),
	(5,0,'something',0,0,0,0);

/*!40000 ALTER TABLE `sometesttable` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

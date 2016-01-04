/*
SQLyog Ultimate v11.24 (32 bit)
MySQL - 5.6.20 : Database - work
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`work` /*!40100 DEFAULT CHARACTER SET utf8 */;

/*Table structure for table `t_daliy_work` */
use work;

CREATE TABLE `t_daliy_work` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `uid` int(10) DEFAULT NULL,
  `content` longtext,
  `year` int(4) DEFAULT NULL COMMENT '年',
  `month` int(2) DEFAULT NULL COMMENT '月',
  `day` int(2) DEFAULT NULL COMMENT '日',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `t_user` */

CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
   `loginName` varchar(20) NOT NULL,
`userName` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `createData` datetime DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

CREATE DATABASE IF NOT EXISTS `sport_track_go`;
USE `sport_track_go`;
DROP TABLE IF EXISTS `Users`; 
CREATE TABLE `Users`(
  `Firstname` varchar(255) DEFAULT NULL,
  `Lastname` varchar(255) DEFAULT NULL,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Uuid` varchar(255) NOT NULL,
  CONSTRAINT UC_User UNIQUE (`Uuid`, `Email`) 
);
INSERT INTO `Users` VALUES ('Cardinal','Tom B. Erichsen','cardinal@gmail.com','password','a922505b-2347-4aa2-9094-10ebf9113d59'),('Wilman Kala','Matti Karttunen','wilman@gmail.com','password','37b0441a-8835-44dd-8999-3bfa9bcf915a');
UNLOCK TABLES;

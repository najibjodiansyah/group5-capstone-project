CREATE DATABASE  IF NOT EXISTS `capstoneproject` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `capstoneproject`;
-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: localhost    Database: capstoneproject
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Table structure for table `applications`
--

DROP TABLE IF EXISTS `applications`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `applications` (
  `id` int NOT NULL AUTO_INCREMENT,
  `employeeid` int DEFAULT NULL,
  `managerid` int DEFAULT NULL,
  `assetid` int DEFAULT NULL,
  `itemid` int DEFAULT NULL,
  `requestDate` datetime DEFAULT CURRENT_TIMESTAMP,
  `returnDate` datetime DEFAULT NULL,
  `activity` varchar(255) DEFAULT NULL,
  `spesification` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `empoyeeid` (`employeeid`),
  KEY `assetid` (`assetid`),
  KEY `applications_ibfk_3_idx` (`managerid`),
  CONSTRAINT `applications_ibfk_1` FOREIGN KEY (`employeeid`) REFERENCES `users` (`id`),
  CONSTRAINT `applications_ibfk_2` FOREIGN KEY (`assetid`) REFERENCES `assets` (`id`),
  CONSTRAINT `applications_ibfk_3` FOREIGN KEY (`managerid`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `applications`
--

LOCK TABLES `applications` WRITE;
/*!40000 ALTER TABLE `applications` DISABLE KEYS */;
INSERT INTO `applications` VALUES (1,1,6,19,33,'2022-02-26 01:34:02','2023-02-26 01:34:03',NULL,'awas','untuk sosyelita','donereturn','2022-02-28 02:57:50'),(2,1,NULL,19,0,'2022-02-26 01:36:49','2023-02-26 01:36:49',NULL,'','untuk sosyelita','','2022-02-28 11:12:01'),(3,1,NULL,19,0,'2022-02-26 01:37:50','2023-02-26 01:37:50',NULL,'','untuk sosyelita','','2022-02-28 11:12:01'),(4,1,NULL,19,0,'2022-02-26 01:38:02','2023-02-26 01:38:03',NULL,'ya setara iphone 13 pro mas','untuk sosyelita','','2022-02-28 11:12:01'),(5,1,NULL,19,0,'2022-02-26 01:40:44','2023-02-26 01:40:44',NULL,'ya setara iphone 13 pro mas','untuk sosyelita','toAdmin','2022-02-28 11:12:01'),(7,1,NULL,19,0,'2022-02-26 07:36:15','2023-02-26 07:36:15',NULL,'ya setara iphone 13 pro mas','ini admin yang assign ke employeeid 1','tomanager','2022-02-27 13:17:21'),(8,1,NULL,22,0,'2022-02-27 18:13:32','2023-01-01 07:00:00',NULL,'nyoba dari admin pake tanggal','untuk hehe','toManager','2022-02-28 11:12:01'),(9,1,NULL,23,45,'2022-03-01 06:57:43','2023-01-01 07:00:00',NULL,'test ','test','inuse','2022-03-01 06:57:17'),(10,3,6,19,33,'2022-03-01 06:57:43','2023-01-01 07:00:00',NULL,'nyoba dari admin pake tanggal','untuk sosyelita','inuse','2022-03-01 06:57:17');
/*!40000 ALTER TABLE `applications` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `assets`
--

DROP TABLE IF EXISTS `assets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `assets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` longtext,
  `categoryid` int DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `categoryid` (`categoryid`),
  CONSTRAINT `assets_ibfk_1` FOREIGN KEY (`categoryid`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `assets`
--

LOCK TABLES `assets` WRITE;
/*!40000 ALTER TABLE `assets` DISABLE KEYS */;
INSERT INTO `assets` VALUES (19,'iphone 12 pro max','buat sosialita',4,5,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645728014.jpeg','2022-02-22 17:05:08'),(21,'AC Samsung anti covid Ruangan ','AC 1PK untuk totalitas karyawan dalam bekerja',4,5,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645728014.jpeg','2022-02-23 10:15:31'),(22,'AC Samsung anti covid Ruangan ','AC 1PK untuk totalitas karyawan dalam bekerja',1,2,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645728014.jpeg','2022-02-25 00:03:38'),(23,'Poster ruangan','biar kamar kerja aestetik',2,2,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645728014.jpeg','2022-02-25 01:40:14'),(24,'Poster ruangan','biar kamar kerja aestetik',2,2,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645955967.jpeg','2022-02-27 16:59:28'),(26,'Poster ruangan','biar kamar kerja aestetik',2,2,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1645956328.jpeg','2022-02-27 17:05:28'),(27,'Toyota Kijang Innova','Digunakan untuk melakukan perjalanan dinas, dalam maupun luar kota.',2,3,'https://capstone-group-5.s3.ap-southeast-1.amazonaws.com/0-1646115064.jpeg','2022-03-01 06:11:04');
/*!40000 ALTER TABLE `assets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Laptop'),(2,'Alat Tulis'),(3,'Kendaraan'),(4,'Lainnya');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `assetid` int DEFAULT NULL,
  `employee` int DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `availablestatus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `assetid` (`assetid`),
  KEY `employee` (`employee`),
  CONSTRAINT `items_ibfk_1` FOREIGN KEY (`assetid`) REFERENCES `assets` (`id`),
  CONSTRAINT `items_ibfk_2` FOREIGN KEY (`employee`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
INSERT INTO `items` VALUES (33,19,1,'iphone 12 pro max 1','tersedia'),(34,19,1,'iphone 12 pro max 2','pemeliharaan'),(35,19,1,'iphone 12 pro max 3','pemeliharaan'),(36,19,1,'iphone 12 pro max 4','pemeliharaan'),(37,19,1,'iphone 12 pro max 5','pemeliharaan'),(38,21,1,'AC Samsung anti covid Ruangan  1','tersedia'),(39,21,1,'AC Samsung anti covid Ruangan  2','tersedia'),(40,21,1,'AC Samsung anti covid Ruangan  3','pemeliharaan'),(41,21,1,'AC Samsung anti covid Ruangan  4','pemeliharaan'),(42,21,1,'AC Samsung anti covid Ruangan  5','tersedia'),(43,22,1,'AC Samsung anti covid Ruangan  1','tersedia'),(44,22,1,'AC Samsung anti covid Ruangan  2','tersedia'),(45,23,1,'Poster ruangan 1','digunakan'),(46,23,1,'Poster ruangan 2','tersedia'),(47,24,1,'Poster ruangan 1','tersedia'),(48,24,1,'Poster ruangan 2','tersedia'),(49,26,1,'Poster ruangan 1','tersedia'),(50,26,1,'Poster ruangan 2','tersedia'),(51,27,1,'Toyota Kijang Innova 1','tersedia'),(52,27,1,'Toyota Kijang Innova 2','tersedia'),(53,27,1,'Toyota Kijang Innova 3','tersedia');
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `procurements`
--

DROP TABLE IF EXISTS `procurements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `procurements` (
  `id` int NOT NULL AUTO_INCREMENT,
  `employeeid` int DEFAULT NULL,
  `managerid` int DEFAULT NULL,
  `assetName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `requestDate` datetime DEFAULT CURRENT_TIMESTAMP,
  `activity` varchar(255) DEFAULT NULL,
  `spesification` varchar(255) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `empoyeeid` (`employeeid`),
  KEY `managerid` (`managerid`),
  CONSTRAINT `procurements_ibfk_1` FOREIGN KEY (`employeeid`) REFERENCES `users` (`id`),
  CONSTRAINT `procurements_ibfk_2` FOREIGN KEY (`managerid`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `procurements`
--

LOCK TABLES `procurements` WRITE;
/*!40000 ALTER TABLE `procurements` DISABLE KEYS */;
/*!40000 ALTER TABLE `procurements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'najib','najib@jodiansyah.com','$2a$10$1G28p8R0TrBdcjFN5TbziOB8ZYE6ICrNhT2WyNhiNIveMDUtiqIOC','hehe.jpg','employee','2022-02-14 23:47:31',NULL),(3,'jodiansyah','jodiansyah@jodiansyah.com','$2a$10$KIbNCKir.1Dh9GDlI/sugOVS5n2Pw4wiG4ZuMCTpu/d4MYf7P/6Kq','jodiansyah.jpg','employee','2022-02-15 06:08:03',NULL),(4,'sukur','sukur@jodiansyah.com','$2a$10$0e8UVYoKDHd.PTkXYMRQ9eeMVUfWMs2ZTBUDS6A1cSY68pEf4Mjsi','','employee','2022-02-15 06:50:24',NULL),(5,'empty','empty@jodiansyah.com','$2a$10$/M4/YffCmHwHVmOB02yAKugUb6OGezjpmBaT/2d.5Stp3QXo5rNsK','empty.jpg','employee','2022-02-15 06:53:21','2022-02-15 07:16:34'),(6,'manager','manager@manager.com','$2a$10$sHTL.NjjARnol25GYSu5yOUrLt/j2bNGHK1G4EwgY1XDqd/7mBXRq','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','manager','2022-02-22 00:17:10',NULL),(7,'admin','admin@admin.com','$2a$10$Jip45xcLd9Os6foDqkEgy.aVu5Di9yM83J6fXKXf91H.cunBFFpcS','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','admin','2022-02-25 01:44:44',NULL),(8,'','jemi.employee@sirclo.com','$2a$10$EEp0yZAeUJnMwpm8q.sx3.cCfRCQ/cxp3uPqSXaELt6gz0DpUcfuy','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','employee','2022-03-01 02:20:20',NULL),(9,'','jemi.admin@sirclo.com','$2a$10$y3Ljf76NKDKKI2XlBNQ8t.9qMbotL4.He3uCHXQMxSdEr.u5jtzF2','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','admin','2022-03-01 02:22:04',NULL),(10,'','jemi.manager@sirclo.com','$2a$10$qzvUqED4Q.g5lTj4ikCtEO1nqxggZr08Uy9cBS4kUt9A.kdd9txZq','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','manager','2022-03-01 02:22:14',NULL),(11,'','jemi1.admin@sirclo.com','$2a$10$qiMfciCbi9lnuwW26qMU0e5M9mkqQtTb1eSXM4HqWhJfKAexyhO0O','https://d11a6trkgmumsb.cloudfront.net/original/3X/d/8/d8b5d0a738295345ebd8934b859fa1fca1c8c6ad.jpeg','admin','2022-03-01 02:23:44',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-02 21:19:14

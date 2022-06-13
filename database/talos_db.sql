-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Εξυπηρετητής: 127.0.0.1
-- Χρόνος δημιουργίας: 17 Μαρ 2022 στις 19:29:26
-- Έκδοση διακομιστή: 10.4.19-MariaDB
-- Έκδοση PHP: 8.0.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Βάση δεδομένων: `talos_db`
--
CREATE DATABASE IF NOT EXISTS `talos_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `talos_db`;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `bme280_list`
--

DROP TABLE IF EXISTS `bme280_list`;
CREATE TABLE IF NOT EXISTS `bme280_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `bme280_values`
--

DROP TABLE IF EXISTS `bme280_values`;
CREATE TABLE IF NOT EXISTS `bme280_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `temperature` double DEFAULT NULL,
  `humidity` double DEFAULT NULL,
  `barometric_pressure` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `dust_sensor_list`
--

DROP TABLE IF EXISTS `dust_sensor_list`;
CREATE TABLE IF NOT EXISTS `dust_sensor_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `dust_sensor_values`
--

DROP TABLE IF EXISTS `dust_sensor_values`;
CREATE TABLE IF NOT EXISTS `dust_sensor_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `dust_density` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq2_list`
--

DROP TABLE IF EXISTS `mq2_list`;
CREATE TABLE IF NOT EXISTS `mq2_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq2_values`
--

DROP TABLE IF EXISTS `mq2_values`;
CREATE TABLE IF NOT EXISTS `mq2_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `h2` double DEFAULT NULL,
  `lpg` double DEFAULT NULL,
  `propane` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq3_list`
--

DROP TABLE IF EXISTS `mq3_list`;
CREATE TABLE IF NOT EXISTS `mq3_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq3_values`
--

DROP TABLE IF EXISTS `mq3_values`;
CREATE TABLE IF NOT EXISTS `mq3_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `alcohol` double DEFAULT NULL,
  `benzene` double DEFAULT NULL,
  `hexane` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq4_list`
--

DROP TABLE IF EXISTS `mq4_list`;
CREATE TABLE IF NOT EXISTS `mq4_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq4_values`
--

DROP TABLE IF EXISTS `mq4_values`;
CREATE TABLE IF NOT EXISTS `mq4_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `ch4` double DEFAULT NULL,
  `co` double DEFAULT NULL,
  `smoke` double NOT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq135_list`
--

DROP TABLE IF EXISTS `mq135_list`;
CREATE TABLE IF NOT EXISTS `mq135_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `mq135_values`
--

DROP TABLE IF EXISTS `mq135_values`;
CREATE TABLE IF NOT EXISTS `mq135_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `co2` double DEFAULT NULL,
  `tolueno` double DEFAULT NULL,
  `nh4` double DEFAULT NULL,
  `acetone` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `rain_sensor_list`
--

DROP TABLE IF EXISTS `rain_sensor_list`;
CREATE TABLE IF NOT EXISTS `rain_sensor_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `rain_sensor_values`
--

DROP TABLE IF EXISTS `rain_sensor_values`;
CREATE TABLE IF NOT EXISTS `rain_sensor_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `sensor_value` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `solar_module_list`
--

DROP TABLE IF EXISTS `solar_module_list`;
CREATE TABLE IF NOT EXISTS `solar_module_list` (
  `sensor_id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `solar_module_values`
--

DROP TABLE IF EXISTS `solar_module_values`;
CREATE TABLE IF NOT EXISTS `solar_module_values` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sensor_id` int(11) DEFAULT NULL,
  `moment` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `solar_charge` tinyint(1) DEFAULT NULL,
  `solar_done` tinyint(1) DEFAULT NULL,
  `solar_warning` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sensor_id` (`sensor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Δομή πίνακα για τον πίνακα `weather_station`
--

DROP TABLE IF EXISTS `weather_station`;
CREATE TABLE IF NOT EXISTS `weather_station` (
  `station_id` int(11) NOT NULL,
  `latitude` varchar(45) NOT NULL,
  `longitude` varchar(45) NOT NULL,
  `bme280_id` int(11) NOT NULL,
  `mq2_id` int(11) NOT NULL,
  `mq3_id` int(11) NOT NULL,
  `mq4_id` int(11) NOT NULL,
  `mq135_id` int(11) NOT NULL,
  `rain_sensor_id` int(11) NOT NULL,
  `dust_sensor_id` int(11) NOT NULL,
  `solar_module_id` int(11) NOT NULL,
  PRIMARY KEY (`station_id`),
  KEY `bme280_id` (`bme280_id`),
  KEY `mq5_id` (`mq3_id`),
  KEY `mq7_id` (`mq4_id`),
  KEY `mq135_id` (`mq135_id`),
  KEY `rain_sensor_id` (`rain_sensor_id`),
  KEY `dust_sensor_id` (`dust_sensor_id`),
  KEY `solar_module_id` (`solar_module_id`),
  KEY `mq2_id` (`mq2_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Περιορισμοί για άχρηστους πίνακες
--

--
-- Περιορισμοί για πίνακα `bme280_values`
--
ALTER TABLE `bme280_values`
  ADD CONSTRAINT `bme280_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `bme280_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `dust_sensor_values`
--
ALTER TABLE `dust_sensor_values`
  ADD CONSTRAINT `dust_sensor_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `dust_sensor_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `mq2_values`
--
ALTER TABLE `mq2_values`
  ADD CONSTRAINT `mq2_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `mq2_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `mq3_values`
--
ALTER TABLE `mq3_values`
  ADD CONSTRAINT `mq3_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `mq3_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `mq4_values`
--
ALTER TABLE `mq4_values`
  ADD CONSTRAINT `mq4_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `mq4_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `mq135_values`
--
ALTER TABLE `mq135_values`
  ADD CONSTRAINT `mq135_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `mq135_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `rain_sensor_values`
--
ALTER TABLE `rain_sensor_values`
  ADD CONSTRAINT `rain_sensor_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `rain_sensor_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `solar_module_values`
--
ALTER TABLE `solar_module_values`
  ADD CONSTRAINT `solar_module_values_ibfk_1` FOREIGN KEY (`sensor_id`) REFERENCES `solar_module_list` (`sensor_id`);

--
-- Περιορισμοί για πίνακα `weather_station`
--
ALTER TABLE `weather_station`
  ADD CONSTRAINT `weather_station_ibfk_1` FOREIGN KEY (`mq2_id`) REFERENCES `mq2_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_2` FOREIGN KEY (`bme280_id`) REFERENCES `bme280_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_3` FOREIGN KEY (`mq3_id`) REFERENCES `mq3_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_4` FOREIGN KEY (`mq4_id`) REFERENCES `mq4_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_5` FOREIGN KEY (`mq135_id`) REFERENCES `mq135_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_6` FOREIGN KEY (`rain_sensor_id`) REFERENCES `rain_sensor_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_7` FOREIGN KEY (`dust_sensor_id`) REFERENCES `dust_sensor_list` (`sensor_id`),
  ADD CONSTRAINT `weather_station_ibfk_8` FOREIGN KEY (`solar_module_id`) REFERENCES `solar_module_list` (`sensor_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

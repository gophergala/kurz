-- phpMyAdmin SQL Dump
-- version 3.4.10.1deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jan 24, 2015 at 05:32 PM
-- Server version: 5.5.40
-- PHP Version: 5.4.36-1+deb.sury.org~precise+2

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `go_kurz`
--

-- --------------------------------------------------------

--
-- Table structure for table `shorturl`
--

CREATE TABLE IF NOT EXISTS `shorturl` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(32) NOT NULL COMMENT 'The short URL itself',
  `domain` int(11) NOT NULL,
  `strategy` varchar(8) NOT NULL DEFAULT 'base',
  `submittedBy` int(11) NOT NULL,
  `submittedInfo` int(11) NOT NULL,
  `isEnabled` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `url` (`url`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1 ;

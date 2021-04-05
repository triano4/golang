-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 25, 2021 at 06:57 AM
-- Server version: 5.7.24
-- PHP Version: 7.2.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_sheet`
--

-- --------------------------------------------------------

--
-- Table structure for table `navs`
--

CREATE TABLE `navs` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `sort_id` int(11) DEFAULT NULL,
  `nav_icon` varchar(255) DEFAULT NULL,
  `nav_path` varchar(255) DEFAULT NULL,
  `nav_acc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `navs`
--

INSERT INTO `navs` (`id`, `name`, `parent_id`, `sort_id`, `nav_icon`, `nav_path`, `nav_acc`) VALUES
(1, 'Home', 0, 1, '', '/', 'public'),
(2, 'User', 0, 1, '', '/Report', 'public'),
(3, 'Setting', 0, 1, '', '/', 'public');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_email` varchar(100) NOT NULL,
  `access` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(100) DEFAULT NULL,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_by` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `user_email`, `access`, `created_at`, `created_by`, `updated_at`, `updated_by`) VALUES
(1, 'celerate.indonesia@gmail.com', 'private', '2021-03-09 13:37:02', '', '2021-03-09 13:37:02', ''),
(2, 'luther@gmail.com', 'public', '2021-03-09 13:37:02', '', '2021-03-09 13:37:02', ''),
(3, 'steven@gmail.com', 'public', '2021-03-09 13:37:02', '', '2021-03-09 13:37:02', '');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nickname`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Steven victor', 'steven@gmail.com', '$2a$10$HRhoSGGHTIWzLHbjgQI/Me76Qoq3bs5Y/JA7N2HsBgL.Ey3g/8U.G', '2021-03-09 13:37:02', '2021-03-09 13:37:02'),
(2, 'Martin Luther', 'luther@gmail.com', '$2a$10$vpVixwfMnilPKAs1iJvLlOZX4KyeBG/1gViOCyi8qltr970CSeoci', '2021-03-09 13:37:02', '2021-03-09 13:37:02'),
(3, 'Admin Celerates', 'celerate.indonesia@gmail.com', '$2a$10$OgAomTlPQrhboURx9yHD4uOjyBvxmmGLFSfpWRpeS7/1Lt8ffItyO', '2021-03-09 13:37:02', '2021-03-09 13:37:02'),
(4, 'sandry', 'sandry@celerates.com', '$2a$10$iVxu38bNTD9QZ29Dmqge8.jrxl9XZzuOpiuYiw.PRbeY35HivOv36', '2021-03-09 13:37:26', '2021-03-09 13:37:26');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `navs`
--
ALTER TABLE `navs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nickname` (`nickname`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `navs`
--
ALTER TABLE `navs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

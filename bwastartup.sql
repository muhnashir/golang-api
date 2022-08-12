-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Aug 12, 2022 at 01:42 PM
-- Server version: 8.0.30-0ubuntu0.20.04.2
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bwastartup`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_description` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int NOT NULL,
  `goal_amount` int NOT NULL,
  `current_amount` int NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 9, 'reland', 'reland', 'desxreanddddddddd', 'peakr,hehe, cuyy', 0, 100, 90, 're-land', '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(2, 9, 'momofin', 'mom', 'desxreanddddddddd', 'peakr', 0, 100, 90, 'momofin', '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(3, 10, 'adaro', 'ada', 'desxreanddddddddd', 'peakr', 0, 100, 90, 'ada', '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(5, 11, 'cek ini adalah nama update hack', 'ini short description,update', 'ini description update', 'perks 1, perks 2, perks 3,update', 0, 250000, 0, 'cek-ini-adalah-nama-11', '2022-07-12 19:58:47', '2022-07-12 21:37:35'),
(6, 11, 'cek ini adalah nama', 'ini short description', 'ini description', 'perks 1, perks 2, perks 3', 0, 200000, 0, 'cek-ini-adalah-nama-11', '2022-07-12 20:47:04', '2022-07-12 20:47:04');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int NOT NULL,
  `campaign_id` int NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `is_primary` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'satu.jpg', 0, '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(2, 1, 'dua.jpg', 1, '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(3, 1, 'tiga.jpg', 0, '2022-06-30 22:11:00', '2022-06-30 22:11:00'),
(4, 6, 'images/11-Black-Ubuntu-20-04-Default-Wallpaper-Dotted-Eyes-Corner-scaled.jpeg', 0, '2022-07-20 15:00:44', '2022-07-20 15:02:47'),
(5, 6, 'images/11-wallpapersden.com_baby-groot-minimalist_1920x1080.jpg', 1, '2022-07-20 15:02:47', '2022-07-20 15:02:47'),
(6, 6, 'images/11-wallpapersden.com_baby-groot-minimalist_1920x1080.jpg', 0, '2022-07-20 15:06:44', '2022-07-20 15:06:44'),
(7, 6, 'images/11-img.jpeg', 0, '2022-07-20 15:08:34', '2022-07-20 15:08:34'),
(8, 6, 'images/11-img.jpeg', 0, '2022-07-20 15:10:06', '2022-07-20 15:10:06');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int NOT NULL,
  `campaign_id` int NOT NULL,
  `user_id` int NOT NULL,
  `amount` int NOT NULL,
  `status` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `payment_url` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `payment_url`, `created_at`, `updated_at`) VALUES
(1, 5, 11, 34000, 'PAID', '', '', '2022-07-20 19:55:01', '2022-07-20 19:55:01'),
(2, 6, 11, 234000, 'PAID', '', '', '2022-07-20 19:56:36', '2022-07-20 19:56:36'),
(9, 1, 11, 3333, 'PENDING', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/b656e3cb-1863-4d72-b774-2c9936dd7ad8', '2022-08-02 13:30:10', '2022-08-02 13:30:11'),
(10, 1, 11, 345000, 'PENDING', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/39267e22-c87c-4ad0-b964-e8f5f32a5b0d', '2022-08-02 13:43:21', '2022-08-02 13:43:22'),
(11, 1, 9, 345000, 'PENDING', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/a2361a32-1a75-41ab-8574-19dc7adcb21f', '2022-08-02 13:48:12', '2022-08-02 13:48:13'),
(12, 1, 9, 245000, 'PENDING', '', 'https://app.sandbox.midtrans.com/snap/v3/redirection/c23eec74-444d-4b40-b3bb-98e8efd461d0', '2022-08-02 14:13:25', '2022-08-02 14:13:26');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'Muhammad Nashir', 'Programmer', 'nashir@gmail.com', 'password', 'avatar.jpg', 'user', '2021-05-14 20:48:54', '2021-05-14 20:48:54'),
(2, 'Glady sunggoro', 'Teacher', 'glady@gmail.com', 'password', 'avatar.jpg', 'user', '2021-05-14 20:48:54', '2021-05-14 20:48:54'),
(3, 'test simpan', '', '', '', '', '', '2021-05-21 01:42:39', '2021-05-21 01:42:39'),
(4, 'Nashir', 'mahasiswa', 'muhammadnashir@gmail.com', '$2a$04$JJfeRZkoEPaeoq6.6YC9iOP.gR5cugDBnmago3hwVs0tr3ZPIJL1.', '', 'user', '2021-05-21 02:37:45', '2021-05-21 02:37:45'),
(5, 'Mujahid', 'Backend developer', 'mujahid@gmail.com', '$2a$04$TUFUFUpqMZkGS9w89WELseJ4WgZVvaRmoEKhoJSIaaGf8bB.Ns64S', '', 'user', '2021-10-07 20:06:10', '2021-10-07 20:06:10'),
(6, 'lonang', 'front end developer', 'lonang@gmail.com', '$2a$04$9leGWC8gEOd5bH0sym6MJe66/55jzx7LNFYrDz/9I0G7H.wSG1MUC', '', 'user', '2021-10-07 20:49:21', '2021-10-07 20:49:21'),
(7, 'lutfi', 'frontend developer', 'lutfi@gmail.com', '$2a$04$lTWTnEagm8ph9b556mc46OYp0S5DlO0riplT1WbBvUAQYJvo7rFl2', '', 'user', '2021-10-07 21:03:16', '2021-10-07 21:03:16'),
(9, 'nashir', 'programmer', 'nashir@transisi.id', '$2a$04$D1rX4JPWGrBVTwWVb5gmWerxrczIrWPKN494CEnzYwz5BKYx9DUVC', '', 'user', '2022-05-17 18:26:36', '2022-07-05 20:13:22'),
(11, 'jwt', 'arsitek', 'jwt@gmail.com', '$2a$04$bhXL.EhTCMQIcEAryoj78.CHqexdejMcuf0GwIKNh.2JmeHKP3oHq', 'images/11-Focal-Fossa-MATE-Wallpaper-Green-scaled.jpeg', 'user', '2022-06-30 22:11:00', '2022-07-05 20:14:11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

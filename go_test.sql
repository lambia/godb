-- phpMyAdmin SQL Dump
-- version 4.6.5.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Creato il: Set 21, 2017 alle 10:29
-- Versione del server: 10.1.21-MariaDB
-- Versione PHP: 7.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_test`
--
CREATE DATABASE IF NOT EXISTS `go_test` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `go_test`;

-- --------------------------------------------------------

--
-- Struttura della tabella `users`
--

CREATE TABLE `users` (
  `ID` mediumint(8) UNSIGNED ZEROFILL NOT NULL,
  `Nome` tinytext NOT NULL,
  `Mail` tinytext NOT NULL,
  `Password` tinytext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Svuota la tabella prima dell'inserimento `users`
--

TRUNCATE TABLE `users`;
--
-- Dump dei dati per la tabella `users`
--

INSERT INTO `users` (`ID`, `Nome`, `Mail`, `Password`) VALUES
(00000001, 'Dio', 'd@i.o', 'godsake'),
(00000026, 'Luca', 'luca@lambia.it', 'mammt'),
(00000027, 'Lambi', 'luca@lambia.com', 'mammt'),
(00000028, 'Soreta', 'luca@lambia.com', 'mammt'),
(00000029, 'Gianfranco', 'gf@marte.ano', 'mammt'),
(00000048, 'Lambi', 'luca@lambia.com', 'mammt'),
(00000054, 'Lambi', 'luca@lambia.com', 'mammt');

--
-- Indici per le tabelle scaricate
--

--
-- Indici per le tabelle `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT per le tabelle scaricate
--

--
-- AUTO_INCREMENT per la tabella `users`
--
ALTER TABLE `users`
  MODIFY `ID` mediumint(8) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=58;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

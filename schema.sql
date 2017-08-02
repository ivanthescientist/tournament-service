DROP DATABASE `tournament_db`;

CREATE DATABASE `tournament_db` COLLATE utf8_general_ci;

USE `tournament_db`;

CREATE TABLE `players` (
  `id` VARCHAR(128) NOT NULL,
  `balance` BIGINT(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tournaments` (
	`id` VARCHAR(128) NOT NULL,
  `deposit` BIGINT(20) NOT NULL DEFAULT 0,
  `winner` VARCHAR(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`winner`) REFERENCES `players` (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tournament_participants` (
  `tournamentId` VARCHAR(128) NOT NULL,
  `participantId` VARCHAR(128) NOT NULL,
  `parentId` VARCHAR(128) DEFAULT NULL,
  FOREIGN KEY (`tournamentId`) REFERENCES `tournaments` (`id`)  ON DELETE CASCADE,
  FOREIGN KEY (`participantId`) REFERENCES `players` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`parentId`) REFERENCES `players` (`id`) ON DELETE CASCADE,
  PRIMARY KEY (`tournamentId`, `participantId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
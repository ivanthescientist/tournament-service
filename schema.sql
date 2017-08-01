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
  `tournament_id` VARCHAR(128) NOT NULL,
  `participant_id` VARCHAR(128) NOT NULL,
  `parent_id` VARCHAR(128) DEFAULT NULL,
  FOREIGN KEY (`tournament_id`) REFERENCES `tournaments` (`tournament_id`)  ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`participant_id`) REFERENCES `players` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (`parent_id`) REFERENCES `players` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (`tournament_id`, `participant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
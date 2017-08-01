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
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tournament_participants` (
  `tournament_id` VARCHAR(128) NOT NULL,
  `participant_id` VARCHAR(128) NOT NULL,
  `deposit` BIGINT(20) NOT NULL DEFAULT 0,
  `parent_id` VARCHAR(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
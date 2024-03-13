-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS `item_log` (
  `item_log_id` int(11) NOT NULL AUTO_INCREMENT,
  `item_id` varchar(36) NOT NULL,
  `action` varchar(50) NOT NULL,
  `activity` TEXT NOT NULL,
  `created` timestamp NULL DEFAULT current_timestamp(),
  `business_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`item_log_id`)
);
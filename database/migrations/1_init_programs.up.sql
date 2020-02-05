USE SCHEMA `mathnavdb`;
DROP TABLE IF EXISTS `programs`;
CREATE TABLE `programs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `deleted_at` timestamp,
  `program_id` varchar(64) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `grade1` tinyint unsigned NOT NULL,
  `grade2` tinyint unsigned NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;

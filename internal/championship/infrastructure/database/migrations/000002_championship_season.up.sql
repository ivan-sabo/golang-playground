CREATE TABLE `golang_playground`.`championship_season` (
  `championship_id` VARCHAR(36) NOT NULL,
  `season_id` VARCHAR(36) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  CONSTRAINT `championship_season_pk` PRIMARY KEY (`championship_id`, `season_id`),
  CONSTRAINT `FK_championship` FOREIGN KEY (`championship_id`) REFERENCES `championships`(`id`),
  CONSTRAINT `FK_season` FOREIGN KEY (`season_id`) REFERENCES `seasons`(`id`)
);
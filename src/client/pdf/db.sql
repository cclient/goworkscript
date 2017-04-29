CREATE SCHEMA `geekbang` DEFAULT CHARACTER SET utf8mb4 ;
USE geekbang;
CREATE TABLE `geekbang`.`pdf` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `engine` CHAR(10) NOT NULL,
  `source_id` CHAR(13) NOT NULL,
  `title` VARCHAR(100) NOT NULL,
  `url` VARCHAR(200) NOT NULL,
  `desc` VARCHAR(500) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `source_id_UNIQUE` (`source_id` ASC));

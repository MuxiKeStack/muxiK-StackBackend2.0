create database kstack;
use kstack;
CREATE TABLE `curriculas`(
    `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `cid`            INT          NOT NULL COMMENT '课程编号',
    `curricula_name` VARCHAR(255) NOT NULL,
    `teacher`        VARCHAR(255) NOT NULL,
    `type`           TINYINT(1) NOT NULL,
    `rate` FLOAT,
    `starts_num` INT UNSIGNED,
    `grade_sample_size` INT UNSIGNED,
    `total_grade` FLOAT,
    `usual_grade` FLOAT,
    `grade_r1` INT UNSIGNED,
    `grade_r2` INT UNSIGNED,
    `grade_r3` INT UNSIGNED,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `collections`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `sid` VARCHAR(10) NOT NULL,
    `c_data_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
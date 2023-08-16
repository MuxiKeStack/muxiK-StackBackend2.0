CREATE TABLE `curriculas`(
    `id`             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `cid`            INT          NOT NULL UNIQUE COMMENT '课程编号',
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
    PRIMARY KEY (`id`),
    UNIQUE KEY `cid_index` (`cid`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
CREATE TABLE `user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `sid`        VARCHAR(10)  NOT NULL UNIQUE COMMENT '学生学号',
    `username`   VARCHAR(25)  NOT NULL ,
    `avatar`     VARCHAR(255)  NOT NULL ,
    `is_blocked` TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '封号',
    `role_grade` INT NOT NULL DEFAULT 0,
    `integral` INT NOT NULL DEFAULT 0 COMMENT'积分',
    `licence`    TINYINT(1)   NOT NULL DEFAULT 0 COMMENT '成绩查看许可',
     PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sid_unique`(`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
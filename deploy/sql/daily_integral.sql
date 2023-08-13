CREATE TABLE `daily_integral`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `sid` VARCHAR(10)  NOT NULL UNIQUE COMMENT '学生学号',
    `like_integral` INT NOT NULL DEFAULT 0 COMMENT'点赞积分',
    `comment_integral` INT NOT NULL DEFAULT 0 COMMENT'评论积分',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_sid_unique`(`sid`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
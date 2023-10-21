CREATE TABLE `reply` (
                           `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                           `comment_id` INT UNSIGNED NOT NULL COMMENT '关联评论的comment_id',
                           `sid` VARCHAR(45) NOT NULL COMMENT '回复者学号',
                           `text` TEXT NOT NULL COMMENT '回复内容',
                           `date` BIGINT NOT NULL COMMENT '回复日期',
                           PRIMARY KEY (`id`)
#                            FOREIGN KEY (`comment_id`) REFERENCES `comments`(`comment_id`),
#                            FOREIGN KEY (`sid`) REFERENCES `user`(`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

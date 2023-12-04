CREATE TABLE `comment` (
                            `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                            `evaluation_pid` INT UNSIGNED NOT NULL COMMENT '关联评价的pid',
                            `sid` VARCHAR(45) NOT NULL COMMENT '评论者学号',
                            `text` TEXT NOT NULL COMMENT '评论内容',
                            `date` BIGINT NOT NULL COMMENT '评论日期',
                            `object_sid` INT UNSIGNED,
                            Primary key(`id`)
#                             FOREIGN KEY (`evaluation_pid`) REFERENCES `evaluation_info`(`pid`),
#                             FOREIGN KEY (`sid`) REFERENCES `user`(`sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

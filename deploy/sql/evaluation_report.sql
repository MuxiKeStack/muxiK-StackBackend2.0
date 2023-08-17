CREATE TABLE `evaluation_report` (
                                     `id` int NOT NULL AUTO_INCREMENT,
                                     `pid` varchar(45) NOT NULL,
                                     `sid` varchar(45) NOT NULL,
                                     `reason` varchar(1000) DEFAULT NULL COMMENT '举报原因',
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
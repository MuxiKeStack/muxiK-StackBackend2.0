CREATE TABLE `evaluation_like` (
                                   `pid` int NOT NULL,
                                   `sid` varchar(45) NOT NULL,
                                   `status` int(10) unsigned zerofill NOT NULL COMMENT '"like:1, nothing:0, dislike:-1 "',
                                   PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
CREATE TABLE `evaluation_info` (
                                   `pid` int NOT NULL AUTO_INCREMENT,
                                   `sid` varchar(45) NOT NULL COMMENT '''学号''',
                                   `cid` varchar(45) NOT NULL COMMENT '''课程号''',
                                   `folded` tinyint(3) unsigned zerofill NOT NULL,
                                   `deleted` tinyint(3) unsigned zerofill NOT NULL,
                                   `info` varchar(1000) NOT NULL COMMENT '''评课内容''',
                                   `liked` int(10) unsigned zerofill DEFAULT NULL,
                                   `disliked` int(10) unsigned zerofill DEFAULT NULL,
                                   `CreatedAt` datetime DEFAULT NULL,
                                   `UpdatedAt` datetime DEFAULT NULL,
                                   `DeletedAt` datetime DEFAULT NULL,
                                   PRIMARY KEY (`pid`),
                                   UNIQUE KEY `pid_UNIQUE` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

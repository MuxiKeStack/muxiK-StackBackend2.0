-- 创建 NormalMessage 表
CREATE TABLE normal_message (
                               id INT UNSIGNED AUTO_INCREMENT,
                               type VARCHAR(255) NOT NULL,
                               sender_sid VARCHAR(40) NOT NULL,
                               object_sid VARCHAR(40) NOT NULL,
                               origin_content_id INT UNSIGNED  NOT NULL,
                               text TEXT,
                               send_at BIGINT NOT NULL,
                               PRIMARY KEY (`id`)
);

CREATE TABLE official_message (
                                 id INT UNSIGNED AUTO_INCREMENT,
                                 title VARCHAR(255) NOT NULL,
                                 text TEXT NOT NULL,
                                 image VARCHAR(255),
                                 object_sid VARCHAR(40),
                                 send_at BIGINT NOT NULL,
                                 PRIMARY KEY (`id`)
);

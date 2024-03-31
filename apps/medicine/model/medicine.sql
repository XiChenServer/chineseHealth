CREATE TABLE chinese_medicine_info (
                                       id INT PRIMARY KEY AUTO_INCREMENT COMMENT '药材ID',
                                       name VARCHAR(255) NOT NULL COMMENT '药名',
                                       alias VARCHAR(255) COMMENT '别名',
                                       taste TEXT COMMENT '性味',
                                       meridian TEXT COMMENT '归经',
                                       efficacy TEXT COMMENT '功效',
                                       usage_dosage TEXT COMMENT '用法用量',
                                       contraindications TEXT COMMENT '禁忌',
                                           PRIMARY KEY (`id`)
) COMMENT='中药基本信息表';
CREATE TABLE medicine_images (
                                 id INT PRIMARY KEY AUTO_INCREMENT COMMENT '图片ID',
                                 medicine_id INT NOT NULL COMMENT '药品ID',
                                 image_url VARCHAR(255) NOT NULL COMMENT '图片链接',
                                     PRIMARY KEY (`id`)
) COMMENT='药品图片链接表';

CREATE TABLE `user_descs` (
                              `id` int unsigned NOT NULL AUTO_INCREMENT,
                              `created_at` datetime DEFAULT NULL,
                              `updated_at` datetime DEFAULT NULL,
                              `deleted_at` datetime DEFAULT NULL,
                              `user_name` varchar(255) DEFAULT NULL,
                              `token` varchar(255) DEFAULT NULL,
                              `gps` varchar(255) DEFAULT NULL,
                              `location` varchar(255) DEFAULT NULL,
                              `status` varchar(255) DEFAULT NULL,
                              `tmp` varchar(255) DEFAULT NULL,
                              `concact` varchar(255) DEFAULT NULL,
                              `pcc` varchar(255) DEFAULT NULL,
                              `auth` int DEFAULT NULL,
                              `email` varchar(255) DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              KEY `idx_user_descs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
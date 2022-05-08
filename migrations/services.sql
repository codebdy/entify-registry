CREATE TABLE `services` (
  `id` int unsigned NOT NULL,
  `name` varchar(500) DEFAULT NULL,
  `url` varchar(500) DEFAULT NULL,
  `type_defs` longtext,
  `version` varchar(100) DEFAULT NULL,
  `is_alive` tinyint(1) DEFAULT NULL,
  `added_time` varchar(45) DEFAULT NULL,
  `updated_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `redirects` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL DEFAULT '',
  `shorten` varchar(8) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_url` (`url`),
  UNIQUE KEY `uniq_shorten` (`shorten`),
  KEY `idx_shorten_url` (`shorten`,`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `dispatch` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `counter` bigint(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_counter` (`counter`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


LOCK TABLES `dispatch` WRITE;  
INSERT INTO `dispatch` (`id`, `counter`)
VALUES
  (1, 0);
UNLOCK TABLES;

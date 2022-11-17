-- create "sites" table
CREATE TABLE `sites` (`id` bigint NOT NULL AUTO_INCREMENT, `url` varchar(2048) NOT NULL, `title` varchar(100) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tags" table
CREATE TABLE `tags` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `count` bigint NOT NULL DEFAULT 0, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(100) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "bookmarks" table
CREATE TABLE `bookmarks` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `note` varchar(1000) NOT NULL, `site_id` bigint NOT NULL, `user_id` bigint NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `bookmarks_sites_bookmark` FOREIGN KEY (`site_id`) REFERENCES `sites` (`id`) ON DELETE NO ACTION, CONSTRAINT `bookmarks_users_bookmark` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;

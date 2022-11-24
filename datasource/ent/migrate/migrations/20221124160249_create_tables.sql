-- create "sites" table
CREATE TABLE `sites` (`id` bigint NOT NULL AUTO_INCREMENT, `url` varchar(2048) NOT NULL, `title` varchar(100) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(100) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "bookmarks" table
CREATE TABLE `bookmarks` (`id` bigint NOT NULL AUTO_INCREMENT, `note` varchar(1000) NOT NULL, `site_id` bigint NOT NULL, `user_id` bigint NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `bookmarks_sites_bookmark_from` FOREIGN KEY (`site_id`) REFERENCES `sites` (`id`) ON DELETE NO ACTION, CONSTRAINT `bookmarks_users_bookmarks` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "tags" table
CREATE TABLE `tags` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `count` bigint NOT NULL DEFAULT 0, `user_id` bigint NOT NULL, PRIMARY KEY (`id`), CONSTRAINT `tags_users_tags` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "bookmark_tags" table
CREATE TABLE `bookmark_tags` (`bookmark_id` bigint NOT NULL, `tag_id` bigint NOT NULL, PRIMARY KEY (`bookmark_id`, `tag_id`), CONSTRAINT `bookmark_tags_bookmark_id` FOREIGN KEY (`bookmark_id`) REFERENCES `bookmarks` (`id`) ON DELETE CASCADE, CONSTRAINT `bookmark_tags_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;

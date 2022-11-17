-- modify "tags" table
ALTER TABLE `tags` ADD COLUMN `user_id` bigint NOT NULL, ADD COLUMN `user_tag` bigint NOT NULL, ADD CONSTRAINT `tags_users_tag` FOREIGN KEY (`user_tag`) REFERENCES `users` (`id`) ON DELETE NO ACTION;

-- Create "products" table
CREATE TABLE `products` (`id` bigint NOT NULL AUTO_INCREMENT, `product_id` bigint NOT NULL, `product_name` varchar(255) NOT NULL, `created_on` timestamp NOT NULL, `identifier` char(36) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `product_id` (`product_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

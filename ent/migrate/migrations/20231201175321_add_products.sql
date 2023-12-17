-- Modify "products" table
ALTER TABLE `products` DROP COLUMN `product_id`;
-- Create "product_categories" table
CREATE TABLE `product_categories` (`id` bigint NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

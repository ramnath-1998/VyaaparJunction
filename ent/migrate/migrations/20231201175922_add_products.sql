-- Modify "product_categories" table
ALTER TABLE `product_categories` ADD COLUMN `product_product_category` bigint NULL, ADD INDEX `product_categories_products_product_category` (`product_product_category`), ADD CONSTRAINT `product_categories_products_product_category` FOREIGN KEY (`product_product_category`) REFERENCES `products` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL;

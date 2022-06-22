create database if not exists demo;
use demo;

drop table if exists `layout`;
drop table if exists `layout_hierarchy`; 

CREATE TABLE `layout` (
    `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `type` ENUM('WIDGET', 'LAYOUT') NOT NULL DEFAULT 'LAYOUT',
    `widget_type` VARCHAR(240) DEFAULT '',
    `active` TINYINT(1) DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `WIDGET_TYPE` (`widget_type`)
);

CREATE TABLE `layout_hierarchy` (
    `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `layout_id` BIGINT(11) UNSIGNED DEFAULT NULL,
    `root_layout_id` BIGINT(11) UNSIGNED DEFAULT NULL,
    `parent_layout_id` BIGINT(11) UNSIGNED DEFAULT NULL,
    `position` INT(11) NOT NULL,
    `repeat_count` INT(11) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `LAYOUT_ID` (`layout_id`),
    FOREIGN KEY (`layout_id`)
        REFERENCES `layout` (`id`),
    FOREIGN KEY (`root_layout_id`)
        REFERENCES `layout` (`id`),
    FOREIGN KEY (`parent_layout_id`)
        REFERENCES `layout` (`id`)
);


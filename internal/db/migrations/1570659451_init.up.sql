CREATE TABLE package (
  `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
  `name` VARCHAR(2048) NOT NULL,
  `workflow_id` VARCHAR(255) NOT NULL,
  `run_id` VARCHAR(36) NOT NULL,
  `aip_id` VARCHAR(36) NOT NULL,
  `location` VARCHAR(2048) NOT NULL,
  `status` TINYINT NOT NULL, -- {new, in progress, done, error, unknown}
  `created_at` TIMESTAMP(6) DEFAULT CURRENT_TIMESTAMP(6) NOT NULL,
  `started_at` TIMESTAMP(6) NULL,
  `completed_at` TIMESTAMP(6) NULL,
  PRIMARY KEY (`id`),
  KEY `package_name_idx` (`name`(50)),
  KEY `package_aip_id_idx` (`aip_id`),
  KEY `package_location_idx` (`location`(50)),
  KEY `package_status_idx` (`status`),
  KEY `package_created_at_idx` (`created_at`),
  KEY `package_started_at_idx` (`started_at`)
);
CREATE TABLE preservation_action (
  `id` INT UNSIGNED AUTO_INCREMENT NOT NULL,
  `action_id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(2048) NOT NULL,
  `status` TINYINT NOT NULL, -- {unspecified, complete, processing, failed}
  `started_at` TIMESTAMP(6) NULL,
  `completed_at` TIMESTAMP(6) NULL,
  `package_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`package_id`) REFERENCES package(`id`)
);

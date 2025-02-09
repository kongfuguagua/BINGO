CREATE TABLE `task` (
    `task_id` VARCHAR(255) NOT NULL COMMENT 'task identifier',
    `create_time` DATETIME NOT NULL COMMENT 'creation time',
    `create_user` VARCHAR(255) NOT NULL COMMENT 'creator user',
    `scene_index` VARCHAR(255) NOT NULL COMMENT 'scene index',
    `task_status` VARCHAR(255) NOT NULL COMMENT 'status of the task',
    `update_time` DATETIME NOT NULL COMMENT 'last update time',
    `dataset_index` VARCHAR(255) NOT NULL COMMENT 'dataset index',
    `task_desc` VARCHAR(255) COMMENT 'description of the task',
    `task_name` VARCHAR(255) NOT NULL COMMENT 'name of the task',
    PRIMARY KEY (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
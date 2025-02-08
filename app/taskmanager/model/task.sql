CREATE TABLE task (
    task_id VARCHAR(255) PRIMARY KEY NOT NULL,
    create_time DATETIME NOT NULL,
    create_user VARCHAR(255) NOT NULL,
    scene_index VARCHAR(255) NOT NULL,
    task_status VARCHAR(255) NOT NULL,
    update_time DATETIME NOT NULL,
    dataset_index VARCHAR(255) NOT NULL,
    task_desc VARCHAR(255),
    task_name VARCHAR(255) NOT NULL
);
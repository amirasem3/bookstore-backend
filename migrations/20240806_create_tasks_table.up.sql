-- migrations/20240806_create_tasks_table.up.sql

CREATE TABLE tasks (
                       id INT PRIMARY KEY IDENTITY(1,1),
                       title NVARCHAR(255) NOT NULL,
                       description NVARCHAR(MAX) NOT NULL,
                       completed BIT NOT NULL
);
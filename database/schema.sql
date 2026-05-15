CREATE DATABASE miniwise;
USE miniwise;
CREATE TABLE tickets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    priority VARCHAR(50),
    tag_name VARCHAR(50),
    status VARCHAR(50)
);
INSERT INTO tickets (title, priority, tag_name, status)
VALUES
('Learn React', 'High', 'React', 'Todo'),
('Build Backend', 'Medium', 'Backend', 'In Progress'),
('Design UI', 'Low', 'UI', 'Done');
SELECT * FROM tickets;
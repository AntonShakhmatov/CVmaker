CREATE DATABASE IF NOT EXISTS test
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE test;

CREATE TABLE resumes (
  id INT AUTO_INCREMENT PRIMARY KEY,
  language VARCHAR(10) DEFAULT 'cs',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE headers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  resume_id INT NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  competence VARCHAR(255) NOT NULL,
  title VARCHAR(255),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_header_resume
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE contacts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  resume_id INT NOT NULL,
  phone VARCHAR(255),
  email VARCHAR(255),
  address_field VARCHAR(255),
  github VARCHAR(255),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_contacts_resume
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE experience (
  id INT AUTO_INCREMENT PRIMARY KEY,
  resume_id INT NOT NULL,
  company VARCHAR(255) NOT NULL,
  location_field VARCHAR(255),
  position VARCHAR(255),
  date_from DATE,
  date_to DATE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_experience_resume
    FOREIGN KEY (resume_id) REFERENCES resumes(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE projects (
  id INT AUTO_INCREMENT PRIMARY KEY,
  experience_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_projects_experience
    FOREIGN KEY (experience_id) REFERENCES experience(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE responsibilities (
  id INT AUTO_INCREMENT PRIMARY KEY,
  project_id INT NOT NULL,
  responsibilities_description TEXT NOT NULL,

  CONSTRAINT fk_responsibilities_project
    FOREIGN KEY (project_id) REFERENCES projects(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE technologies (
  id INT AUTO_INCREMENT PRIMARY KEY,
  technologies_name VARCHAR(100) UNIQUE NOT NULL
) ENGINE=InnoDB;

CREATE TABLE project_technologies (
  project_id INT NOT NULL,
  technology_id INT NOT NULL,

  PRIMARY KEY (project_id, technology_id),

  CONSTRAINT fk_pt_project
    FOREIGN KEY (project_id) REFERENCES projects(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_pt_technology
    FOREIGN KEY (technology_id) REFERENCES technologies(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE links (
  id INT AUTO_INCREMENT PRIMARY KEY,
  project_id INT NOT NULL,
  label VARCHAR(100),
  -- url VARCHAR(255) NOT NULL,

  CONSTRAINT fk_links_project
    FOREIGN KEY (project_id) REFERENCES projects(id)
    ON DELETE CASCADE
) ENGINE=InnoDB;
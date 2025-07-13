CREATE TABLE projects(
  id INT PRIMARY KEY AUTO_INCREMENT, -- poderia ser um CUID, pra eu colocar na URL sem perigo do usuário pular de projeto
  category_id INT FOREIGN KEY(category_id) REFERENCES categories(id),
  name VARCHAR(50) UNIQUE NOT NULL,
  description VARCHAR(255) NOT NULL,
  github_url VARCHAR(255) NOT NULL,
  demo_url VARCHAR(255),
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE technologies(
  id INT PRIMARY KEY AUTO_INCREMENT,
  project_id INT FOREIGN KEY(project_id) REFERENCES projects(id)
  category_id INT FOREIGN KEY(category_id) REFERENCES categories(id),
  name VARCHAR(50) NOT NULL
  created_at TIMESTAMP DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE categories(
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50) UNIQUE NOT NULL
);

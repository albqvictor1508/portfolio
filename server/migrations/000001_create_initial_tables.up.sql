-- Esta função é um "helper" para o banco de dados.
-- O objetivo dela é ser chamada automaticamente (através de um TRIGGER)
-- sempre que uma linha em uma tabela for atualizada.
-- Ela simplesmente atualiza o campo `updated_at` para o horário atual.
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE categories(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE projects(
 id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  category_id INT,
  name VARCHAR(100) UNIQUE NOT NULL,
  -- TEXT é mais apropriado para descrições longas do que VARCHAR.
  description TEXT NOT NULL,
  github_url TEXT NOT NULL,
  demo_url TEXT, -- Pode ser nulo se não houver demo.
  -- TIMESTAMPTZ (timestamp with time zone) é a melhor prática para armazenar datas
  -- e evitar problemas com fusos horários.
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,

  -- Definindo a chave estrangeira para conectar com a tabela `categories`.
  CONSTRAINT fk_category
    FOREIGN KEY(category_id) 
    REFERENCES categories(id)
    -- ON DELETE SET NULL: Se uma categoria for deletada, o campo `category_id`
    -- no projeto se tornará NULO, mas o projeto não será deletado.
    ON DELETE SET NULL
);

-- Um TRIGGER é uma "ação automática" que o banco de dados executa.
-- Este trigger chama a função `trigger_set_timestamp` que criamos antes
-- toda vez que um registro na tabela `projects` for atualizado (UPDATE).
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON projects
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- Tabela para armazenar as tecnologias que você usa (Ex: "React", "Go", "Docker").
-- Esta tabela evita a duplicação de nomes de tecnologias.
CREATE TABLE technologies(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE project_technologies(
  project_id UUID NOT NULL,
  technology_id INT NOT NULL,

  -- A chave primária é composta pelos dois IDs.
  -- Isso garante que você não pode adicionar a mesma tecnologia ao mesmo projeto mais de uma vez.
  PRIMARY KEY (project_id, technology_id),

  -- Chave estrangeira que aponta para a tabela `projects`.
  CONSTRAINT fk_project
    FOREIGN KEY(project_id) 
    REFERENCES projects(id)
    -- ON DELETE CASCADE: Se um projeto for deletado, todas as suas associações
    -- com tecnologias nesta tabela também serão deletadas automaticamente.
    ON DELETE CASCADE,

  -- Chave estrangeira que aponta para a tabela `technologies`.
  CONSTRAINT fk_technology
    FOREIGN KEY(technology_id) 
    REFERENCES technologies(id)
    -- ON DELETE CASCADE: Se uma tecnologia for deletada, todas as suas associações
    -- com projetos nesta tabela também serão deletadas.
    ON DELETE CASCADE
);

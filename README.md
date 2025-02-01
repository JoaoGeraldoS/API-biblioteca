# API Biblioteca

Api feita para o gerenciamento e visualização de livros.
Usuarios podem fazer o cadastro e editar seus perfis.

Somente administradores tem permições de adicionar, remover, deletar e atualizar livros, categorias, autores
Administradores também tem permições para deletar usuarios, e gerenciar os proprios usuarios.

### Tecnologias usadas
- Go
- Gin(Framework)
- MySQL

## Rotas
### Livros
Rota onde cria o livro e categorias, relaciona o autor com o livro.
- #### POST:  /admin/books
   - O envio de dados é feito via Formulário
     - Title
     - Author_id "O autor ja existente"
     - Description
     - Content
     - Caregories "Uma ou uma lista"
     - image "Capa do livro"

Rotas onde visualisa os livros. Rota publica e privada
- Rota privada somente administradores
  - ### GET: /admin/books
- Rota acessivel para todos
  - ### GET: /public/books
~~~
     - id
     - Title
     - Description
     - Content
     - Created_at
     - Updated_at
     - Img "caminho ou url do arquivo"
     - Categories "Lista da categoria com id, nome e data de criação"
     - Author_id
     - Author "Dados do author como id, nome, descrição, foto(Caminho ou url)" 
 ~~~
      
Rota de Atualização do livro passado pelo id
- ### PUT /admin/books/:id
~~~
  - Title
  - Author_id
  - Description
  - Content
~~~
  
Rota para apagar um livro passado pelo id
- ### DELETE /admin/books/:id
   - Sem parametros

### Categorias
Rota para criar categorias
- ### POST /admin/categories
~~~
  - Name
~~~

Rota para visualizar
- ### GET /admin/categories
~~~
   - id
   - Name
   - Created_at
~~~

Rota para deletar categoria
- ### DELETE /admin/categories/:id
   - Sem parametros

### Autores
Rota para adicionar autores via formulario
- ### POST /admin/authors
  - Name
  - Description
  - Photo

Rota de visualizar autores
- ### GET /admin/authors
  ~~~
  - id
  - name
  - Description
  - Photo
  ~~~

Rota para deletar author
- ### DELETE /admin/authors/:id
  - Sem parametros


## Usuarios
Rota para criar usuarios
- ### POST /users
~~~
   - Name
   - Email
   - Username
   - Password
~~~

Rota para visualizar usuarios
- ### GET /admin/users
~~~
  - id
  - Name
  - Email
  - Username
  - Bio
  - Photo
  - Created_at
  - Role
  - Updated_at
~~~

Rota para atualizar usuario via formulario
- ### PUT /users/users/:id
  - Name
  - Bio
  - Photo "Foto do usuario"

Rota para deletar usuario
- ### DELETE /admin/users/:id
  - Sem parametros

Rota de login
- ### POST /login
~~~
  - Username
  - Password
~~~

## Dependencias

Para iniciar o projeto use o comando
~~~
go mod init github.com/JoaoGeraldoS/API-biblioteca
~~~


Instalar as dependencias atravez do comando
~~~
go mod tidy
~~~

Execute o comando 
~~~
docker-compose up -d
~~~
para criar o banco, mais as tabelas terão que ser criadas manualmente.

## Tabelas
- Tabela livros
   ~~~
   CREATE TABLE `books` (
     `id` int NOT NULL AUTO_INCREMENT,
     `title` varchar(100) NOT NULL,
     `description` text NOT NULL,
     `content` text NOT NULL,
     `img` varchar(255) DEFAULT NULL,
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `author_id` int DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `author_id` (`author_id`),
     CONSTRAINT `books_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`)
   )
   ~~~
   
- Tabela categorias
   ~~~
   CREATE TABLE `categories` (
   `id` int NOT NULL AUTO_INCREMENT,
   `name` varchar(100) NOT NULL,
   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`)
   )
   ~~~

- Tabela de relacionamento entre categorias e livros
   ~~~
   CREATE TABLE `intermediary` (
  `id` int NOT NULL AUTO_INCREMENT,
  `book_id` int DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `book_id` (`book_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `intermediary_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE,
  CONSTRAINT `intermediary_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE
   )
   ~~~

- Tabela Autores
   ~~~
   CREATE TABLE `authors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `photo` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
   )
   ~~~

- Tabela usuarios
   ~~~
   CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `bio` text,
  `photo` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `role` enum('user','admin') DEFAULT 'user',
  `username` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
   )
   ~~~





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
Rota onde cria o livro e categorias, relaciona o autor com o livro.
- #### POST:  /admin/books
   - O envio de dados é feito via Formulário
     - Title
     - Author_id "O autor ja existente"
     - Description
     - Content
     - Caregories "Uma ou uma lista"
     - image "Capa do livro"

Rotas onde vizualisa os livros. Rota publica e privada
- Rota privada somente administradores
  - ### GET: /admin/books
- Rota acessivel para todos
  - ### GET: /public/books
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

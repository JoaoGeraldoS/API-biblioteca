// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/author": {
            "post": {
                "description": "Cria um novo autor com os dados fornecidos, incluindo name, descrição, foto",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Cria um Autor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Nome do autor",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Descrição do autor",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Foto do autor",
                        "name": "photo",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Author criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Authors"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/authors": {
            "get": {
                "description": "Ler e restorna os dados dos autores",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Ler os autores",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.GenericResponse-models_Authors"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/authors/{id}": {
            "delete": {
                "description": "Apaga um autor fornecido pelo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authors"
                ],
                "summary": "Apaga um autor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recebe o id do autor",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/books": {
            "get": {
                "description": "Ler e restorna os dados dos livros com categorias e autores, e filtros",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Ler os livros",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.GenericResponse-models_Books"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo livro com os dados fornecidos, incluindo título, descrição, conteúdo, autor, categorias e imagem.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Cria um novo livro",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Título do Livro",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Descrição do Livro",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Conteúdo do Livro",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID do Autor",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Categorias do Livro",
                        "name": "categories",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Imagem do Livro",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Livro criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Books"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/books/{id}": {
            "put": {
                "description": "Atualiza um livro fornecido pelo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Atualiza um novo livro",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recebe o id do livro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.Response"
                        }
                    },
                    "400": {
                        "description": "Requisição nao execultada",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro no servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Apaga um livro fornecido pelo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Apaga um novo livro",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recebe o id do livro",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/categories": {
            "get": {
                "description": "Ler e restorna os dados das categorias",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categorias"
                ],
                "summary": "Ler as categorias",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.GenericResponse-models_Categories"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria uma categoria com os dados fornecidos, nome",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categorias"
                ],
                "summary": "Cria uma categoria",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Cria categoria, parametros: Name",
                        "name": "Categories",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Categories"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Categoria criada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Categories"
                        }
                    },
                    "400": {
                        "description": "Requisição nao execultada",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/categories/{id}": {
            "delete": {
                "description": "Apaga uma categoria fornecido pelo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categorias"
                ],
                "summary": "Apaga uma categoria livro",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recebe o id da categoria",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/relation": {
            "post": {
                "description": "Relaciona livro a categoria com os dados fornecidos, incluindo id do livro, id da categoria",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Relaciona livro a categoria",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Recebe o id do livro e o id da categoria",
                        "name": "ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Intermediary"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Relaciona livro a categoria",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Intermediary"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "description": "Ler e restorna os dados dos usuarios com  filtros",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Ler os Usuarios",
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.GenericResponse-models_Users"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}": {
            "delete": {
                "description": "Apaga um usuario fornecido pelo id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Apaga um usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recebe o id do usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Faz o login do usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Faz o login do usuario",
                "parameters": [
                    {
                        "description": "Dados do login do usuário",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuario criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.Response"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public/books": {
            "get": {
                "description": "Ler e restorna os dados dos livros com categorias e autores, e filtros",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Livros"
                ],
                "summary": "Ler os livros",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Execultada com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.GenericResponse-models_Books"
                        }
                    },
                    "404": {
                        "description": "Dados não existentes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Cria um novo usuario com os dados fornecidos, name, email, senha, username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Cria um novo Usuario",
                "parameters": [
                    {
                        "description": "Cria usuario, Parametros: Name, Email, Password, Username, Role",
                        "name": "Users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Usuario criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Users"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/users/{id}": {
            "put": {
                "description": "Atualiza usuarios com os dados fornecidos, name, bio, foto",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Atualiza usuarios",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Nome do usuario",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Descrição do usuario",
                        "name": "bio",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Foto do usuario",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Livro criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/validacao.ResponseGeneric-models_Users"
                        }
                    },
                    "400": {
                        "description": "Erro nos dados fornecidos",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Erro interno do servidor",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Authors": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                }
            }
        },
        "models.Books": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Authors"
                    }
                },
                "author_id": {
                    "type": "integer"
                },
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Categories"
                    }
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "img": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Categories": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "models.Intermediary": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "category_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "validacao.GenericResponse-models_Authors": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Authors"
                    }
                }
            }
        },
        "validacao.GenericResponse-models_Books": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Books"
                    }
                }
            }
        },
        "validacao.GenericResponse-models_Categories": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Categories"
                    }
                }
            }
        },
        "validacao.GenericResponse-models_Users": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Users"
                    }
                }
            }
        },
        "validacao.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "validacao.ResponseGeneric-models_Authors": {
            "type": "object",
            "properties": {
                "items": {
                    "$ref": "#/definitions/models.Authors"
                }
            }
        },
        "validacao.ResponseGeneric-models_Books": {
            "type": "object",
            "properties": {
                "items": {
                    "$ref": "#/definitions/models.Books"
                }
            }
        },
        "validacao.ResponseGeneric-models_Categories": {
            "type": "object",
            "properties": {
                "items": {
                    "$ref": "#/definitions/models.Categories"
                }
            }
        },
        "validacao.ResponseGeneric-models_Intermediary": {
            "type": "object",
            "properties": {
                "items": {
                    "$ref": "#/definitions/models.Intermediary"
                }
            }
        },
        "validacao.ResponseGeneric-models_Users": {
            "type": "object",
            "properties": {
                "items": {
                    "$ref": "#/definitions/models.Users"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API com JWT",
	Description:      "API com autenticação JWT e Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

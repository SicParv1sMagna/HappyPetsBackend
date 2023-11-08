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
        "/api/pet/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает список всех питомцев пользователя.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Получение списка всех питомцев пользователя.",
                "responses": {
                    "200": {
                        "description": "Список питомцев пользователя",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Pet"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Pet"
                            }
                        }
                    }
                }
            }
        },
        "/api/pet/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Создайте нового питомца с предоставленной информацией.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Создание нового питомца.",
                "parameters": [
                    {
                        "description": "Объект PetCreateRequest в формате JSON",
                        "name": "pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PetCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Питомец успешно создан",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    }
                }
            }
        },
        "/api/pet/delete/{petID}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Удаляет питомца по заданному идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Удаление питомца.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID питомца",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Питомец успешно удален"
                    },
                    "400": {
                        "description": "Неверный запрос"
                    },
                    "404": {
                        "description": "Питомец не найден"
                    }
                }
            }
        },
        "/api/pet/image/remove/{petID}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Удаляет изображение из Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Удаляет изображение из Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pet ID",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Изображение успешно удалено",
                        "schema": {
                            "$ref": "#/definitions/model.RemoveImageResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос, неверные входные данные",
                        "schema": {
                            "$ref": "#/definitions/model.RemoveImageResponse"
                        }
                    },
                    "404": {
                        "description": "Изображение не найдено",
                        "schema": {
                            "$ref": "#/definitions/model.RemoveImageResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.RemoveImageResponse"
                        }
                    }
                }
            }
        },
        "/api/pet/image/upload/{petID}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Загружает изображение в Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Загружает изображение в Minio happypets-image(bucket) определенного домашнего животного, связанного с пользователем.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pet ID",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Изображение успешно загружено",
                        "schema": {
                            "$ref": "#/definitions/model.UploadImageResponse"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос, неверные входные данные",
                        "schema": {
                            "$ref": "#/definitions/model.UploadImageResponse"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.UploadImageResponse"
                        }
                    }
                }
            }
        },
        "/api/pet/update/{petID}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Обновляет информацию о питомце с предоставленными данными.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Обновление информации о питомце.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID питомца",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Объект PetUpdateRequest в формате JSON",
                        "name": "pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PetUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Питомец успешно обновлен",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    }
                }
            }
        },
        "/api/pet/{petID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает информацию о питомце по его ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Питомец"
                ],
                "summary": "Получение информации о питомце по его ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID питомца",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Информация о питомце",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    },
                    "404": {
                        "description": "Питомец не найден",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/model.Pet"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Авторизация пользователя и генерация JWT-токена",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователь"
                ],
                "summary": "Вход пользователя",
                "parameters": [
                    {
                        "description": "Данные для входа",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Регистрация нового пользователя с предоставленной информацией.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Пользователь"
                ],
                "summary": "Регистрация нового пользователя.",
                "parameters": [
                    {
                        "description": "Пользовательский объект в формате JSON",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешно зарегистрированный пользователь",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                }
            }
        },
        "/users/{userID}": {
            "get": {
                "description": "Получение данных пользователя по его идентификатору",
                "produces": [
                    "application/json"
                ],
                "summary": "Получить пользователя по идентификатору",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Обновление данных пользователя по его идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Обновить данные пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Pet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "breed_id": {
                    "type": "integer"
                },
                "color": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "food": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lives_at": {
                    "type": "string"
                },
                "lost": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "spicies": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "model.PetCreateRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "food": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lives_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "spicies": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "model.PetUpdateRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "food": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lives_at": {
                    "type": "string"
                },
                "lost": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "spicies": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "model.RemoveImageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.UploadImageResponse": {
            "type": "object",
            "properties": {
                "imageURL": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "repeat_password": {
                    "type": "string"
                }
            }
        },
        "model.UserLoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "model.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "repeat_password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "HappyPets RestAPI",
	Description:      "API server for Native HappyPets application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

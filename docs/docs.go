// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/v1/accept-income": {
            "post": {
                "description": "Принимает id пользователя, id услуги, id заказа, сумму.",
                "tags": [
                    "Posts"
                ],
                "summary": "Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.acceptRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/append": {
            "post": {
                "description": "Принимает id пользователя и сколько средств зачислить.",
                "tags": [
                    "Posts"
                ],
                "summary": "Метод начисления средств на баланс",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.appendRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get-all-transactions/{date}": {
            "get": {
                "description": "На вход: год-месяц. На выходе ссылка на CSV файл.",
                "tags": [
                    "Gets"
                ],
                "summary": "Метод для получения месячного отчета.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "YYYY-Mmm (example: 2022-Nov)",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.transactionListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get-balance/{id}": {
            "get": {
                "description": "Принимает на вход id пользователя",
                "tags": [
                    "Gets"
                ],
                "summary": "Метод получения баланса пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get-reserve/{id}": {
            "get": {
                "description": "get user reserve",
                "tags": [
                    "Gets"
                ],
                "summary": "Принимает на вход id пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get-transactions-by-date/{id}/{limit}/{offset}": {
            "get": {
                "description": "Принимает id пользователя, количество выводимых строк, количество пропускаемых строк.",
                "tags": [
                    "Gets"
                ],
                "summary": "Метод получения списка транзакция пользователя с сортировкой по дате",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of rows",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of skipped rows",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.transactionListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/get-transactions-by-sum/{id}/{limit}/{offset}": {
            "get": {
                "description": "Принимает id пользователя, количество выводимых строк, количество пропускаемых строк.",
                "tags": [
                    "Gets"
                ],
                "summary": "Метод получения списка транзакция пользователя с сортировкой по сумме",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of rows",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "amount of skipped rows",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.transactionListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/reserve-money": {
            "post": {
                "description": "Принимает id пользователя, id услуги, id заказа, стоимость.",
                "tags": [
                    "Posts"
                ],
                "summary": "Метод резервирования средств с основного баланса на отдельном счете",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.reserveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/transfer-money": {
            "post": {
                "description": "Принимает на вход id пользователя-отправителя, id пользователя-получателя, сумму.",
                "tags": [
                    "Posts"
                ],
                "summary": "Метод перевода средств от пользователя к пользователю",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.transferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        },
        "/v1/unreserve-money": {
            "post": {
                "description": "Принимает id пользователя, id услуги, id заказа, стоимость.",
                "tags": [
                    "Posts"
                ],
                "summary": "Метод разрезервирования средств пользователя",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.unreserveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_controller_http_v1.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "avito_internal_controller_http_v1.acceptRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        },
        "avito_internal_controller_http_v1.appendRequest": {
            "type": "object",
            "properties": {
                "sum": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "avito_internal_controller_http_v1.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "avito_internal_controller_http_v1.reserveRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        },
        "avito_internal_controller_http_v1.transactionListResponse": {
            "type": "object",
            "properties": {
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Transaction"
                    }
                }
            }
        },
        "avito_internal_controller_http_v1.transferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "firstUserUUID": {
                    "type": "string"
                },
                "secondUserUUID": {
                    "type": "string"
                }
            }
        },
        "avito_internal_controller_http_v1.unreserveRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        },
        "entity.Transaction": {
            "type": "object",
            "properties": {
                "money_amount": {
                    "type": "integer"
                },
                "operation_date": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.acceptRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.appendRequest": {
            "type": "object",
            "properties": {
                "sum": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.reserveRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.transactionListResponse": {
            "type": "object",
            "properties": {
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Transaction"
                    }
                }
            }
        },
        "internal_controller_http_v1.transferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "firstUserUUID": {
                    "type": "string"
                },
                "secondUserUUID": {
                    "type": "string"
                }
            }
        },
        "internal_controller_http_v1.unreserveRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "orderUUID": {
                    "type": "string"
                },
                "serviceUUID": {
                    "type": "string"
                },
                "userUUID": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Тестовое задание на позицию стажёра-бэкендера от Avito Tech",
	Description:      "Микросервис для работы с балансом пользователей",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

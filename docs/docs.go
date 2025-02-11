// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Amartha",
            "url": "amartha.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/loans": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loans"
                ],
                "summary": "Endpoint for create loan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Accept",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Content-Type",
                        "in": "header"
                    },
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entity.RequestCreateLoan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerResponseOKDTO"
                        }
                    }
                }
            }
        },
        "/v1/loans/pay": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loans"
                ],
                "summary": "Endpoint for create pay loan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Accept",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Content-Type",
                        "in": "header"
                    },
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entity.RequestPayLoan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SwaggerResponseOKDTO"
                        }
                    }
                }
            }
        },
        "/v1/loans/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Loans"
                ],
                "summary": "Endpoint for get loan detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Accept",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Example: application/json",
                        "name": "Content-Type",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Loan Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ResponseGetLoanDetail"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.LoanSchedule": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "loan_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "week_number": {
                    "type": "integer"
                }
            }
        },
        "entity.RequestCreateLoan": {
            "type": "object",
            "properties": {
                "borrowerId": {
                    "type": "integer"
                },
                "interest": {
                    "type": "number"
                },
                "loanAmount": {
                    "type": "number"
                },
                "tenor": {
                    "type": "integer"
                }
            }
        },
        "entity.RequestPayLoan": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "loanId": {
                    "type": "integer"
                }
            }
        },
        "entity.ResponseGetLoanDetail": {
            "type": "object",
            "properties": {
                "borrowerName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "interestRate": {
                    "type": "number"
                },
                "loanAmount": {
                    "type": "number"
                },
                "loanSchedule": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.LoanSchedule"
                    }
                },
                "status": {
                    "type": "string"
                },
                "totalAmount": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entity.SwaggerResponseOKDTO": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string",
                    "example": "Customer Miscellaneous API"
                },
                "build": {
                    "type": "string",
                    "example": "1"
                },
                "data": {},
                "id": {
                    "type": "string",
                    "example": "16ad78a0-5f8a-4af0-9946-d21656e718b5"
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "version": {
                    "type": "string",
                    "example": "1.0.0"
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

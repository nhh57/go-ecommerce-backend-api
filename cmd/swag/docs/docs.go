// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "github.com/nhh57/go-ecommerce-backend-api",
        "contact": {
            "name": "TEAM TIPSGO",
            "url": "github.com/nhh57/go-ecommerce-backend-api",
            "email": "tipsgo@gmail.com"
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
        "/user/register": {
            "post": {
                "description": "When user is registered send otp to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts management"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/update_pass_register": {
            "post": {
                "description": "UpdatePasswordRegister",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts management"
                ],
                "summary": "UpdatePasswordRegister",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePasswordRegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/verify_account": {
            "post": {
                "description": "Verify OTP Login By User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts management"
                ],
                "summary": "Verify OTP Login By User",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VerifyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.RegisterInput": {
            "type": "object",
            "properties": {
                "verify_key": {
                    "type": "string"
                },
                "verify_purpose": {
                    "type": "string"
                },
                "verify_type": {
                    "type": "integer"
                }
            }
        },
        "model.UpdatePasswordRegisterInput": {
            "type": "object",
            "properties": {
                "user_password": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "model.VerifyInput": {
            "type": "object",
            "properties": {
                "verify_code": {
                    "type": "string"
                },
                "verify_key": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "detail": {},
                "error": {
                    "type": "string"
                }
            }
        },
        "response.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8082",
	BasePath:         "/v1/2024",
	Schemes:          []string{},
	Title:            "API Documentation Ecommerce Backend SHOPDEVGO",
	Description:      "This is a sample server caller server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
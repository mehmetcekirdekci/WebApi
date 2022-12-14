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
        "/api/v1/customer": {
            "post": {
                "tags": [
                    "customer"
                ],
                "summary": "Register the customer",
                "parameters": [
                    {
                        "description": "RegisterCustomerRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RegisterCustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseCustomerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseCustomerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BaseCustomerResponse": {
            "type": "object",
            "properties": {
                "responseMessage": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.RegisterCustomerRequest": {
            "type": "object",
            "properties": {
                "adress": {
                    "type": "string"
                },
                "birthDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

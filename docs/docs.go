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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/auth/login": {
            "post": {
                "description": "이메일과 비밀번호로 로그인하고 인증 토큰을 반환합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "사용자 로그인",
                "parameters": [
                    {
                        "description": "로그인 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidAuth"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "새로운 사용자를 등록하고 인증 토큰을 반환합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "사용자 등록",
                "parameters": [
                    {
                        "description": "등록 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/common.ErrNicknameExists"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        },
        "/api/chats/group": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "그룹 채팅방을 생성합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "그룹 채팅 생성",
                "parameters": [
                    {
                        "description": "채팅방 생성 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateGroupChatRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.ChatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        },
        "/api/chats/private/{user1_id}/{user2_id}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "두 사용자 간의 1:1 채팅을 생성합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "1:1 채팅 생성",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "사용자1 ID",
                        "name": "user1_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "사용자2 ID",
                        "name": "user2_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.ChatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        },
        "/api/messages": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "채팅방에 새로운 메시지를 전송합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "메시지 전송",
                "parameters": [
                    {
                        "description": "메시지 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.SendMessageRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        },
        "/api/messages/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "기존 메시지의 내용을 수정합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "메시지 수정",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "메시지 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "수정할 메시지 내용",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateMessageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.ErrUnauthorizedMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "메시지를 삭제합니다",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "메시지 삭제",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "메시지 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInvalidRequest"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.ErrUnauthorizedMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrInternalServer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.AuthData": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "user": {
                    "$ref": "#/definitions/common.UserData"
                }
            }
        },
        "common.ChatData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2024-03-23T12:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "개발팀 채팅방"
                }
            }
        },
        "common.ChatResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "$ref": "#/definitions/common.ChatData"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "common.ErrEmailExists": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 4004
                },
                "data": {
                    "type": "string",
                    "example": "이미 사용중인 이메일입니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.ErrInternalServer": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 5000
                },
                "data": {
                    "type": "string",
                    "example": "내부 서버 오류가 발생했습니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.ErrInvalidAuth": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 4006
                },
                "data": {
                    "type": "string",
                    "example": "이메일 또는 비밀번호가 올바르지 않습니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.ErrInvalidRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 4000
                },
                "data": {
                    "type": "string",
                    "example": "잘못된 요청입니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.ErrNicknameExists": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 4005
                },
                "data": {
                    "type": "string",
                    "example": "이미 사용중인 닉네임입니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.ErrUnauthorizedMessage": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 4009
                },
                "data": {
                    "type": "string",
                    "example": "메시지에 대한 권한이 없습니다"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "common.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "$ref": "#/definitions/common.AuthData"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "common.MessageData": {
            "type": "object",
            "properties": {
                "chatId": {
                    "type": "integer",
                    "example": 1
                },
                "content": {
                    "type": "string",
                    "example": "안녕하세요"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2024-03-23T12:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "senderId": {
                    "type": "integer",
                    "example": 1
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2024-03-23T12:00:00Z"
                }
            }
        },
        "common.MessageResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "$ref": "#/definitions/common.MessageData"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "common.RegisterResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2001
                },
                "data": {
                    "$ref": "#/definitions/common.AuthData"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "common.UserData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2024-03-23T12:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "nickname": {
                    "type": "string",
                    "example": "홍길동"
                }
            }
        },
        "controllers.CreateGroupChatRequest": {
            "type": "object"
        },
        "controllers.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "controllers.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "nickname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "nickname": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2,
                    "example": "홍길동"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "password123"
                }
            }
        },
        "controllers.SendMessageRequest": {
            "type": "object",
            "required": [
                "chatId",
                "content"
            ],
            "properties": {
                "chatId": {
                    "type": "integer",
                    "example": 1
                },
                "content": {
                    "type": "string",
                    "example": "Hello, how are you?"
                }
            }
        },
        "controllers.UpdateMessageRequest": {
            "type": "object",
            "required": [
                "content"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "example": "Updated message content"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "'Bearer ' 접두사와 함께 JWT 토큰을 입력하세요. 예시: \"Bearer eyJhbGciOi...\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "security": [
        {
            "Bearer": []
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5050",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Realtime Chat API",
	Description:      "실시간 채팅을 위한 RESTful API 서버입니다.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
# API Documentation

This document outlines the endpoints for user authentication and data retrieval.

## Table of Contents
1. [Authentication](#authentication)
    - [Sign Up](#sign-up)
    - [Sign In](#sign-in)
    - [Send OTP](#send-otp)
    - [Reset Password](#reset-password)
2. [User Data](#user-data)
    - [Fetch User Data](#fetch-user-data)

## Authentication

### Sign Up

Create a new user account.

- **Endpoint**: `/auth/sign-up`
- **Method**: `POST`
- **Headers**:
    - `Content-Type: application/json`

#### Request Body

```json
{
  "email": "abc@gmail.com",
  "password": "123456789"
}
```
#### Response

```json
{
  "data": {
    "email": "abc@gmail.com",
    "id": 1719040806,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2NTQ0NjgsInN1YiI6MTcxOTA0MDgwNn0.2_H0ft_-cXN2vN6FJ1TORnhmAwBfwpdDbFVDy7nBUiw"
  },
  "message": "User created successfully"
}
```
---
### Sign In

Sign in existing user.

- **Endpoint**: `auth/sign-in`
- **Method**: `POST`
- **Headers**:
    - `Content-Type: application/json`

#### Request Body

```json
{
    "email": "abc@gmail.com",
    "password": "123456789"
}
```
#### Response

```json
{
  "data": {
    "email": "abc@gmail.com",
    "id": 1719040806,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE2NTQ5MjUsInN1YiI6MTcxOTA0MDgwNn0.NhJNYA-69lANbbRnmPRG2Wz2DsNJ5XQur1Hn1BCxhBE"
  },
  "message": "Signed in successfully"
}
```
---

### Send OTP

Send otp to reset password.

- **Endpoint**: `auth/send-otp`
- **Method**: `POST`
- **Headers**:
    - `Content-Type: application/json`

#### Request Body

```json
{
    "email": "abc@gmail.com"
}
```
#### Response

```json
{
  "message": "Otp sent successfully"
}
```
---
### Reset Password

Verify otp and set new password to reset password.

- **Endpoint**: `auth/reset-passeord`
- **Method**: `POST`
- **Headers**:
    - `Content-Type: application/json`

#### Request Body

```json
{
    "email": "abc@gmail.com",
    "otp": "1234",
    "password": "123456789"
}
```
#### Response

```json
{
  "message": "Password reset successful"
}
```

---
### Fetch User Data

Fetch user data.

- **Endpoint**: `user/fetch-user`
- **Method**: `GET`
- **Headers**:
    - `Content-Type: application/json`
    - `Authorization: Bearer {token}`

#### Response

```json
{
  "data": {
    "email": "abc@gmail.com",
    "id": 1719040806
  },
  "message": "User fetched successfully"
}
```
---


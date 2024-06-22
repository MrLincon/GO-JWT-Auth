# Sign Up

Create a new user.

**Endpoint**: `auth/sign-up`

**Method**: `POST`

## Headers

- `Content-Type: application/json`

## Payload

```json
{
    "email": "abc@gmail.com",
    "password": "123456789"
}
```
## Response

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
# Sign In

Sign in existing user.

**Endpoint**: `auth/sign-in`

**Method**: `POST`

## Headers

- `Content-Type: application/json`

## Payload

```json
{
    "email": "abc@gmail.com",
    "password": "123456789"
}
```
## Response

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
# Fetch User Data

Fetch user data.

**Endpoint**: `user/fetch-user`

**Method**: `GET`

## Headers

- `Content-Type: application/json`
- `Authorization: Bearer {token}`

## Response

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


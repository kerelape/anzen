# Anzen API

## Sign up - Secret

Secret sign-up method allows users to sign-up to the server using client secret (password)
and an email address.

If required - it's possible to enable email verification in the config.

```http
POST /v1/signup/secret
```

| Attribute                | Type     | Required | Description           |
|:-------------------------|:---------|:---------|:----------------------|
| `email`                  | string   | Yes      | Email.                |
| `secret`                 | string   | Yes      | Secret (password).    |

If email verification is enabled, returns `202 Accepted` and sends
an email to the user with verification code. The account is created, but
with the `unverified` status.

> **Note**<br/>
> The verification code is only active for a short period of time, defined
> in the config (6 hours by default).

If email verification is **not** enabled,
returns `201 Created` and creates the account.

If a user with this email already exists, returns
`409 Conflict`.

Example request body:
```json
{
    "email": "user@example.com",
    "secret": "abc123"
}
```

## Sign up - verify

Verify is used to verify the email using the verification code.

```http
POST /v1/signup/verify
```

| Attribute                | Type     | Required | Description           |
|:-------------------------|:---------|:---------|:----------------------|
| `code`                   | string   | Yes      | Verification code.    |

If successful, returns `202 Created` and creates the account.

If the verification code has expired, returns `410 Gone`.

If the verification code is not valid, returns `404 Not Found`.

## Sign in - secret

Sign in to an account using client secret (password).

```http
GET /v1/signin/secret
```

| Attribute                | Type     | Required | Description           |
|:-------------------------|:---------|:---------|:----------------------|
| `email`                  | string   | Yes      | Email.                |
| `secret`                 | string   | Yes      | Secret (password).    |

If successful, returns `200 OK` and the following attributes:

| Attribute                | Type     | Description           |
|:-------------------------|:---------|:----------------------|
| `access_token`           | string   | JWT token.            |
| `refresh_token`          | string   | Refresh token.        |

If the provided email and/or the secret are invalid, returns `403 Forbidden`.

Example request body:
```json
{
    "email": "user@example.com",
    "secret": "abc123"
}
```

Example response body:
```json
{
    "access_token": "...",
    "refresh_token": "..."
}
```

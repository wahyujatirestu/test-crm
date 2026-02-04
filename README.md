# Test CRM API – Golang

Backend API sederhana untuk **CRM Membership & Contact** yang dibangun menggunakan **Golang**, **Gin**, **PostgreSQL**, **sqlx**, **JWT**, dan menerapkan **Clean Architecture**.

Project ini dibuat sebagai **technical test / assessment** untuk mengelola data **Membership** dan **Contact** dengan autentikasi berbasis **JWT**.

---

## Tech Stack

| Component        | Tech                               |
| ---------------- | ---------------------------------- |
| Language         | Go (Golang)                        |
| Framework        | Gin                                |
| Database         | PostgreSQL                         |
| ORM / Query      | sqlx                               |
| Authentication   | JWT                                |
| Architecture     | Clean Architecture                 |
| Password Hashing | MD5 _(mengikuti requirement soal)_ |

---

## Project Structure

```
.
├── config/        # konfigurasi env & app
├── controller/    # HTTP handler (Gin)
├── dto/           # request & response DTO
├── middleware/    # JWT middleware
├── models/        # domain model
├── repository/    # database query (sqlx)
├── routes/        # routing API
├── services/      # business logic
├── sql/           # DDL database
├── utils/         # helper (JWT, hash, dll)
├── api.rest       # collection API (VS Code REST Client)
├── server.go      # setup server & dependency injection
├── main.go        # entry point
└── README.md
```

---

## Database Schema

### **Membership**

| Field         | Type      | Note         |
| ------------- | --------- | ------------ |
| membership_id | SERIAL    | Primary Key  |
| name          | VARCHAR   | Not Null     |
| password      | VARCHAR   | MD5 hash     |
| address       | VARCHAR   | Not Null     |
| is_active     | BOOLEAN   | Default true |
| created_date  | TIMESTAMP | Default now  |
| created_by    | VARCHAR   | Not Null     |
| updated_date  | TIMESTAMP | Nullable     |
| updated_by    | VARCHAR   | Nullable     |

---

### **Contact**

| Field         | Type      | Note            |
| ------------- | --------- | --------------- |
| contact_id    | SERIAL    | Primary Key     |
| membership_id | INT       | FK → membership |
| contact_type  | VARCHAR   | email / phone   |
| contact_value | VARCHAR   | Not Null        |
| is_active     | BOOLEAN   | Default true    |
| created_date  | TIMESTAMP | Default now     |
| created_by    | VARCHAR   | Not Null        |
| updated_date  | TIMESTAMP | Nullable        |
| updated_by    | VARCHAR   | Nullable        |

---

## Authentication

- Login menggunakan **contact_value** (email atau phone)
- Password diverifikasi terhadap **password membership**
- Semua endpoint **wajib JWT**, kecuali login

---

## API Endpoints

### Auth

#### Login

```
POST /api/v1/auth/login
```

Request:

```json
{
    "username": "admin@mail.com",
    "password": "password123"
}
```

---

### Membership

#### Create Membership

```
POST /api/v1/memberships
```

#### Get Active Membership with Contacts (JOIN)

```
GET /api/v1/memberships/with-contacts
```

Filter:

- `membership.is_active = true`
- `contact.is_active = true`

Response menampilkan:

- `contact_id`
- `contact_type`
- `contact_value`

#### Get Membership Detail

```
GET /api/v1/memberships/detail/:id
```

---

### Contact

#### Create Contact

```
POST /api/v1/memberships/:membershipId/contacts
```

Request:

```json
{
    "contact_type": "email",
    "contact_value": "restu@mail.com"
}
```

#### Update / Non-Active Contact

```
PUT /api/v1/contacts/:contactId
```

Request:

```json
{
    "contact_value": "restu.updated@mail.com",
    "is_active": false
}
```

Non-active contact dilakukan dengan **soft delete** (`is_active = false`).

---

## Testing Flow (Recommended)

1. Login
2. Create Membership
3. Create Contact (email & phone)
4. Get Membership with Contacts
5. Update / Non-active Contact
6. Get Membership with Contacts (verifikasi filter)

---

## Design Decisions

- Business logic berada di **service layer**, bukan di controller
- Update hanya mengubah field yang dikirim (partial update)
- Soft delete digunakan untuk contact
- Response JOIN menampilkan `contact_id` untuk menghindari ambigu saat update

---

## Run Application

```bash
go run .
```

Server akan berjalan di:

```
http://localhost:8888
```

---

## Notes

- Seeder digunakan untuk **initial login** karena tidak ada fitur registrasi (sesuai soal)
- **MD5 hanya digunakan karena requirement tes**, bukan best practice production

---

## Status

✔ Semua requirement soal terpenuhi
✔ API siap dites via `api.rest`
✔ Siap untuk submission / demo / interview

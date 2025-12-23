# Dressed™ Platform – Microservices Architecture

This repository contains a microservices-based system designed to support
fashion designers and garment suppliers in managing design submissions,
quotations, orders, and payments.

##  Features Implemented

- **Designer Portal**:
  - Upload clothing designs (images/PDFs) under categories (`Men`, `Women`, `Boy`, `Girl`, `Unisex`)
  - Submit designs for quoting
  - View incoming supplier quotes with price, ETA, and negotiation notes
  - Accept a quote to create an order

- **Supplier Portal**:
  - Browse all submitted designs
  - Submit quotes with price, delivery time, and text notes
  - View quote status

- **Admin Portal**:
  - Read-only dashboard with system metrics (users, designs, quotes, orders)

- **Payment Processing**:
  - Mocked payment initiation and confirmation (`PENDING`/`SUCCESS`/`FAILED`)

> **Assumption**: Real-time negotiation happens externally (e.g., email). The platform captures the final quote as structured text.

---

## 🏗️ Architecture Overview

- **Frontend**: React + TypeScript + Vite + Tailwind CSS (role-based portals)
- **Backend**: Golang microservices (modular, containerized)
- **Database**: PostgreSQL (shared instance with isolated tables)
- **Orchestration**: Docker Compose with API Gateway
- **Authentication**: JWT with role-based access (`designer`, `supplier`, `admin`)

### Microservices
| Service | Port | Responsibility |
|--------|------|----------------|
| `auth-service` | `8001` | User registration, login, JWT issuance |
| `design-service` | `8002` | Design uploads, categories, status management |
| `order-service` | `8003` | Quote submission, order creation |
| `supplier-service` | `8004` | Supplier profiles, public design feed |
| `payment-service` | `8005` | Mocked payment processing |
| `gateway` | `8000` | API Gateway (single entry point) |
| `frontend` | `3000` | React SPA (Designer/Supplier portals) |

![System Architecture](docs/architecture.png)

---

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.22+
- Node.js 20+

### Run the System
```bash
# Clone and enter repo
git clone <your-repo>
cd dressed-platform

# Build and start all services
docker-compose up --build
```

### Access Applications
- **Frontend (Designer/Supplier Portal)**: http://localhost:3000
- **API Gateway**: http://localhost:8000
- **Auth Service Swagger**: http://localhost:8000/auth/swagger/index.html
- **Database**: `localhost:5432` (`dressed_user` / `dressed_password`)

### Sample Accounts (Password: `password123`)
| Role | Email |
|------|-------|
| Designer | `designer1@dressed.com`, `designer2@dressed.com` |
| Supplier | `supplier1@dressed.com`, `supplier2@dressed.com` |

> On first run, the database is auto-seeded with sample data (2 designers, 2 suppliers, 6 designs, 4 quotes, 2 orders).

---

## 📁 Project Structure

```
dressed-platform/
├── docs/                     # Architecture diagrams, docs
├── frontend/                 # React SPA (Vite + TS + Tailwind)
├── services/
│   ├── auth-service/         # JWT auth, user management
│   ├── design-service/       # Design uploads, categories
│   ├── order-service/        # Quotes, orders
│   ├── supplier-service/     # Supplier profiles, public feed
│   ├── payment-service/      # Mocked payments
│   └── gateway-service/      # API Gateway (Go)
├── docker-compose.yml        # Full system orchestration
└── README.md
```

Each service follows:
```
service/
├── cmd/
├── internal/
│   ├── handlers/     # HTTP endpoints
│   ├── services/     # Business logic
│   ├── repositories/ # Data access
│   ├── models/       # Data models
│   └── middleware/   # Auth, roles
├── Dockerfile
└── go.mod
```

---

## 🔒 Security & Best Practices

- **Passwords**: Hashed with `bcrypt`
- **Tokens**: JWT with 15-minute expiry, signed secret
- **CORS**: Restricted to `http://localhost:3000`
- **Input Validation**: File type, category, price, ETA
- **Role-Based Access**: Enforced at API and UI layers
- **Error Handling**: Consistent JSON errors, no stack traces
- **Docker**: Minimal Alpine images, multi-stage builds

---

## 🧪 Testing & Reliability

- **Unit tests** for core service logic
- **Idempotent seed data** for reliable demos
- **Health checks** in Docker Compose
- **Graceful error recovery** in frontend (alerts, loading states)

---

## 📄 Documentation

- `docs/architecture.png`: High-level system diagram
- `docs/user-flow.png`: Designer/supplier interaction flow
- `SUBMISSION_DOCUMENT.md`: Full technical report (assumptions, practices, challenges)

---

## 📬 Contact

Built  by **Umesh Vihanga Nethusara Mallawaarachchi (VK)**
Associate Software Engineer | Robotics & Cybersecurity Researcher

# Dressed™ Platform – Microservices Architecture

This repository contains a microservices-based system designed to support
fashion designers and garment suppliers in managing design submissions,
quotations, orders, and payments.

## Architecture Overview
- Frontend: React (role-based portals)
- Backend: Golang microservices
  - Auth Service
  - Design & Quote Service
  - Order & Payment Service
- Database: PostgreSQL
- Containerization: Docker & Docker Compose

## Services
- auth-service: Authentication and role management
- design-service: Design submissions and supplier quotations
- order-service: Order lifecycle and mocked payment processing

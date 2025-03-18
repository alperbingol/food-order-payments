# 🍔 Food Order & Payment API (Golang + MongoDB + Stripe)

## 🚀 Project Overview
This is a **backend API for food ordering and payment processing**, built with **Golang**, **MongoDB**, and **Stripe** for payment integration.

It supports:
- 📦 **Placing food orders**
- 💳 **Processing payments using Stripe**
- 📄 **Retrieving order details**
- 📡 **Webhook support for real-time payment updates**

The API follows **Clean Architecture** to ensure scalability and maintainability.

---

## 📌 Features
✅ **Order Management**
- `POST /order` → Place a food order
- `GET /order/{id}` → Retrieve order details

✅ **Payment Processing**
- `POST /pay` → Process payment for an order
- `POST /stripe/webhook` → Handle Stripe payment events

✅ **MongoDB Storage**
- Orders are stored in **MongoDB** for persistence.

✅ **Stripe API Integration**
- Handles **real-world payment transactions** securely.

✅ **CI/CD with GitHub Actions**
- Automated **testing, build, and deployment**.

---

## 🔥 Tech Stack
- **Backend**: Golang (`gin` for routing)
- **Database**: MongoDB (`go.mongodb.org/mongo-driver`)
- **Payments**: Stripe (`github.com/stripe/stripe-go`)
- **Deployment**: Docker + GitHub Actions

---

## 📅 Development Roadmap

### ✅ **Phase 1: API & Basic Setup**
- [x] Set up **project structure** & **GitHub repo**
- [x] Implement **basic API routes (`POST /order`, `POST /pay`)**
- [x] Install **REST Client for testing**

### 🚀 **Phase 2: MongoDB Integration (Tomorrow)**
- [ ] Set up **MongoDB connection** (`config/config.go`)
- [ ] Implement **order storage in MongoDB**
- [ ] Modify `POST /order` to **save orders in DB**
- [ ] Implement `GET /order/{id}` to **retrieve orders from DB**

### 💳 **Phase 3: Stripe Payment Processing**
- [ ] Integrate **Stripe API** for real payments
- [ ] Implement `POST /pay` to **charge users**
- [ ] Handle **payment success/failure with webhooks**

### 🛠 **Phase 4: Deployment & CI/CD**
- [ ] Create **Dockerfile** for containerization
- [ ] Set up **GitHub Actions for CI/CD**
- [ ] Deploy to **AWS Lambda or DigitalOcean**


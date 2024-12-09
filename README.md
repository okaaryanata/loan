# Loan Service

A service for managing loans, repayments, and user accounts built using Go and deployed with Docker and Kubernetes.

## Prerequisite Requirements

Before running this project, ensure you have the following tools installed:

1. **Golang** (version 1.22.10 or higher) - for building the application.
2. **Docker** - to containerize the application.
3. **Kubernetes** (with Minikube recommended for local deployments) - to manage containerized services.
4. **kubectl** - to interact with Kubernetes clusters.\
5. **Git** - for version control.
6. **PostgreSQL** - as the database backend.

---

## Project Setup and Usage

### 1. Clone the Repository

```bash
git clone https://github.com/okaaryanata/loan.git
cd loan
```

### 2. Run local

- make sure you already have **PostgreSQL** in your local machine, and create schema `loans`
- create file **.env** base on file **.env.example**

```
go run cmd/laon.main.go
```

### 2. Build and Run with Docker

```bash
# Build the Docker image
docker build -t loan-service .

# Run the Docker container
docker run -d --name loan-service -p 9901:9901 --env-file .env loan-service
```

### 3. Kubernetes Deployment

Ensure Minikube is running, then:

```bash
# Create a ConfigMap for application configuration
kubectl apply -f /deployment/kube-manifest/configmap.yaml

# Deploy the service
kubectl apply -f /deployment/kube-manifest/deployment.yaml

# Deploy the service
kubectl apply -f /deployment/kube-manifest/deployment.yaml
```

---

## API Endpoints

### Service Route

**`{{url}}/svaha-loan`**

### User Endpoints

| Method | Endpoint                     | Description                                              |
| ------ | ---------------------------- | -------------------------------------------------------- |
| `POST` | `/user`                      | Create a user                                            |
| `GET`  | `/users/{id}`                | Get active user by ID                                    |
| `GET`  | `/user?username=okaaryanata` | Get list active user, optional: query param **username** |

### Loan Endpoints

| Method | Endpoint                       | Description                        |
| ------ | ------------------------------ | ---------------------------------- |
| `POST` | `/loan`                        | Create a loan & repayment schedule |
| `GET`  | `/loan/{loanID}/user/{userID}` | Get loan by ID and userID          |
| `GET`  | `/loan/user{userID}`           | Get loans by userID                |

### Repayment Endpoints

| Method | Endpoint                                     | Description                                                                          |
| ------ | -------------------------------------------- | ------------------------------------------------------------------------------------ |
| `POST` | `/repayment`                                 | Make a repayment, will automatically update outstanding_balance, missed_payments etc |
| `GET`  | `/repayment/schedule/user/{userID}?loanID=1` | List all user repayments or specific loanID                                          |

## [Postman Document](https://documenter.getpostman.com/view/7748154/2sAYBd77xw)

## Environment Variables

The application uses the following environment variables (defined in the `.env` file):

```plaintext
APP_HOST=localhost
APP_PORT=9090
DB_INIT_TABLE=true
DB_HOST=
DB_PORT=5432
DB_USER=
DB_PASSWORD=
DB_NAME=svaha
DB_SSL_MODE=disable
```

---

## License

This project is licensed under the MIT License.

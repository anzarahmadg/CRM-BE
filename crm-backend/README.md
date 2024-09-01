## Features
- User Management
- Customer Management
- Interaction Tracking
- Role-Based Access Control
- Dockerized Application


### Prerequisites
- Docker
- Go 1.20+

### Setup
1. Clone the repository
2. Run `docker-compose up` to start the application

### API Documentation
- **POST /api/users** - Create a new user
- **GET /api/users** - Get all users
- **PUT /api/users/:id** - Update a user
- **DELETE /api/users/:id** - Delete a user

- **POST /api/customers** - Create a new customer
- **GET /api/customers** - Get all customers
- **PUT /api/customers/:id** - Update a customer
- **DELETE /api/customers/:id** - Delete a customer

- **POST /api/interactions** - Create a new interaction
- **GET /api/interactions/customer/:customerId** - Get interactions for a specific customer
- **PUT /api/interactions/:id** - Update an interaction
- **DELETE /api/interactions/:id** - Delete an interaction

## Database Schema
- Users: { _id, name, email, password, role }
- Customers: { _id, name, contact_info, company, status, notes }
- Interactions: { _id, customer_id, type, notes, scheduled, resolved }

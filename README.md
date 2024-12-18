
---

# Blog Post Management System  

## Overview  
The Blog Post Management System is a microservices-based application that enables users to create, read, update, and delete blog posts. It includes user management and authentication features, making it secure and scalable. This project is built with Go, PostgreSQL, and the Gin framework.  

---

## Architecture  
Below is the microservice architecture for the system:

(https://github.com/Simret101/Blog_Post_Management_System/blob/main/assets/microservice_architecture%20(1).pdf)

---

## Folder Structure  

This repository is organized into three main services: Blog Service, User Service, and Auth Service. Below is the folder structure for the project:  

```plaintext
blog_post/
├── blog_service/
│   ├── controllers/        
│   ├── models/             
│   ├── repositories/       
│   └── services/           
│
├── user_service/
│   ├── controllers/        
│   ├── models/             
│   ├── repositories/       
│   └── services/           
│
└── auth_service/
    ├── controllers/            
    ├── models/                 
    ├── repositories/           
    └── services/
```  

---

## Features  

- **Blog Management**: Create, read, update, and delete blog posts.  
- **User Management**: Handle user registration, profile updates, and role-based access control.  
- **Authentication**: OAuth-based authentication with token management.  

---

## Technology Stack  

- **Programming Language**: Go  
- **Framework**: Gin  
- **Authentication**: OAuth  

---

## Setup Instructions  

1. Clone the repository:  
   ```bash
   git clone <repository-url>
   ```  
2. Navigate to the project directory:  
   ```bash
   cd Blog_Post_Management_System
   ```  
3. Install dependencies:  
   ```bash
   go mod tidy
   ```  
4. Set up the database:  
   - Configure PostgreSQL.  
   - Run the database migrations for each service.  
5. Start each service:  
   ```bash
   go run blog_service/main.go  
   go run user_service/main.go  
   go run auth_service/main.go  
   ```  
6. Test the endpoints using Postman or any API testing tool.  

---

## Contributing  

1. Fork the repository.  
2. Create a feature branch:  
   ```bash
   git checkout -b feature/your-feature
   ```  
3. Commit your changes:  
   ```bash
   git commit -m "Add your message here"
   ```  
4. Push to the branch:  
   ```bash
   git push origin feature/your-feature
   ```  
5. Create a pull request.  

---

## API Documentation  

Detailed API documentation can be found [here](https://www.postman.com/blogposts/api-fest-nutrition/api/b3821248-5ee4-4a1b-aa5a-7ea179465db6/blog-post-management-api?action=share&creator=40134617).  

---


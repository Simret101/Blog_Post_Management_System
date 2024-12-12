# Blog Post Management System

## Overview
The Blog Post Management System is a microservices-based application that enables users to create, read, update, and delete blog posts. It includes user management and authentication features, making it secure and scalable. This project is built with Go, PostgreSQL, and the Gin framework.
Blog Post Management System Folder Structure:
Blog_Post_Management_System/
│
├── blog_service/        
│   ├── controllers/     
│   │   └── blog_controller.go 
│   ├── models/          
│   │   └── blog.go     
│   ├── repositories/    
│   │   └── blog_repository.go  
│   └── services/        
│       └── blog_service.go  
│
├── user_service/        
│   ├── controllers/     
│   │   └── user_controller.go  
│   ├── models/          
│   │   └── user.go      
│   ├── repositories/    
│   │   └── user_repository.go  
│   └── services/        
│       └── user_service.go  
│
└── auth_service/        
    ├── controllers/     
    │   └── auth_controller.go 
    ├── models/          
    │   └── auth.go      
    ├── repositories/    
    │   └── auth_repository.go  
    └── services/        
        └── auth_service.go  
---

## Features
- Blog Management: Create, read, update, and delete blog posts.
- User Management: Handle user registration, profile updates, and role-based access control.
- Authentication: OAuth-based authentication with token management.

---

## Folder Structure

### Root Directory
Blog_Post_Management_System/
The root directory contains the following subdirectories, each representing a core service:

### Blog Service
Responsible for managing blog-related operations.
blog_service/
├── controllers/      
│   └── blog_controller.go  
├── models/           
│   └── blog.go        
├── repositories/      
│   └── blog_repository.go 
└── services/          
    └── blog_service.go  

### User Service
Manages user-related functionalities like registration, profile management, and role-based access control.
user_service/
├── controllers/       
│   └── user_controller.go  
├── models/            
│   └── user.go        
├── repositories/      
│   └── user_repository.go  
└── services/          
    └── user_service.go  

### Authentication Service
Handles user authentication and token management using OAuth.
auth_service/
├── controllers/       
│   └── auth_controller.go  
├── models/           
│   └── auth.go        
├── repositories/      
│   └── auth_repository.go  
└── services/         
    └── auth_service.go  

---

## Technology Stack
- Programming Language: Go
- Framework: Gin
- Authentication: OAuth

---

## Setup Instructions
1. Clone the repository:
      git clone <repository-url>
   
2. Navigate to the project directory:
      cd Blog_Post_Management_System
   
3. Install dependencies:
      go mod tidy
   
4. Set up the database:
   - Configure PostgreSQL.
   - Run the database migrations for each service.
5. Start each service:
      go run blog_service/main.go
   go run user_service/main.go
   go run auth_service/main.go
   
6. Test the endpoints using Postman or any API testing tool.

---

## API Documentation
Detailed API documentation can be found [here](https://www.postman.com/blogposts/api-fest-nutrition/api/b3821248-5ee4-4a1b-aa5a-7ea179465db6/blog-post-management-api?action=share&creator=40134617).

---

## Contributing
1. Fork the repository.
2. Create a feature branch:
      git checkout -b feature/your-feature
   
3. Commit your changes:
      git commit -m "Add your message here"
   
4. Push to the branch:
      git push origin feature/your-feature
   
5. Create a pull request.

---

## License
This project is licensed under the MIT License.

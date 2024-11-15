# Order Processing System

The Order Processing System is a simple backend solution for managing customers, products, and orders in an e-commerce application. It is built in Golang using PostgreSQL as the database, GORM for ORM, and Gin for handling API requests.

# Setup Instructions

1. Clone the Repository:
  ```git clone https://github.com/lakshay88/lakshay_OrderProcessingSystem.git```

2. Set Up Configuration Variables: 
  Open ```config.yaml``` file update configuration according to your flexiblity else leave it as it is.

3. Install Dependencies:
  ```go mod tidy```

4. Build Docker Image:
  ```docker-compose up -d```

5. Run application 
  ```go run main.go```


# API Endpoints

1. `/api/customers`	GET	Retrieve all customers.
2. `/api/customers/{id}`	GET	Retrieve details for a specific customer, including their orders.
3. `/api/orders`	POST	Create a new order for a customer by specifying customer ID and product IDs.
4. `/api/orders/{id}`	GET	Retrieve order details, including the total price.
5. `/api/products`	POST to add products.
6. `/api/customers`	POST create customers.


# PostMan Collection 
In repo you will find ```order-serving-system.postman_collection.json``` It is a Post man collecation of all the API. Import if in yoyr system and test it. 

# Features

1. `Data Models`: Customer, Product, and Order models represent the core entities, with relationships established using GORM.
2. `Database`: PostgreSQL is used with GORM for object-relational mapping and database migrations.
3. `API Endpoints`: RESTful API endpoints for creating, updating, and retrieving data on customers, products, and orders.
4. `Business Logic`: Includes validation rules, such as preventing a customer from placing a new order if a previous one is unfulfilled.
5. `Condiguration Based system`: It is having a config.yaml file where you can step up your system config.
6. `Gateway service`: It implement server configuration 
7. `Multiple Database support`: You can add different database into your system by defining its methods in database configuration.
8. `Docker`: Currently system is can be run using docker no need to install SQL in system.  

# Future Improvements
  Add user authentication and authorization.
  Add test case(currently not)
  Improve validation and error handling with more detailed messages.
  We can segirigate each service in microservice.
  Add asynchronous processing for order-related tasks.

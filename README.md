# Implementing CRUD in GoLang REST API

## Jubo_intern_RESTful_CRUD
intern_HomeWork
Use Golang & GORMux framework to construct a RESTful API with CRUD function

# Documentation
## How to run the project
1. clone this repository in a root folder
2. install the Golang (1.18.1 version or above)
3. open the "config.json" in root directory, and change "connection_string" to your development machine. I recommand you to use Mysql and follow the string rule by:  
"[username]:[pw]@tcp([HOST]:[port])/[databaseName]"
![connection_string](https://user-images.githubusercontent.com/78125534/232267566-0f177704-581a-47f9-952e-b497b122ec89.png)
4. run "go build ." and "go run ."in CLI, you will see    
![image](https://user-images.githubusercontent.com/78125534/232269109-5263bf4c-3555-4851-91af-3bb69b379377.png)
5. you have started the web server sucessfully!

## Interact with the Web server
### Use "Postman" to interact with the webserver:![image](https://user-images.githubusercontent.com/78125534/232278096-9e382e63-d29b-4f22-a0fe-2cad8b49a649.png)
![get](https://user-images.githubusercontent.com/78125534/232272516-d8f0bb52-866a-4665-9f33-18b350dd3c17.png)

1. [GET]/api/todos：  
Use HTTP verb "GET" and type the url "http://localhost:8080/api/Todos", then click "send" and you will see the total todos:
![get_2](https://user-images.githubusercontent.com/78125534/232272667-c7360462-0748-4b79-8e7e-3f346b236e34.png)

---
2. [GET]/api/todos/: id:  
Use HTTP verb "GET" and type the url "http://localhost:8080/api/Todos/:id", then click "send" and you will see the single todos:
![get_3](https://user-images.githubusercontent.com/78125534/232272992-e8b2803e-1117-4a2d-8c43-e6146e8e3e6e.png)
---
3. [POST]/api/todos：  
Use HTTP verb "POST" and type the url "http://localhost:8080/api/Todos", next type your new Todo information with json in "body", then click "send" and you will see your new todo:
![POST](https://user-images.githubusercontent.com/78125534/232276856-c3660880-9576-41c0-b5a6-2c1857d8f74f.png)
Use HTTP verb "GET" and type the url "http://localhost:8080/api/Todos", then click "send" and you will see the total todos with new Todo:
![POST_2](https://user-images.githubusercontent.com/78125534/232276918-2bf1f5e3-d480-4023-ab8e-e07261c60dd4.png)
---
4. [PUT]/todos/:id：  
Use HTTP verb "PUT" and type the url "http://localhost:8080/api/Todos/:id" with id of Todo you want to edit, next type your editting information with json in "body", then click "send" and you will see your editted todo:
![PUT](https://user-images.githubusercontent.com/78125534/232277414-b3abf261-0435-4eaa-b929-b8768c6ff85a.png)  
Use HTTP verb "GET" and type the url "http://localhost:8080/api/Todos", then click "send" and you will see the total todos with eddited Todo:
![PUT_2](https://user-images.githubusercontent.com/78125534/232277554-e231bf96-ad62-4f63-af7e-bf99468ac2d5.png)
---
5. [DELETE]/api/todos/:id：  
Use HTTP verb "DELETE" and type the url "http://localhost:8080/api/Todos/:id" with id of Todo you want to delete, then click "send" and you will see "Todo deleted successfully":
![DELETE](https://user-images.githubusercontent.com/78125534/232277849-ae1c3359-5902-49d3-add4-13fa20bf2c84.png)
Use HTTP verb "GET" and type the url "http://localhost:8080/api/Todos", then click "send" and you will see the total todos without deleted Todo:
![DELETE_2](https://user-images.githubusercontent.com/78125534/232277975-f3fd197b-88a7-4816-8055-05472c91c777.png)







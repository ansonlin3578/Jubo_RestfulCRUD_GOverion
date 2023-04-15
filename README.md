# Implementing CRUD in GoLang REST API

## Jubo_intern_RESTful_CRUD
intern_HomeWork
Use Golang & GORMux framework to construct a RESTful API with CRUD function

# Documentation
## How to run the project
1. clone this repository in a root folder
2. install the nodeJS (12.2.2 version or above)
3. run "node main.js" in CLI, you will see "Port 3000 in Listening!!"  
![image](https://user-images.githubusercontent.com/78125534/232205410-ebeac99b-826c-4cf5-bb52-7d8e5d07300b.png)  
4. you have started the web server sucessfully!

## Interact with the Web server
1. [GET]/todos：  
Open the chorme tab and type the url "http://localhost:3000/Todos", you will see the Home Page:
![image](https://user-images.githubusercontent.com/78125534/232205800-956c7404-92f5-4c69-9912-bbe6bf5bce2b.png)
---
2. [GET]/todos/: id:  
Click "details" behind the Description each Todo:
![get](https://user-images.githubusercontent.com/78125534/232206046-bcf306a2-410e-408f-bcb1-848701e58fac.png)  
You will route to show.ejs, and the page look like:
![image](https://user-images.githubusercontent.com/78125534/232205894-039a0ab3-fde9-45fd-bd39-560e811f61cc.png)
---
3. [POST]/todos：  
Click the link "Add new Todo" below the home page:
![post](https://user-images.githubusercontent.com/78125534/232206245-cef0c8d3-f5be-4099-b1f1-dc6bf6a28bfe.png)  
You will route to new.ejs, and the page look like:
![image](https://user-images.githubusercontent.com/78125534/232206292-3fea4d2c-3b08-4099-a806-28528445a1a4.png)  
Type the information of your new Todo, then click "submit" button:
![post_2](https://user-images.githubusercontent.com/78125534/232206377-77c2f208-b053-4c90-b50d-68d32e6efb5d.png)  
You will route to home page with new Todo you create:
![post_3](https://user-images.githubusercontent.com/78125534/232206415-70664357-958d-49ff-a1dc-b7345652617a.png)
---
4. [PUT]/todos/:id：  
Click "details" of Todo you want to edit:
![patch](https://user-images.githubusercontent.com/78125534/232206567-5bfaf2a3-6c8a-4cad-ad25-ffd86f0c113d.png)  
Click the link "Edit Todo information", you will route to edit.ejs page:  
![patch_5](https://user-images.githubusercontent.com/78125534/232207437-8f31fa0c-1257-42be-8873-058dba96eeb6.png)
![image](https://user-images.githubusercontent.com/78125534/232206678-545e2a46-47a0-447c-9fc8-f64a9629302a.png)  
Edit the information you want, the click "save" button:
![patch_3](https://user-images.githubusercontent.com/78125534/232206756-f7757cd6-6f57-451f-a65b-3046e8fb5f1d.png)  
You will route to the home page with editted Todo:
![patch_4](https://user-images.githubusercontent.com/78125534/232206815-0a1dbe9f-8223-4dbd-8b70-095d5e9bf54a.png)
---
5. [DELETE]/todos/:id：  
Click "details" route to the single Todo information, and click "Delete" button:  
![Screenshot from 2023-04-15 16-11-52](https://user-images.githubusercontent.com/78125534/232205705-c60ef72f-bcd9-4307-aaad-2aef2ddba355.png)
![Screenshot from 2023-04-15 16-12-09](https://user-images.githubusercontent.com/78125534/232205710-0bfa916e-6e9b-41bc-9254-91699b9d3dc2.png)  
then it will route to the Home page without deleted Todo:  
![Screenshot from 2023-04-15 16-14-31](https://user-images.githubusercontent.com/78125534/232205731-80c35976-f116-4a59-9b85-99f41ce4927b.png)






# Go HTTP Server

Basic Go application that uses CRUD operations to get and manipulate data from MySQL.

## Create MySQL DB and Table
Basic setup instructions are described in `mysql.md`.

## Build Go App and Start 
1. Create root directory  
`$ mkdir Go-App`  
`$ cd Go-App`
2. Set which port to use in main.go (8000 by default)  
`app.Run("localhost:<port>")`
3. Build the App  
`$ go build`  
4. Run executable directory  
`$ ./Go-App`  

## Perform Basic CRUD Operations 
**You can use your Browser for GET, or software such as Postman or Insomnia for POST, PUT, and DELETE**  
* Defined in `handleRoutes()` in app.go ln 152-158  
1. GET all   
**URL**: `localhost:8000/products`  <br><br>
2. GET Product if "id" = x  
**URL**: `localhost:8000/product/1`  <br><br>  
3. Create a new Product  
**Method**: `POST`   
**URL**: `localhost:8000/product`    
**JSON Body**:
```json
{  
   "name": "table",  
   "quantity": 10,  
   "price": 24.99  
}
```
  *Returns*:  
  ```json
{  
   "id": 3,  
   "name": "table",  
   "quantity": 10,  
   "price": 24.99  
}
  ```
<br><br>
4. Update an existing Product  
**Method**: `PUT`    
**URL**: `localhost:8000/product/3`    
**JSON Body**:  
```json
{  
   "name":"table",  
   "quantity":10,  
   "price":18.99  
}
```
*Returns*:  
```json
{  
   "name":"table",  
   "quantity":10,  
   "price":18.99  
}
```
<br><br>
5. Delete a Product  
**Method**: `DELETE`    
**URL**: `localhost:8000/product/3`  
<br>
*Returns*: "Result: Deleted Successfully."  

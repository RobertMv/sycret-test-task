# sycret-test-task

- Add a config file(.env) with the next absolute path C:/test/config/.env
- All generated docs will be saved in the C:/test/ directory
- API request is http://localhost:PortFromYourEnvFile/api/docs/generate
- Body is 
```
 {
    "URLTemplate": "https://sycret.ru/service/apigendoc/forma_025u.doc",
    "RecordId": 30
 }
```
# Example from Postman
Here is an example from Postman
![image](https://user-images.githubusercontent.com/36698814/172055743-f3b168a1-5f43-41a1-b10c-bdc912cb5001.png)

# .env file contains a port to listen
![image](https://user-images.githubusercontent.com/36698814/172057271-713a1823-f16d-40b7-a11e-f1c50ea7da96.png)

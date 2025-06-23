# Go Authentication API

This project is a simple modular REST API built with Golang and Gin framework.  
It handles basic **user authentication** functionalities such as:

- User registration
- User login
- Forgot password with OTP (simulated in terminal)
- Password reset

The project is structured using MVC (Model-View-Controller) pattern with clean code separation.

---


## How to Clone and Use

Make sure you have Golang installed on your device.

#### 1. Clone the repository
```
git clone https://github.com/ranandasatria/fgo24-be-auth.git
```

#### 2. Navigate into the project directory
```
cd fgo24-be-auth
```

#### 3. Run the program
```
go run main.go
```

#### 4. Example HTTP request

Register:
```
POST /auth/register
Content-Type: application/x-www-form-urlencoded

name=Rananda&email=ran@mail.com&password=1234
```
Login:
```
POST /auth/login
Content-Type: application/x-www-form-urlencoded

email=ran@mail.com&password=1234
```
Forgot Password:
```
POST /auth/forgot-password
Content-Type: application/x-www-form-urlencoded

email=ran@mail.com
```
Reset Password:
```
POST /auth/reset-password
Content-Type: application/x-www-form-urlencoded

email=ran@mail.com&otp=123456&newPass=abcd&confPass=abcd
```

## 📄 License

This project is licensed under the **MIT License**.  

## ©️ Copyright

&copy; 2025 Kodacademy
# JWE Manager

JWE Manager is a HTTP server who expose tree main functionalities: 

  - Create a pair of asymmetrical keys (RSA)
  - Encrypt data using the JEW specification (RFC 7516)
  - JWE decrypt

The main technologies used to built this project was **GoLang** and **Redis**. The project is structured in layers, which layer has there own responsibility.
# JWE Manager

JWE Manager is a HTTP server who expose five main functionalities: 

  - Create a pair of asymmetrical keys (RSA)
  - Get the public key
  - Encrypt data using the JWE specification (RFC 7516)
  - JWE decrypt
  - Health check

The main technologies used to built this project was **GoLang** and **Redis**. The project is structured in layers, which layer has there own responsibility.

TODO:
 - Gin Recovery strategy
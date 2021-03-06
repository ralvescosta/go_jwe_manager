# JWE Manager

[![CircleCI](https://circleci.com/gh/ralvescosta/go_jwe_manager/tree/main.svg?style=svg)](https://circleci.com/gh/ralvescosta/go_jwe_manager/tree/main)
[![codecov](https://codecov.io/gh/ralvescosta/go_jwe_manager/branch/main/graph/badge.svg?token=PCFNRLU9Y3)](https://codecov.io/gh/ralvescosta/go_jwe_manager)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ralvescosta_go_jwe_manager&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ralvescosta_go_jwe_manager)

JWE Manager is a HTTP server who expose five main functionalities: 

  - Create a pair of asymmetrical keys (RSA)
  - Get the public key
  - Encrypt data using the JWE specification (RFC 7516)
  - JWE decrypt
  - Health check

The main technologies used to built this project was **GoLang** and **Redis**. The project is structured in layers, which layer has there own responsibility.
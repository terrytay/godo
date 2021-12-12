# "Todo" Microservice Example

[![codecov](https://codecov.io/gh/terrytay/godo/branch/main/graph/badge.svg?token=DIB9PG9PUL)](https://codecov.io/gh/terrytay/godo)

## Project Motivation

The motivation for this project comes from the need to understand microservices and production systems. There is a big difference between monolith applications and microservices. Monolith applications have services that live in the same repository and communicates directly with each other. Microservices seek to decouple the services such that it increases scalability, reliability and maintainability in a large production system.

With that, the architecture has to be sound and carefully thought after. After which, observability has to be used as it is one aspect that the system must have in order for development and debugging.


This project is created for educational and learning purposes. As such, it does not represent an accurate system that can be used for production.

## Introduction

This application is a Todo application that aims to create itself in a microservice architectural way. The ideas and code largely come from [Mario Carrion](https://mariocarrion.com). The business logic is kept simple to focus on designing the microservice.

The following are the things this project seeks to achieve.

- Hexagonal Architecture
- Domain Driven Design
- more to be added...

## Getting Started

Prerequisites:

- Docker
- Docker Compose

To run the webserver:

```sh
docker-compose up --build
```

## Plans

- [ ] OpenTelemetry
- [ ] Redis
- [ ] Memcache
- [ ] Apache Kafka
- [ ] ELK
- [ ] HTTPS & Certificate
- [ ] Better CICD
- [x] Basic Hexagonal Architecture
- [x] Docker and compose scripts needed to run
- [x] Connection to Postgresql

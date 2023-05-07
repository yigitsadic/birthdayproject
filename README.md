# 🎁 Birthday Project

Birhday project is a project for companies to send birthday emails to their employees on their birthdays automatically.

## Tech Stack

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Rails](https://img.shields.io/badge/rails-%23CC0000.svg?style=for-the-badge&logo=ruby-on-rails&logoColor=white)
![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=for-the-badge&logo=vite&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![Next JS](https://img.shields.io/badge/Next-black?style=for-the-badge&logo=next.js&logoColor=white)

```mermaid
---
title: CakedayToday Stack
---

classDiagram

API --> Postgres Database
Web App --> API
Admin App --> Postgres Database
Admin App --> Mongo Database
Email Sender --> Postgres Database
API --> RabbitMQ
Admin App --> RabbitMQ
RabbitMQ --> Audit Logger
Audit Logger --> Mongo Database
Static Site --> API

API : Golang
Web App : React + TypeScript
Email Sender : Golang
Audit Logger : Golang
Admin App : Ruby on Rails
Static Site : Nextjs

```

## Flows

```mermaid
---
title : Audit Logging Flow
---

sequenceDiagram
    API ->> RabbitMQ : Publishes row changes via messages
    Admin App ->> RabbitMQ : Publishes row changes via messages

    Audit Logger ->> RabbitMQ : Fetches row changes via messages
    Audit Logger ->> RabbitMQ : Fetches row changes via messages
    Audit Logger ->> Mongo DB : Writes changes to

    Admin App ->> AuditLogger : Fetches row changes via HTTP
```


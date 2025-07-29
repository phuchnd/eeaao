# eeaao
Everything Everywhere All at Once

## Introduction
This is my playground repository where I experiment with new technologies, concepts, and ideas. It serves as both a sandbox for learning and a knowledge base to help me remember what I've learned.

## Purpose
- **Experimentation**: A safe space to try out new programming languages, frameworks, and tools
- **Learning**: Document my learning journey and discoveries
- **Reference**: Store code snippets, examples, and notes for future reference

## Organization
This repository is structured as a monorepo, containing multiple services written in various programming languages. Each service or project is organized in its own directory, allowing for:

- **Multiple Languages**: Services written in different programming languages coexist in the same repository
- **Independent Development**: Each service can be developed and deployed independently
- **Shared Resources**: Common utilities and configurations can be shared across services
- **Centralized Management**: All code is managed in one place, simplifying version control

## Repository Structure
```
.
├── backing-services/       # Infrastructure components
│   ├── kafka/              # Kafka message broker service
│   └── mariadb/            # MariaDB database service
└── services/               # Application services
    └── go/                 # Go language services
        ├── common/         # Shared utilities and libraries
        ├── rule-engine/    # Business rules processing service
        ├── user/           # User management core service
        └── user-api/       # User-facing API service
```

The repository is organized into two main sections:

- **backing-services**: Contains infrastructure components that support the application services
  - Kafka for event streaming and message processing (with its own docker-compose.yml)
  - MariaDB for relational database storage (with its own docker-compose.yml)
  - Each backing service has its own Docker Compose file in its respective directory

- **services**: Contains the actual application services, currently focused on Go implementations
  - Common utilities and shared code
  - Rule engine for business logic processing
  - User management service
  - User API for external interactions

## Current Explorations
*This section will be updated as new experiments are added*

---
Last updated: 2025-07-29

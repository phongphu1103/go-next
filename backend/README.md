Repositories talk to DB models

Services talk to domain entities

Controllers/handlers map domain → response DTOs

#### Below is the final architecture, where:

# Handler

Validates input

Parses query / params / body

Shapes HTTP response (status, envelope)

# Controller

Orchestrates use cases

Calls services

Converts domain → DTO

# Service

Business logic

# Repository

Model → Domain conversion
# Medical API

The Medical API is a RESTful API designed to manage patients and doctors for a clinic or medical facility.

## Table of Contents

- [Medical API](#medical-api)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
- [Create a new patient](#create-a-new-patient)
- [Retrieve all patients](#retrieve-all-patients)
- [Update a patient record](#update-a-patient-record)
- [Delete a patient record](#delete-a-patient-record)

## Introduction

The Medical API provides endpoints for creating, retrieving, updating, and deleting patient and doctor records. It aims to streamline the management of patient information and facilitate communication between medical staff and patients.

## Features

- **Patient Management**: Create, retrieve, update, and delete patient records.
- **Doctor Management**: Create, retrieve, update, and delete doctor records.
- **Authentication and Authorization**: Secure endpoints with authentication and role-based access control.
- **Validation**: Validate input data to ensure data integrity and consistency.
- **Documentation**: Detailed documentation for API endpoints.

## Installation

To run the Medical API locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/TechLateef/medical_api.git
    ```

2. Navigate to the project directory:

    ```bash
    cd medical_api
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Set up the database:

    - Create a PostgreSQL database.
    - Update the database configuration in the `.env` file.

5. Build and run the application:

    ```bash
    go build
    ./medical-api
    ```

## Usage

Once the Medical API is running, you can interact with it using HTTP requests. You can use tools like cURL or Postman to send requests to the API endpoints.

For example:

```bash
# Create a new patient
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "password": "password", "Email": "john@gmail.com"}' http://localhost:8080/api/v1/patients

# Retrieve all patients
curl http://localhost:8080/api/v1/patients

# Update a patient record
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Jane Doe"}' http://localhost:8080/api/v1/patients/1

# Delete a patient record
curl -X DELETE http://localhost:8080/api/v1/patients/1



API Endpoints

The Medical API exposes the following endpoints:

Patients:
POST /api/v1/patients: Create a new patient.
GET /api/v1/patients: Retrieve all patients.
GET /api/v1/patients/{id}: Retrieve a specific patient.
PUT /api/v1/patients/{id}: Update a patient.
DELETE /api/v1/patients/{id}: Delete a patient.
Doctors:
POST /api/v1/doctors: Create a new doctor.
GET /api/v1/doctors: Retrieve all doctors.
GET /api/v1/doctors/{id}: Retrieve a specific doctor.
PUT /api/v1/doctors/{id}: Update a doctor.
DELETE /api/v1/doctors/{id}: Delete a doctor.
For detailed documentation of each endpoint, refer to the API documentation.


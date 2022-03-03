# Employee Asset Management

Application to facilitate requests and monitoring of employee assets in the company.
Facilitate the process of checking asset availability, requesting asset loans, and monitoring assets, where company assets will be supervised by the admin and approved by the manager.

<!-- Aplikasi untuk memudahkan permintaan dan monitoring asset karyawan pada perusahaan.
Memudahkan proses pengecekan ketersediaan asset, permintaan peminjaman asset, dan monitoring asset, dimana asset perusahaan akan di supervise oleh admin dan di setujui oleh manager. -->

## Acknowledgements

- [Wireframe](https://drive.google.com/drive/folders/146w2Xl-i1A9sDIYG0D2PrpAfKN8rtFGa?usp=sharing)
- [UI](https://www.figma.com/file/RIUf4ssTfGccwwN3zRpANU/Capstone)
- [ERD](https://drive.google.com/file/d/1BoHoIu2F-heUm3OG4G03r613qPgjvVGT/view?usp=sharing)
- [TaskDocs](https://docs.google.com/presentation/d/1nQiBvc3hWbIXPMPp0cOSmURZxOKcfCy1SHTHbQK9ZU0/edit#slide=id.g115a0663133_0_137)

## Tech Stack

- Go Language
- Godotenv
- Labstack Echo
- Mysql
- Docker
- AWS

## Features

- Bcrypt
- JWT
- REST API
- CICD (Github Action)
- AWS S3 Bucket
- Role Base Login User
- Create, Read, Update Application
- Create, Read, Update Procurement
- Create, Read Asset
- Read, Update Item
- Read Employee

- **Scope of Limitation:**

- Category : Laptop, Alat Tulis Kerja, Kendaraan, Lainnya
- Role : Employee, Admin, Manager

## Documentations

- [Swagger Open API](https://app.swaggerhub.com/apis-docs/iswanulumam/EventPlanningApp/1.0.0)

## Deployment

- https://beeldy.site/api/v1

## Authors

- [@eldyhidayat](https://www.github.com/eldyhidayat)
- [@najibjodiansyah](https://www.github.com/najibjodiansyah)

## Environment Variables

To run this project, you will need to set up the following environment variables to your .env file

## Run Locally

Clone the project

```bash
  git clone https://github.com/najibjodiansyah/group5-capstone-project.git
```

Go to the project directory

```bash
  cd group5-capstone-project
```

Install dependencies

```bash
  go mod tidy
  go get
```

Setup .env from template.env

```bash
    nano/vim template.env
```

Start the server

```bash
  go run ./app/server.go
```

## Running Tests

To run tests, run the following command

```bash
  go test ./...
```

## Used By

This project is used by the following companies:

- Group 5 Avenger Employee Management Asset

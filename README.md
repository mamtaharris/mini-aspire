# Mini Aspire

It is an app that allows authenticated users to go through a loan application. It doesn’t have to contain too many fields, but at least “amount
required” and “loan term.” All the loans will be assumed to have a “weekly” repayment frequency.
After the loan is approved, the user must be able to submit the weekly loan repayments. It can be a simplified repay functionality, which won’t
need to check if the dates are correct but will just set the weekly amount to be repaid.

## Getting Started

Follow below steps to get the service running on your local machine.

### Prerequisites

- Golang
- PostgreSQL

### Local setup

Open terminal in your project root directory and run below command

```bash
make server
```

## Project Overview

- cmd - This package contains different commands that the service provides. Currently, we can start server, migrate the database and seed some user info.

- config - This directory contains configuration files or code related to your application's configuration, such as database connections or environment variables.

- internal - This directory is added at the root level and contains packages that should not be imported or accessed by external packages. It can include packages like internal models, repositories, services, or any other internal implementation details.

- pkg - This directory is used to store packages that can be potentially reused across multiple applications. In this case, we have two subdirectories: /database, /logger, /server and /middlewares.

- main.go - This is the entry point of the application which calls all the initialization scripts and starts the http server.

- Makefile - This file defines the set of tasks to be executed.

## Makefile

Run the application

```bash
make server
```

Migrate the database

```bash
make migrate
```

Seed the database with user info

```bash
make seed
```

Run linters

```bash
make lint
```

Run all tests

```bash
make test
```

Generates mock files for unit testing

```bash
make mock
```

## Go API on EKS Logging with Logrus and Fluent Bit on AWS CloudWatch
This repository illustrates a simple `Go` microservice that allows a user to register, login, add and query books from their library. The API makes use of `JWT` to control authorization for the Add-a-book and List-books endpoints. All data in `PostgreSQL` database is persisted in a PostgreSQL container. The Go `Logrus` library is used for structured logging to record API access and transactional activities.

### Local Deployment
1. Grab the repository
   ```sh
   $ git clone https://github.com/bensonmacharia/book_store_api.git
   ```
2. Configure the environment variables
   ```sh
   $ cd book_store_api
   $ cp .env.example .env
   ```
   Make sure to edit the variables inside <<>> to set up your custom configuration.
3. Run the application
   ```sh
   $ docker-compose up -d
   #Confirm the containers have been created and are running
   $ docker ps
   ```

### Deployment on AWS
![AWS Deployment](https://miro.medium.com/v2/resize:fit:800/format:webp/1*CC1RYKj7nXI6WwETVpXMjw.png)
> Follow this guide to deploy this application on AWS - https://lnkd.in/d3qTv74h
> 
> The guide will illustrate how to deploy the application on `AWS EKS`, persist the database in an `AWS EFS` and ship logs to `AWS CloudWatc`h with the aid of the `Fluent Bit` daemon.
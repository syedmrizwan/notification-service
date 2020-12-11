# Notification Service

# Overview
This repo contains implementation of Notification Service

![Architecture](https://github.com/syedmrizwan/notification-service/blob/master/images/architecture.png)


# Running the project
```
cd deployment
docker-compose up -d --build
```

# API documentation
After running the project, navigate to [Swagger Docs](http://localhost:8080/swagger/index.html) to view the APIs

# Testing the Project
Open logs for Email Handler
`docker logs -f email-handler`

In another terminal, POST a notification with type Email using cURL
`curl -X POST "http://localhost:8080/api/v1/bulk-notifications" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"notification_mode\": \"Email\", \"notification_text\": \"Promotional Message\", \"user_ids\": [ 1,2,3,4,5 ]}"`

Notification detail will be displayed on the Email Handler log(Email vendor can use the message for further processing)

# Todos
Add a service for Notification Tracking

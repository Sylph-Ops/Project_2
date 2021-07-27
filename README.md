# Project_2
Simple API bank/wallet

### Task:

 develop a microservice for personal expenses management (wallet);

 the service should provide users with the following functionality:

• crediting funds;

• write-off of funds;

• obtaining a balance sheet;

• obtaining a list of transactions for the selected period of time.


 the service must support different currencies;

 the service must maintain a write-off limit for one day for an individual
user. The limit value must be passed to the application using a parameter
configuration at startup;

 The parameter must be used in the queries to identify the user
user_id string type. User authentication is not required;

 each transaction (crediting, writing off) must have a field with a description;

 example of transaction parameters: user_id, amount, currency, description;

 The service must provide HTTP API and receive / send requests / responses in the format
JSON;

 the service must return user-understandable errors in case of invalid requests,
insufficient balance of funds, exceeding the write-off limit.

Requirements:
 any development language, preference is given to Go;

 you can use any frameworks and libraries;

 data should be stored in the database of the candidate's choice;

 The solution must be provided as a github repository with a readme file from
startup instructions and sample query / response examples.

# Management System

Having fairly simple requirements, the client needed a web-based application that can manage their AWS infrastructure, that was friendlier and cleaner than the AWS console. 

### User Interface

The user interface needed an area to create, view and edit for both users and their "queue buster" domain. The user interface also needed a dashboard/homepage which displayed a list of "active" and "disabled" queue busters. 

A login page, a change password page, error and maintanence pages were also required. The error and maintenance pages were just basic HTML and CSS to display to a user that a particular event has occured.

User's needed to be able to change their own password and logout. On every page in the interface, there is an "Account" dropdown which provides an option to the user, to change their password - as well as, "Logout".

### API

A RESTful API was required to allow the user interface to manage the client's infrastructure; this needed to be done with AWS Lambda which is Amazon's approach to serverless. 

The RESTful API was built using a collection of AWS Lambda functions through an API Gateway.

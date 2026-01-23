# Bookings Web Application

## Description

This is a simple web application for handling hotel bookings.
A go backend is used and is inspired by following along
tsawler's go web applications udemy course. His own repository
can be found at https://github.com/tsawler/bookings-app

This project is mostly focused on showcasing the golang 
backend features including the following backend methods...
* Creation and handling of http routes
* Creation and implementation of Middleware for routes
* Serving static files to the website
* Implementing a template rendering function for go tmpl files
* Integrating javascript and css to be served as static files
* Writing tests for all backend functions
* Connecting to a Postgress database to persist user data
* Implementing sending mail notifications
* Authenticating users and setting up an Admin user

Anywhere along the way I implemented any personal changes on
my end based on my own knowledge of web design/development.
Such as

* Using tailwindcss for generating the styles of the website
* Using Maizzle to generate styled email templates
* Docker for generating the PostgreSQL database

## Getting Started

### Dependencies

* [scs](https://github.com/alexedwards/scs) v2.9.0 (for session management)
* [chi](https://github.com/go-chi/chi) v1.5.5 (for a routing framework)
* [nosurf](https://github.com/justinas/nosurf) v1.2.0 (for CSRF protection)
* [tailwindcss](https://github.com/tailwindlabs/tailwindcss) v4.1.18 (as a css framework tool)
* [govalidator](https://github.com/asaskevich/govalidator) v11.0.1 (for form validation)
* [pgx](https://github.com/jackc/pgx) v5.8.0 (PostgreSql Driver)
* [Go Simple Mail](https://github.com/xhit/go-simple-mail) v2.16.0 (for handling smtp email)
* [Maizzle](https://github.com/maizzle/maizzle/tree/master) v5.4.0 (for email templating)

### Installing manually

TODO: Add more information on installing and running the source code as the project progresses

### Authors

Max Kranker

## Version History

* 0.1
  * Initial Release

## License

This project is licensed under the MIT License - see the LICENSE.md file for details

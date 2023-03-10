# Project Proposal
Goal Tracker. Users can keep track of their daily, weekly and even monthly goals. 

## Technologies Required
- No third party API require at the moment
- DB outline - https://dbdiagram.io/d/64056790296d97641d859c47

## Data Model
- User table, id bigserial [pk], first_name varchar, last_name varchar, created_at datetime
- Daily goal table, id bigserial [pk], discriptions varchar, completed bool, time_created datetime, user_id int [fk]
- Monthly goal table, id bigserial [pk], discriptions varchar, completed bool, time_created datetime, user_id int [fk]
- Weekly goal table, id bigserial [pk], discriptions varchar, completed bool, time_created datetime, user_id int [fk]
- possibly another table added but I'm not sure what it will hold

## Roadmap

### MVP
- User can view/edit/delete their daily goals
- Track each completed and incomplete goals
 

### 2.0
- Create some form of a countdown if the goal is not completed to notify the user
- Implement a user and password interface for security

### Nice-to-haves
- Evolve into a website or phone app

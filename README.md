# Appointment-Manager

- Appointment Booking app typically for managers with a lot of appointments throughout the day 
- Used go fiber framework, gorm for database models (ORM) with sqlite, jwt for authentication
- The app has a very simple UI for users written in HTML, CSS and js

  # How the app works:
- The user first identifies the position as there are different pages for each user:
- The employees(receptionist or secretary) have more options as they create or update the appointments
- Also, the employee views all appointments as the app is designed to have more than one person as a manager
- So, one employee can manage appointments for many managers
- The managers only can see the appointments assigned to them and then react to them by approving, rejecting or updating
- when a secretary creates an appointment a notification will appear to the manager to notify of a new appointment


- There are also some advanced searching or filtration the user can make like:
- showing only the approved appointments, declined appointments
- all the appointments from one specific time to another 
- all the appointments related to a specific client and more

# Employee page (for creating appointments)
![employee page](https://github.com/Mostafa-Elshahawy/Appointment-Booker_Go/assets/147338664/440185a1-2445-4916-97bc-bf524f0130d6)

# Manager page (approval or canceling appointments)
![manager page](https://github.com/Mostafa-Elshahawy/Appointment-Booker_Go/assets/147338664/3e019098-edec-4ad3-b103-4ec4cd62ce7d)

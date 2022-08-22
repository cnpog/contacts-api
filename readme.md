
## Challenge
Create the „Contacts API“. It’s a simple API, where a user can get a quick overview over all contacts resources like person, skills...
The following use cases should be implemented:
#### UC1

Create an CRUD endpoint for managing contacts. A contact should have at least the following attributes and appropriate validation:
- Firstname
- Lastname
- Fullname
- Address
- Email
- Mobile phone number

#### UC2
Create a CRUD endpoint for skills. A contact can have multiple skills and a skill can belong to multiple contacts. A skill should have the following attributes and appropriate validation:
- Name
- Level (expertise)

#### UC3
Document your API with Swagger.

#### UC4 (optional)
Implement the following security aspects. All bullet points are optional and can be implemented partially. • Authentication
- Authorization
- Users can only change their contact
- Have checks for skills changes (e.g. the current user can’t change skills for other users)
## Notes
- oauth is for authorization
- user grants access on his resources to a third party (me as a user, I grant access to my google calendar to a time tracking app)
- Flow:
    - user clicks on "login with google" on the time tracking app (authentication)
    - time tracking app redirects user to google
    - google asks user if he wants to grant access to his calendar through consent screen (authorization)
    - user grants access
    - google redirects user to time tracking app
    - time tracking app can now access the user's calendar

## Setup google oauth
- create a project in the google cloud console
- api & services -> oauth consent screen
    - application type: web application
    - application homepage link: http://localhost:8080
    - authorized domains: localhost
    - authorized redirect URIs: http://localhost:3000/auth/google/callback
    - scopes: 
        - https://www.googleapis.com/auth/drive: View and manage the files in your Google Drive
        - https://www.googleapis.com/auth/userinfo.email
        - https://www.googleapis.com/auth/userinfo.profile
- api & services -> credentials
    - create credentials -> oauth client id
    - application type: web application
    - name: Web client 1
    - authorized javascript origins: http://localhost:8080
    - authorized redirect URIs: 
        - http://localhost:8080/callback
        - http://localhost:8080/drivecallback

## Q/A
- Given app only needs email/userinfo access scopes. What happens after server retrieving userinfo?
    - How server authenticates subsequent requests?
    - Do server needs to persist userinfo?
    - Do server needs to persist access token/refresh token? When?
- How passport-oauth2 handles this process?
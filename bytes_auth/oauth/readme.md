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
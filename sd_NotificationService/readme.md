# Notification Service

## `templates` table
```sql
CREATE TABLE `templates` (
    code TEXT PRIMARY KEY,
    description TEXT,
    channel TEXT,
    title TEXT
    body TEXT
);
INSERT INTO `templates` (code, description, channel, title, title, body)
VALUES 
    ('WELCOME_NEW_USER', 
    'Welcome new user', 
    'email', 
    'Welcome {{.Username}}',
    'Hi {{.Username}}. Great to see you here!'),
    ('HAPPY_BIRTHDAY',
    'Happy birthday user',
    'sms',
    'Happy birthday {{.FirstName}}'
    'Today is {{.CurrentDate}}. Happy birthday {{.FirstName}}');
```
# Personalised Email
A simple script for me to send personalised emails to my students, addressed by name.

## Required environment variables
```dotenv
SMTP_HOST=
SMTP_PORT=
SMTP_USERNAME= # because my NUS email is forwarded to Gmail
SMTP_PASSWORD= # gmail has app password, use it
FROM=
TEST=true # set to false after testing
```

## Example Email
message.yaml
```yaml
subject: "[CS2040C] Welcome to B07"
html_file: "welcome.html"
```
welcome.html
```html
<h2>Welcome to CS2040C!!</h2>
```
users.yaml
```yaml
- fullname: John Smith # not used
  name: John
  email: a@b.c

- fullname: Jane Maria Doe
  name: Jane
  email: x@y.z
```
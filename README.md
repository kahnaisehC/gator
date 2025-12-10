# gator 

Gator is a CLI that allows the users to follow feeds and display them as a list in the terminal

## Installation
### Requirements
- Postgresql 15.13
- Go 1.23



### Configuration
gator stores its configuration on ~/.gatorconfig.json
it looks like this:

```js
{
    "db_url":"postgres://postgres:<your-password>@<your-host>:<your-port>/gator?sslmode=disable",
    "current_user_name":""
}
```

here, db_url is the connection url for the database
current_user_name represents the user that is currently logged in the app. Don't worry about it for now and just leave it empty

### Database
gator needs access to a postgres database with the schema represented in the /sql/schema directory
```sql



```


### Installation




## Usage






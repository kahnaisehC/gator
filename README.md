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
gator needs access to a postgres database with the following schema:
```sql
CREATE TABLE users (
	id UUID PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	name TEXT UNIQUE NOT NULL
);

CREATE TABLE feeds(
	feed_id SERIAL PRIMARY KEY,
	name VARCHAR(1024) NOT NULL,
	url VARCHAR(1024) UNIQUE NOT NULL,
	user_id UUID REFERENCES users ON DELETE CASCADE NOT NULL,
	last_fetched_at TIMESTAMP
);

CREATE TABLE feed_follows (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID REFERENCES users ON DELETE CASCADE NOT NULL,
	feed_id INTEGER REFERENCES feeds ON DELETE CASCADE NOT NULL,
	UNIQUE(user_id, feed_id)
);

CREATE TABLE posts (
	id SERIAL PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	title VARCHAR(1024) NOT NULL,
	url VARCHAR(1024) NOT NULL UNIQUE,
	description TEXT,
	published_at TIMESTAMP,
	feed_id INTEGER REFERENCES feeds NOT NULL
);
```

### Installation




## Usage






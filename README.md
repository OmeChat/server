<div align="center">
    <h1>OmeChat server</h1>
    <hr>
    <strong>
    The official backend of the OmeChat chat app.
</strong><br><br>
<img src="https://img.shields.io/github/license/OmeChat/server?style=for-the-badge">
<img src="https://img.shields.io/github/go-mod/go-version/OmeChat/server?style=for-the-badge">
<img src="https://img.shields.io/github/last-commit/OmeChat/server?style=for-the-badge">
</div>

<div align="center">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" height="100">
</div>
<hr>

# information

OmeChat is an open-source chat app. It is inspired by the platform Omegle. I always wanted to get into chat apps and 
securing them the right way. Therefore I started this project. In my opinion it is a great opportunity for finding new
people at the same age in an anonymous and secure way. You can search for people at your age with a maximum tolerance
of 5 years and send them messages. It is an easy way to meet new people and find new friends.

# privacy and anonymity

OmeChat does not require any personal data. The only data that is required to create an account and use OmeChat is an
username (for example a pseudonym) and your age. OmeChat is not logging any data and only saves random IDs as your
userHash (userID) and your clientHash (clientID). The messages are end-to-end encrypted. But unlike other chat apps
the encryption layer only bases on the client instead of the server. That means the server does not have any private 
encryption keys and is not able to read any messages. 

# personal security

personal security means the security of the own identity and dignity. But also the protection against pedophiles 
and other "bad" people is part of the personal security. OmeChat implements or is going to implement different systems
to guarantee your personal security. There are systems like the ATSS (age-tolerance-security-system) or the
MLBS (multi-layer-banning-system) to do exactly this.

# performance and traffic reduction

I am doing everything to improve the performance and reduce the network traffic of the OmeChat backend.
For example, I am using different caching technics to improve the performance. All message blocks of chats
are cached into the RAM of the server to improve the backend performance and reduce the number of file reads.

# Special concepts

- <strong>Secrets instead of passwords:</strong>  Because many people are using insecure and
  bad passwords, I am using secrets for the user verification. The secret is a 64 character long
  random string used as the user passwords. The user is not choosing his password itself.
  Therefore it could not be insecure. The secret acts less than a password and more like a
  token. This is the reason why it is called `secret`. The secret is completely hashed with
  the argon2ID algorithm. To put it in a nutshell, the secret concept has a much higher security
  level, because it prevents bad passwords and social engineering attacks. Furthermore it
  is like every other password.

<hr>

# stored data

The data is stored in two different files. I am using files at this point, because a database is too overkill for
the tiny amount of data I am storing. 

All the users are stored into a file called `user.json`
```json
"<userHash>": {
    "clients": [
      "<clientHash>"
    ],
    "age": 16,
    "secret": "WvXXIGTwxLxsYIQguXUbIQDlmMiplmLJ",
    "username": "Mathis Burger1"
  }
```
This is the basic structure of a user stored into the `user.json` file.

| key      	| explainastion                                                                         	|
|----------	|---------------------------------------------------------------------------------	|
| userHash 	| Identifier of the user. It is used for all identification processes of the user 	|
| clients  	| An Array of all clients (hashes) existing for the user                          	|
| age      	| The age of the user                                                             	|
| secret   	| The secret (password) of the user                                               	|
| username 	| the unique username of the user                                                 	|

All the clients are stored into a file called `client.json`
```json
"<clientHash>": {
    "owner": "<userHash>",
    "access_token": "EDCtNaTadxEpKMibFXSkLCfiBOKwlIBZMpIzOTywGJovSLFFYXgrTBTocuePNWFl"
  }
```
This is the basic structure of a user stored into the `client.json` file.

| key          	| explaination                                                  	|
|--------------	|---------------------------------------------------------------	|
| clientHash   	| The clientHash is the identifier of the client                	|
| owner        	| The hash-identifier of the user who owns the client           	|
| access_token 	| The accessToken of the user required to access the websocket. 	|
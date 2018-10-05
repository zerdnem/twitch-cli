# twitch-cli
commandline interface for twitch

# Dependencies
```
go get github.com/catsby/go-twitch/service/kraken
go get gopkg.in/AlecAivazis/survey.v1
```

# How to get twitch api access_token
```
create new app https://glass.twitch.tv/console/apps/create

curl -X GET 'https://id.twitch.tv/oauth2/authorize?response_type=token+id_token&client_id=YOUR_APP_CLIENT_ID&redirect_uri=http://localhost&scope=user_read+openid&state=c3ab8aa609ea11e793ae92361f002671'

click result link

login then authorize app

get acccess_token from url

export TWITCH_ACCESS_TOKEN="YOUR ACCESS TOKEN"
```

# Screenshot
![ss](https://github.com/zerdnem/twitch-cli/blob/master/carbon.png)

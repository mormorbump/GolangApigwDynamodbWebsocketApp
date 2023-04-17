# API Gateway Websocket with Golang

## infra Architecture

- Golang
- API Gateway(WebSocket)
- API Gateway(RestAPI)
- AWS lambda
- Terraform

## Architecture

層の外側から順に

- src
  - lambda <- ここが実質のルート。
    - connection
    - disconnection
    - send_message
  - infrastructure
    - router
    - handlers
  - adapter
    - controllers
    - database
    - interfaces
    - protocols
  - usecase
    - interfaces <- controllerやrepositoryをDIする用
  - domain
    - models
  
  
## init

- terraform/config.tfを設定
- terraform/variables.tfの通りにterraform/terraform.tfvarsを設定

```
cd app
make init
```

## Build and Deploy

```
cd app
make deploy
```

## Verification

use [wscat](https://github.com/websockets/wscat)

### connection

stringParameterでtoken(Bearer), roomId, userId, ownerIdを取りたいので以下のようにパラメタつける 

```
wscat -c wss://<api-id>.execute-api.<region>.amazonaws.com/<stage>?<params1=value1&params2=value2...>

```

### send message

こっちはRestAPIなのでbodyに諸々情報詰める
t
`POST https://qb5tb8b4n3.execute-api.ap-northeast-1.amazonaws.com/dev/send_message`

リク
```
{
    "platformUserId": 111,
    "roomId": "XXXXXXXXXXXXX",
    "message": "hello world",
    "name": "dummy", 
    "imageUrl": "<cdnUrl>",
}
```

# muse

Telegram бот для сбора песен от участников в плейлист на дискотеку.

## Начало работы
1. Установите docker
2. Скопируйте ```compose.yml```
3. Создайте папку ```settings``` там же, где и ```compose.yml```
4. Добавьте данные от Google Cloud Auth, нужен scope="https://www.googleapis.com/auth/spreadsheets" (```settings/credentials.json``` и ```settings/token.json```)
5. Добавьте ```settings/config.json```. Смотрите примеры авторизации на Go от Google Cloud Auth. 
```json
{
    "db": {
        "redis": "redis:6379",
        "sheet": "google sheet id"
    },
    "bot": {
        "token": "bot token",
        "webhook": {
            "url": "webhook url",
            "secret": "webhook secret"
        },
        "admins": [ # telegram ids of admins
            7952812
        ]
    },
    "yandex": {
        "token": "Some oauth token",
        "userId": 7952812
    },
    "tracksLimit": 7,
    "groups": [
        {
            "name": "group name",
            "playlistId": "yandex music playlist uuid",
            "sheetListName": "google sheets list name"
        }
    ]
}
```
### ВАЖНО
Если вы не хотите использовать webhook, то поставьте DEBUG=TRUE в переменных среды
# Nectar Landing Page Backend

a simple backend for nectar landing page

## Routes

|      API Path       | Method |     What it does      |
| :-----------------: | :----: | :-------------------: |
|   /api/characters   |  GET   |   Fetch characters    |
| /api/character/{id} |  GET   | Fetch character by id |
|   /api/fantasies    |  GET   |    Fetch fantasies    |
|  /api/fantasy/{id}  | DELETE |  Fetch fantasy by ID  |

## Environment Variables

create .env file then set the ff

|    ENV     |     Description     |
| :--------: | :-----------------: |
|  DB_HOST   | postgresl hostname  |
|  DB_PORT   |   postgresql port   |
|  DB_USER   | postgresql username |
|  DB_PASS   | postgresql username |
|  DB_NAME   | postgresql DB name  |
| DB_SSLMODE |      SSL Mode       |
|  APP_PORT  |      App port       |

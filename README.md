# test-task

Сбилдить базу и сервер в докере можно командой `make docker-build`<br>Предварительно надо переименовать `docker/local.env.example` в `docker/local.env` и проставить там желаемые параметры

Так же есть ряд других команд для работы с докером, типа `make docker-logs/docker-rebuild` и прочие, а так же для других
задач, вроде генерации grpc файлов<br
Все остальные Make команды так же сделаны, как указано в доке

Для корректной работы линтера необходимо переименовать `.golangci.yml.example` в `.golangci.yml` и проставить там
желаемые настройки линтинга

Выполнение миграций делается командой `make migrate`, обратите внимание, что предварительно надо переименовать
`env.example` в `local.env`, если вы же вы запускаете сервис в докере, то миграции будут исполнены автоматически

Что касается тестов, ключевой функционал сложно выделить в такой небольшой задаче, поэтому я покрыл тестом требуемый метод

Так же не очень понял, каждое сохранение курса должно новой строкой в базу ложиться или заменять предыдущее?
Не уверен, что для конкретной биржи есть смысл хранить историю, но я решил пока просто сохранять каждый запрос в базу, если что, не сердитесь, делать On Conflict я умею

По метрикам, я добавил парочку каких-то, чтобы посмотреть, как чудесно они работают, доступны по `/metrics`<br>
Хэлсчек доступен по `/health`
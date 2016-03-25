# aneks-api #

## Настройки ##
Настройки приложения производятся в файле config.yaml

## API ##

### Index ###
    endpoint/aneks?page=4&count=20

- _page_ - номер страницы (по-умолчанию == 0)
- _count_ - колличество элементов (по-умолчанию == 20)

### Get ###

    endpoint/aneks/1
Получение анека по id

### Random ###
    endpoint/aneks/random
Получение случайного анекдота

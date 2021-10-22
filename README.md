![](./retranslator.png)

#### Задание 2

1. Создать репозиторий именование которого указано в таблице прогресса

2. Описать сущность `{domain}.{Subdomain}` и `{domain}.{Subdomain}Event` в **internal/model/{subdomain}.go**

3. Реализовать паттерн consumer-producer из **db** в **kafka** на основе интерфейсов`EventRepo` и `EventSender` для одного типа события **Created**

4. Написать тесты

5. Создавать задачи у workerpool по обработке батчевых идентификаторов записей событий 💎

6. Поддержать несколько типов событий учитывая корректный порядок 💎

7. Реализовать доставку, как минимум один раз 💎

**Рецепт**

[omp-demo-api](https://github.com/ozonmp/omp-demo-api)

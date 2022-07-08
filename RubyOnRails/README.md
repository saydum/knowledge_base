```bash
# Создание нового проекта
rails new name_app

# Запуск локальног сервера
rails s # rails serve

# Генерирукм модель
rails generate model modelname title:text user_id:integer


# Запуск миграции
rails db:migrate

# Запуск консоли
rails c # rails console

# CREATE
modelname = ModelName.create(body: 'Hello World!', user_id: 1)
```

CRUD CREATE / READ / UPDATE / DELETE

```bash
# Create
Modelname.create(...)

# Read
modelname = Modelname.find(...)
modelname = Modelname.where(...)

# Update
modelname.update_attributes(...)

# Delete
modelname.destroy
```
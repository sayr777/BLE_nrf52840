# BLE_nrf52840
Этот код настраивает BLE-адаптер на плате nice!nano и начинает рекламу BLE-устройства с именем "NiceNano". Он также включает светодиод для индикации работы BLE. Если возникает ошибка, функция must выводит сообщение об ошибке.

## Компиляция и загрузка

1. Установите TinyGo: https://tinygo.org/getting-started/
2. Подключите nice!nano
3. Выполните в терминале:

```bash
tinygo flash -target=nicenano -programmer=uf2 main.go
```


Установка зависимостей:

bash
Copy
# Для Linux/macOS
tinygo get tinygo.org/x/bluetooth

# Для Windows

go get tinygo.org/x/bluetooth
2. Обновление TinyGo (если нужно):

bash
Copy
# Для пользователей Homebrew (macOS/Linux)
brew upgrade tinygo

# Для других платформ скачайте свежую версию:
# https://github.com/tinygo-org/tinygo/releases
3. Проверка установки:
Убедитесь, что в вашей среде:

bash
Copy
tinygo env
В выводе должна быть строка с TINYGOROOT, указывающая на корректную установку.


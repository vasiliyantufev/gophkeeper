# Менеджер паролей GophKeeper

-----
## Общие требования

GophKeeper представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

Сервер должен реализовывать следующую бизнес-логику:

<ul>
  <li>регистрация, аутентификация и авторизация пользователей;</li>
  <li>хранение приватных данных;</li>
  <li>синхронизация данных между несколькими авторизованными клиентами одного владельца;</li>
  <li>передача приватных данных владельцу по запросу.</li>
</ul>

Клиент должен реализовывать следующую бизнес-логику:

<ul>
  <li>аутентификация и авторизация пользователей на удалённом сервере;</li>
  <li>доступ к приватным данным по запросу.</li>
</ul>

Функции, реализация которых остаётся на усмотрение исполнителя:

<ul>
  <li>создание, редактирование и удаление данных на стороне сервера или клиента;</li>
  <li>формат регистрации нового пользователя;</li>
  <li>выбор хранилища и формат хранения данных;</li>
  <li>обеспечение безопасности передачи и хранения данных;</li>
  <li>протокол взаимодействия клиента и сервера;</li>
  <li>механизмы аутентификации пользователя и авторизации доступа к информации.</li>
</ul> 

Дополнительные требования:

<ul>
  <li>клиент должен распространяться в виде CLI-приложения с возможностью запуска на платформах Windows, Linux и Mac OS;</li>
  <li>клиент должен давать пользователю возможность получить информацию о версии и дате сборки бинарного файла клиента.</li>
</ul> 

## Типы хранимой информации

<ul>
  <li>пары логин/пароль;</li>
  <li>произвольные текстовые данные;</li>
  <li>произвольные бинарные данные;</li>
  <li>данные банковских карт.</li>
</ul>

Для любых данных должна быть возможность хранения произвольной текстовой метаинформации (принадлежность данных к веб-сайту, личности или банку, списки одноразовых кодов активации и прочее).

## Абстрактная схема взаимодействия с системой
Ниже описаны базовые сценарии взаимодействия пользователя с системой. Они не являются исчерпывающими — решение отдельных сценариев (например, разрешение конфликтов данных на сервере) остаётся на усмотрение исполнителя.
Для нового пользователя:

 <ol>
  <li>Пользователь получает клиент под необходимую ему платформу.</li>
  <li>Пользователь проходит процедуру первичной регистрации.</li>
  <li>Пользователь добавляет в клиент новые данные.</li>
  <li>Клиент синхронизирует данные с сервером.</li>
</ol>  

Для существующего пользователя:

 <ol>
  <li>Пользователь получает клиент под необходимую ему платформу.</li>
  <li>Пользователь проходит процедуру аутентификации.</li>
  <li>Клиент синхронизирует данные с сервером.</li>
  <li>Пользователь запрашивает данные.</li>
  <li>Клиент отображает данные для пользователя.</li>
</ol>  

## Тестирование и документация
Код всей системы должен быть покрыт юнит-тестами не менее чем на 80%. Каждая экспортированная функция, тип, переменная, а также пакет системы должны содержать исчерпывающую документацию.
## Необязательные функции
Перечисленные ниже функции необязательны к имплементации, однако позволяют лучше оценить степень экспертизы исполнителя. Исполнитель может реализовать любое количество из представленных ниже функций на свой выбор:

 <ul>
  <li>поддержка данных типа OTP (one time password);</li>
  <li>поддержка терминального интерфейса (TUI — terminal user interface);</li>
  <li>использование бинарного протокола;</li>
  <li>наличие функциональных и/или интеграционных тестов;</li>
  <li>описание протокола взаимодействия клиента и сервера в формате Swagger.</li>
</ul>

-----
## Принципиальная схема

![alt text](https://github.com/vasiliyantufev/gophkeeper/web/assets/schema.png?raw=true)

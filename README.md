# go-kafka

Перед вами мой pet проект в котором я учусь работать с kafka на практике.
Пока что тут нету ничего сложного, возможно позже я дополню данный проект чем то еще, но пока тут только основы. 

У нас есть producer, который записывает в топик order в партицию 0 14 записей заказов. В свою очередь consumer читает из этой очереди заказы, далее сохраняет их в базе данных, после чего выводит список сохраненных записей.

Для того чтобы запустить проект вам достаточно перейти в makefile:
1. Запустить команду docker-compose up
2. Запустить команду producer-up
3. Запустить команду consumer-up

В консоле вывведутся заказы которые были добавлены в очередь и сохранены в базе данных
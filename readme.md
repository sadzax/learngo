# Общий план:
1. Составление базового роадмапа
2. Выполнение leetcode-type задач на Go
3. Написание простенького сервиса на Gin
4. Определение практического модуля для изучения
5. ...

### Базовый роадмап:
- Изучить синтаксис языка
- Познакомиться с основными типами данных
- Освоить функции
- Разобраться со структурами
- Подключить и использовать базы данных
- Реализовать простые веб-взаимодействия
- Изучить основы конкурентности (goroutines, каналы)
- (Научиться работать с файлами)

### Текущий лог:

| ph0 | ph1 | ph2 | subj      | topic     | res     | link                                             |
| --- | --- | --- | --------- | --------- | ------- | ------------------------------------------------ |
| 1   |     |     | env       |           | ok      |                                                  |
|     |     |     | syntax    |           |         |                                                  |
|     | 1   | 1   | package   |           | ?       | [1.1](./tour/1_basics/1_hello/main.go)           |
|     |     |     | import    |           | ?       | [1.1](./tour/1_basics/1_hello/main.go)           |
|     |     | 2   | export    |           | ?       | [1.2](./tour/1_basics/2_export/main.go)          |
|     |     | 3   | syntax    | func      | ok + ?  | [1.3](./tour/1_basics/3_func/main.go)            |
|     |     | 4   | syntax    | var       | ok      | [1.4](./tour/1_basics/4_var/main.go)             |
|     |     | 5   | datatypes | basic     | ok      | [1.5](./tour/1_basics/5_basic_datatypes/main.go) |
|     |     | 6   | syntax    | constant  | ok + ?  | [1.6](./tour/1_basics/6_const/main.go)           |
|     | 2   | 1   | syntax    | loop      | ok      | [2.1](./tour/2_flowcontrol/1_for/main.go)        |
|     |     | 2   | syntax    | condition | ok      | [2.2](./tour/2_flowcontrol/2_if/main.go)         |
|     |     |     | syntax    | switch    | ok      | [2.2](./tour/2_flowcontrol/2_if/main.go)         |
|     |     |     | syntax    | defer     | ok      | [2.2](./tour/2_flowcontrol/2_if/main.go)         |
|     |     | 3   |           |           | xrsz    | [2.3](./tour/2_flowcontrol/3_xrsz/main.go)       |
|     |     | 4   | pointer   |           | ok + ?  | [2.4](./tour/2_flowcontrol/4_pointer/main.go)    |
|     | 3   | 1   | datatypes | struct    | ok + ?  | [3.1](./tour/3_moretypes/1_struct/main.go)       |
|     |     | 2   | datatypes | array     | ok + ?  | [3.2](./tour/3_moretypes/2_array/main.go)        |
|     |     | 3   | datatypes | slice     | ok + ?  | [3.3](./tour/3_moretypes/3_slice/main.go)        |
|     |     |     | datatypes | append    | ok + ?  | [3.3](./tour/3_moretypes/3_slice/main.go)        |
|     |     | 4   | datatypes | nil       | ok + ?  | [3.4](./tour/3_moretypes/4_nil/main.go)          |
|     |     |     | syntax    | make      | ok      | [3.4](./tour/3_moretypes/4_nil/main.go)          |
|     |     | 5   | syntax    | range     | ok + ?  | [3.5](./tour/3_moretypes/5_range/main.go)        |
|     |     | 6   |           |           | xrsz ?  | [3.6](./tour/3_moretypes/6_xrsz/main.go)         |
|     |     | 7   | datatypes | map       | ok + ?  | [3.7](./tour/3_moretypes/7_map/main.go)          |
|     |     |     |           |           | xrsz    | [3.7](./tour/3_moretypes/7_map/main.go)          |
|     |     | 8   | syntax    | funcvalue | ok + ?? | [3.8](./tour/3_moretypes/8_funcvalue/main.go)    |
|     |     | 9   |           |           | xrsz    | [3.9](./tour/3_moretypes/9_xrsz/main.go)         |
|     | 4   | 1   | syntax    | methods   | ok + ?  | [4.1](./tour/4_methods/1_method/main.go)         |
|     |     |     |           |           |         | [.](./tour///main.go)                            |
|     |     |     |           |           |         | [.](./tour///main.go)                            |
|     |     |     |           |           |         | [.](./tour///main.go)                            |


link_2 = 'https://codeshare.io/5z8jkG'

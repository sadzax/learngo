# Общий план:
1. Составление базового роадмапа
2. Выполнение leetcode-type задач на Go
3. Получение обратной связи (вопросы по п.2. + корректировка п.1.)
4. Определение практического модуля для чтения
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

| ph0 | ph1 | ph2 | src         | subj      | topic      | meth     | res    | link                                              |
|-----|-----|-----|-------------|-----------|------------|----------|--------|---------------------------------------------------|
| 1   |     |     | go.dev      | env       |            |          | ok     |                                                   |
|     |     |     | go.dev/tour | syntax    |            |          |        |                                                   |
|     | 1   | 1   |             | package   |            | examples | ?      | [1.1](./tour/1_basics/1_hello/main.go)            |
|     |     |     |             | import    |            | examples | ?      | [1.1](./tour/1_basics/1_hello/main.go)            |
|     |     | 2   |             | export    |            | examples | ?      | [1.2](./tour/1_basics/2_export/main.go)           |
|     |     | 3   |             | syntax    | func       | tasks    | ok + ? | [1.3.](./tour/1_basics/3_func/main.go)            |
|     |     | 4   |             | syntax    | var        | tasks    | ok     | [1.4.](./tour/1_basics/4_var/main.go)             |
|     |     | 5   |             | datatypes | basic      |          | ok + ? | [1.5.](./tour/1_basics/5_basic_datatypes/main.go) |
|     |     | 6   |             | syntax    | constants  |          | ok + ? | [1.6.](./tour/1_basics/6_const/main.go)           |
|     | 2   | 1   |             | syntax    | loop       | tasks    | ok     | [2.1.](./tour/2_flowcontrol/1_for/main.go)        |                                                  |
|     |     | 2   |             | syntax    | conditions | tasks    | ok     |                                                   |
|     |     |     |             | syntax    | switch     | tasks    | ok     |                                                   |
|     |     |     |             | syntax    | defer      | examples | ok + ? | [2.2.](./tour/2_flowcontrol/2_if/main.go)         |
|     |     | 3   |             | xrsz      |            |          | ok     |                                                   |
|     |     |     |             | pointer   |            |          | ok + ? | [2.4.](./tour/2_flowcontrol/4_pointer/main.go)    |
|     | 3   | 1   |             | datatypes | struct     | examples | ok + ? | [3.1.](./tour/3_moretypes/1_struct/main.go)       |
|     |     |     |             | datatypes | arrays     | tasks    | ok     | [3.2.](./tour/3_moretypes/2_array/main.go)        |
|     |     |     |             | datatypes | slices     | tasks    | ok     | [3.3.](./tour/3_moretypes/3_slice/main.go)        |                                            |


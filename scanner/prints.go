package scanner

import (
	"ToDoList/todo"
	"fmt"

	"github.com/k0kubun/pp"
)

func printResult(result string) {
	fmt.Println("Резудьтат выполнения команды:", result)
	fmt.Println("")
}

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена")
	fmt.Println("")
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println("")
}

func printDone(title string) {
	fmt.Println("Задача'" + title + "' помечена как выполненая")
	fmt.Println("")
}

func printDel(title string) {
	fmt.Println("Задача'" + title + "'успешно удалена")
	fmt.Println("")
}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("1. help")
	fmt.Println("— эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("")
	fmt.Println("2. add")
	fmt.Println("{заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("")
	fmt.Println("3. list")
	fmt.Println("эта команда позволяет получить полный список всех задач")
	fmt.Println("")
	fmt.Println("4. del")
	fmt.Println("{заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("")
	fmt.Println("5. done")
	fmt.Println("{заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("")
	fmt.Println("6. events")
	fmt.Println("— эта команда позволяет получить список всех событий")
	fmt.Println("")
	fmt.Println("7. exit")
	fmt.Println(" — эта команда позволяет завершить выполнение программы")
	fmt.Println("")
}

func printEvents(events []Event) {
	pp.Println("События:", events)
	fmt.Println("")
}

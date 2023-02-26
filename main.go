package main

import (
	"fmt"
)

type Task struct {
	name      string
	desc      string
	completed bool
}
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	// siu quiero eliminar el indice 2
	// del slice, hago un nuevo slice con elemento 0 y 1 y de 3 en adelante
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripcion: %s \nCompletado : %t\n", t.name, t.desc, t.completed)
}

func main() {

	dias := make(map[int]string)
	dias[1] = "Lunes"
	dias[2] = "Martes"
	fmt.Println(dias)
	dias[7] = "otro"
	fmt.Println(dias)
	fmt.Println(len(dias))

	estudiantes := make(map[string][]int)
	estudiantes["romina"] = []int{13, 23, 256, 23}
	estudiantes["ivan"] = []int{43, 5465, 22, 44}

	fmt.Println(estudiantes)

	fmt.Println(estudiantes["romina"])
	fmt.Println(estudiantes["romina"][1])

	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar el curso udemy de go",
		completed: false}

	t2 := Task{
		name:      "curso de flutter",
		desc:      "completar el curso de flutter",
		completed: false,
	}

	lista := TaskList{}

	lista.appendTask(&t1)
	lista.appendTask(&t2)
	t1.toPrint()
	t2.toPrint()

	t3 := Task{
		name:      "Curso de Go",
		desc:      "Completar el curso udemy de go",
		completed: false,
	}
	lista.appendTask(&t3)
	fmt.Println(lista)
	for i, task := range lista.tasks {
		fmt.Println(i, task)
	}

}

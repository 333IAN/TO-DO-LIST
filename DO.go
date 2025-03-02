package main
import("fmt"
       "bufio"
	    "os")
var todos[]string //variables added are of slice todos
func addToDos(){  //function for addition of tasks and queueing
	reader:=bufio.NewReader(os.Stdin) //creates buffered reader that reads input from standard input
	var task string //string datatype of task variable
	fmt.Print("Enter the task description:") //Asks for input from user
	task,err:=reader.ReadString('\n')  //reads input from reader and stores in the task variable eliminating any errors during reading
	if err!=nil{
		fmt.Println(err)  //any error found is displayed
		return
	}
	todos=append(todos,task)//adding new task stored in task to the todos slice
	fmt.Println("Task added successfully")
}
func listToDos(){ //another function for listing added tasks
	if len(todos)==0{ //len checks if any task is added
		fmt.Println("No tasks Found..")//message displayed
		return
	}
	for i,todo:=range todos{
		fmt.Printf("%d, %s\n",i+1,todo)
	}
}
func main(){
	fmt.Println("TO-DO List APP")
	for {
		var choice int //choice is of integer datatype
		fmt.Println("1.Add To-Do item")
		fmt.Println("2.List To-Dos")
		fmt.Println("3.Exit")
		fmt.Print("Enter your choice:") //asks user to select between the three choices above
		_,err:=fmt.Scanln(&choice) //stores input from user in the choice variable declared earlier
		if err!=nil{
			fmt.Println(err)
			continue
		}
		switch choice{
		case 1:
			addToDos()// first function call if user chooses 1
		case 2:
			listToDos()// second function call if user chooses 2
		case 3:
			fmt.Println("Exiting...")// if user chooses choice 3 shows this message
		default:
			fmt.Println("Invalid Option.Try Again.")
		}
		fmt.Println() //prints an extra free line for tidiness.
	}
}


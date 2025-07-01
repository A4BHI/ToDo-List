package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Task struct{
	title string
	done bool
}
var tasks []Task = make([]Task,0)



func addTask(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm()
		var task string = r.FormValue("task")
		newTask := Task{title: task,done: false}
		tasks = append(tasks,newTask)

		w.Header().Set("content-type" , "text/plain")
		fmt.Fprint(w,task)
		fmt.Println(task)
		fmt.Println(tasks)
		
	}
}

func onreload(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodGet{
		w.Header().Set("content-type" , "text/plain")
		fmt.Println(tasks)
		// joined := strings.Join(tasks, "|")
		for _,t := range tasks{
			tt := t.title
			dd := t.done
			fmt.Fprint(w,tt+"+"+ strconv.FormatBool(dd) + "|")
			fmt.Println(tt,dd)

		}
		// fmt.Fprint(w,tasks)
	
		
		
	}else{
		fmt.Println(r.Method)
	}
}

func finishTask(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost{
		
		bytes,_:=io.ReadAll(r.Body)
		task:= string(bytes)
		w.Header().Set("content-type","text/plain")
		fmt.Println(task)
		var count int
		for index := range tasks{
			if(tasks[index].title==task){
				count++
				if(!tasks[index].done){
					
					tasks[index].done=true
						
				}
			}
			
		}
		if(count==1){
			fmt.Fprint(w,"finish")
			
			
		}else if(count>1){
			fmt.Fprintf(w,"reload")
		}

		
	}
}
func deleteTask(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost{
		duptask:=[]Task{}
		bytes,_:=io.ReadAll(r.Body)
		task:=string(bytes)
		for i := range tasks{
			if(tasks[i].title!=task){
				duptask=append(duptask,tasks[i])
			}
		}
		tasks=duptask
		fmt.Println(tasks)

	}
}
func main(){
	fmt.Println("Web Server Started.")

	http.Handle("/" ,http.FileServer(http.Dir("./html")))
	http.HandleFunc("/add",addTask)
	http.HandleFunc("/reload",onreload)
	http.HandleFunc("/finish",finishTask)
	http.HandleFunc("/delete",deleteTask)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Print("Error ",err)
	}

}
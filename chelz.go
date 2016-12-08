package main 
import "reflect"
import (
	"ethos/syscall"
	"ethos/ethos"
	ethosLog "ethos/log"
	"ethos/efmt"
)
//import "encoding/json"
func main(){
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("cs587_csrini2")
	if status != syscall.StatusOk {
	efmt.Fprintf(syscall.Stderr, "Error opening %v: %v\n", path, status)
	syscall.Exit(syscall.StatusOk)
}

// 3 objects
	obj1 := SimLog{"Chellam","Jan",10}
	obj2 := SimLog{"Arvindh", "Feb", 20}
	obj3 := SimLog{"Janu", "Mar", 40}
	arr_objects := [3] SimLog{obj1, obj2, obj3}

//core part of idea made easy with reflect package

	efmt.Println("Name of field\t->\tValue of field\n\n")
	for obj:=0;obj<len(arr_objects); obj ++{
		curr_obj := arr_objects[obj]
		curr_obj_handle := reflect.ValueOf(&curr_obj).Elem()
		typeOfT := curr_obj_handle.Type()
		efmt.Fprintf(syscall.Stderr, "Object %v", obj +1)
		for i:=0; i<curr_obj_handle.NumField(); i++ {
			field_val := curr_obj_handle.Field(i)
			efmt.Fprintf(syscall.Stderr, "\n%v\t->\t%v\n",typeOfT.Field(i).Name,field_val)
		} 
		efmt.Println("\n")
}

	fd, status := ethos.OpenDirectoryPath(path)
	obj1.Write(fd)
	efmt.Fprint(syscall.Stderr, "\n\nOver\n\n")


}

package err

import ("log")

// Stop a function that stop the software if the error isn't null
func RaiseError(err error){
	if err!=nil{
		log.Fatal(err)
	}
}
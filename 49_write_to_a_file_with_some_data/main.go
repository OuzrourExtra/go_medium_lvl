package main

import(
	"fmt"
	"os"
	"encoding/json"
	"log"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Mail string `json:"mail"`
	Password string `json:"password"`
}

func WriteJSON(fileName string , data []Person){
	_ , err := os.Stat(fileName)
	if err == nil { // file exist
		AppendJson(fileName,data)
	}else{ // file don't exist
		persons := []Person{Person{"ilyas",28,"ilyas.ouzrour@gmail.com","123456lol"},
	Person{"admin",19,"admin@ouzrour.com","i'mtheadmin"}}
		jsonData ,errJson := json.MarshalIndent(persons,"","	")
		if errJson!=nil{
			log.Fatal(errJson)
		}
		errWriting := os.WriteFile(fileName,jsonData,0644)
		if errWriting != nil{
			log.Fatal(errWriting)
		}
		fmt.Printf("\n JSON Created !")
	} 
}

func AppendJson(filename string,dataAfter []Person){
	dataBefore , errReading := os.ReadFile(filename)
	if errReading!=nil{
		log.Fatal(errReading)
	}
	var jsonBefore []Person
	errUnmarshal := json.Unmarshal(dataBefore,&jsonBefore)
	if errUnmarshal != nil{
		log.Fatal(errUnmarshal)
	}
	jsonAfter := append(jsonBefore, dataAfter...)
	lastJson , errJsonConvertion := json.MarshalIndent(jsonAfter,"","	")
	if errJsonConvertion != nil {
		log.Fatal(errJsonConvertion)
	}
	os.WriteFile(filename,lastJson,0644)
	fmt.Printf("\nJSON Already Exist ! Modification Done !")

}

func ReadJSON(fileName string){
	_ , err := os.Stat(fileName)
	
	if err == nil { // file exist
		data , err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		var jsonData interface{}
		json.Unmarshal(data,&jsonData)
		fmt.Printf("The Data %v",jsonData)
	}else{ // file don't exist
	fmt.Printf("\n don't Exist ! ")
	} 
}

func main(){
	persons := []Person{Person{"ilyas",28,"ilyas.ouzrour@gmail.com","123456lol"},
	Person{"admin",19,"admin@ouzrour.com","i'mtheadmin"}}
	WriteJSON("unfounded.json",persons)
	ReadJSON("exist.txt")
	

	newList := []Person{
	{"Ilyas", 28, "ilyas.ouzrour@gmail.com", "123456lol"},
	{"Admin", 19, "admin@ouzrour.com", "i'mtheadmin"},
	{"Alice Johnson", 32, "alice.j@example.com", "a1iceP@ss"},
	{"Bob Lee", 29, "bob.lee@corp.org", "b0bSecure!"},
	{"Carla Mendes", 24, "carla.mendes@events.com", "event24!carla"},
	{"Derek Nguyen", 27, "derek.n@events.com", "derekPass123"},
	{"Emma Watson", 30, "emma.watson@mail.com", "hermioneGr@ng3r"},
	{"Frank Castle", 41, "punisher@marvel.com", "revenge!007"},
	{"Grace Hopper", 85, "ghopper@navy.mil", "cobolLegend"},
	{"Hassan Ali", 26, "hassan.ali@devmail.com", "code4life"},
	{"Ivy Chen", 22, "ivy.chen@edu.net", "studyHard123"},
	{"Jack Sparrow", 38, "jack@blackpearl.com", "whereIsMyRum"},
	{"Kelly Brown", 35, "kelly.b@workplace.org", "safePass!1"},
	{"Liam Smith", 21, "liam.smith@uni.edu", "LiamUni2024"},
	{"Mona Lisa", 500, "mona@art.com", "davinci@smile"},
	{"Nashit Khan", 27, "nashit@freelance.dev", "nashitRules"},
	{"Olivia Taylor", 33, "olivia.taylor@media.com", "oLiveMeDia"},
	{"Paul Walker", 40, "paul@fastcars.com", "rideOrDie"},
	{"Quincy Adams", 29, "q.adams@usa.org", "historyBuff"},
	{"Rania Khalil", 31, "rania.k@designhub.net", "creative2025"},
	{"Steve Rogers", 105, "captain@avengers.com", "shield#America"},
	{"Tina Fey", 45, "tina.fey@funny.tv", "comedyQueen"},
	{"Usman Tariq", 27, "usman.t@remote.io", "remoteWork!"},
	{"Valentina Rossi", 36, "valentina.r@motors.com", "vr46Speed"},
	{"William Zhang", 34, "will.zhang@techlab.com", "AI_2024"},
	{"Ximena Cruz", 25, "ximena.cruz@startup.mx", "holaMundo!"},
	{"Youssef El Amrani", 30, "youssef@morocco.ma", "marhaba123"},
	{"Zoey Martin", 28, "zoey.m@producthub.com", "launchItNow"},
	{"Léa Dupont", 26, "lea.dupont@france.fr", "bonjourParis"},
	{"Tariq Ben Ali", 31, "tariq.ali@projet.ma", "tariqDevGo"},
	{"Fatima Zahra", 24, "fatima.z@elearning.org", "learn&grow"},
	{"Nina Patel", 29, "nina.patel@ngo.org", "serve2lead"},
	}
	AppendJson("exist.json",newList)
	ReadJSON("exist.json")
}	
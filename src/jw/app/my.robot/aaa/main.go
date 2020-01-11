package main

func main() {

	self := make([]int32, 10)
	for i, _ := range self {



		if i < len(self){

			self=append(self[:i],self[i+1:]...)
		}



	}


}

package main
import ("fmt")
import ("github.com/Greetings")

func main() {
	fmt.Println("Hello, playground")
	slicee := []byte{'a','b','c','d','y'}
	fmt.println(ConcatSlice(slicee))  
fmt.Println(string(Encrypt(slicee,3)))  
str=EZGreetings("Uzair")
  fmt.Printf("%s",str)


 
}

func ConcatSlice(sliceToConcat[]byte) string{
    var ret string=""
    for _, v := range sliceToConcat {
        
        ret+=string(v)
    }
 return ret 


func Encrypt(sliceToEncrypt[]byte,ceaserCount int) []byte{
    for i:=0;i<len(sliceToEncrypt);i++{
	
	if (sliceToEncrypt[i]+byte(ceaserCount))>=122 {
	     sliceToEncrypt[i]=((sliceToEncrypt[i]+byte(ceaserCount))%122)+97
	} else {
	sliceToEncrypt[i]=sliceToEncrypt[i]+byte(ceaserCount)
	}
	
  	
    }
return sliceToEncrypt
 

}
func EZGreetings(name string) string{
  return greetings.PrintGreetings(name)
}

 
}
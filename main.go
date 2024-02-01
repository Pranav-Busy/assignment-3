package main

import (
	"fmt"
	"reflect"
)


func merger(helper interface{},helper1 interface{})(interface{},error){

  if helper==nil||helper1==nil{
    if helper==nil{
      return helper1,nil
    }else 
    {
      return helper,nil
    }
  }else 
  {
    var result interface{}
     var final[] interface{}

     first :=reflect.ValueOf(helper)
     
     switch first.Kind(){

     case reflect.Slice:
      a := first.Interface().([]int)

      for _,val:=range a{
        final = append(final, val)
      }
      break;
     case reflect.Int:

      a := first.Interface().(int)
      final = append(final, a)
      break;


     case reflect.String:
      a := first.Interface().(string)
      final = append(final, a)
      break;
     }
     second :=reflect.ValueOf(helper1)
     
     switch second.Kind(){

     case reflect.Slice:
      a := second.Interface().([]int)

      for _,val:=range a{
        final = append(final, val)
      }
      break;
     case reflect.Int:

      a := second.Interface().(int)
      final = append(final, a)
      break;


     case reflect.String:
      a := second.Interface().(string)
      final = append(final, a)
      break;
     }

     result = final
     return result,nil
  }
  



}
func main()  {
  
  var helper interface{}

  helper=[]int{1,2,3,4,5}

  

  var helper2 interface{}
  helper2="pranav"

  result,err:=merger(helper,helper2)

  if err!=nil {
    fmt.Println(err);
  }else 
  {
    fmt.Println(result)
  }

}
package main

import (
    "fmt"
    "errors"
    )
    
    
    func task1()error{
        return errors.New("task one is failed")
    }
    
    func task2()error{
        return nil
    }
    
    func task3()error{
        return errors.New("Task three is failed")
    }
    
    
func main(){
    
    var errs [] error
    
    if err:=task1(); err!=nil{
        errs=append(errs,err)
    }
    if err:=task2(); err!=nil{
        errs=append(errs,err)
    }
    if err:=task3(); err!=nil{
        errs=append(errs,err)
    }
    
    
    if len(errs)>0{
       for i,e:=range errs{
            fmt.Printf("%d, %s\n", i+1, e)
        }
        
        joined:=errors.Join(errs...)
        fmt.Println(joined)
    }else{
        fmt.Println("all errors is occurred successfully")
    }
    
}
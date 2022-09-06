package main

import (
	"fmt"
)

func binarysearch(array[]int, size int, num int) int{
   var len_n   int = len(array)
   var l_index int = 0
   var r_index int = len_n-1
   var m_index int = 0
   for l_index<=r_index {                                    // we don't have while loop in go lang 

         m_index = l_index +(r_index -l_index)/2
         if array[m_index] == num{
            return m_index
            }
         if array[m_index] < num{
            l_index = m_index +1
            }else{
                  r_index = m_index -1   
            }
         
      }
    
   return -1


}

func main(){
   fmt.Println("Enter the size of the array")
   var size int
   fmt.Scan(&size) 

   fmt.Println("Enter the capacity ")
   var capacity int
   fmt.Scan(&capacity)

   var num int
   array:=make([]int, size, capacity)
    
 
   for i :=0;i< size; i++{
	 fmt.Println("Enter the element:")
	 fmt.Scan(&array[i])

   }


   fmt.Println("Enter the number to be searched")
   fmt.Scan(&num)

    var search_found_index int
    search_found_index= binarysearch(array,size,num)
    if search_found_index >= 0{
		fmt.Println("Number found at position :", search_found_index)


	}else{fmt.Println("Not found number in any of the index") }


}
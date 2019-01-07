package shapes
 
 type square struct{
        l float64
 }

 func(s square) area()float64{

    return s.l * s.l
 }
  
 
func(s square) perimeter()float64{

    return 4 * s.l
}   

func(s *square) setLength(length float64){

    s.l = length
}   

func(s square) length() float64{
    return s.l
}


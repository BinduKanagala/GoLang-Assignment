package shapes
 
 type Rect struct{
        l,b float64
 }

 func(r Rect) area()float64{

    return r.l * r.b
 }
  
 
func(r Rect) perimeter()float64{

    return 2 * (r.l + r.b)
}   

func(r *Rect) setLength(length float64){

    r.l = length
}   

func(r Rect) length() float64{
    return r.l
}

func(r *Rect) setBreadth(breadth float64){

    r.b = breadth
}   

func(r Rect) breadth() float64{
    return r.b
}
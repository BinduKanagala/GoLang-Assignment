package shapes

import "math"
 
 type circle struct{
    	r float64
 }

 func(c *circle) area()float64{

 	return math.pi * c.r * c.r
 }
  
func(c *circle) perimeter()float64{

	return 2 * math.pi * c.r
}	

func(c *circle) setRadius(radius float64){

	c.r= radius
}	

func(c circle) radius() float64{
	return c.r
}
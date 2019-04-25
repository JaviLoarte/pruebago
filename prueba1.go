// version 1.2

package main
        import "fmt"
        import "os"
        import "os/exec"
        import "log"
        import "math"
        import "time"
        
        func sum (a, b int) int { return a + b }
        
        var x int 
        
        func sumador (delta int) int {
            x += delta 
            return x
        }
        
        func fp (a[3] int) {
             fmt.Println (a[0], a[1], a[2])
         }
        
        type Punto struct { x, y float64 }
        func (p Punto) Abs() float64 {
             return math.Sqrt(p.x*p.x + p.y*p.y)
        }
        
        type IntVector []int
        func (v IntVector) Sum() (s int) {
           for x := range v {
              s = s + x }
        return }
        
        func (v IntVector) Sum2() (s int) {
           for x := range v {
              s = s + v[x] }
        return }
          
        type MyFloat float32
        func (f MyFloat) Abs() MyFloat {
             if f < 0 { return -f }
               return f
        }  
         
        func IsReady (what string, minutes int64) { 
            time.Sleep (6 * time.Second) 
            fmt.Println (what, "esta listo")
        } 
           
        func main() {
            fmt.Printf("Hello World!\n")
            var a = 2;
            var b = 3;
            var c = sum (a, b) ;
            fmt.Printf("%d\n", c) ;
            a *= b
            fmt.Printf("%d\n", a) 
        
        fmt.Printf("%d\n",sumador (1))
        fmt.Printf("%d\n",sumador (20))
        fmt.Printf("%d\n",sumador (200))
        
        var y int = 2
        var punt *int
        fmt.Fprintln (os.Stdout, "La variable y que vale %d, esta en la posicion %p", y, &punt)        
     
        fmt.Println(os.Hostname())
                
        cmd := exec.Command ("/bin/ls", "-l")
        stdout, err := cmd.Output()
        if err != nil {
           log.Printf("error: %v\n", err)
        }
        
        // print(string(stdout))
        fmt.Println(string(stdout))
        
        var ar = [3] int {1,2,3}
        fp(ar)
         
        m := map[string] float32 { "1":1.0, "pi":3.1415 }
        for key, value := range m {
           fmt.Printf("clave %s, valor %g\n", key, value)
        } 
        
        type Point struct { x, y float32 }
        var p Point
        p.x = 7
        p.y = 23.4
        fmt.Printf("x %g, y %g\n", p.x, p.y)
        
        var pp *Point = new(Point)
        *pp = p
        fmt.Printf("x %g, y %g\n", pp.x, pp.y)
        rr := new(Point);
        rr = &p
        fmt.Printf("x %g, y %g\n", rr.x, rr.y)
        
        
        p2 := new(Punto)
        p2 = &Punto { 3, 4 } 
        fmt.Printf("%g\n", p2.Abs())
        
        fmt.Println(IntVector { 1, 2, 5 }.Sum())
        fmt.Println(IntVector { 1, 2, 5 }.Sum2())
        
        var valor1 MyFloat
        var valor2 MyFloat
        valor1 = -64
        fmt.Println(valor1.Abs())
        valor2 = 64
        fmt.Println(valor2.Abs())
        
        fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
        time.Sleep(2 * time.Second)
        fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
        
        fmt.Printf("Preparandose cafe\n")
        go IsReady("cafe",1);
        fmt.Printf("Preparandose te\n")
        go IsReady("te",1);
        time.Sleep(10 * time.Second)
        
     }


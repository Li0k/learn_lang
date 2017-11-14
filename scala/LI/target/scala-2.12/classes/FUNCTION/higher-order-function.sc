def apply(f:Int => String,v:Int) = f(v)

class Decorator(left: String,right: String){
  def layout[A](x:A) = left + x.toString + right
}


val decorator = new Decorator("[", "]")
println(apply(decorator.layout, 7))

def filter(xs: List[Int], proc: Int => Boolean): List[Int] = {
  if (xs.isEmpty) xs
  else if (proc(xs.head)) xs.head :: filter(xs.tail, proc)
  else filter(xs.tail, proc)
}

//val add = (x: Int) => println(x)
def modN(n:Int)(x:Int)= (n % x) == 0

val nums = List(1, 2, 3, 4, 5, 6, 7, 8)
println(filter(nums,modN(2)))
println(filter(nums,modN(3)))


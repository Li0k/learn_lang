def factorial(x:Int):Int = {
  def fact_iter(x:Int,accumulator: Int):Int = {
    if (x <= 1) accumulator
    else fact_iter(x - 1, x * accumulator)
  }

  fact_iter(x,1)
}

print(factorial(5))
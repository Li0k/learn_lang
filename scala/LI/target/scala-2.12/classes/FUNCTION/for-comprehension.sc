case class User(val name: String, val age: Int)

val userBase = List(new User("Travis", 28),
  new User("Kelly", 33),
  new User("Jennifer", 44),
  new User("Dennis", 23))

val twentySomethings = for (user <- userBase if (user.age >= 20 && user.age < 30))
  yield user.name

twentySomethings.foreach(name => println(name))

def foo(n: Int, v: Int) = for (i <- 0 until n;
                               j <- i until n if i + j == v)
  yield (i,j)


//{case 2 => "OK"} 连同花括号整体是一个lambda(函数字面量)。
//普通的方法都是完全函数，即 f(i:Int) = xxx 是将所有Int类型作为参数的，是对整个Int集的映射；而偏函数则是对部分数据的映射，比如上面{case 2=> "OK" }就仅仅只对2做了映射。偏函数的实现都是通过模式匹配来表达的。
//偏函数是通过 { case x => y } 这种特殊的方式来描述的
foo(10,10) foreach {
  case (i,j) => println(i,j)
}

//println(foo(10,10))


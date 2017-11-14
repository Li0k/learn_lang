//懵逼
//将类的static放入伴生对象中，不懂
class X {
  import X._
  def blah = foo
}
object X {
  private def foo = 42
}

//实现工厂模式
class Bar(foo:                                                                    String)

object Bar {
  def apply(foo: String) = new Bar(foo)
}
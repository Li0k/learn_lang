class Stack[A] {
  private var elements:List[A] = Nil
  def push(x:A) = {elements = x::elements}
  def peek:A = elements.head
  def pop():A = {
    val currentTop = peek
    elements = elements.tail
    currentTop
  }
}


//泛型类型的子类型是不变的
val stack = new Stack[Int]
stack.push(1)
stack.push(2)
println(stack.peek)
println(stack.pop())
println(stack.pop())


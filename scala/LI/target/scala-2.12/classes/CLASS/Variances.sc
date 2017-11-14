//协变 (Covariance) 意味着对于两个类型A和B，且A是B的子类型，则List[A]也是List[B]的子类型。
//class Foo[+A] // A covariant class

abstract class Animal {
  def name: String
}

case class Cat(name: String) extends Animal

case class Dog(name: String) extends Animal

def printAnimal(animals: List[Animal]) = {
  animals foreach (animal => println(animal.name))
}

//List[+A]
val cats: List[Cat] = List(Cat("Whiskers"), Cat("Tom"))
val dogs: List[Dog] = List(Dog("Fido"), Dog("Rex"))

printAnimal(cats)
printAnimal(dogs)

//逆变（Contravariance）意味着对于两个类型A和B，且A是B的子类型，则Writer[B]是Writer[A]的子类型。

abstract class Printer[-A] {
  def print(value: A): Unit
}

class AnimalPrinter extends Printer[Animal] {
  def print(animal: Animal): Unit =
    println("The animal's name is: " + animal.name)
}

class CatPrinter extends Printer[Cat] {
  def print(cat: Cat): Unit =
    println("The cat's name is: " + cat.name)
}

val myCat: Cat = Cat("Boots")
def printMyCat(printer: Printer[Cat]): Unit = {
  printer.print(myCat)
}

val catPrinter: Printer[Cat] = new CatPrinter
val animalPrinter: Printer[Animal] = new AnimalPrinter
printMyCat(catPrinter)
printMyCat(animalPrinter)

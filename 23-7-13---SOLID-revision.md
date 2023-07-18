## SOLID

S - SRP

O - Open to addition - closed to modification

L - Liskov's substituion principle

I - interface segregation 

D - dependency inversion

### SRP

----------

### OpenClosed Principle

eg: instead of creating if else condition --> make classes and make them implement the method
eg. shape calculate area..

----------

### L


----------

### Interface Segregation Principle

`A client should never be forced to implement an interface that it doesn’t use, or clients shouldn’t be forced to depend on methods they do not use.`

```java
interface ShapeInterface {
    public function area();

    public function volume();
}
-> not useful for flat shapes-like squares and circle
        Break the interface :

interface ShapeInterface {
    public function area();
}

interface ThreeDimensionalShapeInterface {
    public function volume();
}

```

------------

### Dependency inversion:

`We want our classes to be open to extension, so we have reorganized our dependencies to depend on interfaces instead of concrete classes.`

`Entities must depend on abstractions, not on concretions. It states that the high-level module must not depend on the low-level module, but they should depend on abstractions.`

from https://martinfowler.com/articles/dipInTheWild.html
When your system uses the String class, you probably do not invert that dependency. You could, for example if you think
of String as a primitive (strictly not, but close enough), then manipulating a number of Strings starts to
resemble `Primitive Obsession`.

If you introduce a type around the Strings, and add methods that make sense to your use of the those Strings, rather
than simply exposing String methods, that starts to look like a kind of Dependency Inversion, so long as the resulting
type is closer to your domain than a String.


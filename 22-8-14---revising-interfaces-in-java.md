
## Revising interfaces
#### Aug 15th 2022

Noting down things I had forgotten about Java 

- `final`
  - var --> const
  - func --> cant override
  - class --> cant inherit
    - eg. Wrapper classes - Integer, Float - cant be inherited
----------
- `interface`
  - like any other language
  - only one used - Public.
  - used to achieve total abstraction
  - All methods are public and abstract.
  - All fields are public, static, and final
  - Post `JDK 8` --> can define `default implementation` of func -> For backward compatibility in the situation : `add a new function in an existing interface`

----------

- `marker interface`
  - empty interface (no field or methods)
  - provides run-time type information about objects, so the compiler and JVM have additional information about the object.
  - same effect can be achieved by annotation
  - but marker interface - will mark sub-classes of marked class too
  - Example :
    - `Cloneable` -> in `java.lang` package
    - `Serializable` -> in `java.io` package
    - `Remote` --> in `java.rmi` for RMI
----------
- `FunctionalInterface`
  - can contain `only one abstract method`
  - However, they can include `many default` and `static` methods`
  - EXAMPLE :
    - Runnable -> contains only `run()` method
    - ActionListener -> contains only `actionPerformed()` method
    - Comparable -> contains only `compareTo()` method
    - Callable -> contains only `call()` method
    - `Lambda expressions` as the instance of functional interface
    - java.util.function package contains many built-in functional interfaces in Java 8.
    - good ref : https://www.geeksforgeeks.org/functional-interfaces-java/?ref=lbp
----------
    
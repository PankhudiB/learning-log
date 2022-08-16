
## Reading Headfirst design pattern 
###Observer Pattern
#### Aug 16th 2022

### Internet-based Weather Monitoring Station

WeatherData - temperature, humidity, and barometric pressure
application - with 3 display elements - current conditions, weather statistics and a simple forecast

- Strive for loosely coupled designs between objects that interact.
- Observer - defines many to one dependency between objects so that 1 object changes state, all dependants are notified
- Don’t depend on a specific order of notification for your Observers.
- Java has several implementations of the Observer Pattern --> including `java.util. Observable.`
  - Advantages : 
    - boiler plat code of add / delete observer is taken care of 
  - Disadvantages : 
    - Observable is a class -> which means u have to extend it | cant implement 
    - setChanged is protected -> This means you can’t even create an instance of the Observable class and compose it with your own objects, you have to subclass. 
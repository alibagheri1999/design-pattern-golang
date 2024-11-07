# Design Patterns in Go

This repository contains various design pattern implementations in Go. Each file represents a specific design pattern, including descriptions and usage examples.

## 1. Adapter Pattern - `adapter.go`
**Description**: The Adapter pattern allows incompatible interfaces to work together by wrapping an existing class with a new interface. 

**Explanation**: In this implementation, two payment gateways (`SamanPayment` and `AsanPardakhtPayment`) are adapted to a unified interface (`IPaymentGateway`). The adapters (`SamanPaymenentAdapter` and `AsanPardakhtPaymentAdapter`) wrap the existing payment classes, making them compatible with the `IPaymentGateway` interface.

## 2. Builder Pattern - `builder.go`
**Description**: The Builder pattern separates the construction of a complex object from its representation, allowing the same construction process to create different representations.

**Explanation**: This implementation constructs a `Car` object step-by-step using the `CarBuilder` struct. The builder allows for the customization of each `Car` property (brand, model, color, year) and returns a complete `Car` object through the `Build` method.

## 3. Decorator Pattern - `decorator.go`
**Description**: The Decorator pattern allows behavior to be added to an individual object, dynamically, without affecting the behavior of other objects from the same class.

**Explanation**: Here, a basic `Coffee` interface is extended by decorators like `WithMilk` and `WithSuger`, which add additional features (milk, sugar) and alter the price. The decorators wrap the `SimpleCoffee` object, enhancing it with additional properties.

## 4. Factory Pattern - `factory.go`
**Description**: The Factory pattern provides a way to create objects without specifying the exact class of object that will be created.

**Explanation**: In this example, `PaymentFactory` creates instances of different payment gateways (`Saman` or `Mellat`) based on an input type. This pattern encapsulates the creation logic and returns an interface (`PaymentGateway`) to the caller.

## 5. Observer Pattern - `observer.go`
**Description**: The Observer pattern defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

**Explanation**: The `Subject` struct manages a list of `Observer` instances (such as `SMSService` and `EmailService`) that listen for updates. When the `Notify` method is called, each observer executes its `Update` method, responding to changes in the subject.

## 6. Singleton Pattern - `singleton.go`
**Description**: The Singleton pattern ensures a class has only one instance and provides a global point of access to it.

**Explanation**: This code defines a `CacheManager` as a singleton using `sync.Once` to ensure only one instance is created. `GetCacheManager` initializes the instance only once and returns the same instance each time it is called, ensuring a single, globally accessible cache.

## 7. Strategy Pattern - `strategy.go`
**Description**: The Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable. This pattern lets the algorithm vary independently from clients that use it.

**Explanation**: In this example, `Order` can use different shipping strategies (`PeykShipping`, `PostShipping`) based on a chosen method. By setting different strategies, `Order` can dynamically change its shipping behavior at runtime.

---

Each design pattern is implemented with a simple use case to illustrate how it can be applied in Go. These patterns offer a foundation for scalable and maintainable code by allowing separation of concerns, encapsulation of behaviors, and improved modularity.

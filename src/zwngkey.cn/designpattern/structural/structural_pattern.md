# 结构型模式

结构型模式是软件设计模式的一类，关注于如何组合类和对象以获得更大的结构。它们可以让你构建更复杂的类和对象，以满足你的业务需求。
常见的结构型模式包括适配器模式、桥接模式、组合模式和装饰器模式。这些模式都可以帮助你更好地组织和管理你的代码，使其更容易维护和扩展。

- 结构型模式(Structural Pattern)关注如何将现有类或对象组织在一起形成更加强大的结构
- 不同的结构型模式从不同的角度组合类或对象，它们在尽可能满足各种面向对象设计原则的同时为类或对象的组合提供一系列巧妙的解决方案

> 类结构型和对象结构型

类结构型模式

- 关心类的组合，由多个类组合成一个更大的系统，在类结构型模式中一般只存在继承关系和实现关系


对象结构型模式

- 关心类与对象的组合，通过关联关系，在一个类中定义另一个类的实例对象，然后通过该对象调用相应的方法

> 结构型模式一览表

| **模式名称**                    | **定** **义**                                                | **学习难度** | **使用频率** |
| ------------------------------- | ------------------------------------------------------------ | ------------ | ------------ |
| **适配器模式(Adapter Pattern)** | **将一个类的接口转换成客户希望的另一个接口。适配器模式让那些接口不兼容的类可以一起工作。**允许使用不兼容的接口之间的类互相协作。 | ★★☆☆☆        | ★★★★☆        |
| **桥接模式(Bridge Pattern)**    | **将抽象部分与它的实现部分解耦，使得两者都能够独立变化。**   | ★★★☆☆        | ★★★☆☆        |
| **组合模式(Composite Pattern)** | **组合多个对象形成树形结构，以表示具有部分-整体关系的层次结构。组合模式让客户端可以统一对待单个对象和组合对象。** | ★★★☆☆        | ★★★★☆        |
| **装饰模式(Decorator Pattern)** | **动态地给一个对象增加一些额外的职责。就扩展功能而言，装饰模式提供了一种比使用子类更加灵活的替代方案。** | ★★★☆☆        | ★★★☆☆        |
| **外观模式(Facade Pattern)**    | **为子系统中的一组接口提供一个统一的入口。外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。** | ★☆☆☆☆        | ★★★★★        |
| **享元模式(Flyweight Pattern)** | **运用共享技术有效地支持大量细粒度对象的复用。**             | ★★★★☆        | ★☆☆☆☆        |
| **代理模式(Proxy Pattern)**     | **给某一个对象提供一个代理或占位符，并由代理对象来控制对原对象的访问。** | ★★★☆☆        | ★★★★☆        |


> 适配器模式案例:

例如，假设你有一个旧系统，它使用一个特定的接口来读取数据，但你有一个新的系统，它使用另一种接口来读取数据。在这种情况下，你可以使用适配器模式，通过创建一个适配器类来翻译旧系统的接口，使它与新系统的接口匹配。这样，两个系统就可以在一起工作，而你无需修改旧系统的代码。


> 桥接模式案例

例如，假设你有一个图形绘制库，它可以绘制不同类型的图形，例如线条、圆形和矩形。如果你将抽象图形类与它的实现分离开来，那么你就可以很容易地扩展该库以支持新的图形类型，而无需修改现有的代码。


> 装饰器模式案例

例如，假设你有一个用于显示图像的类，你可能希望为该类添加旋转、缩放和添加边框的功能。你可以使用装饰器模式来为该类创建一个装饰器，以便你可以在不更改该类的情况下为它添加新功能。


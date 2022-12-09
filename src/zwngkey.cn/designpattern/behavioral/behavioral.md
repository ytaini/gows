# 行为型模式

行为型模式是软件设计模式的一类，关注于对象之间的通信和协作。它们可以让你更好地控制对象之间的交互，以实现更复杂的业务逻辑。

- 行为型模式(Behavioral Pattern) 关注系统中对象之间的交互，研究系统在运行时对象之间的相互通信与协作，进一步明确对象的职责
- 行为型模式：不仅仅关注类和对象本身，还重点关注它们之间的相互作用和职责划分

> 类行为型模式

使用继承关系在几个类之间分配行为，主要通过多态等方式来分配父类与子类的职责

> 对象行为型模式

使用对象的关联关系来分配行为，主要通过对象关联等方式来分配两个或多个类的职责

> 行为型模式一览表

| **模式名称**                                    | **定** **义**                                                | **学习难度** | **使用频率** |
| ----------------------------------------------- | ------------------------------------------------------------ | ------------ | ------------ |
| **职责链模式(Chain of Responsibility Pattern)** | **避免将一个请求的发送者与接收者耦合在一起，让多个对象都有机会处理请求。将接收请求的对象连接成一条链，并且沿着这条链传递请求，直到有一个对象能够处理它为止。** | ★★★☆☆        | ★★☆☆☆        |
| **命令模式(Command Pattern)**                   | **将一个请求封装为一个对象，从而让你可以用不同的请求对客户进行参数化，对请求排队或者记录请求日志，以及支持可撤销的操作。** | ★★★☆☆        | ★★★★☆        |
| **解释器模式(Interpreter Pattern)**             | **给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。** | ★★★★★        | ★☆☆☆☆        |
| **迭代器模式(IteratorPattern)**                 | **提供一种方法顺序访问一个聚合对象中的各个元素，且不用暴露该对象的内部表示。** | ★★★☆☆        | ★★★★★        |
| **中介者模式(Mediator Pattern)**                | **定义一个对象来封装一系列对象的交互。中介者模式使各对象之间不需要显式地相互引用，从而使其耦合松散，而且让你可以独立地改变它们之间的交互。** | ★★★☆☆        | ★★☆☆☆        |
| **备忘录模式(Memento Pattern)**                 | **在不破坏封装的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态，这样可以在以后将对象恢复到原先保存的状态。** | ★★☆☆☆        | ★★☆☆☆        |
| **观察者模式(Observer Pattern)**                | **定义对象之间的一种一对多依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象都得到通知并被自动更新。** | ★★★☆☆        | ★★★★★        |
| **状态模式(State Pattern)**                     | **允许一个对象在其内部状态改变时改变它的行为。对象看起来似乎修改了它的类。** | ★★★☆☆        | ★★★☆☆        |
| **策略模式(Strategy Pattern)**                  | **定义一系列算法，将每一个算法封装起来，并让它们可以相互替换，策略模式让算法可以独立于使用它的客户变化。** | ★☆☆☆☆        | ★★★★☆        |
| **模板方法模式(Template Method Pattern)**       | **定义一个操作中算法的框架，而将一些步骤延迟到子类中。模板方法模式使得子类不改变一个算法的结构即可重定义该算法的某些特定步骤。** | ★★☆☆☆        | ★★★☆☆        |
| **访问者模式(Visitor Pattern)**                 | **表示一个作用于某对象结构中的各个元素的操作。访问者模式让你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。** | ★★★★☆        | ★☆☆☆☆        |


例如，假设你正在开发一个应用程序，该应用程序需要在用户提交表单时进行多个验证。在这种情况下，你可以使用责任链模式来处理验证过程。该模式允许你将验证过程分解成多个独立的处理器，每个处理器都负责一个特定的验证步骤。然后，你可以将这些处理器链接起来，当用户提交表单时，每个处理器都会按顺序执行验证，直到所有验证都通过或者遇到验证失败的情况。这种方法可以让你更好地控制验证过程，并且可以更方便地扩展.
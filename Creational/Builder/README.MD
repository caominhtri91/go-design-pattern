# Builder design pattern - reusing an algorithm to create many implementations of an interface

## A Builder design pattern tries to:

- Abstract complex creations so that object creation is separated from the object user
- Create an object step by step by filling its fields and creating the embedded objects
- Reuse the object creation algorithm between many objects

## Description

The Builder design pattern helps us maintain an unpredictable number of products by using a common construction algorithm that is used by the director. The construction process is always abstracted from the user of the product.

At the same time, having a defined construction pattern helps when a newcomer to our source code needs to add a new product to the pipeline. The BuildProcess interface specifies what he must comply to be part of the possible builders.

However, try to avoid the Builder pattern when you are not completely sure that the algorithm is going to be more or less stable because any small change in this interface will affect all your builders and it could be awkward if you add a new method that some of your builders need and others Builders do not.

## My Description

Mày ko cần phải khởi tạo struct và set giá trị cho nó bằng tay, chỉ cần gọi builder là nó sẽ trả lại cho mày struct có giá trị. Nó abstract toàn bộ logic khởi tạo struct, mày ko cần phải biết vì sao cái struct đó có giá trị như vậy, chỉ cần biết gọi builder là nó trả về cho mày 1 struct có thể dùng được

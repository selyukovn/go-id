
# Id

### TL;DR

This package contains basic code for artificial `Id`s based on `Uint` and `UUID`-strings.

### Here's the thing:

Unlike, for example, the type `OrderNumber`,  
whose set of values is defined by business rules in accordance with the real world  
(e.g., `{store-code}-{date:yyyyMMdd}-{daily-order-sequence:6-digits}` — `EB-20251231-000001`)  
and differs in each specific case,  
the type `Id` is most often based on a set dictated by data storage systems.

Usually these are sets of `UUID` or `unsigned int` values.

> Just in case:
>
> Using UUID or Uint types inside `Id` does not equate these types with `Id`,  
> but merely restricts the set of valid values for this type,  
> as if it were defined by business rules, similar to an order number.
>
> UUID is NOT an implementation of `Id`, but simply a set of tools, like regex, which helps validate values.  
> Generation of new values in this case does not rely on, for example, a database and a table with an incremental field,  
> but on some algorithm that produces random values, unique in more than 99.99...99% of cases.  
> Here, UUID is just a ready‑made and "completely accidentally perfect for us" implementation of such an algorithm.

Often, objects require precisely such artificial identifiers due to the lack of natural ones.  
The code in this case will always be the same — i.e., duplicated.

A solution to the duplication problem is to move the type `Id` into the `common` package as a shared type, like `Email`.  
Here, another problem arises: `account.Id` and `product.Id` logically should be different types,  
but using `common.Id` violates this natural logic, which risks errors.  
For example, it is easy to swap arguments when calling the function `addToCart(aId account.Id, pId product.Id)`  
and learn about it not at compile time, but during testing or even from a user bug report.

A better option is to move the repeated code into a separate package,  
but use it as a base for identifiers of different object types, not directly as a type.  
In this case, you will also avoid duplication, but `account.Id` and `product.Id` will remain different types,  
a mistaken substitution of which will be immediately caught by the IDE, and even more so by the compiler.

### Example:

```go
package account

import "github.com/selyukovn/go-id/like_uuid"

// Id
// ------------------------

type Id like_uuid.Id

func IdFromString(value string) (Id, error) {
    id, err := like_uuid.IdFromString(value)
    return Id(id), err
}

func IdFromStringMust(value string) Id {
    return Id(like_uuid.IdFromStringMust(value))
}

func (id Id) String() string {
    return like_uuid.Id(id).String()
}

// Generator
// ------------------------

type IdGenerator interface {
    Generate() (Id, error)
}

// Generator implementation
// ------------------------

type IdGeneratorImplUniqueRandom struct {
    internal like_uuid.IdGeneratorUniqueRandom
}

func (g *IdGeneratorImplUniqueRandom) Generate() (Id, error) {
    id, err := g.internal.Generate()
    return Id(id), err
}
```

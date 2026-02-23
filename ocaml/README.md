# Data Query

A small scripting language for querying datasets.

```
$columns['Category'].
```

## API

- `List`
    - `.where(Func(Any) -> Bool) -> List`: returns a list of matches
- `Map`
    - `[String] -> Any`: retrieve value from map
    - `[String] =(Any) -> Void`: assign a value in a map
    - `.each(Func(String, Any) -> Void) -> Void`:

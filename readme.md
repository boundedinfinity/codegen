# Canonical specification

-   `types[]`:

    List of types.

-   `data[]`:

    List of data to pre-population for a given set of types.  The data
    specification is:
    -   `type`: Reference to a type
    -   An instance of the type.

-   `operations[]`:

    List of operations.  Operations have the following format:
    -   `id`: 
    
        The ID of the operation
    -   `name`:

        The name of the operation.
    -   `inputs[]`:

        List of type references for the inputs to the operation.
    -   `outputs[]`:

        List of type references for the outputs of the operation.

# Type specification

## Common type specifications

-   `id`:

    The type ID.

    required if this is a type definition.  

-   `name`: 

    The name of the type.
    
    required: **false**

    If the name is not given. The name of this type if required
    will be generated from the type ID.

-   `type`: 

    The type ID of the the type.

    required: **true**

    The type ID can be one of the following built-in types basic types:
    - `integer`
    - `float`
    - `boolean`
    - `string`
    - `time`
    - `date`
    - `date-time`
    - `duration`
    - `uuid`

    or one of the compound types:
    - `object`
    - `array`
    - `enum`
    - `range`
    - `ref`

    or a URL encoded reference to another type's `type` parameter.

-   `version`: 

    Version of this item.

    required: **false**
-   `upgrade`:  

    A reference to another specification as the upgraded, or
    the next version, of this type

    required: **false**
-   `downgrade`:  

    A reference to another specification as the downgraded, or
    previous version, of this type

    required: **false**
-   `required`: 

    A boolean which determines if this item is required or not

    required: **false**
-   `default`: 

    If provided, the type specific value which will be assigned
    if not value is given.

    required: **false**

-   `example[]`: 

    If provided, list of the type specific examples for a given value.

    required: **false**

## Bool type specification

There are no additional items for this type

## Integer type speficiation

-   `ranges[]` 
    
    The list of range constraints for the number

    This contains a list of the following sub-constraints:

    -   `min`: 
        
        The inclusive minimum constraint for the number

        This item is mutually exclusive with the `exclusive-min` constraint
    -   `max`: 
        
        The inclusive maximum for the number

        This item is mutually exclusive with the `exclusive-max` constraint
    -   `exclusive-min`: The exclusive minimum for the number

        This item is mutually exclusive with the `min` constraint
    -   `exclusive-max`: The exclusive maximum for the number

        This item is mutually exclusive with the `max` constraint
-   `multiple-of`: Constrains the number to be a multiple of this value.
-   `negative`: Constrains the number to be less than `0`

    This constraints is mutually exclusive with the `positive` constraint.
-   `positive`: Constrains the number to be greater than `0`

    This constraints is mutually exclusive with the `negative` constraint.
-   `none-of`: 

    List of integers which are not value.

    NOTE: These values are combined with values for from other contraints.
-   `one-of`: 

    List of integrars which are considered valid.

    NOTE: These values are combined with values for from other contraints.

## Float type speficiation

The `float` constraints are the same as for the `integer` type with the
following additions:
-   `tolerance`:

    Is an equality constraint which calculates the difference
    between another constraint and if the difference is below the `tolerance`
    the equality is assumed to be true.
-   `precision`:

    Comparse a number down to a certain number of places below
    the period.

## String type specification
-   `min`: 

    The inclusive minimum length of the string
-   `max`: 

    The inclusive maximum lenght of the string
-   `regex`: 

    A regular expression to which the string must conform.
-   `abnf`: 

    An ABNF expression to which the string must conform
-   `includes`: 

    List of sub-string items which must appear in the string
-   `excludes`: 

    List of sub-string items which must not appear in the string
-   `one-of`: 

    Value must match a value in the list.
-   `none-of`: 

    List of sub-string items which must not appear in the string

## Enum type specification
-   `items`:

    List if sub-items in these enumeration

    -   `name`:
    
        The name of the item
    -   `values[]`:

        The list of values that represent this item.  
        
        The first item in the list should be the primary item.
    -   `description`:

        The description of this item.

## Array type specification

-   `min`: 

    The inclusive minimum length of the array
-   `max`: 

    The inclusive maximum lenght of the array


## Object type specification

-   `min`: 

    The inclusive minimum length of the array
-   `max`: 

    The inclusive maximum lenght of the array

# Reference:
- https://www.learnjsonschema.com/2020-12/core/
- https://json-schema.org/understanding-json-schema/reference/string
- https://zod.dev/?id=table-of-contents
- https://entgo.io/
- https://spec.openapis.org/oas/v3.1.0
- https://json-schema.org/understanding-json-schema/
- https://www.asyncapi.com/
- https://www.baeldung.com/raml-restful-api-modeling-language-tutorial
- https://apiblueprint.org/
- https://grpc.io/
- https://developers.google.com/protocol-buffers/
- https://www.uml.org/
- https://raml.org/
- https://avro.apache.org/
- https://avro.apache.org/docs/1.11.1/specification/
- https://typeschema.org/
- https://github.com/Masterminds/sprig
- https://developers.google.com/protocol-buffers
- https://go-ozzo.github.io/ozzo-validation/
- https://www.jsonrpc.org/

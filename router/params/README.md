# Params options

## Source types (`from:""`)

* `url`: The param is part of the URL as in `/item/your-param`.
* `query`: The param is part of the query string as in `item?id=your-param`.
* `form`: The param is part of the body of the request. It can be from a JSON payload or a basic form-urlencoded payload.
* `file`: The param is a file sent using `multipart/form-data`. The param type MUST be a `*formfile.FormFile`.

## Params type (`params:""`)

### global options (works on most of the types)

* `required`: The field is required and an error will be returned if the field is empty.

### String specific params

* `trim`: The value will be trimmed of its trailing spaces.
* `uuid`: The value is required to be a valid UUIDv4.
* `url`: The value is required to be a valid http(s) url.
* `email`: The value is required to be a valid email.

### Pointers specific params

* `noempty`: The value cannot be empty. The pointer can be nil, but if a value is provided it cannot be an empty string. the difference with `required` is that `required` does not accept nil pointer.

### Files specific params

* `image`: The provided file is required to be an image.

## Ignoring and naming `json:""`

* Use `json:"_"` to prevent a field to be altered or checked.
* Use `json:"field_name"` to name a field.

## Default value

Use `default:"my_value"` to set a default value. The default value will be use if nothing is found in the payload or if the provided value is an empty string.

## Maxlen of a string

Use `maxlen:"255"` to make sure the len of a string is not bigger than 255 char. Any invalid values (including `0`) will be ignored

## Custom Validator

You can add a custom validator by implementing `params.CustomValidation`.

## Examples

```golang
type UpdateParams struct {
  ID        string  `from:"url" json:"id" params:"required,uuid"`
  Name      *string `from:"form" json:"name" params:"trim,noempty"`
  ShortName *string `from:"form" json:"short_name" params:"trim"`
  Website   *string `from:"form" json:"website" params:"url,trim"`
  InTrash   *bool   `from:"form" json:"in_trash"`
}
```
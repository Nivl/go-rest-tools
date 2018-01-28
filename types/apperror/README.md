# AppError

## How to deal with errors

A known and common way to deal with errors is to [handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully). Unfortunately
in an API this is not that easy and the error still needs a **context** (Status Code, field name, etc.) to let the clients know why a request failed. The best way to handle this is by dealing with errors in the lowest possible function (as opposite to the highest function), because the function failing knows the best _why_ it's failing.

If a higher level function wants to handle the error, `apperror.Is*` can be used to determine the kind of the error.

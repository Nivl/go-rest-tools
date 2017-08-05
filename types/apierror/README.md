# APIError

## How to deal with errors

A known and common way to deal with errors is to [handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully). Unfortunately
in a REST API this is not that easy an the error still needs a **context** (HTTP Status, Internal code, field name, etc.) to let the clients know why it failed. The best way to handle this is by dealing with errors in the lowest possible function (as opposite to the highest function), because the function failing knows the best _why_ it's failing.

Not that a function should **only** create an error using the `apierror.New*` functions, and should never set the HTTP Status code itself. This is because it makes sense for a function to return `apierror.NotFound()`, but not to know what `NotFound()` actually implies for the response. If no `apierror.New*` suite your need you should then just return the raw error.

If a higher level function wants to handle the error, `apierror.Is*` can be used to determine the kind of the error.
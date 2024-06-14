using System.Net;
using Microsoft.AspNetCore.Mvc;
using Results.Application;
using Results.Domain;

namespace RentalModule.API.Utilities;

public static class ControllerHelper
{
    public static IActionResult ErrorResult(this ControllerBase controller, IErrorResult errorResult) => errorResult switch
    {
        EntityNotFoundErrorResult => controller.NotFound(errorResult),
        EntityAlreadyExistsErrorResult => controller.Conflict(errorResult),
        ValidationErrorResult => controller.BadRequest(errorResult),
        _ => controller.StatusCode((int)HttpStatusCode.InternalServerError, errorResult)
    };
}
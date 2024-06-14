using System.Net;
using Microsoft.AspNetCore.Mvc;
using Results.Application;
using Results.Domain;
using Results.Infrastructure;

namespace BlogModule.API.Utilities;

public static class ControllerHelper
{
    public static IActionResult ErrorResult(this ControllerBase controller, IErrorResult errorResult) => errorResult switch
    {
        EntityNotFoundErrorResult => controller.NotFound(errorResult),
        EntityAlreadyExistsErrorResult => controller.Conflict(errorResult),
        ValidationErrorResult => controller.BadRequest(errorResult),
        TaskCanceledErrorResult => controller.StatusCode(418, errorResult),
        _ => controller.StatusCode((int)HttpStatusCode.InternalServerError, errorResult)
    };
}
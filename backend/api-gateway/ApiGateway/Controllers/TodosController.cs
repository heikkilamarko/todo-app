using ApiGateway.Models;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using ApiGateway.Services;

namespace ApiGateway.Controllers
{
    [ApiController]
    [Route("todos")]
    public class TodosController : ControllerBase
    {
        private readonly TodoApiClient _api;

        public TodosController(TodoApiClient api)
        {
            _api = api;
        }

        [HttpGet]
        public async Task<ActionResult> GetTodos([FromQuery] GetTodosQuery query)
        {
            var data = await _api.GetTodos(Request.QueryString.Value);

            return Ok(data);
        }

        [HttpPost]
        public async Task<ActionResult> CreateTodo([FromBody] NewTodo todo)
        {
            await _api.CreateTodo(todo);

            return Ok();
        }
    }
}
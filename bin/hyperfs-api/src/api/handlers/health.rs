extern crate serde;
extern crate serde_json;
extern crate iron;

use iron::{Handler, Request, Response, IronResult, status};
use iron::mime::Mime;

#[derive(Serialize, Debug)]
struct HealthResponse {
    status: bool,
}

pub struct HealthHandler;

impl Handler for HealthHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let health = HealthResponse { status: true };

        let response = serde_json::to_string(&health).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

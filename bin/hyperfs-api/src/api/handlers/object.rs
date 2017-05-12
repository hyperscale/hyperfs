extern crate serde;
extern crate serde_json;
extern crate iron;

use iron::{Handler, Request, Response, IronResult, status};
use iron::mime::Mime;

#[derive(Serialize, Debug)]
struct ObjectResponse {
    status: bool,
}

pub struct ObjectListHandler;

impl Handler for ObjectListHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let object = ObjectResponse { status: true };

        let response = serde_json::to_string(&object).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct ObjectCreateHandler;

impl Handler for ObjectCreateHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let object = ObjectResponse { status: true };

        let response = serde_json::to_string(&object).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct ObjectDetailHandler;

impl Handler for ObjectDetailHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let object = ObjectResponse { status: true };

        let response = serde_json::to_string(&object).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct ObjectDeleteHandler;

impl Handler for ObjectDeleteHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let object = ObjectResponse { status: true };

        let response = serde_json::to_string(&object).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

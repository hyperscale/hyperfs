extern crate serde;
extern crate serde_json;
extern crate iron;

use iron::{Handler, Request, Response, IronResult, status};
use iron::mime::Mime;

#[derive(Serialize, Debug)]
struct BucketResponse {
    status: bool,
}

pub struct BucketListHandler;

impl Handler for BucketListHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct BucketCreateHandler;

impl Handler for BucketCreateHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct BucketEditHandler;

impl Handler for BucketEditHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct BucketDetailHandler;

impl Handler for BucketDetailHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct BucketDeleteHandler;

impl Handler for BucketDeleteHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

pub struct BucketDeleteBulkHandler;

impl Handler for BucketDeleteBulkHandler {
    fn handle(&self, _: &mut Request) -> IronResult<Response> {

        let bucket = BucketResponse { status: true };

        let response = serde_json::to_string(&bucket).unwrap();

        let content_type = "application/json".parse::<Mime>().unwrap();
        Ok(Response::with((content_type, status::Ok, response)))
    }
}

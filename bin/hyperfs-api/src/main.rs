#[macro_use]
extern crate clap;
#[macro_use]
extern crate serde_derive;
extern crate config;
extern crate iron;
extern crate router;
extern crate iron_cors;

mod api;

use self::api::handlers::health::HealthHandler;
use self::api::handlers::bucket::{
    BucketListHandler,
    BucketCreateHandler,
    BucketEditHandler,
    BucketDetailHandler,
    BucketDeleteHandler,
    BucketDeleteBulkHandler
};
use self::api::handlers::object::{
    ObjectListHandler,
    ObjectCreateHandler,
    ObjectDetailHandler,
    ObjectDeleteHandler
};
use std::env;
use config::{Config, File, FileFormat, Environment};
use clap::App;
use iron::{Iron, Request, Response, IronResult, Chain, status};
// use iron::prelude::*;
use router::Router;
use iron_cors::CorsMiddleware;

fn main() {
    // Create a new local configuration
    let mut c = Config::new();

    c.set_default("server.host", "127.0.0.1").unwrap();
    c.set_default("server.port", "8080").unwrap();

    // Add 'Settings.toml'
    c.merge(File::new("config", FileFormat::Toml).required(false))
        .unwrap();

    // Add 'Settings.$(RUST_ENV).toml`
    let name = format!("config.{}", env::var("env").unwrap_or("development".into()));

    c.merge(File::new(&name, FileFormat::Toml).required(false))
        .unwrap();

    // Add environment variables that begin with HYPERFS_
    c.merge(Environment::new("HYPERFS")).unwrap();

    App::new(crate_name!())
        .version(crate_version!())
        .author(crate_authors!("\n"))
        .about(crate_description!())
        .get_matches();

    let mut router = Router::new();

    // Health
    router.get("/health", HealthHandler, "health");

    // Bucket
    router.get("/buckets", BucketListHandler, "bucket_list");
    router.post("/buckets", BucketCreateHandler, "bucket_create");
    router.put("/buckets/:id", BucketEditHandler, "bucket_edit");
    router.get("/buckets/:id", BucketDetailHandler, "bucket_detail");
    router.delete("/buckets/:id", BucketDeleteHandler, "bucket_delete");
    router.delete("/buckets", BucketDeleteBulkHandler, "bucket_delete_bulk");

    // Object
    router.get("/:bucket", ObjectListHandler, "object_list");
    router.post("/:bucket/*object", ObjectCreateHandler, "object_create");
    router.get("/:bucket/*object", ObjectDetailHandler, "object_detail");
    router.delete("/:bucket/*object", ObjectDeleteHandler, "object_delete");

    router.get("/", handler, "index");
    router.get("/:query", handler, "query");

    let cors_middleware = CorsMiddleware::with_allow_any(true);

    // Setup chain with middleware
    let mut chain = Chain::new(router);
    chain.link_around(cors_middleware);

    let listen = format!("{}:{}",
                         c.get_str("server.host").unwrap(),
                         c.get_str("server.port").unwrap());

    println!("Starting new server on {:?}...", listen);
    Iron::new(chain).http(listen).unwrap();

    fn handler(req: &mut Request) -> IronResult<Response> {
        let ref query = req.extensions
            .get::<Router>()
            .unwrap()
            .find("query")
            .unwrap_or("/");
        Ok(Response::with((status::Ok, *query)))
    }
}

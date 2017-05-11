#[macro_use]
extern crate clap;

use clap::{App, SubCommand};

mod command;

use command::status;

fn main() {
    let matches = App::new(crate_name!())
        .version(crate_version!())
        .author(crate_authors!("\n"))
        .about(crate_description!())
        .subcommand(SubCommand::with_name("status")
            .about("gets information on the running workers"))
        .get_matches();

    match matches.subcommand() {
        ("status", Some(_)) => {
            status();
        }
        _ => println!("unknown subcommand"),
    }
}

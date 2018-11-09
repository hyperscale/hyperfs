#[macro_use]
extern crate clap;

use clap::{
    App,
    SubCommand,
    AppSettings,
    Arg
};

mod command;

use self::command::cluster::status::{
    status_command
};

use self::command::file::upload::{
    upload_command
};

fn main() {
    let matches = App::new(crate_name!())
        .global_setting(AppSettings::ColoredHelp)
        .global_setting(AppSettings::ColorAuto)
        .setting(AppSettings::SubcommandRequired)
        .setting(AppSettings::AllowExternalSubcommands)
        .version(crate_version!())
        .author(crate_authors!("\n"))
        .about(crate_description!())
        .subcommand(SubCommand::with_name("file")
            .setting(AppSettings::SubcommandRequired)
            .about("manage file")
            .subcommand(SubCommand::with_name("upload")
                .about("Upload file to cluster")
                .arg(Arg::with_name("from")
                    .takes_value(true)
                    .help("the path of local file")
                    .required(true))
                .arg(Arg::with_name("to")
                    .takes_value(true)
                    .help("the path of remote file")
                    .required(true))))
        .subcommand(SubCommand::with_name("cluster")
            .setting(AppSettings::SubcommandRequired)
            .about("manage cluster")
            .subcommand(SubCommand::with_name("status")
                .about("gets information on the cluster")))
        .get_matches();

    match matches.subcommand() {
        ("file", Some(file)) => {
            match file.subcommand() {
                ("upload", Some(upload)) => {
                    upload_command(upload);
                }
                _ => println!("unknown command for file")
            }
        }
        ("cluster", Some(cluster)) => {
            match cluster.subcommand() {
                ("status", Some(_)) => {
                    status_command();
                }
                _ => println!("unknown command for cluster")
            }
        }
        _ => println!("unknown subcommand")
    }
}

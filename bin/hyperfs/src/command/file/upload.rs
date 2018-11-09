
use std::fs::File;
use std::io::BufReader;
use std::io::prelude::*;
use clap::{
    ArgMatches
};

pub fn upload_command(matches: &ArgMatches) {
    let from = matches.value_of("from").unwrap();
    let to = matches.value_of("to").unwrap();

    println!("Value for from: {}", from);
    println!("Value for to: {}", to);

    let mut file = File::open(from).expect("file not found");
    // let mut buf_reader = BufReader::new(file);

    let mut buffer = [0; 10];

    // read up to 10 bytes
    file.read(&mut buffer[..]).expect("cannot read file");
}

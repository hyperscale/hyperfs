extern crate tree_magic;

use std::fs::File;
use std::path::Path;


include!(concat!(env!("OUT_DIR"), "/file.rs"));

impl<'a> From<String> for Metadata {
    fn from(filename: String) -> Metadata {
        let p = Path::new(&filename);

        let file = File::open(p).expect("open file");

        let mut metadata = Self::from(file);

        metadata.filename = filename.clone();

        return metadata;
    }
}

impl<'a> From<File> for Metadata {
    fn from(f: File) -> Metadata {
        let md = f.metadata().expect("metadata");

        let result = tree_magic::from_u8(f.bytes());

        let metadata = Metadata {
            filename: String::from(""),
            size: md.len(),
            mime_type: result,
            chunks: Vec::new(),
        };

        return metadata;
    }
}

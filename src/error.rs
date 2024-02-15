use std::{num::ParseFloatError, string::FromUtf8Error};


#[derive(Debug,PartialEq,Eq)]
pub struct Error {
    msg: String,
}

impl Error {
    pub fn new(msg: String) -> Error {
        Error { msg }
    }
}

impl std::fmt::Display for Error {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_str(&self.msg)
    }
}

impl std::error::Error for Error {}

impl std::convert::From<ParseFloatError> for Error {
    fn from(value: ParseFloatError) -> Self {
        Error {
            msg: format!("{}", value),
        }
    }
}

impl std::convert::From<FromUtf8Error> for Error {
    fn from(value: FromUtf8Error) -> Self {
        Error {
            msg: format!("{}", value),
        }
    }
}

impl std::convert::From<std::io::Error> for Error {
    fn from(value: std::io::Error) -> Self {
        Error {
            msg: format!("{}", value),
        }
    }
}

pub type Result<T> = std::result::Result<T, Error>;

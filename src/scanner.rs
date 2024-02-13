use std::io::Read;
use crate::token::*;

pub type Result = std::io::Result<Token>;

pub struct Scanner<R: std::io::Read> {
    r: std::io::BufReader<R>,
    current: u8,
    next: u8,
    number_buffer: Vec<u8>,
    first_time: bool,
}

impl<R: std::io::Read> Scanner<R> {
    pub fn new(r: R) -> Scanner<R> {
        Scanner {
            r: std::io::BufReader::new(r),
            current: 0,
            next: 0,
            number_buffer: Vec::new(),
            first_time: true, 
        }
    }

    fn advance(&mut self) -> std::io::Result<()> {
        let mut buf: [u8; 1] = [0; 1];

        let r = self.r.read(&mut buf)?;

        self.current = self.next;    
        if r == 0 {
            self.next = 0;
        } else {
            self.next = buf[0];
        }

        Ok(())
    }

    fn resolve_number(&mut self) -> Option<Result> {
        if self.number_buffer.len() == 0 {
            return None;
        }

        let b = self.number_buffer.clone();
        self.number_buffer.clear();

        match String::from_utf8(b) {
            Err(e) => Some(Err(std::io::Error::new(
                std::io::ErrorKind::InvalidData,
                format!("Invalid utf8 bytes: {}", e),
            ))),
            Ok(s) => match s.parse::<f64>() {
                Err(e) => Some(Err(std::io::Error::new(
                    std::io::ErrorKind::InvalidData,
                    format!("Invalid number: {}", e),
                ))),
                Ok(val) => Some(Ok(Token::Number(val))),
            },
        }
    }
}

impl<R: std::io::Read> Iterator for Scanner<R> {
    type Item = Result;

    fn next(&mut self) -> Option<Self::Item> {
        if self.first_time {
            if let Err(e) = self.advance() {
                return Some(Err(e));
            }

            self.first_time = false;
        }

        loop {
            if let Err(e) = self.advance() {
                return Some(Err(e));
            }

            if self.current == 0 {
                return None
            }

            match self.current {
                b' ' | b'\t' | b'\r' | b'\n' => {                    
                    if self.number_buffer.len() == 0 {
                        // Skip whitespace if there is nothing left in the number buffer.
                        continue;
                    }
                    // If the number buffer contains some chars, handle them.
                    return self.resolve_number();
                },

                b'+' => return Some(Ok(Token::Operator(Operator::Add))),
                b'-' => return Some(Ok(Token::Operator(Operator::Sub))),
                b'*' => return Some(Ok(Token::Operator(Operator::Mul))),
                b'/' => return Some(Ok(Token::Operator(Operator::Div))),
                b'(' => return Some(Ok(Token::Parenthesis(Paren::Left))),
                b')' => return Some(Ok(Token::Parenthesis(Paren::Right))),

                x if (x >= b'0' && x <= b'9') || x == b'.' => {
                    self.number_buffer.push(x);
                    if (self.next < b'0' || self.next > b'9') && self.next != b'.' {
                        return self.resolve_number();   
                    }
                }

                x => {
                    return Some(Err(std::io::Error::new(
                        std::io::ErrorKind::InvalidData,
                        format!("unexpected input char: {x}"),
                    )))
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    macro_rules! scanner_test {
        ($name:ident, $input:expr $(, $t:expr)*) => {
            #[test]
            fn $name() {
                let b = $input.as_bytes();
                let mut s = Scanner::<&[u8]>::new(b);
                $(
                    {
                        let v = s.next();
                        match v {
                            Some(r) => match r {
                                Ok(t) => assert_eq!(t, $t),
                                Err(e) => assert!(false, "expected Ok({:?}) but got Err({:?})", stringify!($t), e),
                            },
                            None => assert!(false, "expected Some(token) but got None"),
                        }
                    }
                )*;
                {
                    if let Some(t) = s.next() {
                        assert!(false, "Expected no more tokens but got {:?}", t);
                    }
                };
            }
        };
    }

    scanner_test! {operator_add, "+", Token::Operator(Operator::Add)}
    scanner_test! {operator_add_w_spaces, " +  ", Token::Operator(Operator::Add)}
    scanner_test! {operator_sub, "-", Token::Operator(Operator::Sub)}
    scanner_test! {operator_mul, "*", Token::Operator(Operator::Mul)}
    scanner_test! {operator_div, "/", Token::Operator(Operator::Div)}
    scanner_test! {operator_lparen, "(", Token::Parenthesis(Paren::Left)}
    scanner_test! {operator_rparen, ")", Token::Parenthesis(Paren::Right)}
    scanner_test! {operator_number, "2.34", Token::Number(2.34)}
    scanner_test! {simple_expr_0, "2+3", 
        Token::Number(2.0),
        Token::Operator(Operator::Add),
        Token::Number(3.0)
    }
    scanner_test! {simple_expr_1, " 2   + 3  *  4  ", 
        Token::Number(2.0),
        Token::Operator(Operator::Add),
        Token::Number(3.0),
        Token::Operator(Operator::Mul),
        Token::Number(4.0)
    }

    scanner_test! {paren_expr, "2*(3+4)", 
        Token::Number(2.0),
        Token::Operator(Operator::Mul),
        Token::Parenthesis(Paren::Left),
        Token::Number(3.0),
        Token::Operator(Operator::Add),
        Token::Number(4.0),
        Token::Parenthesis(Paren::Right)
    }


    // #[test]
    // fn test_scanner_iter() {
    //     let b = "+".as_bytes();
    //     let mut s = Scanner::<&[u8]>::new(b);
    //     let v = s.next();
    //     match v {
    //         Some(r) => match r {
    //             Ok(t) => assert_eq!(t, Token::Operator(Operator::Add)),
    //             _ => assert!(false),
    //         },
    //         None => assert!(false),
    //     }
    // }
}

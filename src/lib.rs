
pub mod token;
pub mod scanner;

struct InfixToRPNTranslator<I>
where I: Iterator<Item=scanner::Result> {
    iter: I,
    out: Vec<token::Token>,
    operators: Vec<token::Token>,
}

impl <I> InfixToRPNTranslator<I>
where I: Iterator<Item=scanner::Result> {
    fn new(i: I) -> InfixToRPNTranslator<I> {
        InfixToRPNTranslator{ 
            iter: i,
            out: Vec::new(),
            operators: Vec::new(),
        }
    }
}

impl <I> Iterator for InfixToRPNTranslator<I>
where I: Iterator<Item=scanner::Result> {
    type Item = scanner::Result;

    fn next(&mut self) -> Option<Self::Item> {
        if !self.out.is_empty() {
            return Some(Ok(self.out.remove(0)));
        }

        match self.iter.next() {            
            None => {
                if !self.operators.is_empty() {
                    while let Some(t) = self.operators.pop() {
                        self.out.push(t);
                    }
                    return self.next()
                } else {
                    None
                }
            },
            Some(res) => match res {
                Err(e) => Some(Err(e)),
                Ok(tok) => match tok {
                    token::Token::Number(_) => Some(Ok(tok)),
                    token::Token::Operator(op) => {
                        if !self.operators.is_empty() {
                            match self.operators[self.operators.len()-1] {
                                token::Token::Operator(op2) => {
                                    if op2.precedence() >= op.precedence() {
                                        self.out.push(self.operators.pop().unwrap());
                                    }
                                }
                                _ => panic!("not implemented"),
                            }
                        }
                        self.operators.push(tok);
                        self.next()
                    },
                    _ => panic!("not implemented"),
                },
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::scanner::Scanner;

    macro_rules! rpn_test {
        ($name:ident, $input:expr $(, $tok:expr)+) => {
            #[test]
            fn $name () {
                let b = $input.as_bytes();
                let s = Scanner::new(b);
                let mut rpn = InfixToRPNTranslator::new(s);
                $(
                    {
                        match rpn.next() {
                            None => assert!(false, "expected Some({}) but got none", stringify!($tok)),
                            Some(t) => match t {
                                Ok(t) => assert_eq!(t, $tok),
                                Err(e) => assert!(false, "Expected {} but got Err({})", stringify!($tok), e),
                            }
                        }
                    }
                )*;
                {
                    if let Some(t) = rpn.next() {
                        assert!(false, "Expected no more tokens but got {:?}", t);
                    }
                };                
            }
        };
    }

    rpn_test!{simple, "1+2", 
        token::Token::Number(1.0),
        token::Token::Number(2.0),
        token::Token::Operator(token::Operator::Add)
    }
    rpn_test!{precedence, "1+2*3", 
        token::Token::Number(1.0),
        token::Token::Number(2.0),
        token::Token::Number(3.0),
        token::Token::Operator(token::Operator::Mul),
        token::Token::Operator(token::Operator::Add)
    }
}

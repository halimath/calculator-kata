pub mod error;
pub mod scanner;
pub mod token;

pub fn eval<I: std::io::Read>(input: I) -> error::Result<f64> {
    let mut stack = Vec::<f64>::new();

    let s = scanner::Scanner::new(input);
    let rpn = InfixToRPNTranslator::new(s);

    for tok in rpn {
        match tok? {
            token::Token::Number(n) => stack.push(n),
            token::Token::Operator(op) => {
                if stack.len() < 2 {
                    return Err(error::Error::new("stack underrun".to_owned()));
                }
                let right = stack.pop().unwrap();
                let left = stack.pop().unwrap();

                match op {
                    token::Operator::Add => stack.push(left + right),
                    token::Operator::Mul => stack.push(left * right),
                    token::Operator::Sub => stack.push(left - right),
                    token::Operator::Div => stack.push(left / right),
                }
            },
            _ => panic!("not implemented"),
        }
    }

    stack.pop().ok_or(error::Error::new("stack underrun".to_owned()))
}

type Result = error::Result<token::Token>;

struct InfixToRPNTranslator<I>
where
    I: Iterator<Item = Result>,
{
    iter: I,
    out: Vec<token::Token>,
    operators: Vec<token::Token>,
}

impl<I> InfixToRPNTranslator<I>
where
    I: Iterator<Item = Result>,
{
    fn new(i: I) -> InfixToRPNTranslator<I> {
        InfixToRPNTranslator {
            iter: i,
            out: Vec::new(),
            operators: Vec::new(),
        }
    }
}

impl<I> Iterator for InfixToRPNTranslator<I>
where
    I: Iterator<Item = Result>,
{
    type Item = Result;

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
                    return self.next();
                } else {
                    None
                }
            }

            Some(res) => match res {
                Err(e) => Some(Err(e)),
                Ok(tok) => match tok {
                    token::Token::Number(_) => Some(Ok(tok)),
                    token::Token::Operator(op) => {
                        if !self.operators.is_empty() {
                            while !self.operators.is_empty() {
                                match self.operators[self.operators.len() - 1] {
                                    token::Token::Operator(op2) => {
                                        if op2.precedence() >= op.precedence() {
                                            self.out.push(self.operators.pop().unwrap());
                                        } else {
                                            break;
                                        }
                                    }
                                    _ => break,
                                }
                            }
                        }
                        self.operators.push(tok);
                        self.next()
                    }
                    token::Token::Parenthesis(token::Paren::Left) => {
                        self.operators.push(tok);
                        self.next()
                    }
                    token::Token::Parenthesis(token::Paren::Right) => {
                        while let Some(tok) = self.operators.pop() {
                            if let token::Token::Parenthesis(token::Paren::Left) = tok {
                                return self.next();
                            }
                            self.out.push(tok);
                        }

                        Some(Err(error::Error::new("missing (".to_owned())))
                    }
                },
            },
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::scanner::Scanner;
    use crate::token::*;

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
                            None => assert!(false, "expected Some({:?}) but got none", stringify!($tok)),
                            Some(t) => match t {
                                Ok(t) => assert_eq!(t, $tok, "expected {:?} but got {:?}", stringify!($tok), t),
                                Err(e) => assert!(false, "Expected {:?} but got Err({:?})", stringify!($tok), e),
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

    rpn_test! {simple, "1+2",
        Token::Number(1.0),
        Token::Number(2.0),
        Token::Operator(Operator::Add)
    }
    rpn_test! {precedence, "1+2*3",
        Token::Number(1.0),
        Token::Number(2.0),
        Token::Number(3.0),
        Token::Operator(Operator::Mul),
        Token::Operator(Operator::Add)
    }

    rpn_test! {chain, "2+3*4-5",
        Token::Number(2.0),
        Token::Number(3.0),
        Token::Number(4.0),
        Token::Operator(Operator::Mul),
        Token::Operator(Operator::Add),
        Token::Number(5.0),
        Token::Operator(Operator::Sub)
    }
    rpn_test! {parenthesis, "2+3*(4-5)",
        Token::Number(2.0),
        Token::Number(3.0),
        Token::Number(4.0),
        Token::Number(5.0),
        Token::Operator(Operator::Sub),
        Token::Operator(Operator::Mul),
        Token::Operator(Operator::Add)
    }

    macro_rules! eval_test {
        ($name:ident, $input:expr, $want:expr) => {
            #[test]
            fn $name() {
                let got = eval($input.as_bytes()).unwrap();
                assert_eq!($want, got);
            }
        };
    }

    eval_test!(add, "2+3", 5.0);
    eval_test!(add_mul, "2+3*4", 14.0);
    eval_test!(paren, "(2+3)*4", 20.0);
    eval_test!(add_sub, "2+3-4", 1.0);
    eval_test!(add_div, "2+12/4", 5.0);
    eval_test!(funny_twitter_term, "8/2*(2+2)", 16.0);
}

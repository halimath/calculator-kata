#[derive(PartialEq, Eq, Clone, Copy, Debug)]
pub enum Operator {
    Add,
    Sub,
    Mul,
    Div,
}

impl Operator {
    pub fn precedence (&self) -> usize {
        match self {
            Operator::Add | Operator::Sub => 1,
            Operator::Mul | Operator::Div => 2,
        }
    }
}

#[derive(PartialEq, Eq, Clone, Copy, Debug)]
pub enum Paren {
    Left,
    Right,
}

#[derive(PartialEq, Clone, Debug)]
pub enum Token {
    Number(f64),
    Operator(Operator),
    Parenthesis(Paren),
}
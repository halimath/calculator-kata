use calc::eval;

fn main() {
    println!("{:.5}", eval(std::io::stdin()).unwrap())
}
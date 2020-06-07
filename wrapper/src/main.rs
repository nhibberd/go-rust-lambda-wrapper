use std::io::Write;

fn main() {
    loop {
        let mut buffer = String::new();
        std::io::stdin().read_line(&mut buffer).unwrap();
        std::io::stdout().write(buffer.as_bytes()).unwrap();
    }
}

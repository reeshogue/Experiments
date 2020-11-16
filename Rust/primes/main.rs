fn eratosthenes(limit : usize) -> Vec<usize> {
    let mut primes = vec![true; limit+1];
    primes[0] = false;
    if limit >= 1 {
        primes[1] = false;
    }
    for n in 2..limit {
        if primes[n] {
            let mut mult = n * n;
            while mult <= limit {
                primes[mult] = false;
                mult += n
            }
        }
    }

    let mut prime_list = Vec::new();
    for n in 0..limit {
        if primes[n] {
            prime_list.push(n);
        }
    }
    return prime_list;
}

fn main() {
    print!("{:?}", eratosthenes(100))
}
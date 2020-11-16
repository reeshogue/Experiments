struct Ant {
    x: usize,
    y: usize,
    dir: Direction 
}

#[derive(Copy, Clone)]
enum Direction {
    North,
    South,
    East,
    West
}

use Direction::*;
impl Ant {
    fn mv(&mut self, grid: &mut Vec<Vec<u8>>) {
        let pointer = &mut grid[self.x][self.y];
        match *pointer {
            0 => self.dir = self.dir.right(),
            1 => self.dir = self.dir.left(),
            _ => panic!(),
        }

        if *pointer == 0 {
            *pointer = 1;
        } else if *pointer == 1 {
            *pointer = 0;
        }

        match self.dir {
            North => self.y += 1,
            South => self.y -= 1,
            East => self.x -= 2,
            West => self.x += 2,
        }
    }
}

impl Direction {
    fn right(&self) -> Direction {
        match self {
            North => East,
            East => South,
            South => West,
            West => North,
        }
    }

    fn left(&self) -> Direction {
        self.right().right().right()
    }
}

fn main() {
    let mut grid: Vec<Vec<u8>> = vec![vec![0;120];120];
    let mut ant = Ant {
        x: 50,
        y: 50,
        dir: Direction::North
    };

    while ant.x < 120 && ant.y < 120 && ant.x > 0 && ant.y > 0 {
        ant.mv(&mut grid)
    }

    for i in grid.iter() {
        let printed = i.iter()
        .map(|&x| if x == 0 {" "} 
        else{"X"})
        .fold(String::new(), |x, y| x+y);
        println!("{}", printed);
    }
    print!("{}[2J", 27 as char);
}
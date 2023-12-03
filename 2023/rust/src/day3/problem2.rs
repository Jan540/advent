use std::collections::HashMap;

#[derive(Hash, Eq, PartialEq)]
struct Gear {
    x: usize,
    y: usize,
}

pub fn solve(lines: Vec<&str>) -> String {
    let mut sum = 0;
    let mut gears: HashMap<Gear, Vec<i32>> = HashMap::new();

    for (i, line) in lines.iter().enumerate() {
        let mut temp = String::new();

        for (j, char) in line.chars().enumerate() {
            if char.is_digit(10) {
                temp.push(char);
            }

            if (!char.is_digit(10) || j == line.len() - 1) && temp.len() > 0 {
                if let Some(gear) = is_adjecent_to_gear(&temp, j - temp.len(), i, &lines) {
                    let num = temp.parse::<i32>().expect("Failed to parse number :(");

                    match gears.get_mut(&gear) {
                        Some(nums) => nums.push(num),
                        None => {
                            gears.insert(gear, vec![num]);
                        }
                    }
                }

                temp.clear();
            }
        }
    }

    for (_, nums) in gears {
        if nums.len() == 2 {
            sum += nums[0] * nums[1];
        }
    }

    return sum.to_string();
}

fn is_adjecent_to_gear(num: &String, x: usize, y: usize, lines: &Vec<&str>) -> Option<Gear> {
    for dy in -1..2 {
        for dx in -1..num.len() as i32 + 1 {
            let x = x as i32 + dx;
            let y = y as i32 + dy;

            if y < 0 || x < 0 {
                continue;
            }

            let x = x as usize;
            let y = y as usize;

            if y >= lines.len() || x >= lines[y].len() {
                continue;
            }

            let char = lines[y].chars().nth(x).unwrap();

            if char == '*' {
                return Some(Gear { x, y });
            }
        }
    }

    None
}

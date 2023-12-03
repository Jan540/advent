pub fn solve(lines: Vec<&str>) -> String {
    let mut sum = 0;

    for (i, line) in lines.iter().enumerate() {
        let mut temp = String::new();

        for (j, char) in line.chars().enumerate() {
            if char.is_digit(10) {
                temp.push(char);
            }

            if (!char.is_digit(10) || j == line.len() - 1) && temp.len() > 0 {
                if is_adjecent_to_symbol(&temp, j - temp.len(), i, &lines) {
                    let num = temp.parse::<i32>().expect("Failed to parse number :(");
                    sum += num;
                }

                temp.clear();
            }
        }
    }

    return sum.to_string();
}

fn is_adjecent_to_symbol(num: &String, x: usize, y: usize, lines: &Vec<&str>) -> bool {
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

            if !char.is_digit(10) && char != '.' {
                return true;
            }
        }
    }

    false
}

pub fn solve(lines: Vec<&str>) -> String {
    let mut sum = 0;

    for line in lines.iter() {
        sum += process_card(line, &lines);
    }

    sum.to_string()
}

fn process_card(line: &str, lines: &Vec<&str>) -> i32 {
    let card: Vec<&str> = line.split(": ").collect();
    let id: i32 = card[0].split_whitespace().last().unwrap().parse().unwrap();
    let allnums: Vec<&str> = card[1].split(" | ").collect();
    let winning: Vec<&str> = allnums[0].split_whitespace().collect();

    let mut sum = 1;
    let mut count = id as usize;

    for num in allnums[1].split_whitespace() {
        if winning.contains(&num) {
            sum += process_card(lines[count], lines);
            count += 1;
        }
    }

    sum
}

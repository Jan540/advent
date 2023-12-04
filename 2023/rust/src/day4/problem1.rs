pub fn solve(lines: Vec<&str>) -> String {
    let mut sum = 0;

    for line in lines.iter() {
        let card: Vec<&str> = line.split(": ").collect();
        let allnums: Vec<&str> = card[1].split(" | ").collect();

        let winning: Vec<&str> = allnums[0].split_whitespace().collect();

        let mut points = 0;
        for num in allnums[1].split_whitespace() {
            if winning.contains(&num) {
                if points == 0 {
                    points = 1;
                } else {
                    points *= 2;
                }
            }
        }

        sum += points;
    }

    sum.to_string()
}

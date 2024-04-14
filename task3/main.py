import csv
import os

if __name__ == "__main__":
    # Read CSV files
    with open("day1.csv", newline='') as file1:
        reader1 = csv.reader(file1)
        records1 = list(reader1)

    with open("day2.csv", newline='') as file2:
        reader2 = csv.reader(file2)
        records2 = list(reader2)

    # Store visits in sets
    visits1 = set()
    visits2 = set()

    for record in records1:
        visits1.add(record[0] + "," + record[1])

    for record in records2:
        visits2.add(record[0] + "," + record[1])

    # Find users who visited on both days and visited a new page on the second day
    result = []

    for key in visits1:
        if key in visits2:
            continue  # User visited on both days

        # User visited only on the first day, check if they visited a new page on the second day
        user_id, _ = key.split(",")  
        new_page_visited = False
        for record in records2:
            if record[0] == user_id:
                new_key = record[0] + "," + record[1]
                if new_key not in visits1:
                    new_page_visited = True
                    break
        if new_page_visited:
            result.append(user_id)

    # Sort and print result
    result.sort()
    print("Users who visited some pages on both days and a new page on the second day:")
    for user_id in result:
        print(user_id)

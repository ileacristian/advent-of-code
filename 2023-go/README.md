To fetch the AOC input files I use the fetch_input_data.sh script that uses the following command:

wget --header "Cookie: session=your-session-cookie" -O dayXX.txt https://adventofcode.com/2023/day/XX/input

To get the session cookie:

- In your browser, open the developer tools (usually F12 or right-click and select "Inspect").
- Go to the "Network" tab.
- Reload the Advent of Code input page.
- Click on the first request (usually the top one in the list), which should be for the input page.
- Look for the "Request Headers" section and find the Cookie header.
- Copy the value of the session cookie. It should be a long string of characters.
{list1, list2} =
  File.stream!("./day1/day1.txt", :line)  # Stream the file line by line
    |> Stream.map(&String.trim/1)
    |> Stream.map(&String.split(&1, "   "))
    |> Enum.reduce({[], []}, fn [a, b], {list1, list2} ->
      {[String.to_integer(a) | list1], [String.to_integer(b) | list2]}         # Accumulate the two lists
    end)

#part 1
part1_result =Enum.sort(list1)
  |> Enum.zip(Enum.sort(list2))
  |> Enum.map(fn {a, b} -> abs(a - b) end)
  |> Enum.sum

IO.puts("Part1: #{part1_result}")

#part 2
part2_result = Enum.map(list1, fn elem -> elem * Enum.count(list2, &(&1 == elem)) end)
  |> Enum.sum

IO.puts("Part2: #{part2_result}")

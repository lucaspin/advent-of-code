count_tuple = File.read!("../input.txt")
  |> String.split(",")
  |> Enum.map(&String.to_integer/1)
  |> Enum.reduce(Tuple.duplicate(0, 9), fn number, tuple ->
    put_elem(tuple, number, elem(tuple, number) + 1)
  end)

Enum.reduce(1..256, count_tuple, fn _day, count_tuple ->
  {
    elem(count_tuple, 1),
    elem(count_tuple, 2),
    elem(count_tuple, 3),
    elem(count_tuple, 4),
    elem(count_tuple, 5),
    elem(count_tuple, 6),
    elem(count_tuple, 7) + elem(count_tuple, 0),
    elem(count_tuple, 8),
    elem(count_tuple, 0)
  }
end)
|> Tuple.sum()
|> IO.inspect()
